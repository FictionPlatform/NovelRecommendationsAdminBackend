package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
)

type SysDictData struct {
	api.Api
}

// GetPage admin-获取字典数据分页列表
// @Summary 获取字典数据分页列表
// @Description 获取字典数据分页列表
// @Tags 系统字典管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param dictLabel query string false "字典标签"
// @Param dictValue query string false "字典键值"
// @Param dictType query string false "字典类型"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dict/data [get]
func (e SysDictData) GetPage(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataQueryReq{}
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

// Get admin-获取字典数据详情
// @Summary 获取字典数据详情
// @Description 获取字典数据详情
// @Tags 系统字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典数据编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dict/data/{id} [get]
func (e SysDictData) Get(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataGetReq{}
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

// Insert admin-新增字典数据
// @Summary 新增字典数据
// @Description 新增字典数据
// @Tags 系统字典管理
// @Accept json
// @Produce json
// @Param body body dto.SysDictDataInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dict/data [post]
func (e SysDictData) Insert(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataInsertReq{}
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

// Update admin-更新字典数据
// @Summary 更新字典数据
// @Description 更新字典数据
// @Tags 系统字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典数据编号"
// @Param body body dto.SysDictDataUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dict/data/{id} [put]
func (e SysDictData) Update(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataUpdateReq{}
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

// Delete admin-删除字典数据
// @Summary 删除字典数据
// @Description 删除字典数据
// @Tags 系统字典管理
// @Accept json
// @Produce json
// @Param body body dto.SysDictDataDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dict/data [delete]
func (e SysDictData) Delete(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataDeleteReq{}
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

// GetList admin-获取字典数据全部列表
// @Summary 获取字典数据全部列表
// @Description 获取字典数据全部列表
// @Tags 系统字典管理
// @Accept json
// @Produce json
// @Param dictLabel query string false "字典标签"
// @Param dictValue query string false "字典键值"
// @Param dictType query string false "字典类型"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dict/data/select [get]
func (e SysDictData) GetList(c *gin.Context) {
	s := service.SysDictData{}
	req := dto.SysDictDataQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err.Error())
		return
	}
	list, respCode, err := s.GetList(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(list, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
