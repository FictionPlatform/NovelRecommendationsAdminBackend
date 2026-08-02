package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/app/user/service"
	"go-admin/app/app/user/service/dto"
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

type UserCountryCode struct {
	api.Api
}

// GetPage app-获取国家区号管理分页列表
// @Summary 获取国家区号管理分页列表
// @Description 获取国家区号管理分页列表
// @Tags 国家区号管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param country query string false "国家地区"
// @Param code query string false "区号"
// @Param status query string false "状态(1-可用 2-停用)"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-country-code [get]
func (e UserCountryCode) GetPage(c *gin.Context) {
	req := dto.UserCountryCodeQueryReq{}
	s := service.UserCountryCode{}
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

// Get app-获取国家区号管理详情
// @Summary 获取国家区号管理详情
// @Description 获取国家区号管理详情
// @Tags 国家区号管理
// @Accept json
// @Produce json
// @Param id path int true "区号编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-country-code/{id} [get]
func (e UserCountryCode) Get(c *gin.Context) {
	req := dto.UserCountryCodeGetReq{}
	s := service.UserCountryCode{}
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

// Insert app-新增国家区号管理
// @Summary 新增国家区号管理
// @Description 新增国家区号管理
// @Tags 国家区号管理
// @Accept json
// @Produce json
// @Param body body dto.UserCountryCodeInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-country-code [post]
func (e UserCountryCode) Insert(c *gin.Context) {
	req := dto.UserCountryCodeInsertReq{}
	s := service.UserCountryCode{}
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

// Update app-更新国家区号管理
// @Summary 更新国家区号管理
// @Description 更新国家区号管理
// @Tags 国家区号管理
// @Accept json
// @Produce json
// @Param id path int true "区号编号"
// @Param body body dto.UserCountryCodeUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-country-code/{id} [put]
func (e UserCountryCode) Update(c *gin.Context) {
	req := dto.UserCountryCodeUpdateReq{}
	s := service.UserCountryCode{}
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

// Delete app-删除国家区号管理
// @Summary 删除国家区号管理
// @Description 删除国家区号管理
// @Tags 国家区号管理
// @Accept json
// @Produce json
// @Param body body dto.UserCountryCodeDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-country-code [delete]
func (e UserCountryCode) Delete(c *gin.Context) {
	s := service.UserCountryCode{}
	req := dto.UserCountryCodeDeleteReq{}
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

// Export app-导出国家区号管理
// @Summary 导出国家区号管理
// @Description 导出国家区号管理
// @Tags 国家区号管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param country query string false "国家地区"
// @Param status query string false "状态(1-可用 2-停用)"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user-country-code/export [get]
func (e UserCountryCode) Export(c *gin.Context) {
	req := dto.UserCountryCodeQueryReq{}
	s := service.UserCountryCode{}
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
	fileName := "user-country-code_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
