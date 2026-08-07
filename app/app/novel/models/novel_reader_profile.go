package models

import (
	"time"

	"gorm.io/gorm"
)

type NovelReaderProfile struct {
	Id                   int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId               int64      `json:"userId" gorm:"column:user_id;type:int;uniqueIndex;comment:用户编号"`
	Nickname             string     `json:"nickname" gorm:"column:nickname;type:varchar(64);comment:昵称"`
	Avatar               string     `json:"avatar" gorm:"column:avatar;type:varchar(500);comment:头像"`
	Bio                  string     `json:"bio" gorm:"column:bio;type:varchar(500);comment:简介"`
	Email                string     `json:"-" gorm:"column:email;type:varchar(128);comment:邮箱(加密)"`
	EmailValue           string     `json:"email" gorm:"-"`
	PreferredCategories  string     `json:"-" gorm:"column:preferred_categories;type:json;comment:偏好分类"`
	PreferredCategoriesL []string   `json:"preferredCategories" gorm:"-"`
	NotifyComment        int        `json:"notifyComment" gorm:"column:notify_comment;type:tinyint;comment:评论消息通知"`
	NotifyBookUpdate     int        `json:"notifyBookUpdate" gorm:"column:notify_book_update;type:tinyint;comment:书籍更新通知"`
	BanPostUntil         *time.Time `json:"banPostUntil" gorm:"column:ban_post_until;type:datetime;comment:禁止发帖截止时间(空=未禁言)"`
	BanReason            string     `json:"banReason" gorm:"column:ban_reason;type:varchar(255);comment:禁言原因"`
	CreatedAt            *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt            *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (NovelReaderProfile) TableName() string {
	return "app_novel_reader_profile"
}

// BeforeCreate 空 JSON 列兜底：Postgres 的 JSON 类型不接受空字符串，插入前统一转为 []。
func (e *NovelReaderProfile) BeforeCreate(_ *gorm.DB) error {
	if e.PreferredCategories == "" {
		e.PreferredCategories = "[]"
	}
	return nil
}

// BeforeUpdate 同上，避免 map update 误写空字符串
func (e *NovelReaderProfile) BeforeUpdate(_ *gorm.DB) error {
	if e.PreferredCategories == "" {
		e.PreferredCategories = "[]"
	}
	return nil
}
