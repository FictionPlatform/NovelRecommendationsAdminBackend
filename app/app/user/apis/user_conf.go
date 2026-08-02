package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/user/service"
	"go-admin/app/app/user/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
)

type UserConf struct {
	api.Api
}

// GetPage app-获取用户配置管理分页列表
// @Summary 获取用户配置管理分页列表
// @Description 获取用户配置管理分页列表
// @Tags 用户配置管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-conf [get]
func (e UserConf) GetPage(c *gin.Context) {
	req := dto.UserConfQueryReq{}
	s := service.UserConf{}
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
	req.ShowInfo = false
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get app-获取用户配置管理详情
// @Summary 获取用户配置管理详情
// @Description 获取用户配置管理详情
// @Tags 用户配置管理
// @Accept json
// @Produce json
// @Param id path int true "配置编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-conf/{id} [get]
func (e UserConf) Get(c *gin.Context) {
	req := dto.UserConfGetReq{}
	s := service.UserConf{}
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

// Update app-更新用户配置管理
// @Summary 更新用户配置管理
// @Description 更新用户配置管理
// @Tags 用户配置管理
// @Accept json
// @Produce json
// @Param id path int true "配置编号"
// @Param body body dto.UserConfUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-conf/{id} [put]
func (e UserConf) Update(c *gin.Context) {
	req := dto.UserConfUpdateReq{}
	s := service.UserConf{}
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
