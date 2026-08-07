package service

import (
	"errors"
	"strings"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"gorm.io/gorm"
)

// NovelUser 后台读者管理服务
type NovelUser struct {
	service.Service
}

// NewNovelUserService app-实例化读者管理服务
func NewNovelUserService(s *service.Service) *NovelUser {
	var srv = new(NovelUser)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 后台-分页查询读者（读者资料 + app_user 用户名/状态/手机号）
func (e *NovelUser) GetPage(c *dto.NovelUserQueryReq) ([]dto.NovelUserItem, int64, int, error) {
	var data models.NovelReaderProfile
	var list []dto.NovelUserItem
	var count int64
	err := e.Orm.Model(&data).
		Select("app_novel_reader_profile.*, app_user.user_name, app_user.status, app_user.mobile").
		Joins("left join app_user on app_user.id = app_novel_reader_profile.user_id").
		Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("app_novel_reader_profile.created_at desc, app_novel_reader_profile.id desc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// ChangeStatus 后台-启用/禁用读者账户（status=2 后 jwtauth 实时拒绝其全部请求）
func (e *NovelUser) ChangeStatus(c *dto.NovelUserStatusReq) (int, error) {
	if c.Id <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Status != global.SysStatusOk && c.Status != global.SysStatusNotOk {
		return baseLang.NovelUserStatusErrCode, lang.MsgErr(baseLang.NovelUserStatusErrCode, e.Lang)
	}
	u := &userModels.User{}
	err := e.Orm.First(u, c.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	err = e.Orm.Model(&userModels.User{}).Where("id = ?", c.Id).
		Updates(map[string]interface{}{"status": c.Status, "updated_at": time.Now()}).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// BanPost 后台-禁止/解除禁止读者发帖（banUntil 为空则解除）
func (e *NovelUser) BanPost(c *dto.NovelBanPostReq) (int, error) {
	if c.Id <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.BanUntil != nil && !c.BanUntil.After(time.Now()) {
		return baseLang.NovelBanTimeErrCode, lang.MsgErr(baseLang.NovelBanTimeErrCode, e.Lang)
	}
	if len([]rune(c.Reason)) > 255 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	u := &userModels.User{}
	err := e.Orm.First(u, c.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 资料不存在则自动创建（与 reader-profile 保持一致）
	profile := &models.NovelReaderProfile{}
	perr := e.Orm.Where("user_id = ?", c.Id).First(profile).Error
	if perr != nil && !errors.Is(perr, gorm.ErrRecordNotFound) {
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, perr)
	}
	if errors.Is(perr, gorm.ErrRecordNotFound) {
		now := time.Now()
		profile = &models.NovelReaderProfile{
			UserId:           c.Id,
			Nickname:         u.UserName,
			Avatar:           u.Avatar,
			NotifyComment:    1,
			NotifyBookUpdate: 1,
			CreatedAt:        &now,
			UpdatedAt:        &now,
		}
		if err := e.Orm.Create(profile).Error; err != nil {
			return baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
		}
	}
	updates := map[string]interface{}{
		"ban_post_until": c.BanUntil,
		"ban_reason":     strings.TrimSpace(c.Reason),
		"updated_at":     time.Now(),
	}
	if err := e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", c.Id).Updates(updates).Error; err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// NovelPostAdmin 后台帖子管理服务
type NovelPostAdmin struct {
	service.Service
}

// NewNovelPostAdminService app-实例化帖子管理服务
func NewNovelPostAdminService(s *service.Service) *NovelPostAdmin {
	var srv = new(NovelPostAdmin)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 后台-分页查询帖子（全部状态，含已禁止访问）
func (e *NovelPostAdmin) GetPage(c *dto.NovelPostAdminQueryReq) ([]models.NovelPost, int64, int, error) {
	var data models.NovelPost
	var list []models.NovelPost
	var count int64
	err := e.Orm.Model(&data).
		Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("created_at desc, id desc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// ChangeStatus 后台-禁止访问/恢复帖子（status=2 后读者端列表与详情均不可见）
func (e *NovelPostAdmin) ChangeStatus(c *dto.NovelPostAdminStatusReq) (int, error) {
	if c.Id <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Status != global.SysStatusOk && c.Status != global.SysStatusNotOk {
		return baseLang.NovelPostStatusErrCode, lang.MsgErr(baseLang.NovelPostStatusErrCode, e.Lang)
	}
	post := &models.NovelPost{}
	err := e.Orm.First(post, c.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelPostNotExistCode, lang.MsgErr(baseLang.NovelPostNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	err = e.Orm.Model(&models.NovelPost{}).Where("id = ?", c.Id).
		Updates(map[string]interface{}{"status": c.Status, "updated_at": time.Now()}).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
