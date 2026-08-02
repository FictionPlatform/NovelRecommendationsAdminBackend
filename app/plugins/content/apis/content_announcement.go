package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/plugins/content/service"
	"go-admin/app/plugins/content/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/dateutils"
	"time"
)

type ContentAnnouncement struct {
	api.Api
}

// GetPage plugins-获取公告管理分页列表
// @Summary 获取公告管理分页列表
// @Description 获取公告管理分页列表
// @Tags 内容管理-公告管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/content/content-announcement [get]
func (e ContentAnnouncement) GetPage(c *gin.Context) {
	req := dto.ContentAnnouncementQueryReq{}
	s := service.ContentAnnouncement{}
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

// Get plugins-获取公告管理详情
// @Summary 获取公告管理详情
// @Description 获取公告管理详情
// @Tags 内容管理-公告管理
// @Accept json
// @Produce json
// @Param id path int true "公告编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/content/content-announcement/{id} [get]
func (e ContentAnnouncement) Get(c *gin.Context) {
	req := dto.ContentAnnouncementGetReq{}
	s := service.ContentAnnouncement{}
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

// Insert plugins-新增公告管理
// @Summary 新增公告管理
// @Description 新增公告管理
// @Tags 内容管理-公告管理
// @Accept json
// @Produce json
// @Param body body dto.ContentAnnouncementInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/content/content-announcement [post]
func (e ContentAnnouncement) Insert(c *gin.Context) {
	req := dto.ContentAnnouncementInsertReq{}
	s := service.ContentAnnouncement{}
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

// Update plugins-更新公告管理
// @Summary 更新公告管理
// @Description 更新公告管理
// @Tags 内容管理-公告管理
// @Accept json
// @Produce json
// @Param id path int true "公告编号"
// @Param body body dto.ContentAnnouncementUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/content/content-announcement/{id} [put]
func (e ContentAnnouncement) Update(c *gin.Context) {
	req := dto.ContentAnnouncementUpdateReq{}
	s := service.ContentAnnouncement{}
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

// Delete plugins-删除公告管理
// @Summary 删除公告管理
// @Description 删除公告管理
// @Tags 内容管理-公告管理
// @Accept json
// @Produce json
// @Param body body dto.ContentAnnouncementDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/content/content-announcement [delete]
func (e ContentAnnouncement) Delete(c *gin.Context) {
	s := service.ContentAnnouncement{}
	req := dto.ContentAnnouncementDeleteReq{}
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

// Export plugins-导出公告管理
// @Summary 导出公告管理
// @Description 导出公告管理
// @Tags 内容管理-公告管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /plugins/content/content-announcement/export [get]
func (e ContentAnnouncement) Export(c *gin.Context) {
	req := dto.ContentAnnouncementQueryReq{}
	s := service.ContentAnnouncement{}
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
	fileName := "content-announcement_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
