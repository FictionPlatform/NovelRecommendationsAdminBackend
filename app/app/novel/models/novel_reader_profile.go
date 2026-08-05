package models

import "time"

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
	CreatedAt            *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt            *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (NovelReaderProfile) TableName() string {
	return "app_novel_reader_profile"
}
