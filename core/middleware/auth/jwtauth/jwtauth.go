package jwtauth

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/appleboy/gin-jwt/v3/core"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	jwtIn "github.com/golang-jwt/jwt/v5"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/response"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/middleware/auth/casbin"
	"go-admin/core/runtime"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	JWTLoginPrefix      = "admin:jwt:login"
	JWTBlacklistPrefix  = "admin:jwt:blacklist"
	JWTDevicesPrefix    = "admin:jwt:devices"
	JWTActivityPrefix   = "admin:jwt:activity"
	JWTPwdChangedPrefix = "admin:jwt:pwdchanged"
)

// 设备锁管理器，防止竞态条件
type deviceLockManager struct {
	locks map[string]*sync.Mutex
	mu    sync.RWMutex
}

func newDeviceLockManager() *deviceLockManager {
	return &deviceLockManager{
		locks: make(map[string]*sync.Mutex),
	}
}

func (d *deviceLockManager) getLock(userID string) *sync.Mutex {
	d.mu.Lock()
	defer d.mu.Unlock()

	if lock, exists := d.locks[userID]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	d.locks[userID] = lock
	return lock
}

type JwtAuth struct {
	mw                *jwt.GinJWTMiddleware
	enableDeviceCheck bool
	enableBlacklist   bool
	maxDevices        int
	deviceLocks       *deviceLockManager
}

// validateSecrets 启动时校验 JWT/AES 密钥：必须非空、不得使用默认值或已泄露的旧值
func validateSecrets() error {
	invalid := func(name string) error {
		return fmt.Errorf("auth.%s 必须设置为自定义密钥，禁止使用空值、占位符或历史默认值", name)
	}

	secret := config.AuthConfig.Secret
	if secret == "" || secret == PlaceholderJwtSecret || secret == LegacyDefaultSecret {
		return invalid("secret")
	}

	secretAes := config.AuthConfig.SecretAes
	if secretAes == "" || secretAes == PlaceholderAesSecret || secretAes == LegacyDefaultSecret || secretAes == config.AuthConfig.Secret {
		return invalid("secretAes")
	}
	if len(secretAes) != 16 && len(secretAes) != 24 && len(secretAes) != 32 {
		return fmt.Errorf("auth.secretAes 长度必须为 16/24/32 字节（AES-128/192/256），当前为 %d 字节", len(secretAes))
	}
	return nil
}

// 已泄露到仓库历史/文档中的旧默认密钥，必须更换
const (
	LegacyDefaultSecret  = "admin-api-20231019-jason"
	PlaceholderJwtSecret = "CHANGE_ME_JWT_SECRET"
	PlaceholderAesSecret = "CHANGE_ME_AES_SECRET"
)

func NewJwtAuth() (*JwtAuth, error) {
	if err := validateSecrets(); err != nil {
		return nil, err
	}

	timeout := time.Duration(config.AuthConfig.Timeout) * time.Second
	if timeout <= 0 {
		timeout = 7200 * time.Second
	}

	maxRefresh := time.Duration(config.AuthConfig.MaxRefresh) * time.Second
	if maxRefresh <= 0 {
		maxRefresh = 604800 * time.Second
	}

	jwtAuth := &JwtAuth{
		enableDeviceCheck: config.AuthConfig.EnableDeviceCheck,
		enableBlacklist:   config.AuthConfig.EnableBlacklist,
		maxDevices:        max(1, config.AuthConfig.MaxDeviceCount),
		deviceLocks:       newDeviceLockManager(),
	}

	mw, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            config.ApplicationConfig.Name,
		SigningAlgorithm: "HS256",
		Key:              []byte(config.AuthConfig.Secret),
		Timeout:          timeout,
		MaxRefresh:       maxRefresh,

		IdentityKey:     authdto.LoginUserId,
		Authenticator:   jwtAuth.Authenticator,
		Authorizer:      jwtAuth.Authorizer,
		PayloadFunc:     jwtAuth.PayloadFunc,
		IdentityHandler: jwtAuth.IdentityHandler,

		LoginResponse:   jwtAuth.LoginResponse,
		LogoutResponse:  jwtAuth.LogoutResponse,
		RefreshResponse: jwtAuth.RefreshResponse,
		Unauthorized:    jwtAuth.Unauthorized,

		// 仅从 Authorization 头/cookie 取 token；不支持 query 传参（会泄露到日志/Referer/历史记录）
		TokenLookup:   "header: Authorization, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    false,
		TimeFunc:      time.Now,
	})
	if err != nil {
		return nil, err
	}

	jwtAuth.mw = mw
	return jwtAuth, nil
}

