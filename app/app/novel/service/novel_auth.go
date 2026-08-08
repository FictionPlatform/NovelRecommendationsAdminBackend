package service

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	sysConfigModels "go-admin/app/admin/sys/models"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/utils/idgen"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NovelAuth struct {
	service.Service
}

// 用户唯一9位数字ID序列（admin_sys_config, config_type=3）
const (
	AppUserUniqueIDConfigType = "3"
	AppUserUniqueIDConfigKey  = "app_user_unique_id_seq"
	AppUserUniqueIDInit       = 100000101
)

// NewNovelAuthService app-实例化读者注册登录服务
func NewNovelAuthService(s *service.Service) *NovelAuth {
	var srv = new(NovelAuth)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// LoginVerify app-读者登录验证（失败统一返回“用户名或密码错误”，防用户名枚举）
// 禁言(status=2)可登录；注销(status=3)不可登录，同样返回通用错误
func (e *NovelAuth) LoginVerify(login *dto.NovelAuthLoginReq) (*userModels.User, int, error) {
	user := &userModels.User{}
	err := e.Orm.Where("user_name = ? AND status <> ?", login.Username, global.SysStatusCancelled).First(user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("novel reader login query user error:%s", err.Error())
		}
		return nil, baseLang.NovelUserPwdErrCode, lang.MsgErr(baseLang.NovelUserPwdErrCode, e.Lang)
	}
	if !user.CheckPwd(login.Password) {
		return nil, baseLang.NovelUserPwdErrCode, lang.MsgErr(baseLang.NovelUserPwdErrCode, e.Lang)
	}
	return user, baseLang.SuccessCode, nil
}

// Register app-读者注册（自动创建读者扩展资料，注册成功即可登录）
// 注册时从 admin_sys_config(config_type=3, app_user_unique_id_seq) 取号分配唯一9位数字ID：
// 事务内行锁取号+1并回写；若 unique_id 冲突（并发或序列落后），回滚后自动修复序列并重试一次
func (e *NovelAuth) Register(req *dto.NovelAuthRegisterReq) (*userModels.User, int, error) {
	var count int64
	if err := e.Orm.Model(&userModels.User{}).Where("user_name = ?", req.Username).Count(&count).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if count > 0 {
		return nil, baseLang.NovelUsernameExistCode, lang.MsgErr(baseLang.NovelUsernameExistCode, e.Lang)
	}

	var user *userModels.User
	for attempt := 0; attempt < 2; attempt++ {
		now := time.Now()
		user = &userModels.User{
			LevelId:   1,
			UserName:  req.Username,
			TrueName:  "- -",
			Money:     decimal.NewFromInt(0),
			RefCode:   idgen.InviteId(),
			Pwd:       req.Password,
			Status:    global.SysStatusOk,
			TreeLeaf:  global.SysStatusOk,
			CreateBy:  0,
			UpdateBy:  0,
			CreatedAt: &now,
			UpdatedAt: &now,
		}
		txErr := e.Orm.Transaction(func(tx *gorm.DB) error {
			seq, err := e.nextUniqueID(tx)
			if err != nil {
				return err
			}
			user.UniqueId = &seq
			// email/mobile/mobile_title 未提供时保持 NULL（唯一索引下多个空串会互相冲突），ref_code 已生成随机邀请码
			if err := tx.Omit("email", "mobile", "mobile_title").Create(user).Error; err != nil {
				return err
			}
			profile := &models.NovelReaderProfile{
				UserId:           user.Id,
				Nickname:         req.Username,
				Avatar:           req.Avatar,
				NotifyComment:    1,
				NotifyBookUpdate: 1,
				CreatedAt:        &now,
				UpdatedAt:        &now,
			}
			return tx.Create(profile).Error
		})
		if txErr == nil {
			// 播种 3 条种子系统公告（欢迎/新功能/社区规范）
			NewNovelNotificationService(&e.Service).Seed(user.Id)
			return user, baseLang.SuccessCode, nil
		}
		if !isUniqueIDConflict(txErr) {
			return nil, baseLang.NovelRegisterFailCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.NovelRegisterFailCode, baseLang.DataInsertLogCode, txErr)
		}
		// unique_id 冲突：自动修复一次最高值后重试
		if err := e.repairUniqueIDMax(); err != nil {
			return nil, baseLang.NovelRegisterFailCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.NovelRegisterFailCode, baseLang.DataInsertLogCode, err)
		}
	}
	return nil, baseLang.NovelRegisterFailCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.NovelRegisterFailCode, baseLang.DataInsertLogCode, errors.New("unique id allocate retry exceeded"))
}

