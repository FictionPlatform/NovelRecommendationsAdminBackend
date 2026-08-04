package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth/jwtauth"
	"go-admin/core/runtime"
	"go-admin/core/utils/dateutils"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/strutils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"

	cDto "go-admin/core/dto"
)

type SysUser struct {
	service.Service
}

// NewSysUserService admin-实例化用户管理
func NewSysUserService(s *service.Service) *SysUser {
	var srv = new(SysUser)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取系统用户管理分页列表
func (e *SysUser) GetPage(c *dto.SysUserQueryReq, p *middleware.DataPermission) ([]models.SysUser, int64, int, error) {
	var list []models.SysUser
	var data models.SysUser
	var count int64

	err := e.Orm.Model(&data).Preload("Dept").Preload("Role").Preload("Post").Order("created_at desc").
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

// Get admin-获取系统用户管理详情
func (e *SysUser) Get(id int64, p *middleware.DataPermission) (*models.SysUser, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysUser{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne admin-获取系统用户管理一条记录
func (e *SysUser) QueryOne(queryCondition *dto.SysUserQueryReq, p *middleware.DataPermission) (*models.SysUser, int, error) {
	data := &models.SysUser{}
	err := e.Orm.Model(&models.SysUser{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取系统用户管理数据总数
func (e *SysUser) Count(c *dto.SysUserQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysUser{}).
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

// checkRoleAssignable 校验角色可分配性：角色必须存在且未停用；admin 角色仅限 admin 操作者分配（防垂直越权提权）
func (e *SysUser) checkRoleAssignable(targetRoleId, opUserId int64) (int, error) {
	if targetRoleId <= 0 || opUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var role models.SysRole
	err := e.Orm.First(&role, targetRoleId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if role.Status != global.SysStatusOk {
		return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
	}
	if role.RoleKey == constant.RoleKeyAdmin {
		opRoleKey, respCode, err := e.getRoleKeyByUserId(opUserId)
		if err != nil {
			return respCode, err
		}
		if opRoleKey != constant.RoleKeyAdmin {
			return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
		}
	}
	return baseLang.SuccessCode, nil
}

// getRoleKeyByUserId 获取用户的角色标识
func (e *SysUser) getRoleKeyByUserId(userId int64) (string, int, error) {
	var roleKey string
	err := e.Orm.Table("admin_sys_user").
		Select("admin_sys_role.role_key").
		Joins("left join admin_sys_role on admin_sys_role.id = admin_sys_user.role_id").
		Where("admin_sys_user.id = ?", userId).
		Scan(&roleKey).Error
	if err != nil {
		return "", baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return roleKey, baseLang.SuccessCode, nil
}

// checkDeptAssignable 校验目标部门在操作者数据权限范围内（防越权跨部门授权）
func (e *SysUser) checkDeptAssignable(deptId, opUserId int64) (int, error) {
	if deptId <= 0 || opUserId <= 0 {
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
	// admin 或全部数据权限：不做部门限制
	if op.RoleKey == constant.RoleKeyAdmin || op.DataScope == "" || op.DataScope == constant.DataScope1 {
		return baseLang.SuccessCode, nil
	}
	switch op.DataScope {
	case constant.DataScope2:
		// 自定义数据权限：目标部门须在角色部门授权内
		var cnt int64
		if err := e.Orm.Table("admin_sys_role_dept").Where("role_id = ? AND dept_id = ?", op.RoleId, deptId).Count(&cnt).Error; err != nil {
			return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if cnt == 0 {
			return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
		}
	case constant.DataScope3:
		// 本部门：目标部门必须与操作者部门一致
		if op.DeptId != deptId {
			return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
		}
	case constant.DataScope4:
		// 本部门及以下：目标部门必须在操作者部门子树内（parent_ids 含操作者部门）
		var cnt int64
		if err := e.Orm.Table("admin_sys_dept").
			Where("id = ? AND (parent_ids LIKE ? OR id = ?)", deptId, "%,"+strconv.FormatInt(op.DeptId, 10)+",%", op.DeptId).
			Count(&cnt).Error; err != nil {
			return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if cnt == 0 {
			return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
		}
	case constant.DataScope5:
		// 仅本人：不允许创建用户或为他人分配部门
		return baseLang.ForbitErr, lang.MsgErr(baseLang.ForbitErr, e.Lang)
	}
	return baseLang.SuccessCode, nil
}

// Insert admin-新增系统用户管理
func (e *SysUser) Insert(c *dto.SysUserInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return 0, baseLang.SysUserNameEmptyCode, lang.MsgErr(baseLang.SysUserNameEmptyCode, e.Lang)
	}
	if c.NickName == "" {
		return 0, baseLang.SysNickNameEmptyCode, lang.MsgErr(baseLang.SysNickNameEmptyCode, e.Lang)
	}
	if c.Phone == "" {
		return 0, baseLang.SysUserPhoneEmptyCode, lang.MsgErr(baseLang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return 0, baseLang.SysUserEmailEmptyCode, lang.MsgErr(baseLang.SysUserEmailEmptyCode, e.Lang)
	}
	if c.DeptId <= 0 {
		return 0, baseLang.SysUserDeptEmptyCode, lang.MsgErr(baseLang.SysUserDeptEmptyCode, e.Lang)
	}
	if c.Password == "" {
		return 0, baseLang.SysUserPwdEmptyCode, lang.MsgErr(baseLang.SysUserPwdEmptyCode, e.Lang)
	}

	// 防垂直越权：校验角色可分配性与部门数据权限范围
	if c.RoleId > 0 {
		if respCode, err := e.checkRoleAssignable(c.RoleId, c.CurrUserId); err != nil {
			return 0, respCode, err
		}
	}
	if respCode, err := e.checkDeptAssignable(c.DeptId, c.CurrUserId); err != nil {
		return 0, respCode, err
	}

	if c.Username != "" {
		query := dto.SysUserQueryReq{}
		query.Username = c.Username
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, baseLang.SysUserNameExistCode, lang.MsgErr(baseLang.SysUserNameExistCode, e.Lang)
		}
	}
	if c.NickName != "" {
		query := dto.SysUserQueryReq{}
		query.NickName = c.NickName
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, baseLang.SysUserNickNameExistCode, lang.MsgErr(baseLang.SysUserNickNameExistCode, e.Lang)
		}
	}
	if c.Phone != "" {
		query := dto.SysUserQueryReq{}
		query.Phone = c.Phone
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, baseLang.SysUserPhoneExistCode, lang.MsgErr(baseLang.SysUserPhoneExistCode, e.Lang)
		}
	}
	if c.Email != "" {
		query := dto.SysUserQueryReq{}
		query.Email = c.Email
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, baseLang.SysUserEmailExistCode, lang.MsgErr(baseLang.SysUserEmailExistCode, e.Lang)
		}
	}

	if c.Avatar == "" {
		sysConfService := NewSysConfigService(&e.Service)
		defaultAvatar, respCode, err := sysConfService.GetWithKeyStr("admin_sys_user_default_avatar")
		if err != nil {
			return 0, respCode, err
		}
		c.Avatar = defaultAvatar
	}

	// insert data
	now := time.Now()
	data := models.SysUser{}
	data.Username = c.Username
	data.Password = c.Password
	data.NickName = c.NickName
	data.Phone = c.Phone
	data.RoleId = c.RoleId
	data.Avatar = c.Avatar
	data.Sex = c.Sex
	data.Email = c.Email
	data.DeptId = c.DeptId
	data.PostId = c.PostId
	data.Status = c.Status
	data.Remark = c.Remark
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update admin-更新系统用户管理
func (e *SysUser) Update(c *dto.SysUserUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return false, baseLang.SysUserNameEmptyCode, lang.MsgErr(baseLang.SysUserNameEmptyCode, e.Lang)
	}
	//if c.NickName == "" {
	//	return false, baseLang.SysNickNameEmptyCode, lang.MsgErr(baseLang.SysNickNameEmptyCode, e.Lang)
	//}
	if c.Phone == "" {
		return false, baseLang.SysUserPhoneEmptyCode, lang.MsgErr(baseLang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return false, baseLang.SysUserEmailEmptyCode, lang.MsgErr(baseLang.SysUserEmailEmptyCode, e.Lang)
	}
	/*	if c.DeptId <= 0 {
		return false, baseLang.SysUserDeptEmptyCode, lang.MsgErr(baseLang.SysUserDeptEmptyCode, e.Lang)
	}*/

	authChange := false //检查影响用户登录认证的字段是否发生变化，若发生变化，则需要强制该用户退出登录

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Username != "" && data.Username != c.Username {
		req := dto.SysUserQueryReq{}
		req.Username = c.Username
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserNameExistCode, lang.MsgErr(baseLang.SysUserNameExistCode, e.Lang)
		}
		updates["username"] = c.Username
	}
	if c.NickName != "" && data.NickName != c.NickName {
		req := dto.SysUserQueryReq{}
		req.NickName = c.NickName
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserNickNameExistCode, lang.MsgErr(baseLang.SysUserNickNameExistCode, e.Lang)
		}
		updates["nick_name"] = c.NickName
	}
	if c.Phone != "" && data.Phone != c.Phone {
		req := dto.SysUserQueryReq{}
		req.Phone = c.Phone
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserPhoneExistCode, lang.MsgErr(baseLang.SysUserPhoneExistCode, e.Lang)
		}
		updates["phone"] = c.Phone
	}
	if c.RoleId > 0 && data.RoleId != c.RoleId {
		// 防垂直越权：角色变更须可分配（存在/未停用，admin 角色仅限 admin 操作者分配）
		if respCode, err := e.checkRoleAssignable(c.RoleId, c.CurrUserId); err != nil {
			return false, respCode, err
		}
		updates["role_id"] = c.RoleId
		authChange = true //角色切换，需要重新登录
	}
	if c.Avatar != "" && data.Avatar != c.Avatar {
		updates["avatar"] = c.Avatar
	}
	if c.Sex != "" && data.Sex != c.Sex {
		updates["sex"] = c.Sex
	}
	if c.Email != "" && data.Email != c.Email {
		if !strutils.VerifyEmailFormat(c.Email) {
			return false, baseLang.SysUserEmailFormatErrCode, lang.MsgErr(baseLang.SysUserEmailFormatErrCode, e.Lang)
		}
		req := dto.SysUserQueryReq{}
		req.Email = c.Email
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserEmailExistCode, lang.MsgErr(baseLang.SysUserEmailExistCode, e.Lang)
		}
		updates["email"] = c.Email
	}
	if c.DeptId > 0 && data.DeptId != c.DeptId {
		// 防越权：部门变更须在操作者数据权限范围内
		if respCode, err := e.checkDeptAssignable(c.DeptId, c.CurrUserId); err != nil {
			return false, respCode, err
		}
		updates["dept_id"] = c.DeptId
	}
	if c.PostId > 0 && data.PostId != c.PostId {
		updates["post_id"] = c.PostId
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
	}
	if c.Remark != "" && data.Remark != c.Remark {
		updates["remark"] = c.Remark
	}
	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}

		if authChange {
			// 角色变更：强制该用户重新登录，旧 token 立即失效
			jwtauth.MarkPwdChanged(c.Id)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// UpdateStatus admin-更新系统用户状态
func (e *SysUser) UpdateStatus(c *dto.SysUserStatusUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Status == "" {
		return false, baseLang.SysUserStatusEmptyCode, lang.MsgErr(baseLang.SysUserStatusEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.UserId, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Status != "" && u.Status != c.Status {
		updates["status"] = c.Status
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.UserId).Updates(updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// ResetPwd admin-重置系统用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	var err error
	u, respCode, err := e.Get(c.UserId, p)
	if err != nil {
		return false, respCode, err
	}

	//新密码与旧密码一致时跳过（bcrypt 哈希与明文不可直接比较，用校验函数判断）
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(c.Password)) != nil {
		now := time.Now()
		err = e.Orm.Where("id=?", c.UserId).Updates(&models.SysUser{
			Password:  c.Password,
			UpdatedAt: &now,
			UpdateBy:  c.CurrUserId,
		}).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		// 密码变更后，该用户所有旧 token 立即失效
		jwtauth.MarkPwdChanged(c.UserId)
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Delete admin-删除系统用户管理
func (e *SysUser) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}

	//find if have admin account,not allow delete
	for _, id := range ids {
		u, respCode, err := e.Get(id, p)
		if err != nil {
			return respCode, err
		}
		if u.Username == constant.RoleKeyAdmin {
			return baseLang.SysAdminUserNotAllowDeleteErrCode, lang.MsgErr(baseLang.SysAdminUserNotAllowDeleteErrCode, e.Lang)
		}
	}
	var err error
	var data models.SysUser
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// GetProfile admin-获取系统登录用户信息
func (e *SysUser) GetProfile(userId int64) (*dto.SysUserResp, int, error) {
	if userId <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	user := &models.SysUser{}
	err := e.Orm.Preload("Dept").Preload("Post").Preload("Role").First(user, userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.SysUserNoExistCode, lang.MsgErr(baseLang.SysUserNoExistCode, e.Lang)
	}

	if user.Role.RoleKey == "" {
		return nil, baseLang.SysUserNoRoleErrCode, lang.MsgErr(baseLang.SysUserNoRoleErrCode, e.Lang)
	}

	respUser := &dto.SysUserResp{}
	respUser.Id = user.Id
	respUser.Email = user.Email
	respUser.Phone = user.Phone
	respUser.Username = user.Username
	respUser.Avatar = user.Avatar
	respUser.CreatedAt = dateutils.ConvertToStrByPrt(user.CreatedAt, -1)
	respUser.Sex = user.Sex
	respUser.DeptName = user.Dept.DeptName
	respUser.RoleName = user.Role.RoleName

	if user.Role.RoleKey == constant.RoleKeyAdmin {
		respUser.Permissions = []string{"*:*:*"}
	} else {
		roleService := NewSysRoleService(&e.Service)
		list, _, _ := roleService.GetPermissionsByRoleId(int64(user.RoleId))
		respUser.Permissions = list
	}
	respUser.RoleKyes = []string{user.Role.RoleKey}
	return respUser, baseLang.SuccessCode, nil
}

// UpdateProfile admin-更新系统登录用户信息
func (e *SysUser) UpdateProfile(c *dto.SysUserUpdateReq) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return false, baseLang.SysUserNameEmptyCode, lang.MsgErr(baseLang.SysUserNameEmptyCode, e.Lang)
	}
	if c.Phone == "" {
		return false, baseLang.SysUserPhoneEmptyCode, lang.MsgErr(baseLang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return false, baseLang.SysUserEmailEmptyCode, lang.MsgErr(baseLang.SysUserEmailEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, nil)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Sex != "" && data.Sex != c.Sex {
		updates["sex"] = c.Sex
	}
	if c.Username != "" && data.Username != c.Username {
		req := dto.SysUserQueryReq{}
		req.Username = c.Username
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserNameExistCode, lang.MsgErr(baseLang.SysUserNameExistCode, e.Lang)
		}
		updates["username"] = c.Username
	}
	if c.Phone != "" && data.Phone != c.Phone {
		req := dto.SysUserQueryReq{}
		req.Phone = c.Phone
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserPhoneExistCode, lang.MsgErr(baseLang.SysUserPhoneExistCode, e.Lang)
		}
		updates["phone"] = c.Phone
	}
	if c.Email != "" && data.Email != c.Email {
		if !strutils.VerifyEmailFormat(c.Email) {
			return false, baseLang.SysUserEmailFormatErrCode, lang.MsgErr(baseLang.SysUserEmailFormatErrCode, e.Lang)
		}
		req := dto.SysUserQueryReq{}
		req.Email = c.Email
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && resp.Id != data.Id {
			return false, baseLang.SysUserEmailExistCode, lang.MsgErr(baseLang.SysUserEmailExistCode, e.Lang)
		}
		updates["email"] = c.Email
	}
	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// LoginVerify admin-登录验证
func (e *SysUser) LoginVerify(login *dto.LoginReq) (*models.SysUser, int, error) {
	user := &models.SysUser{}
	// 所有账号（含 admin）统一校验 status=正常：被禁用的 admin 账号同样禁止登录
	err := e.Orm.Preload("Dept").Preload("Post").Preload("Role").Where("username = ? and status = ?", login.Username, global.SysStatusOk).First(user).Error
	if err != nil {
		// 记录真实原因用于排查，但对外统一返回“用户名或密码错误”，避免用户名枚举
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("LoginVerify query user error:%s", err.Error())
		}
		return nil, baseLang.SysUserPwdErrCode, lang.MsgErr(baseLang.SysUserPwdErrCode, e.Lang)
	}
	if !strutils.CompareHashAndPassword(user.Password, login.Password) {
		return nil, baseLang.SysUserPwdErrCode, lang.MsgErr(baseLang.SysUserPwdErrCode, e.Lang)
	}
	// 角色被删/未分配时 Preload 得到 nil，下游 userResp.Role.RoleKey 会 panic，此处直接拦截
	if user.Role == nil || user.Role.RoleKey == "" {
		return nil, baseLang.SysUserNoRoleErrCode, lang.MsgErr(baseLang.SysUserNoRoleErrCode, e.Lang)
	}
	return user, baseLang.SuccessCode, nil
}

// UpdateProfileAvatar admin-更新系统登录用户头像
func (e *SysUser) UpdateProfileAvatar(c *dto.SysUserAvatarUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Avatar != "" && u.Avatar != c.Avatar {
		updates["avatar"] = c.Avatar
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.CurrUserId).Updates(updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// UpdateProfilePwd admin-更新系统登录用户密码
func (e *SysUser) UpdateProfilePwd(c dto.UpdateSysUserPwdReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.NewPassword == "" {
		return false, baseLang.SysUserNewPwdEmptyCode, lang.MsgErr(baseLang.SysUserNewPwdEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, p)
	if err != nil {
		return false, respCode, err
	}

	if !strutils.CompareHashAndPassword(u.Password, c.OldPassword) {
		return false, baseLang.SysUserPwdErrCode, lang.MsgErr(baseLang.SysUserPwdErrCode, e.Lang)
	}

	if !strutils.CompareHashAndPassword(u.Password, c.NewPassword) {
		now := time.Now()
		u.Password = c.NewPassword
		u.UpdateBy = c.CurrUserId
		u.UpdatedAt = &now
		err = e.Orm.Where("id=?", c.CurrUserId).Updates(&models.SysUser{
			Password:  c.NewPassword,
			UpdatedAt: &now,
			UpdateBy:  c.CurrUserId,
		}).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		// 密码变更后，该用户所有旧 token 立即失效（含其他设备）
		jwtauth.MarkPwdChanged(c.CurrUserId)
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// LoginLogToDB admin-登录日志记录到数据库
func (e *SysUser) LoginLogToDB(c *gin.Context, status string, msg string, userId int64) {
	if !config.LoggerConfig.EnabledDB {
		return
	}
	l := make(map[string]interface{})

	clientIP := iputils.GetClientIP(c)
	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = clientIP
	//用于定位ip所在城市
	l["loginLocation"] = iputils.GetLocation(clientIP, config.ApplicationConfig.AmpKey)
	l["loginTime"] = strutils.GetCurrentTime()
	l["status"] = status
	l["agent"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["userId"] = userId
	l["remark"] = msg

	q := runtime.RuntimeConfig.GetMemoryQueue(c.Request.Host)
	message, err := runtime.RuntimeConfig.GetStreamMessage("", global.LoginLog, l)
	if err != nil {
		if e.Log != nil {
			e.Log.Errorf("SysUserService LoginLogToDB error:%s", err)
		}
		//日志报错错误，不中断请求
		return
	}
	err = q.Append(message)
	if err != nil && e.Log != nil {
		e.Log.Errorf("SysUserService LoginLogToDB error:%s", err)
	}
}

// LoginFailToDB admin-记录登录失败日志（用于审计暴力破解尝试）
func (e *SysUser) LoginFailToDB(c *gin.Context, username string, msg string) {
	remark := msg
	if username != "" {
		remark = "登录失败[账号:" + username + "]:" + msg
	}
	e.LoginLogToDB(c, constant.UserLoginFailStatus, remark, 0)
}
