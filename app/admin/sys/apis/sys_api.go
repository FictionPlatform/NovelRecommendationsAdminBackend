package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/service"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/dateutils"
	"time"
)

type SysApi struct {
	api.Api
}

// GetPage admin-获取接口管理分页列表
// @Summary 获取接口管理分页列表
// @Description 获取接口管理分页列表
// @Tags 系统接口管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param description query string false "功能描述"
// @Param path query string false "地址"
// @Param method query string false "请求方法"
// @Param apiType query string false "接口类型"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api [get]
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiQueryReq{}
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

// GetList admin-获取接口管理全部列表
// @Summary 获取接口管理全部列表
// @Description 获取接口管理全部列表
// @Tags 系统接口管理
// @Accept json
// @Produce json
// @Param description query string false "功能描述"
// @Param path query string false "地址"
// @Param method query string false "请求方法"
// @Param apiType query string false "接口类型"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api/list [get]
func (e SysApi) GetList(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiQueryReq{}
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
	list, _, respCode, err := s.GetList(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(list, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get admin-获取接口管理详情
// @Summary 获取接口管理详情
// @Description 获取接口管理详情
// @Tags 系统接口管理
// @Accept json
// @Produce json
// @Param id path int true "接口编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api/{id} [get]
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
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

// Update admin-更新接口管理
// @Summary 更新接口管理
// @Description 更新接口管理
// @Tags 系统接口管理
// @Accept json
// @Produce json
// @Param id path int true "接口编号"
// @Param body body dto.SysApiUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api/{id} [put]
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
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

// Delete admin-删除接口管理
// @Summary 删除接口管理
// @Description 删除接口管理
// @Tags 系统接口管理
// @Accept json
// @Produce json
// @Param body body dto.SysApiDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api [delete]
func (e SysApi) Delete(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
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

// Export admin-导出接口管理
// @Summary 导出接口管理
// @Description 导出接口管理
// @Tags 系统接口管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param description query string false "功能描述"
// @Param apiType query string false "接口类型"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api/export [get]
func (e SysApi) Export(c *gin.Context) {
	req := dto.SysApiQueryReq{}
	s := service.SysApi{}
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
	if err != nil || maxSize <= 0 {
		//配置缺失/非法时使用兜底值，避免导出数据量异常
		maxSize = constant.DefaultExportMaxSize
	}
	if maxSize > constant.ExportMaxSizeLimit {
		maxSize = constant.ExportMaxSizeLimit
	}

	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = maxSize
	req.PageSizeLimit = maxSize
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.Export(list)
	fileName := "api_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

// Sync admin-同步接口数据
// @Summary 同步接口数据
// @Description 将当前程序中的所有接口录入数据库
// @Tags 系统接口管理
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-api/sync [get]
func (e SysApi) Sync(c *gin.Context) {
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	respCode, err := s.Sync()
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
