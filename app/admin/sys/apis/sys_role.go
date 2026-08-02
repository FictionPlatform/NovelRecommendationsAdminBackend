package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/casbin"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/runtime"
)

type SysRole struct {
	api.Api
}

// GetList admin-获取角色管理全部列表
// @Summary 获取角色管理全部列表
// @Description 获取角色管理全部列表
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param id query int false "角色编码"
// @Param roleName query string false "角色名称"
// @Param status query string false "状态 1-正常 2-停用"
// @Param roleKey query string false "角色代码"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role/list [get]
func (e SysRole) GetList(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleQueryReq{}
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

// GetPage admin-获取角色管理分页列表
// @Summary 获取角色管理分页列表
// @Description 获取角色管理分页列表
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param id query int false "角色编码"
// @Param roleName query string false "角色名称"
// @Param status query string false "状态 1-正常 2-停用"
// @Param roleKey query string false "角色代码"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role [get]
func (e SysRole) GetPage(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleQueryReq{}
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

// Get admin-获取角色管理详情
// @Summary 获取角色管理详情
// @Description 获取角色管理详情
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色编码"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role/{id} [get]
func (e SysRole) Get(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetReq{}
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

// Insert admin-新增角色管理
// @Summary 新增角色管理
// @Description 新增角色管理
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param body body dto.SysRoleInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role [post]
func (e SysRole) Insert(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleInsertReq{}
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
	cb := runtime.RuntimeConfig.GetCasbinKey(c.Request.Host)
	id, respCode, err := s.Insert(&req, cb)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	_, _ = mycasbin.LoadPolicy(c)
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update admin-更新角色管理
// @Summary 更新角色管理
// @Description 更新角色管理
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色编码"
// @Param body body dto.SysRoleUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role/{id} [put]
func (e SysRole) Update(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleUpdateReq{}
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

	cb := runtime.RuntimeConfig.GetCasbinKey(c.Request.Host)
	b, respCode, err := s.Update(&req, cb)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	_, _ = mycasbin.LoadPolicy(c)
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete admin-删除角色管理
// @Summary 删除角色管理
// @Description 删除角色管理
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param body body dto.SysRoleDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role [delete]
func (e SysRole) Delete(c *gin.Context) {
	s := new(service.SysRole)
	req := dto.SysRoleDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	respCode, err := s.Delete(req.Ids)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	_, _ = mycasbin.LoadPolicy(c) //把最新决策加载到内存
	e.OK(req.Ids, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// UpdateStatus admin-更新角色管理状态
// @Summary 更新角色管理状态
// @Description 更新角色管理状态
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param body body dto.UpdateStatusReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role/role-status [put]
func (e SysRole) UpdateStatus(c *gin.Context) {
	s := service.SysRole{}
	req := dto.UpdateStatusReq{}
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
	respCode, err := s.UpdateStatus(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(req.RoleId, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// UpdateDataScope admin-更新角色管理数据权限
// @Summary 更新角色管理数据权限
// @Description 更新角色管理数据权限
// @Tags 系统角色管理
// @Accept json
// @Produce json
// @Param body body dto.RoleDataScopeReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-role/role-data-scope [put]
func (e SysRole) UpdateDataScope(c *gin.Context) {
	s := service.SysRole{}
	req := dto.RoleDataScopeReq{}
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
	b, respCode, err := s.UpdateDataScope(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	_, _ = mycasbin.LoadPolicy(c)
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
