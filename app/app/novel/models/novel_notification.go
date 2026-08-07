package models

import "time"

type NovelNotification struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:int;index;comment:接收用户编号"`
	Title     string     `json:"title" gorm:"column:title;type:varchar(100);comment:通知标题"`
	Content   string     `json:"content" gorm:"column:content;type:text;comment:通知内容"`
	IsRead    string     `json:"isRead" gorm:"column:is_read;type:char(1);comment:0-未读 1-已读"`
	Source    string     `json:"source" gorm:"column:source;type:char(16);comment:system-系统 feedback-反馈 complaint-投诉 notice-公告"`
	NoticeId  *int64     `json:"noticeId" gorm:"column:notice_id;type:int;comment:来源公告编号"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (NovelNotification) TableName() string {
	return "app_novel_notification"
}
