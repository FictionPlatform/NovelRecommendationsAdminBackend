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

// ChangeStatus 后台-变更读者状态：1-正常 2-禁言（可带截止时间/原因）3-注销（终态，用户名改为「用户已注销」）
// 注销为终态不可恢复：目标已是注销用户时仅允许保持 status=3
func (e *NovelUser) ChangeStatus(c *dto.NovelUserStatusReq) (int, error) {
	if c.Id <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Status != global.SysStatusOk && c.Status != global.SysStatusBanned && c.Status != global.SysStatusCancelled {
		return baseLang.NovelUserStatusErrCode, lang.MsgErr(baseLang.NovelUserStatusErrCode, e.Lang)
	}
	if c.Status == global.SysStatusBanned {
		if c.BanUntil != nil && !c.BanUntil.After(time.Now()) {
			return baseLang.NovelBanTimeErrCode, lang.MsgErr(baseLang.NovelBanTimeErrCode, e.Lang)
		}
		if len([]rune(c.Reason)) > 255 {
			return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
		}
	}
	u := &userModels.User{}
	err := e.Orm.First(u, c.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 已注销为终态，禁止变更到其他状态
	if u.Status == global.SysStatusCancelled && c.Status != global.SysStatusCancelled {
		return baseLang.NovelUserCancelledCode, lang.MsgErr(baseLang.NovelUserCancelledCode, e.Lang)
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":     c.Status,
		"updated_at": now,
	}
	// 注销：用户名直接改为「用户已注销」
	if c.Status == global.SysStatusCancelled && u.UserName != NovelCancelledUserName {
		updates["user_name"] = NovelCancelledUserName
	}
	// 禁言：同步禁言截止时间/原因；正常：清空禁言字段
	if c.Status == global.SysStatusBanned {
		if respCode := e.ensureProfile(u, c.Id); respCode != baseLang.SuccessCode {
			return respCode, lang.MsgErr(baseLang.DataQueryCode, e.Lang)
		}
		err = e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", c.Id).
			Updates(map[string]interface{}{
				"ban_post_until": c.BanUntil,
				"ban_reason":     strings.TrimSpace(c.Reason),
				"updated_at":     now,
			}).Error
		if err != nil {
			return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
	} else if c.Status == global.SysStatusOk {
		err = e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", c.Id).
			Updates(map[string]interface{}{
				"ban_post_until": nil,
				"ban_reason":     "",
				"updated_at":     now,
			}).Error
		if err != nil {
			return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
	}
	err = e.Orm.Model(&userModels.User{}).Where("id = ?", c.Id).Updates(updates).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// ensureProfile 后台-确保读者资料存在（不存在则按 app_user 快照创建）
func (e *NovelUser) ensureProfile(u *userModels.User, userId int64) int {
	var count int64
	if err := e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return baseLang.DataQueryLogCode
	}
	if count > 0 {
		return baseLang.SuccessCode
	}
	now := time.Now()
	profile := &models.NovelReaderProfile{
		UserId:           userId,
		Nickname:         u.UserName,
		Avatar:           u.Avatar,
		NotifyComment:    1,
		NotifyBookUpdate: 1,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}
	if err := e.Orm.Create(profile).Error; err != nil {
		return baseLang.DataInsertLogCode
	}
	return baseLang.SuccessCode
}

// BanPost 后台-禁止/解除禁止读者发帖（banUntil 为空则解除；设置禁言时同步 status=2，解除时恢复 status=1）
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
	// 已注销为终态，禁止任何状态变更
	if u.Status == global.SysStatusCancelled {
		return baseLang.NovelUserCancelledCode, lang.MsgErr(baseLang.NovelUserCancelledCode, e.Lang)
	}
	// 资料不存在则自动创建（与 reader-profile 保持一致）
	if respCode := e.ensureProfile(u, c.Id); respCode != baseLang.SuccessCode {
		msg := baseLang.DataQueryCode
		if respCode == baseLang.DataInsertLogCode {
			msg = baseLang.DataInsertCode
		}
		return respCode, lang.MsgErr(msg, e.Lang)
	}
	updates := map[string]interface{}{
		"ban_post_until": c.BanUntil,
		"ban_reason":     strings.TrimSpace(c.Reason),
		"updated_at":     time.Now(),
	}
	if err := e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", c.Id).Updates(updates).Error; err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	// 状态联动：禁言=2，解除=1
	targetStatus := global.SysStatusOk
	if c.BanUntil != nil {
		targetStatus = global.SysStatusBanned
	}
	if u.Status != targetStatus {
		if err := e.Orm.Model(&userModels.User{}).Where("id = ?", c.Id).
			Updates(map[string]interface{}{"status": targetStatus, "updated_at": time.Now()}).Error; err != nil {
			return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
	}
	return baseLang.SuccessCode, nil
}

// NovelPostAdmin 后台话题管理服务
type NovelPostAdmin struct {
	service.Service
}

// NewNovelPostAdminService app-实例化话题管理服务
func NewNovelPostAdminService(s *service.Service) *NovelPostAdmin {
	var srv = new(NovelPostAdmin)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 后台-分页查询话题（全部状态，含已禁止访问）
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

// ChangeStatus 后台-禁止访问/恢复话题（status=2 后读者端列表与详情均不可见）
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