// nextUniqueID 事务内行锁取号：读取当前最高值+1，回写 admin_sys_config，返回新ID
func (e *NovelAuth) nextUniqueID(tx *gorm.DB) (int64, error) {
	var cfg struct{ ConfigValue string }
	if err := tx.Table("admin_sys_config").
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("config_type = ? AND config_key = ?", AppUserUniqueIDConfigType, AppUserUniqueIDConfigKey).
		Scan(&cfg).Error; err != nil {
		return 0, err
	}
	if cfg.ConfigValue == "" {
		// 序列配置缺失时自动初始化（幂等：并发下靠 config_key 唯一索引兜底）
		now := time.Now()
		seed := &sysConfigModels.SysConfig{
			ConfigName:  "App-用户唯一ID序列",
			ConfigKey:   AppUserUniqueIDConfigKey,
			ConfigValue: strconv.FormatInt(AppUserUniqueIDInit, 10),
			ConfigType:  AppUserUniqueIDConfigType,
			IsFrontend:  "2",
			Remark:      "app读者注册唯一9位数字ID：当前已分配的最高值，每次注册时+1",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		}
		if err := tx.Create(seed).Error; err != nil {
			return 0, err
		}
		cfg.ConfigValue = strconv.FormatInt(AppUserUniqueIDInit, 10)
	}
	next, err := strconv.ParseInt(cfg.ConfigValue, 10, 64)
	if err != nil {
		return 0, err
	}
	next++
	if err := tx.Table("admin_sys_config").
		Where("config_type = ? AND config_key = ?", AppUserUniqueIDConfigType, AppUserUniqueIDConfigKey).
		Update("config_value", strconv.FormatInt(next, 10)).Error; err != nil {
		return 0, err
	}
	return next, nil
}

// repairUniqueIDMax unique_id 冲突后修复：序列落后于表内最高值时，回拨到 MAX(unique_id)
func (e *NovelAuth) repairUniqueIDMax() error {
	var maxID *int64
	if err := e.Orm.Model(&userModels.User{}).Select("MAX(unique_id)").Scan(&maxID).Error; err != nil {
		return err
	}
	if maxID == nil {
		return nil
	}
	var cur struct{ ConfigValue string }
	if err := e.Orm.Table("admin_sys_config").
		Where("config_type = ? AND config_key = ?", AppUserUniqueIDConfigType, AppUserUniqueIDConfigKey).
		Scan(&cur).Error; err != nil {
		return err
	}
	curVal, err := strconv.ParseInt(cur.ConfigValue, 10, 64)
	if err != nil {
		return err
	}
	if curVal >= *maxID {
		return nil
	}
	return e.Orm.Table("admin_sys_config").
		Where("config_type = ? AND config_key = ?", AppUserUniqueIDConfigType, AppUserUniqueIDConfigKey).
		Update("config_value", strconv.FormatInt(*maxID, 10)).Error
}

// isUniqueIDConflict 判断错误是否为 unique_id 唯一索引冲突
func isUniqueIDConflict(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	if !strings.Contains(s, "unique_id") {
		return false
	}
	return strings.Contains(s, "duplicate key") ||
		strings.Contains(s, "23505") ||
		strings.Contains(s, "Duplicate entry")
}

// CancelAccount app-读者主动注销（终态：验证登录密码，置 status=3 并改用户名为「用户已注销」）
func (e *NovelAuth) CancelAccount(req *dto.NovelCancelAccountReq) (int, error) {
	if req.CurrUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	user := &userModels.User{}
	err := e.Orm.First(user, req.CurrUserId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if user.Status == global.SysStatusCancelled {
		return baseLang.NovelUserCancelledCode, lang.MsgErr(baseLang.NovelUserCancelledCode, e.Lang)
	}
	// 验证登录密码，防误操作
	if !user.CheckPwd(req.Password) {
		return baseLang.NovelUserPwdWrongCode, lang.MsgErr(baseLang.NovelUserPwdWrongCode, e.Lang)
	}
	now := time.Now()
	updates := map[string]interface{}{
		"status":     global.SysStatusCancelled,
		"updated_at": now,
	}
	if user.UserName != NovelCancelledUserName {
		updates["user_name"] = NovelCancelledUserName
	}
	if err := e.Orm.Model(&userModels.User{}).Where("id = ?", user.Id).Updates(updates).Error; err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