func (j *JwtAuth) AuthMiddlewareFunc() gin.HandlerFunc {
	return j.mw.MiddlewareFunc()
}

// AuthCheckRoleMiddlewareFunc 权限检查中间件
func (j *JwtAuth) AuthCheckRoleMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleKey := c.GetString(authdto.RoleKey)

		rLog := log.GetRequestLogger(c)
		var err error
		// 管理员直接通过
		if roleKey == constant.RoleKeyAdmin {
			c.Next()
			return
		}
		// 角色不存在或已停用：禁止访问（禁用角色即时生效）
		if !j.checkRoleEnabled(c, roleKey) {
			rLog.Warnf("role %s not exists or disabled", roleKey)
			response.Error(c, baseLang.ForbitErr, lang.MsgByCode(baseLang.ForbitErr, lang.GetAcceptLanguage(c)))
			c.Abort()
			return
		}
		for _, i := range casbin.CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				rLog.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
				c.Next()
				return
			}
		}

		e := runtime.RuntimeConfig.GetCasbinKey(c.Request.Host)
		if e == nil {
			rLog.Error("Casbin instance not found")
			response.Error(c, baseLang.ServerErr,
				lang.MsgByCode(baseLang.ServerErr, lang.GetAcceptLanguage(c)))
			return
		}
		res, err := e.Enforce(roleKey, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			rLog.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, baseLang.ServerErr, lang.MsgByCode(baseLang.ServerErr, lang.GetAcceptLanguage(c)))
			return
		}

		if !res {
			rLog.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, roleKey, c.Request.Method, c.Request.URL.Path, "The current request has no permission. Please confirm it!")
			response.Error(c, baseLang.ForbitErr, lang.MsgByCode(baseLang.ForbitErr, lang.GetAcceptLanguage(c)))
			c.Abort()
			return
		}

		rLog.Infof("isTrue: %v role: %s method: %s path: %s", res, roleKey, c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func (j *JwtAuth) Login(c *gin.Context) {
	j.mw.LoginHandler(c)
}

func (j *JwtAuth) Logout(c *gin.Context) {
	j.mw.LogoutHandler(c)
}

func (j *JwtAuth) RefreshToken(c *gin.Context) {
	j.mw.RefreshHandler(c)
}

func (j *JwtAuth) LoginResponse(c *gin.Context, token *core.Token) {
	lg := lang.GetAcceptLanguage(c)

	userIDStr, errCode, err := j.GetUserIdStr(c)
	if err != nil {
		j.unauthorized(c, http.StatusUnauthorized, errCode, err.Error())
		return
	}

	if config.ApplicationConfig.IsSingleLogin {
		j.revokeAllTokens(userIDStr)

		// set 用于单点登录，更新登录信息
		err := runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTLoginPrefix,
			userIDStr,
			token.AccessToken,
			config.AuthConfig.Timeout,
		)
		if err != nil {
			j.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
			return
		}
	}

	// 添加设备信息
	if j.enableDeviceCheck {
		deviceFP := j.extractDeviceFingerprint(c)
		// 记录设备
		if deviceFP == "" {
			log.GetRequestLogger(c).Warn("device fingerprint empty")
		} else {
			j.recordDevice(userIDStr, deviceFP)
		}
	}

	// 缓存记录用户最新登录状态
	j.updateLastActivity(userIDStr)

	// 无密码变更记录时，写入本次登录时间作为基准（不覆盖已有记录，避免多设备互踢）
	if cached, _ := runtime.RuntimeConfig.GetCacheAdapter().Get(JWTPwdChangedPrefix, userIDStr); cached == "" {
		_ = runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTPwdChangedPrefix,
			userIDStr,
			strconv.FormatInt(time.Now().Unix(), 10),
			config.AuthConfig.MaxRefresh,
		)
	}

	response.OK(c, gin.H{
		"token":    token.AccessToken,
		"username": c.GetString(authdto.UserName),
		"expire":   token.ExpiresAt,
	}, http.StatusOK, lang.MsgByCode(baseLang.SuccessCode, lg))
}

