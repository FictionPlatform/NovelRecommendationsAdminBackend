package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/plugins/msg/service"
	"go-admin/app/plugins/msg/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
)

type MsgCode struct {
	api.Api
}

// GetPage plugins-获取验证码管理分页列表
// @Summary 获取验证码管理分页列表
// @Description 获取验证码管理分页列表
// @Tags 消息管理-验证码管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/msg/msg-code [get]
func (e MsgCode) GetPage(c *gin.Context) {
	req := dto.MsgCodeQueryReq{}
	s := service.MsgCode{}
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

// Get plugins-获取验证码管理详情
// @Summary 获取验证码管理详情
// @Description 获取验证码管理详情
// @Tags 消息管理-验证码管理
// @Accept json
// @Produce json
// @Param id path int true "验证码编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/msg/msg-code/{id} [get]
func (e MsgCode) Get(c *gin.Context) {
	req := dto.MsgCodeGetReq{}
	s := service.MsgCode{}
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
