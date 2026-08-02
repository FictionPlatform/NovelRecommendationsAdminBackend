package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/utils/captchautils"
	"go-admin/core/utils/fileutils"
	"go-admin/core/utils/idgen"
	"io"
	"net/http"
	"strings"
)

// maxAvatarUploadSize 头像上传大小上限 2MB
const maxAvatarUploadSize = 2 << 20

type SysUser struct {
	api.Api
}

// GetPage admin-获取系统用户管理分页列表
// @Summary 获取系统用户管理分页列表
// @Description 获取系统用户管理分页列表
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param userId query int false "用户ID"
// @Param username query string false "用户名"
// @Param phone query string false "手机号"
// @Param email query string false "邮箱"
// @Param postId query int false "岗位"
// @Param deptId query int false "部门"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user [get]
func (e SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get admin-获取系统用户管理详情
// @Summary 获取系统用户管理详情
// @Description 获取系统用户管理详情
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/{id} [get]
func (e SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert admin-新增系统用户管理
// @Summary 新增系统用户管理
// @Description 新增系统用户管理
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param body body dto.SysUserInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user [post]
func (e SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
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
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update admin-更新系统用户管理
// @Summary 更新系统用户管理
// @Description 更新系统用户管理
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户编号"
// @Param body body dto.SysUserUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/{id} [put]
func (e SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.Update(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete admin-删除系统用户管理
// @Summary 删除系统用户管理
// @Description 删除系统用户管理
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param body body dto.SysUserDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user [delete]
func (e SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Delete(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// UpdateStatus admin-更新系统用户状态
// @Summary 更新系统用户状态
// @Description 更新系统用户状态
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param body body dto.SysUserStatusUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/update-status [put]
func (e SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserStatusUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
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

	//数据权限检查
	p := middleware.GetPermissionFromContext(c)

	b, respCode, err := s.UpdateStatus(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// ResetPwd admin-重置系统用户密码
// @Summary 重置系统用户密码
// @Description 重置系统用户密码
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param body body dto.ResetSysUserPwdReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/pwd-reset [put]
func (e SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
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

	//数据权限检查
	p := middleware.GetPermissionFromContext(c)

	b, respCode, err := s.ResetPwd(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// UpdateProfileAvatar admin-更新系统登录用户头像
// @Summary 更新系统登录用户头像
// @Description 更新系统登录用户头像
// @Tags 系统用户管理
// @Accept multipart/form-data
// @Produce json
// @Param avatar formData file true "头像文件"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/profile/avatar [post]
func (e SysUser) UpdateProfileAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserAvatarUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil || form == nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	files := form.File["avatar"]
	if len(files) == 0 {
		e.Error(baseLang.SysUseAvatarUploadErrCode, lang.MsgByCode(baseLang.SysUseAvatarUploadErrCode, e.Lang))
		return
	}
	file := files[0]
	// 头像大小限制 2MB
	if file.Size <= 0 || file.Size > maxAvatarUploadSize {
		e.Error(baseLang.SysUseAvatarUploadErrCode, lang.MsgByCode(baseLang.SysUseAvatarUploadErrCode, e.Lang))
		return
	}
	// 校验图片魔数，防止上传 html/脚本等非图片内容
	src, err := file.Open()
	if err != nil {
		e.Error(baseLang.SysUseAvatarUploadErrCode, lang.MsgByCode(baseLang.SysUseAvatarUploadErrCode, e.Lang))
		return
	}
	defer src.Close()
	head := make([]byte, 512)
	n, _ := io.ReadFull(src, head)
	if !strings.HasPrefix(http.DetectContentType(head[:n]), "image/") {
		e.Error(baseLang.SysUseAvatarUploadErrCode, lang.MsgByCode(baseLang.SysUseAvatarUploadErrCode, e.Lang))
		return
	}
	guid := idgen.UUID()
	reqPath := config.ApplicationConfig.FileRootPath + "admin/avatar/"
	err = fileutils.IsNotExistMkDir(reqPath)
	if err != nil {
		e.Error(baseLang.SysUseAvatarUploadErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.SysUseAvatarUploadErrCode, baseLang.SysUseAvatarUploadErrLogCode, err).Error())
		return
	}
	filPath := reqPath + guid + ".jpg"
	// 上传文件至指定目录
	err = c.SaveUploadedFile(file, filPath)
	if err != nil {
		e.Error(baseLang.SysUseAvatarUploadErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.SysUseAvatarUploadErrCode, baseLang.SysUseAvatarUploadErrLogCode, err).Error())
		return
	}
	// 数据权限检查
	req.Avatar = global.RouteRootPath + "/" + filPath

	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid

	p := middleware.GetPermissionFromContext(c)
	b, respCode, err := s.UpdateProfileAvatar(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(req.Avatar, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// UpdateProfilePwd admin-更新系统登录用户密码
// @Summary 更新系统登录用户密码
// @Description 更新系统登录用户密码
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param body body dto.UpdateSysUserPwdReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/profile/pwd [put]
func (e SysUser) UpdateProfilePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	// 数据权限检查
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.UpdateProfilePwd(req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// GetProfile admin-获取系统登录用户信息
// @Summary 获取系统登录用户信息
// @Description 获取系统登录用户信息
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/profile [get]
func (e SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	err := e.MakeContext(c).
		MakeOrm().
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

	user, respCode, err := s.GetProfile(uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(user, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// UpdateProfile admin-更新系统登录用户信息
// 当前登录用户才能更新自己的信息
// 受限的子账户登录时，为了数据安全，不能让用户通过Update方法/接口来修改自己账户
// @Summary 更新系统登录用户信息
// @Description 更新系统登录用户信息
// @Tags 系统用户管理
// @Accept json
// @Produce json
// @Param body body dto.SysUserUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/profile [put]
func (e SysUser) UpdateProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
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
	req.Id = uid
	b, respCode, err := s.UpdateProfile(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Login admin-登录系统
// @Summary 登录系统
// @Description 登录系统
// @Tags 系统登录
// @Accept json
// @Produce json
// @Param body body dto.LoginReq true "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /login [post]
func (e SysUser) Login(c *gin.Context) {
	req := dto.LoginReq{}
	s := service.SysUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	if req.Code == "" || req.Password == "" || req.Username == "" {
		e.Error(baseLang.ParamErrCode, lang.MsgByCode(baseLang.ParamErrCode, e.Lang))
		return
	}

	if config.ApplicationConfig.Mode != "dev" {
		if !captchautils.Verify(req.UUID, req.Code, true) {
			e.Error(baseLang.SysUseCapErrLogCode, lang.MsgByCode(baseLang.SysUseCapErrLogCode, e.Lang))
			return
		}
	}

	userResp, respCode, err := s.LoginVerify(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}

	c.Set(authdto.LoginUserId, userResp.Id)
	c.Set(authdto.UserName, userResp.Username)
	c.Set(authdto.RoleKey, userResp.Role.RoleKey)
	s.LoginLogToDB(c, constant.UserLoginStatus, lang.MsgByCode(baseLang.SysUseLoginOpCode, e.Lang), userResp.Id)
	auth.Auth.Login(c)
}

// LogOut admin-退出系统
// @Summary 退出系统
// @Description 退出系统
// @Tags 系统登录
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/logout [get]
func (e SysUser) LogOut(c *gin.Context) {
	s := new(service.SysUser)
	err := e.MakeContext(c).
		MakeOrm().
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
	s.LoginLogToDB(c, constant.UserLogoutStatus, lang.MsgByCode(baseLang.SysUseLoginOpCode, e.Lang), uid)
	auth.Auth.Logout(c)
}

// RefreshToken admin-刷新token
// @Summary 刷新token
// @Description 刷新token
// @Tags 系统登录
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-user/refresh-token [get]
func (e SysUser) RefreshToken(c *gin.Context) {
	auth.Auth.RefreshToken(c)
}

// GenCaptcha admin-获取图形验证码
// @Summary 获取图形验证码
// @Description 获取图形验证码
// @Tags 系统登录
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /captcha [get]
func (e SysUser) GenCaptcha(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	id, b64s, _, err := captchautils.DriverDigitFunc()
	if err != nil {
		e.Error(baseLang.SysUseGenCaptchaErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.SysUseGenCaptchaErrCode, baseLang.SysUseGenCaptchaErrLogCode, err).Error())
		return
	}
	resp := map[string]string{
		"data": b64s,
		"id":   id,
	}
	e.OK(resp, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