func (j *JwtAuth) RefreshResponse(c *gin.Context, token *core.Token) {
	if config.ApplicationConfig.IsSingleLogin {
		userIDStr, _, _ := j.GetUserIdStr(c)
		if userIDStr != "" {
			_ = runtime.RuntimeConfig.GetCacheAdapter().Set(
				JWTLoginPrefix,
				userIDStr,
				token.AccessToken,
				config.AuthConfig.Timeout,
			)
			j.updateLastActivity(userIDStr)
		}
	}
	response.OK(c, gin.H{
		"token":  token.AccessToken,
		"expire": token.ExpiresAt,
	}, http.StatusOK, lang.MsgByCode(baseLang.SuccessCode, lang.GetAcceptLanguage(c)))
}

func (j *JwtAuth) LogoutResponse(c *gin.Context) {
	_, err := j.revokeToken(c)
	if err != nil {
		log.GetRequestLogger(c).Warnf("Failed to revoke token during logout: %v", err)
	}
	response.OK(c, nil, baseLang.SysUseLogoutSuccessCode, lang.MsgByCode(baseLang.SysUseLogoutSuccessCode, lang.GetAcceptLanguage(c)))
}

// revokeToken 撤销Token
func (j *JwtAuth) revokeToken(c *gin.Context) (int, error) {
	claims, err := j.mw.GetClaimsFromJWT(c)
	if err != nil {
		return baseLang.AuthErr, lang.MsgErr(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	}
	userIDStr, errCode, err := j.GetUserIdStr(c)
	if err != nil {
		return errCode, err
	}

	// 1. 清除单点登录缓存
	if config.ApplicationConfig.IsSingleLogin {
		j.revokeAllTokens(userIDStr)
	}

	// 2. 将token加入黑名单
	if j.enableBlacklist {
		jit, ok := claims[authdto.JTI].(string)
		if ok && jit != "" {
			runtime.RuntimeConfig.GetCacheAdapter().Set(
				JWTBlacklistPrefix,
				jit,
				"1",
				config.AuthConfig.MaxRefresh,
			)
		}
	}

	// 3. 清除设备记录中的当前设备
	if j.enableDeviceCheck {
		currentDeviceFP := j.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[authdto.DeviceFingerprint].(string)

		if ok && currentDeviceFP != "" && savedDeviceFP != "" && currentDeviceFP == savedDeviceFP {
			j.removeDevice(userIDStr, currentDeviceFP)
		}
	}

	return http.StatusOK, nil
}

func (j *JwtAuth) GetUserIdStr(c *gin.Context) (string, int, error) {
	userID := c.GetInt64(authdto.LoginUserId)
	if userID <= 0 {
		return "", baseLang.AuthErr, lang.MsgErrf(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	}
	return strconv.FormatInt(userID, 10), baseLang.SuccessCode, nil
}

func (j *JwtAuth) GetUserId(c *gin.Context) (int64, int, error) {
	userID := c.GetInt64(authdto.LoginUserId)
	if userID <= 0 {
		return 0, baseLang.AuthErr, lang.MsgErrf(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	}
	return userID, baseLang.SuccessCode, nil
}

func (j *JwtAuth) GetRoleKey(c *gin.Context) string {
	return c.GetString(authdto.RoleKey)
}

// GetUserDevices 获取用户的所有设备
func (j *JwtAuth) GetUserDevices(userID string) []string {
	devices, err := runtime.RuntimeConfig.GetCacheAdapter().Get(JWTDevicesPrefix, userID)
	if err != nil || devices == "" {
		return []string{}
	}

	return strings.Split(devices, ",")
}

func (j *JwtAuth) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		authdto.LoginUserId:       claims[authdto.LoginUserId],
		authdto.RoleKey:           claims[authdto.RoleKey],
		authdto.JTI:               claims[authdto.JTI],
		authdto.DeviceFingerprint: claims[authdto.DeviceFingerprint],
	}
}

func (j *JwtAuth) PayloadFunc(data interface{}) jwtIn.MapClaims {
	claims := jwtIn.MapClaims{}
	if v, ok := data.(map[string]interface{}); ok {
		claims[authdto.LoginUserId] = v[authdto.LoginUserId]
		claims[authdto.RoleKey] = v[authdto.RoleKey]
		claims[authdto.JTI] = v[authdto.JTI]
		claims[authdto.DeviceFingerprint] = v[authdto.DeviceFingerprint]
	}
	return claims
}

func (j *JwtAuth) Authenticator(c *gin.Context) (interface{}, error) {
	userId, _, _ := j.GetUserId(c)
	if userId <= 0 {
		return nil, errors.New("incorrect Username or Password")
	}
	roleKey, _ := c.Get(authdto.RoleKey)
	resp := map[string]interface{}{
		authdto.LoginUserId:       userId,
		authdto.RoleKey:           roleKey,
		authdto.JTI:               idgen.UUID(),
		authdto.DeviceFingerprint: j.extractDeviceFingerprint(c),
	}
	return resp, nil
}

func (j *JwtAuth) Authorizer(c *gin.Context, data interface{}) bool {
	if v, ok := data.(map[string]interface{}); ok {
		userId, _ := v[authdto.LoginUserId]
		if userId != nil {
			c.Set(authdto.LoginUserId, int64(userId.(float64)))
		}
		roleKey, _ := v[authdto.RoleKey]
		if roleKey != nil {
			c.Set(authdto.RoleKey, roleKey)
		}
		return j.authCheck(c)
	}
	return false
}

func (j *JwtAuth) Unauthorized(c *gin.Context, httpCode int, message string) {
	_, _ = j.revokeToken(c)
	temp := strings.SplitN(message, "_", 1)
	errCode := httpCode
	if len(temp) == 2 {
		code, err := strconv.ParseInt(temp[0], 10, 64)
		if err == nil {
			errCode = int(code)
			message = temp[1]
		}
	}
	log.GetRequestLogger(c).Errorf("Unauthorized failed: %d-%s", errCode, message)
	response.ErrorByHttpCode(c, httpCode, errCode, message)
}

// removeDevice 移除设备
func (j *JwtAuth) removeDevice(userID string, deviceFP string) {
	lock := j.deviceLocks.getLock(userID)
	lock.Lock()
	defer lock.Unlock()

	deviceList := j.GetUserDevices(userID)
	var newList []string

	for _, d := range deviceList {
		if d != deviceFP {
			newList = append(newList, d)
		}
	}

	if len(newList) > 0 {
		runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTDevicesPrefix,
			userID,
			strings.Join(newList, ","),
			config.AuthConfig.MaxRefresh,
		)
	} else {
		runtime.RuntimeConfig.GetCacheAdapter().Del(JWTDevicesPrefix, userID)
	}
}

