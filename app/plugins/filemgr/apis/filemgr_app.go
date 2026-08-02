package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/plugins/filemgr/service"
	"go-admin/app/plugins/filemgr/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/dateutils"
	"mime/multipart"
	"time"
)

type FilemgrApp struct {
	api.Api
}

// GetPage plugins-获取APP管理分页列表
// @Summary 获取APP管理分页列表
// @Description 获取APP管理分页列表
// @Tags 文件管理-APP管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param version query string false "版本号"
// @Param platform query string false "平台 (1-安卓 2-苹果)"
// @Param appType query string false "版本(1-默认)"
// @Param downloadType query string false "下载类型(1-本地 2-外链 3-oss )"
// @Param status query string false "状态（1-已发布 2-待发布）"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app [get]
func (e FilemgrApp) GetPage(c *gin.Context) {
	req := dto.FilemgrAppQueryReq{}
	s := service.FilemgrApp{}
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

// Get plugins-获取APP管理详情
// @Summary 获取APP管理详情
// @Description 获取APP管理详情
// @Tags 文件管理-APP管理
// @Accept json
// @Produce json
// @Param id path int true "App编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app/{id} [get]
func (e FilemgrApp) Get(c *gin.Context) {
	req := dto.FilemgrAppGetReq{}
	s := service.FilemgrApp{}
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

// Insert plugins-新增APP管理
// @Summary 新增APP管理
// @Description 新增APP管理
// @Tags 文件管理-APP管理
// @Accept json
// @Produce json
// @Param body body dto.FilemgrAppInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app [post]
func (e FilemgrApp) Insert(c *gin.Context) {
	req := dto.FilemgrAppInsertReq{}
	s := service.FilemgrApp{}
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

// Delete plugins-删除APP管理
// @Summary 删除APP管理
// @Description 删除APP管理
// @Tags 文件管理-APP管理
// @Accept json
// @Produce json
// @Param body body dto.FilemgrAppDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app [delete]
func (e FilemgrApp) Delete(c *gin.Context) {
	s := service.FilemgrApp{}
	req := dto.FilemgrAppDeleteReq{}
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

// Upload  plugins-上传APP
// @Summary 上传APP
// @Description 上传APP
// @Tags 文件管理-APP管理
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "APP安装包文件"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app/upload [post]
func (e FilemgrApp) Upload(c *gin.Context) {
	s := service.FilemgrApp{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	form, err := e.Context.MultipartForm()
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	//获取上传文件信息
	var filePath string
	file := &multipart.FileHeader{}

	respCode, err := s.GetSingleUploadFileInfo(form, file, &filePath)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}

	//保存上传文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		e.Error(baseLang.AppUploadLogCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.AppUploadCode, baseLang.AppUploadLogCode, err).Error())
		return
	}
	e.OK(filePath, lang.MsgByCode(baseLang.AppUploadSuccessCode, e.Lang))
}

// Update  plugins-更新APP管理
// @Summary 更新APP管理
// @Description 更新APP管理
// @Tags 文件管理-APP管理
// @Accept json
// @Produce json
// @Param id path int true "App编号"
// @Param body body dto.FilemgrAppUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app/{id} [put]
func (e FilemgrApp) Update(c *gin.Context) {
	req := dto.FilemgrAppUpdateReq{}
	s := service.FilemgrApp{}
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

// Export  plugins-导出APP管理
// @Summary 导出APP管理
// @Description 导出APP管理
// @Tags 文件管理-APP管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param version query string false "版本号"
// @Param platform query string false "平台 (1-安卓 2-苹果)"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/filemgr/filemgr-app/export [get]
func (e FilemgrApp) Export(c *gin.Context) {
	req := dto.FilemgrAppQueryReq{}
	s := service.FilemgrApp{}
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
	fileName := "filemgr-app_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
