package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/service"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"
	"time"
)

type SysOperLog struct {
	api.Api
}

// GetPage admin-获取操作日志分页列表
// @Summary 获取操作日志分页列表
// @Description 获取操作日志分页列表
// @Tags 系统日志管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param title query string false "操作模块"
// @Param method query string false "函数"
// @Param requestMethod query string false "请求方式"
// @Param operUrl query string false "访问地址"
// @Param operIp query string false "客户端ip"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-oper-log [get]
func (e SysOperLog) GetPage(c *gin.Context) {
	s := service.SysOperLog{}
	req := dto.SysOperLogQueryReq{}
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

// Get admin-获取操作日志详情
// @Summary 获取操作日志详情
// @Description 获取操作日志详情
// @Tags 系统日志管理
// @Accept json
// @Produce json
// @Param id path int true "日志编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-oper-log/{id} [get]
func (e SysOperLog) Get(c *gin.Context) {
	s := new(service.SysOperLog)
	req := dto.SysOperLogGetReq{}
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

// Delete admin-删除操作日志
// @Summary 删除操作日志
// @Description 删除操作日志
// @Tags 系统日志管理
// @Accept json
// @Produce json
// @Param body body dto.SysOperLogDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-oper-log [delete]
func (e SysOperLog) Delete(c *gin.Context) {
	s := new(service.SysOperLog)
	req := dto.SysOperLogDeleteReq{}
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

// Export admin-导出操作日志
// @Summary 导出操作日志
// @Description 导出操作日志
// @Tags 系统日志管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param title query string false "操作模块"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-oper-log/export [get]
func (e SysOperLog) Export(c *gin.Context) {
	req := dto.SysOperLogQueryReq{}
	s := service.SysOperLog{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	sysConfService := adminService.NewSysConfigService(&s.Service)
	maxSize, respCode, err := sysConfService.GetWithKeyInt("admin_sys_max_export_size")
	if err != nil {
		e.Error(respCode, err.Error())
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = maxSize
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.Export(list)
	fileName := "operlog_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
