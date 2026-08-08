package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/service"
	"go-admin/app/app/novel/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/loginlock"
	"go-admin/core/utils/reglimit"
)

type Auth struct {
	api.Api
}

// Login app-读者登录
// @Summary 读者登录
// @Description 读者账号密码登录（返回 JWT，RoleKey=reader）
// @Tags 读者认证
// @Accept json
// @Produce json
// @Param body body dto.NovelAuthLoginReq true "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/user/login [post]
func (e Auth) Login(c *gin.Context) {
	req := dto.NovelAuthLoginReq{}
	s := service.NovelAuth{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		loginlock.RecordFail("", iputils.GetClientIP(c))
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	clientIP := iputils.GetClientIP(c)
	if loginlock.Check(req.Username, clientIP) {
		msg := lang.MsgByCode(baseLang.SysUseLoginLockedCode, e.Lang)
		e.ErrorByHttpCode(http.StatusTooManyRequests, baseLang.SysUseLoginLockedCode, msg)
		return
	}
	if req.Username == "" || req.Password == "" {
		loginlock.RecordFail(req.Username, clientIP)
		e.Error(baseLang.ParamErrCode, lang.MsgByCode(baseLang.ParamErrCode, e.Lang))
		return
	}

	userResp, respCode, err := s.LoginVerify(&req)
	if err != nil {
		loginlock.RecordFail(req.Username, clientIP)
		e.Error(respCode, err.Error())
		return
	}
	loginlock.Clear(req.Username, clientIP)
	c.Set(authdto.LoginUserId, userResp.Id)
	c.Set(authdto.UserName, userResp.UserName)
	c.Set(authdto.RoleKey, constant.RoleKeyReader)
	auth.Auth.Login(c)
}

// Register app-读者注册
// @Summary 读者注册
// @Description 注册读者账号（自动创建读者资料，注册成功即签发 token）
// @Tags 读者认证
// @Accept json
// @Produce json
// @Param body body dto.NovelAuthRegisterReq true "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/user/register [post]
func (e Auth) Register(c *gin.Context) {
	req := dto.NovelAuthRegisterReq{}
	s := service.NovelAuth{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	if req.Username == "" || req.Password == "" {
		e.Error(baseLang.ParamErrCode, lang.MsgByCode(baseLang.ParamErrCode, e.Lang))
		return
	}
	// IP 注册频率限制（每天每 IP 最多 N 次，settings.yml 可配置）
	clientIP := iputils.GetClientIP(c)
	if reglimit.Check(clientIP) {
		e.Error(baseLang.NovelRegLimitCode, lang.MsgByCode(baseLang.NovelRegLimitCode, e.Lang))
		return
	}
	user, respCode, err := s.Register(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	reglimit.Record(clientIP)
	c.Set(authdto.LoginUserId, user.Id)
	c.Set(authdto.UserName, user.UserName)
	c.Set(authdto.RoleKey, constant.RoleKeyReader)
	auth.Auth.Login(c)
}

// CancelAccount app-读者主动注销（终态，需验证登录密码）
// @Summary 读者主动注销
// @Description 验证登录密码后注销当前账号（置为已注销，立即拒绝一切请求）
// @Tags 读者认证
// @Accept json
// @Produce json
// @Param body body dto.NovelCancelAccountReq true "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/user/cancel [post]
func (e Auth) CancelAccount(c *gin.Context) {
	req := dto.NovelCancelAccountReq{}
	s := service.NovelAuth{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	if req.Password == "" {
		e.Error(baseLang.ParamErrCode, lang.MsgByCode(baseLang.ParamErrCode, e.Lang))
		return
	}
	respCode, err := s.CancelAccount(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
