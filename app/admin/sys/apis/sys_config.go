package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/service"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/dateutils"
	"time"
)

type SysConfig struct {
	api.Api
}

// GetPage admin-获取配置管理分页列表
// @Summary 获取配置管理分页列表
// @Description 获取配置管理分页列表
// @Tags 系统配置管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param configName query string false "参数名称"
// @Param configKey query string false "参数键名"
// @Param configType query string false "参数类型"
// @Param isFrontend query string false "是否前台"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config [get]
func (e SysConfig) GetPage(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigQueryReq{}
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

// Get admin-获取配置管理详情
// @Summary 获取配置管理详情
// @Description 获取配置管理详情
// @Tags 系统配置管理
// @Accept json
// @Produce json
// @Param id path int true "配置编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config/{id} [get]
func (e SysConfig) Get(c *gin.Context) {
	req := dto.SysConfigGetReq{}
	s := service.SysConfig{}
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

// Insert admin-新增配置管理
// @Summary 新增配置管理
// @Description 新增配置管理
// @Tags 系统配置管理
// @Accept json
// @Produce json
// @Param body body dto.SysConfigInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config [post]
func (e SysConfig) Insert(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigInsertReq{}
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

// Update admin-更新配置管理
// @Summary 更新配置管理
// @Description 更新配置管理
// @Tags 系统配置管理
// @Accept json
// @Produce json
// @Param id path int true "配置编号"
// @Param body body dto.SysConfigUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config/{id} [put]
func (e SysConfig) Update(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigUpdateReq{}
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

// Delete admin-删除配置管理
// @Summary 删除配置管理
// @Description 删除配置管理
// @Tags 系统配置管理
// @Accept json
// @Produce json
// @Param body body dto.SysConfigDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config [delete]
func (e SysConfig) Delete(c *gin.Context) {
	s := service.SysConfig{}
	req := dto.SysConfigDeleteReq{}
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

// Export admin-导出配置管理
// @Summary 导出配置管理
// @Description 导出配置管理
// @Tags 系统配置管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param configName query string false "参数名称"
// @Param configKey query string false "参数键名"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config/export [get]
func (e SysConfig) Export(c *gin.Context) {
	req := dto.SysConfigQueryReq{}
	s := service.SysConfig{}
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
	fileName := "config_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}

// GetByKey admin-根据Key获取配置值
// @Summary 根据Key获取配置值
// @Description 根据Key获取配置值
// @Tags 系统配置管理
// @Accept json
// @Produce json
// @Param configKey path string true "参数键名"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-config/key/{configKey} [get]
func (e SysConfig) GetByKey(c *gin.Context) {
	var s = new(service.SysConfig)
	var req = new(dto.SysConfigByKeyReq)
	var resp = new(dto.SysConfigByKeyResp)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	resp, respCode, err := s.GetByKey(req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(resp, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
