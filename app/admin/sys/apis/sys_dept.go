package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
)

type SysDept struct {
	api.Api
}

// GetTree admin-获取部门管理树
// @Summary 获取部门管理树
// @Description 获取部门管理树
// @Tags 系统部门管理
// @Accept json
// @Produce json
// @Param deptName query string false "部门名称"
// @Param leader query string false "负责人"
// @Param phone query string false "手机"
// @Param email query string false "邮箱"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dept/dept-tree [get]
func (e SysDept) GetTree(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	list, respCode, err := s.GetTreeList(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(list, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get admin-获取部门管理详情
// @Summary 获取部门管理详情
// @Description 获取部门管理详情
// @Tags 系统部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dept/{id} [get]
func (e SysDept) Get(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetReq{}
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

// Insert admin-新增部门管理
// @Summary 新增部门管理
// @Description 新增部门管理
// @Tags 系统部门管理
// @Accept json
// @Produce json
// @Param body body dto.SysDeptInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dept [post]
func (e SysDept) Insert(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptInsertReq{}
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

// Update admin-更新部门管理
// @Summary 更新部门管理
// @Description 更新部门管理
// @Tags 系统部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门编号"
// @Param body body dto.SysDeptUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dept/{id} [put]
func (e SysDept) Update(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptUpdateReq{}
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

// Delete admin-删除部门管理
// @Summary 删除部门管理
// @Description 删除部门管理
// @Tags 系统部门管理
// @Accept json
// @Produce json
// @Param body body dto.SysDeptDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dept [delete]
func (e SysDept) Delete(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptDeleteReq{}
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

// GetDeptTreeByRole admin-根据角色获取部门
// @Summary 根据角色获取部门
// @Description 根据角色获取部门
// @Tags 系统部门管理
// @Accept json
// @Produce json
// @Param roleId path int true "角色编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /admin/sys/sys-dept/role-dept-tree-select/{roleId} [get]
func (e SysDept) GetDeptTreeByRole(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SelectDeptRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	result, respCode, err := s.GetTreeList(&dto.SysDeptQueryReq{})
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	deptIds := make([]int64, 0)
	if req.RoleId != 0 {
		sysRoleService := service.NewSysRoleService(&s.Service)
		deptIds, respCode, err = sysRoleService.GetDeptIdsByRole(req.RoleId)
		if err != nil {
			e.Error(respCode, err.Error())
			return
		}
	}
	resp := dto.DeptTreeRoleResp{
		Depts:       result,
		CheckedKeys: deptIds,
	}
	e.OK(resp, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
