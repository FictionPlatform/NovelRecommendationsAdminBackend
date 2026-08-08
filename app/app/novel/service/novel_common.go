package service

import (
	"encoding/json"
	"errors"
	"time"

	"go-admin/app/app/novel/models"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"gorm.io/gorm"
)

// NovelClickCachePrefix 详情点击当日去重缓存前缀（key=book:{bookId}:{userId}，TTL 24h）
const NovelClickCachePrefix = "novelClick"

// NovelCancelledUserName 注销用户默认用户名（后台/主动注销时覆盖 user_name）
const NovelCancelledUserName = "用户已注销"

// getUserSnapshot 读取读者昵称/头像快照（app_user 复用）
func getUserSnapshot(e *service.Service, userId int64) (string, string, int, error) {
	u := &userModels.User{}
	err := e.Orm.First(u, userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return u.UserName, u.Avatar, baseLang.SuccessCode, nil
}

// CheckReaderWritePermission 校验读者写操作权限（发书评/话题/回复/互动/关注等主动操作）：
//   - 注销(status=3)：拒绝
//   - 禁言中（status=2 或 ban_post_until 未过期）：拒绝
//   - 禁言截止已过期：惰性恢复（清空禁言字段，status=2 恢复为 1）后放行
// 返回 SuccessCode 表示放行，否则返回对应错误码与错误信息
func CheckReaderWritePermission(e *service.Service, userId int64) (int, error) {
	if userId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	u := &userModels.User{}
	err := e.Orm.First(u, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if u.Status == global.SysStatusCancelled {
		return baseLang.NovelUserCancelledCode, lang.MsgErr(baseLang.NovelUserCancelledCode, e.Lang)
	}

	profile := &models.NovelReaderProfile{}
	perr := e.Orm.Where("user_id = ?", userId).First(profile).Error
	if perr != nil && !errors.Is(perr, gorm.ErrRecordNotFound) {
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, perr)
	}
	var banUntil *time.Time
	if perr == nil {
		banUntil = profile.BanPostUntil
	}
	now := time.Now()
	// 禁言判定：截止时间未过期，或 status=2 且未设置截止时间（永久禁言）
	banned := (banUntil != nil && banUntil.After(now)) ||
		(u.Status == global.SysStatusBanned && banUntil == nil)
	if banned {
		return baseLang.NovelPostBannedCode, lang.MsgErr(baseLang.NovelPostBannedCode, e.Lang)
	}
	// 禁言截止已过期：惰性恢复
	if banUntil != nil {
		updates := map[string]interface{}{
			"ban_post_until": nil,
			"ban_reason":     "",
			"updated_at":     now,
		}
		if err := e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", userId).Updates(updates).Error; err != nil {
			e.Log.Errorf("lazy clear reader ban fields error:%s", err.Error())
		}
		if u.Status == global.SysStatusBanned {
			if err := e.Orm.Model(&userModels.User{}).Where("id = ?", userId).
				Updates(map[string]interface{}{"status": global.SysStatusOk, "updated_at": now}).Error; err != nil {
				e.Log.Errorf("lazy restore reader status error:%s", err.Error())
			}
		}
	}
	return baseLang.SuccessCode, nil
}

// fillShelfBooks 批量填充书架书本信息
func fillShelfBooks(e *service.Service, list []models.NovelBookshelf) {
	if len(list) == 0 {
		return
	}
	bookIds := make([]int64, 0, len(list))
	for _, s := range list {
		bookIds = append(bookIds, s.BookId)
	}
	var books []models.NovelBook
	err := e.Orm.Where("id in (?)", bookIds).Find(&books).Error
	if err != nil {
		return
	}
	bookMap := map[int64]models.NovelBook{}
	for _, b := range books {
		if b.Tags != "" {
			_ = json.Unmarshal([]byte(b.Tags), &b.TagList)
		}
		bookMap[b.Id] = b
	}
	for i := range list {
		if b, ok := bookMap[list[i].BookId]; ok {
			list[i].Book = &b
		}
	}
}
