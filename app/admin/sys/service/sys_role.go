package service

import (
	"errors"
	"fmt"
	"go-admin/config/base/constant"
	mycasbin "go-admin/core/casbin"

	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth/jwtauth"
	"gorm.io/gorm/clause"
	"time"

	"github.com/casbin/casbin/v2"

	"gorm.io/gorm"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	cDto "go-admin/core/dto"
)

type SysRole struct {
	service.Service
}

// NewSysRoleService admin-实例化角色管理
func NewSysRoleService(s *service.Service) *SysRole {
	var srv = new(SysRole)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetList admin-获取角色管理全部列表
func (e *SysRole) GetList(c *dto.SysRoleQueryReq, p *middleware.DataPermission) ([]models.SysRole, int64, int, error) {
	var list []models.SysRole
	var data models.SysRole
	var count int64

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// GetPage admin-获取角色管理分页列表
func (e *SysRole) GetPage(c *dto.SysRoleQueryReq, p *middleware.DataPermission) ([]models.SysRole, int64, int, error) {
	var list []models.SysRole
	var data models.SysRole
	var count int64

	err := e.Orm.Model(&data).Preload("SysMenu").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get admin-获取角色管理详情
func (e *SysRole) Get(id int64, p *middleware.DataPermission) (*models.SysRole, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysRole{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}

	menuIds, respCode, err := e.GetMenuIdsByRole(data.Id)
	if err != nil {
		return nil, respCode, err
	}
	deptIds, respCode, err := e.GetDeptIdsByRole(data.Id)
	if err != nil {
		return nil, respCode, err
	}
	data.MenuIds = menuIds
	data.DeptIds = deptIds
	return data, baseLang.SuccessCode, nil
}

// QueryOne admin-获取角色管理一条记录
func (e *SysRole) QueryOne(queryCondition *dto.SysRoleQueryReq, p *middleware.DataPermission) (*models.SysRole, int, error) {
	data := &models.SysRole{}
	err := e.Orm.Model(&models.SysRole{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Preload("SysMenu").First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取角色管理数据总数
func (e *SysRole) Count(c *dto.SysRoleQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysRole{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert admin-新增角色管理
func (e *SysRole) Insert(c *dto.SysRoleInsertReq, cb *casbin.SyncedEnforcer) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	if c.RoleName == "" {
		return 0, baseLang.SysRoleNameEmptyCode, lang.MsgErr(baseLang.SysRoleNameEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, baseLang.SysRoleStatusEmptyCode, lang.MsgErr(baseLang.SysRoleStatusEmptyCode, e.Lang)
	}
	if c.RoleKey == "" {
		return 0, baseLang.SysRoleKeyEmptyCode, lang.MsgErr(baseLang.SysRoleKeyEmptyCode, e.Lang)
	}
	if c.RoleSort < 0 {
		return 0, baseLang.SysRoleSortEmptyCode, lang.MsgErr(baseLang.SysRoleSortEmptyCode, e.Lang)
	}

	//确保角色key不存在
	req := dto.SysRoleQueryReq{}
	req.RoleKey = c.RoleKey
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.SysRoleKeyExistCode, lang.MsgErr(baseLang.SysRoleKeyExistCode, e.Lang)
	}

	//获取菜单
	sysMenuService := NewSysMenuService(&e.Service)
	sysMenuQueryReq := dto.SysMenuQueryReq{}
	sysMenuQueryReq.MenuIds = c.MenuIds
	sysMens, respCode, err := sysMenuService.GetList(&sysMenuQueryReq, true)
	if err != nil {
		return 0, respCode, err
	}

	//插入数据 菜单和权限
	now := time.Now()
	data := models.SysRole{}
	data.RoleName = c.RoleName
	data.Status = c.Status
	data.RoleKey = c.RoleKey
	data.RoleSort = c.RoleSort
	data.Remark = c.Remark
	data.DataScope = c.DataScope
	data.SysMenu = &sysMens
	data.SysDept = c.SysDept
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now

	err = e.Orm.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(&data).Error
		if err != nil {
			return err
		}
		adapter, err := mycasbin.NewAdapterByDB(tx) // 使用事务对象
		if err != nil {
			return fmt.Errorf("failed to create Casbin adapter: %w", err)
		}
		enforcer, err := casbin.NewSyncedEnforcer(cb.GetModel(), adapter)
		if err != nil {
			return fmt.Errorf("failed to create Casbin enforcer: %w", err)
		}
		for _, menu := range sysMens {
			for _, item := range menu.SysApi {
				_, err = enforcer.AddNamedPolicy("p", data.RoleKey, item.Path, item.Method)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update admin-更新角色管理
func (e *SysRole) Update(c *dto.SysRoleUpdateReq, cb *casbin.SyncedEnforcer) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error

	//确保角色key不存在
	req := dto.SysRoleQueryReq{}
	req.RoleKey = c.RoleKey
	role, respCode, err := e.QueryOne(&req, nil)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return false, respCode, err
	}
	if respCode == baseLang.SuccessCode && role.Id != c.Id {
		return false, baseLang.SysRoleKeyExistCode, lang.MsgErr(baseLang.SysRoleKeyExistCode, e.Lang)
	}
	if role != nil && role.RoleKey == constant.RoleKeyAdmin {
		return false, baseLang.SysRoleAdminNoOpCode, lang.MsgErr(baseLang.SysRoleAdminNoOpCode, e.Lang)
	}

	var data = models.SysRole{}
	err = e.Orm.Preload("SysMenu").First(&data, c.Id).Error
	if err != nil {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// admin 角色不可停用
	if data.RoleKey == constant.RoleKeyAdmin && c.Status != "" && c.Status != global.SysStatusOk {
		return false, baseLang.SysRoleAdminNoOpCode, lang.MsgErr(baseLang.SysRoleAdminNoOpCode, e.Lang)
	}

	//获取选中菜单的id对应的菜单完整信息
	var mlist = make([]models.SysMenu, 0)
	if err = e.Orm.Preload("SysApi").Where("id in ?", c.MenuIds).Find(&mlist).Error; err != nil {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	roleUpdates := map[string]interface{}{}
	//更新角色信息
	now := time.Now()
	if c.RoleName != "" && data.RoleName != c.RoleName {
		roleUpdates["role_name"] = c.RoleName
	}
	if c.Status != "" && data.Status != c.Status {
		roleUpdates["status"] = c.Status
	}
	// role_key 禁止修改：重命名后 casbin 策略仍挂旧 key，新 key 用户全部 403
	if c.RoleKey != "" && data.RoleKey != c.RoleKey {
		return false, baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
	}
	if c.RoleSort >= 0 && data.RoleSort != c.RoleSort {
		roleUpdates["role_sort"] = c.RoleSort
	}
	if c.Remark != "" && data.Remark != c.Remark {
		roleUpdates["remark"] = c.Remark
	}

	needUpdate := false
	err = e.Orm.Transaction(func(tx *gorm.DB) error {
		if len(roleUpdates) > 0 {
			data.UpdatedAt = &now
			data.UpdateBy = c.CurrUserId
			needUpdate = true
			if err = tx.Model(&data).Where("id=?", data.Id).Updates(roleUpdates).Error; err != nil {
				return err
			}
		}

		needUpdateMenu := len(*data.SysMenu) != len(mlist)
		if !needUpdateMenu {
			menuMap := make(map[int64]bool) // 假设 SysApi 的主键 ID 类型是 int64
			for _, api := range *data.SysMenu {
				menuMap[api.Id] = true // 用 ID 作为键值存储 SysApi
			}

			// 遍历 alist，检查每个 SysApi 是否在 data.SysApi 中存在
			for _, m := range mlist {
				if _, exists := menuMap[m.Id]; !exists {
					needUpdateMenu = true
					break
				}
			}
		}
		if needUpdateMenu {
			needUpdate = true
			data.SysMenu = &mlist
			if err = tx.Model(&data).Association("SysMenu").Replace(data.SysMenu); err != nil {
				return err // 如果更新失败，事务将自动回滚
			}
			if _, err = e.UpdateCasbin(mlist, data.RoleKey, tx, cb); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	if needUpdate {
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Delete admin-删除角色管理
func (e *SysRole) Delete(ids []int64) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userService := NewSysUserService(&e.Service)
	userReq := dto.SysUserQueryReq{}
	userReq.RoleIds = ids
	count, respCode, err := userService.Count(&userReq)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return baseLang.SysRoleUserExistNoDeleteCode, lang.MsgErr(baseLang.SysRoleUserExistNoDeleteCode, e.Lang)
	}

	for _, id := range ids {
		var role = models.SysRole{}
		err = tx.Preload("SysMenu").Preload("SysDept").First(&role, id).Error
		if err != nil {
			return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}

		if role.RoleKey == constant.RoleKeyAdmin {
			return baseLang.SysRoleAdminNoOpCode, lang.MsgErr(baseLang.SysRoleAdminNoOpCode, e.Lang)
		}
		err = tx.Select(clause.Associations).Delete(&role).Error
		if err != nil {
			return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
		}
	}
	return baseLang.SuccessCode, nil
}

// GetMenuIdsByRole admin-获取角色对应的菜单编号集合
func (e *SysRole) GetMenuIdsByRole(roleId int64) ([]int64, int, error) {
	if roleId <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	menuIds := make([]int64, 0)
	model := models.SysRole{}
	model.Id = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].Id)
	}
	return menuIds, baseLang.SuccessCode, nil
}

// GetDeptIdsByRole admin-获取角色对应的的部门编号集合
func (e *SysRole) GetDeptIdsByRole(roleId int64) ([]int64, int, error) {
	deptIds := make([]int64, 0)
	deptList := make([]dto.SysRoleDeptResp, 0)
	if err := e.Orm.Table("admin_sys_role_dept").
		Select("admin_sys_role_dept.dept_id").
		Joins("LEFT JOIN admin_sys_dept on admin_sys_dept.id=admin_sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" admin_sys_role_dept.dept_id not in(select admin_sys_dept.parent_id from admin_sys_role_dept LEFT JOIN admin_sys_dept on admin_sys_dept.id=admin_sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds, baseLang.SuccessCode, nil
}

// UpdateDataScope admin-更新角色管理数据权限
func (e *SysRole) UpdateDataScope(c *dto.RoleDataScopeReq) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	//查找角色id所属部门
	var data = models.SysRole{}
	if err := e.Orm.Preload("SysDept").First(&data, c.Id).Error; err != nil {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 禁止修改 admin 角色数据权限
	if data.RoleKey == constant.RoleKeyAdmin {
		return false, baseLang.SysRoleAdminNoOpCode, lang.MsgErr(baseLang.SysRoleAdminNoOpCode, e.Lang)
	}
	// 角色管辖范围校验：目标角色现有/新数据权限不得宽于操作者自身，自定权限部门须在操作者数据权限内
	if respCode, err := e.checkRoleScopeAssignable(data.DataScope, c.DataScope, c.DeptIds, c.CurrUserId); err != nil {
		return false, respCode, err
	}

	var dlist = make([]models.SysDept, 0)
	if err = e.Orm.Where("id in ?", c.DeptIds).Find(&dlist).Error; err != nil { //查找所选部门id对应的部门信息
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	dUpdates := map[string]interface{}{}
	now := time.Now()
	if c.DataScope != "" && data.DataScope != c.DataScope {
		dUpdates["data_scope"] = c.DataScope
	}

	needUpdate := false
	err = e.Orm.Transaction(func(tx *gorm.DB) error {
		if len(dUpdates) > 0 {
			data.UpdatedAt = &now
			data.UpdateBy = c.CurrUserId
			needUpdate = true
			if err = tx.Model(&data).Where("id=?", data.Id).Updates(dUpdates).Error; err != nil {
				return err
			}
		}

		needUpdateDept := len(data.SysDept) != len(dlist)
		if !needUpdateDept {
			deptMap := make(map[int64]bool)
			for _, item := range data.SysDept {
				deptMap[item.Id] = true
			}

			// 遍历 alist，检查每个 SysApi 是否在 data.SysApi 中存在
			for _, m := range dlist {
				if _, exists := deptMap[m.Id]; !exists {
					needUpdateDept = true
					break
				}
			}
		}
		if needUpdateDept {
			needUpdate = true
			data.SysDept = dlist
			if err = tx.Model(&data).Association("SysDept").Replace(data.SysDept); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	if needUpdate {
		// 数据权限变更：强制该角色下全部用户重新登录
		var userIds []int64
		if err = e.Orm.Table("admin_sys_user").Where("role_id = ?", c.Id).Pluck("id", &userIds).Error; err == nil {
			for _, uid := range userIds {
				jwtauth.MarkPwdChanged(uid)
			}
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// checkRoleScopeAssignable 校验操作者对目标角色数据权限的管辖范围：
// 目标角色现有/新数据权限均不得宽于操作者自身；自定数据权限的部门须在操作者数据权限范围内
func (e *SysRole) checkRoleScopeAssignable(targetDataScope, newDataScope string, deptIds []int64, opUserId int64) (int, error) {
	if opUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var op struct {
		RoleKey   string
		DataScope string
		DeptId    int64
		RoleId    int64
	}
	err := e.Orm.Table("admin_sys_user").
		Select("admin_sys_role.role_key", "admin_sys_role.data_scope", "admin_sys_user.dept_id", "admin_sys_user.role_id").
		Joins("left join admin_sys_role on admin_sys_role.id = admin_sys_user.role_id").
		Where("admin_sys_user.id = ?", opUserId).
		Scan(&op).Error
	if err != nil {
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// admin 或全部数据权限：可管理任意角色数据权限
	if op.RoleKey == constant.RoleKeyAdmin || op.DataScope == "" || op.DataScope == constant.DataScope1 {
		return baseLang.SuccessCode, nil
	}
	// 目标角色现有及新数据权限均不得宽于操作者自身
	if dataScopeRank(targetDataScope) > dataScopeRank(op.DataScope) {
		return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
	}
	if newDataScope != "" && dataScopeRank(newDataScope) > dataScopeRank(op.DataScope) {
		return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
	}
	// 自定数据权限：所选部门须在操作者数据权限范围内
	if newDataScope == constant.DataScope2 && len(deptIds) > 0 {
		userService := NewSysUserService(&e.Service)
		for _, deptId := range deptIds {
			if respCode, err := userService.checkDeptAssignable(deptId, opUserId); err != nil {
				return respCode, err
			}
		}
	}
	return baseLang.SuccessCode, nil
}

// dataScopeRank 数据权限宽窄排序：值越大范围越宽（1全部 > 2自定 > 4本部门及以下 > 3本部门 > 5仅本人）
func dataScopeRank(scope string) int {
	switch scope {
	case constant.DataScope1:
		return 5
	case constant.DataScope2:
		return 4
	case constant.DataScope4:
		return 3
	case constant.DataScope3:
		return 2
	default:
		return 1
	}
}

// UpdateStatus admin-更新角色管理状态
func (e *SysRole) UpdateStatus(c *dto.UpdateStatusReq) (int, error) {
	if c.RoleId <= 0 || c.CurrUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var model = models.SysRole{}
	err = tx.First(&model, c.RoleId).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	// admin 角色不可停用
	if model.RoleKey == constant.RoleKeyAdmin && c.Status != global.SysStatusOk {
		return baseLang.SysRoleAdminNoOpCode, lang.MsgErr(baseLang.SysRoleAdminNoOpCode, e.Lang)
	}
	if model.Status == c.Status {
		return baseLang.SuccessCode, nil
	}
	model.Status = c.Status
	err = tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	// 停用角色：强制该角色下全部用户重新登录，旧 token 立即失效
	if c.Status != global.SysStatusOk {
		var userIds []int64
		if err := e.Orm.Table("admin_sys_user").Where("role_id = ?", c.RoleId).Pluck("id", &userIds).Error; err == nil {
			for _, uid := range userIds {
				jwtauth.MarkPwdChanged(uid)
			}
		}
	}
	return baseLang.SuccessCode, nil
}

// GetWithName admin-根据角色名获取角色详情
func (e *SysRole) GetWithName(d *dto.SysRoleQueryReq) (*models.SysRole, int, error) {
	model := &models.SysRole{}
	err := e.Orm.Where("role_name = ?", d.RoleName).First(model).Error
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	menuIds, respCode, err := e.GetMenuIdsByRole(model.Id)
	if err != nil {
		return nil, respCode, err
	}

	deptIds, respCode, err := e.GetDeptIdsByRole(model.Id)
	if err != nil {
		return nil, respCode, err
	}

	model.MenuIds = menuIds
	model.DeptIds = deptIds
	return model, baseLang.SuccessCode, nil
}

// GetPermissionsByRoleId admin-根据角色获取权限
func (e *SysRole) GetPermissionsByRoleId(roleId int64) ([]string, int, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.Id = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return permissions, baseLang.SuccessCode, nil
}

// UpdateCasbin 更新casbin
func (e *SysRole) UpdateCasbin(mList []models.SysMenu, roleKey string, tx *gorm.DB, cb *casbin.SyncedEnforcer) (int, error) {
	adapter, err := mycasbin.NewAdapterByDB(tx) // 使用事务对象
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	enforcer, err := casbin.NewSyncedEnforcer(cb.GetModel(), adapter)
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	_, err = enforcer.RemoveFilteredPolicy(0, roleKey)
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	for _, menu := range mList {
		for _, api := range menu.SysApi {
			_, err = enforcer.AddNamedPolicy("p", roleKey, api.Path, api.Method)
			if err != nil {
				return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
			}
		}
	}
	return baseLang.SuccessCode, nil
}

// RebuildCasbinByRoles 按角色的当前菜单/API 重建 casbin 策略（权限回收）。
// 用于菜单/API 删除后，将受影响角色的策略重算为剩余权限，消除已删资源的残留策略。
// 必须在删除事务内调用（tx 保证读取删除后的菜单集合）。
func (e *SysRole) RebuildCasbinByRoles(roleIds []int64, tx *gorm.DB, cb *casbin.SyncedEnforcer) (int, error) {
	if len(roleIds) <= 0 || cb == nil {
		return baseLang.SuccessCode, nil
	}
	roleService := NewSysRoleService(&e.Service)
	roleService.Orm = tx
	for _, roleId := range roleIds {
		var role models.SysRole
		if err := tx.Preload("SysMenu.SysApi").First(&role, roleId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue // 角色已删除，其策略残留由角色删除流程负责
			}
			return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if role.RoleKey == "" {
			continue
		}
		mlist := make([]models.SysMenu, 0)
		if role.SysMenu != nil {
			mlist = append(mlist, (*role.SysMenu)...)
		}
		if respCode, err := roleService.UpdateCasbin(mlist, role.RoleKey, tx, cb); err != nil {
			return respCode, err
		}
	}
	return baseLang.SuccessCode, nil
}
