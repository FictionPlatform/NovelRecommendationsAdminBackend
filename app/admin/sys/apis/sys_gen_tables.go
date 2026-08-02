package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
)

type SysTables struct {
	api.Api
}

// GetPage admin-获取表管理分页列表
// @Summary 获取表管理分页列表
// @Description 获取表管理分页列表
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param tableName query string false "表名"
// @Param tableComment query string false "表别名"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table [get]
func (e SysTables) GetPage(c *gin.Context) {
	s := service.SysGenTable{}
	req := dto.SysGenTableQueryReq{}
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

// Get admin-获取表管理详情
// @Summary 获取表管理详情
// @Description 获取表管理详情
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param id path int true "表编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/{id} [get]
func (e SysTables) Get(c *gin.Context) {
	req := dto.SysGenTableGetReq{}
	s := service.SysGenTable{}
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

// Insert admin-新增表管理
// @Summary 新增表管理
// @Description 新增表管理
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param body body dto.SysGenTableInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table [post]
func (e SysTables) Insert(c *gin.Context) {
	s := service.SysGenTable{}
	req := dto.SysGenTableInsertReq{}
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
	respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update admin-更新表管理
// @Summary 更新表管理
// @Description 更新表管理
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param id path int true "表编号"
// @Param body body dto.SysGenTableUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/{id} [put]
func (e SysTables) Update(c *gin.Context) {
	req := dto.SysGenTableUpdateReq{}
	s := service.SysGenTable{}

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

// Delete admin-删除表管理
// @Summary 删除表管理
// @Description 删除表管理
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param body body dto.SysGenTableDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table [delete]
func (e SysTables) Delete(c *gin.Context) {
	req := dto.SysGenTableDeleteReq{}
	s := service.SysGenTable{}
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

// GetDBTablePage admin-获取表管理的DB表分页列表
// @Summary 获取表管理的DB表分页列表
// @Description 获取表管理的DB表分页列表
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param tableName query string false "表名"
// @Param tableComment query string false "表别名"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/db-tables [get]
func (e SysTables) GetDBTablePage(c *gin.Context) {
	s := service.SysGenTable{}
	req := dto.DBTableQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	list, count, respCode, err := s.GetDBTablePage(req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Preview admin-预览表管理的代码页面
// @Summary 预览表管理的代码页面
// @Description 预览表管理的代码页面
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param id path int true "表编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/preview/{id} [get]
func (e SysTables) Preview(c *gin.Context) {
	s := service.SysGenTable{}
	req := dto.SysGenTableGenCodeReq{}
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
	resp, respCode, err := s.Preview(req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(resp, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// GenCode admin-生成表管理的代码
// @Summary 生成表管理的代码
// @Description 生成表管理的代码
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param id path int true "表编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/gen/{id} [get]
func (e SysTables) GenCode(c *gin.Context) {
	s := service.SysGenTable{}
	req := dto.SysGenTableGenCodeReq{}
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
	req.IsDownload = global.SysStatusNotOk
	_, respCode, err := s.GenCode(req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// DownloadCode admin-表管理下载代码
// @Summary 表管理下载代码
// @Description 表管理下载代码
// @Tags 系统代码生成
// @Accept json
// @Produce application/zip
// @Param id path int true "表编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/gen/download/{id} [get]
func (e SysTables) DownloadCode(c *gin.Context) {
	s := service.SysGenTable{}
	req := dto.SysGenTableGenCodeReq{}
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
	req.IsDownload = global.SysStatusOk
	resp, respCode, err := s.GenCode(req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	//resp不为空，则表示下载代码
	e.DownloadZip("code.zip", resp.Bytes())
}

// GenDB admin-表管理中生成菜单数据
// @Summary 表管理中生成菜单数据
// @Description 表管理中生成菜单数据
// @Tags 系统代码生成
// @Accept json
// @Produce json
// @Param id path int true "表编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-table/gen/db/{id} [get]
func (e SysTables) GenDB(c *gin.Context) {
	req := dto.SysGenTableGetReq{}
	s := service.SysGenTable{}
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

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.GenDB(req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