func (j *JwtAuth) unauthorized(c *gin.Context, httpCode, errCode int, message string) {
	c.Header("WWW-Authenticate", "JWT realm=\""+j.mw.Realm+"\"")
	if !j.mw.DisabledAbort {
		c.Abort()
	}

	j.mw.Unauthorized(c, httpCode, strconv.Itoa(errCode)+"_"+message)
}

// recordDevice 记录设备
func (j *JwtAuth) recordDevice(userIDStr string, deviceFP string) {
	lock := j.deviceLocks.getLock(userIDStr)
	lock.Lock()
	defer lock.Unlock()

	// 获取现有设备列表
	deviceList := j.GetUserDevices(userIDStr)

	// 添加新设备
	for _, d := range deviceList {
		if d == deviceFP {
			return
		}
	}

	deviceList = append(deviceList, deviceFP)

	// 限制设备数量
	if len(deviceList) > j.maxDevices {
		deviceList = deviceList[len(deviceList)-j.maxDevices:]
	}

	runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTDevicesPrefix,
		userIDStr,
		strings.Join(deviceList, ","),
		config.AuthConfig.MaxRefresh,
	)
}

// addDevice 校验并追加设备：加锁后读改写（原子化，避免并发丢失更新/突破上限）。
// 设备已存在返回 true；设备数已达上限返回 false（拒绝访问）；否则追加并写回。
// 说明：单实例内存缓存下 per-user 互斥锁即保证原子性；多实例 Redis 部署需改用 Lua 脚本保证跨进程原子
func (j *JwtAuth) addDevice(userIDStr string, deviceFP string) bool {
	lock := j.deviceLocks.getLock(userIDStr)
	lock.Lock()
	defer lock.Unlock()

	deviceList := j.GetUserDevices(userIDStr)

	for _, d := range deviceList {
		if d == deviceFP {
			return true
		}
	}

	if len(deviceList) >= j.maxDevices {
		return false
	}

	deviceList = append(deviceList, deviceFP)
	_ = runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTDevicesPrefix,
		userIDStr,
		strings.Join(deviceList, ","),
		config.AuthConfig.MaxRefresh,
	)
	return true
}

