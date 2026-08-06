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
	user, respCode, err := s.Register(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	c.Set(authdto.LoginUserId, user.Id)
	c.Set(authdto.UserName, user.UserName)
	c.Set(authdto.RoleKey, constant.RoleKeyReader)
	auth.Auth.Login(c)
}
