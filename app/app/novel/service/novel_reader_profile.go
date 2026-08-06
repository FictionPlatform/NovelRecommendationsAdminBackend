package service

import (
	"encoding/json"
	"errors"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/utils/dberr"
	"go-admin/core/utils/encrypt"
	"gorm.io/gorm"
)

type NovelReaderProfile struct {
	service.Service
}

// NewNovelReaderProfileService app-实例化读者资料服务
func NewNovelReaderProfileService(s *service.Service) *NovelReaderProfile {
	var srv = new(NovelReaderProfile)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// Get app-查询我的读者资料（不存在则自动创建）
func (e *NovelReaderProfile) Get(userId int64) (*models.NovelReaderProfile, int, error) {
	if userId <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelReaderProfile{}
	err := e.Orm.Where("user_id = ?", userId).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 自动创建默认资料，昵称/头像同步 app_user
		userName, userAvatar, _, getUserErr := getUserSnapshot(&e.Service, userId)
		if getUserErr != nil {
			return nil, baseLang.DataNotFoundCode, getUserErr
		}
		now := time.Now()
		data.UserId = userId
		data.Nickname = userName
		data.Avatar = userAvatar
		data.NotifyComment = 1
		data.NotifyBookUpdate = 1
		data.CreatedAt = &now
		data.UpdatedAt = &now
		err = e.Orm.Create(data).Error
		if err != nil {
			// 并发下唯一索引兜底：重查已存在资料，幂等返回
			if dberr.IsDuplicateKey(err) {
				existing := &models.NovelReaderProfile{}
				if qErr := e.Orm.Where("user_id = ?", userId).First(existing).Error; qErr == nil {
					data = existing
					err = nil
				} else {
					return nil, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
				}
			} else {
				return nil, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
			}
		}
	} else if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 邮箱解密
	if data.Email != "" {
		if email, err := encrypt.AesDecrypt(data.Email, []byte(config.AuthConfig.SecretAes)); err == nil {
			data.EmailValue = email
		} else {
			data.EmailValue = data.Email
		}
	} else {
		data.EmailValue = ""
	}
	// 偏好分类
	if data.PreferredCategories != "" {
		_ = json.Unmarshal([]byte(data.PreferredCategories), &data.PreferredCategoriesL)
	}
	return data, baseLang.SuccessCode, nil
}

// Update app-更新我的读者资料
func (e *NovelReaderProfile) Update(c *dto.NovelProfileUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 获取/创建资料
	data, respCode, err := e.Get(c.CurrUserId)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Nickname != "" && data.Nickname != c.Nickname {
		if len([]rune(c.Nickname)) > 64 {
			return false, baseLang.NovelNicknameTooLongCode, lang.MsgErr(baseLang.NovelNicknameTooLongCode, e.Lang)
		}
		updates["nickname"] = c.Nickname
	}
	if c.Avatar != "" && data.Avatar != c.Avatar {
		updates["avatar"] = c.Avatar
	}
	if c.Bio != "" && data.Bio != c.Bio {
		if len([]rune(c.Bio)) > 500 {
			return false, baseLang.NovelBioTooLongCode, lang.MsgErr(baseLang.NovelBioTooLongCode, e.Lang)
		}
		updates["bio"] = c.Bio
	}
	if c.Email != data.EmailValue {
		if len([]rune(c.Email)) > 128 {
			return false, baseLang.NovelEmailTooLongCode, lang.MsgErr(baseLang.NovelEmailTooLongCode, e.Lang)
		}
		if c.Email == "" {
			// 支持清空邮箱
			updates["email"] = ""
		} else {
			email, err := encrypt.AesEncrypt(c.Email, []byte(config.AuthConfig.SecretAes))
			if err != nil {
				return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
			}
			updates["email"] = email
		}
	}
	if c.PreferredCategories != nil {
		categories, _ := json.Marshal(c.PreferredCategories)
		updates["preferred_categories"] = string(categories)
	}
	if c.NotifyComment != data.NotifyComment {
		updates["notify_comment"] = c.NotifyComment
	}
	if c.NotifyBookUpdate != data.NotifyBookUpdate {
		updates["notify_book_update"] = c.NotifyBookUpdate
	}
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.NovelReaderProfile{}).Where("user_id = ?", data.UserId).Updates(updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}