// extractDeviceFingerprint 提取设备指纹
func (j *JwtAuth) extractDeviceFingerprint(c *gin.Context) string {
	userAgent := c.Request.UserAgent()
	clientIP := c.ClientIP()

	// 如果这些为空，生成默认值
	if userAgent == "" {
		userAgent = "unknown"
	}
	if clientIP == "" {
		clientIP = "0.0.0.0"
	}

	// 确保有足够的数据生成指纹
	data := userAgent + "|" + clientIP

	acceptLanguage := c.GetHeader("Accept-Language")
	if acceptLanguage != "" {
		data += "|" + acceptLanguage
	}

	acceptEncoding := c.GetHeader("Accept-Encoding")
	if acceptEncoding != "" {
		data += "|" + acceptEncoding
	}

	return encrypt.SHA256VString(data)
}

// revokeAllTokens 撤销用户的所有Token
func (j *JwtAuth) revokeAllTokens(userID string) {
	// 使用锁确保原子性
	lock := j.deviceLocks.getLock(userID)
	lock.Lock()
	defer lock.Unlock()

	keys := []string{JWTLoginPrefix, JWTDevicesPrefix, JWTActivityPrefix}
	for _, prefix := range keys {
		runtime.RuntimeConfig.GetCacheAdapter().Del(prefix, userID)
	}
}

func (j *JwtAuth) authCheck(c *gin.Context) bool {
	rLog := log.GetRequestLogger(c)
	claims, err := j.mw.CheckIfTokenExpire(c)
	if err != nil {
		rLog.Error(err)
		return false
	}

	userIDStr, _, err := j.GetUserIdStr(c)
	if err != nil {
		rLog.Error(err.Error())
		return false
	}
	userID := c.GetInt64(authdto.LoginUserId)

	roleKeyToken, _ := claims[authdto.RoleKey].(string)

	// 校验用户状态与角色一致性：用户已删除/已停用，或角色已变更（当前 role_key 与 token 声明不一致）时旧 token 立即失效
	if !j.checkUserStatus(c, userID, roleKeyToken) {
		rLog.Error("user not exists, disabled or role changed")
		return false
	}

	// 校验密码变更：token 签发时间（iat）早于最近一次密码变更时间则失效
	if !j.checkPwdChanged(claims, userIDStr) {
		rLog.Error("token issued before password change")
		return false
	}

	token, err := j.mw.ParseToken(c)
	if err != nil {
		rLog.Error(err)
		return false
	}

	if config.ApplicationConfig.IsSingleLogin {
		// 从缓存获取该用户最新的token
		cacheToken, _ := runtime.RuntimeConfig.GetCacheAdapter().Get(JWTLoginPrefix, userIDStr)
		if cacheToken != token.Raw {
			rLog.Error(errors.New("only support single login"))
			j.revokeAllTokens(userIDStr)
			return false
		}
	}

	// 3. check blocklist
	if j.enableBlacklist {
		jti, ok := claims[authdto.JTI].(string)
		if ok && jti != "" {
			isBlacklisted, _ := runtime.RuntimeConfig.GetCacheAdapter().Get(JWTBlacklistPrefix, jti)
			if isBlacklisted != "" {
				rLog.Error(errors.New("token is blacklisted"))
				return false
			}
		}
	}

	// 4. device check
	if j.enableDeviceCheck {
		currentDeviceFP := j.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[authdto.DeviceFingerprint].(string)
		if !ok || currentDeviceFP == "" || savedDeviceFP == "" || currentDeviceFP != savedDeviceFP {
			rLog.Error(errors.New("device check error"))
			return false
		}

		// 设备追加统一走加锁的 addDevice（读改写原子化），
		// 避免并发请求互相覆盖设备列表或突破 maxDeviceCount 上限
		if !j.addDevice(userIDStr, currentDeviceFP) {
			rLog.Error(errors.New("too many devices login"))
			return false
		}
		// 设备已在列表中，允许访问
	}

	// 4. 更新最后活动时间
	j.updateLastActivity(userIDStr)
	return true
}

// MarkPwdChanged 记录密码变更时间，使该用户所有已签发 token 立即失效（需重新登录）
func MarkPwdChanged(userID int64) {
	_ = runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTPwdChangedPrefix,
		strconv.FormatInt(userID, 10),
		strconv.FormatInt(time.Now().Unix(), 10),
		config.AuthConfig.MaxRefresh,
	)
}

// checkUserStatus 校验用户与角色状态：用户已删除（记录不存在）/已停用（status != 正常），
// 或角色已变更（当前 role_key 与 token 声明不一致，即角色被重新分配/降权）时拒绝访问。
// 实时查库比对替代可过期的 JwtRolePrefix 缓存，避免缓存过期后旧角色声明恢复生效
func (j *JwtAuth) checkUserStatus(c *gin.Context, userID int64, tokenRoleKey string) bool {
	db := runtime.RuntimeConfig.GetDbByKey(c.Request.Host)
	if db == nil {
		return false
	}
	var user struct {
		RoleKey string
	}
	if err := db.Table("admin_sys_user").
		Select("admin_sys_role.role_key").
		Joins("left join admin_sys_role on admin_sys_role.id = admin_sys_user.role_id").
		Where("admin_sys_user.id = ? AND admin_sys_user.status = ?", userID, global.SysStatusOk).
		Scan(&user).Error; err != nil {
		return false
	}
	return user.RoleKey == tokenRoleKey
}

// checkRoleEnabled 校验角色状态：角色存在且未停用（status != 正常）时拒绝访问
func (j *JwtAuth) checkRoleEnabled(c *gin.Context, roleKey string) bool {
	db := runtime.RuntimeConfig.GetDbByKey(c.Request.Host)
	if db == nil {
		return false
	}
	var count int64
	if err := db.Table("admin_sys_role").
		Where("role_key = ? AND status = ?", roleKey, global.SysStatusOk).
		Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// checkPwdChanged 校验密码变更时间：token 签发时间（iat）早于最近一次密码变更时间则失效
func (j *JwtAuth) checkPwdChanged(claims jwtIn.MapClaims, userIDStr string) bool {
	changedAt, _ := runtime.RuntimeConfig.GetCacheAdapter().Get(JWTPwdChangedPrefix, userIDStr)
	if changedAt == "" {
		return true
	}
	ts, err := strconv.ParseInt(changedAt, 10, 64)
	if err != nil {
		return true
	}
	iat, ok := claims["iat"].(float64)
	if !ok {
		return false
	}
	return int64(iat) >= ts
}

// updateLastActivity 更新最后活动时间
func (j *JwtAuth) updateLastActivity(userID string) {
	// 可以记录用户最后活动时间，用于会话管理
	runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTActivityPrefix,
		userID,
		strconv.FormatInt(time.Now().Unix(), 10),
		config.AuthConfig.MaxRefresh,
	)
}
