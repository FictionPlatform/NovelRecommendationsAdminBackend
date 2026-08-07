package models

import "time"

type NovelFeedback struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:int;index;comment:提交用户编号"`
	UserName  string     `json:"userName" gorm:"column:user_name;type:varchar(64);comment:昵称快照"`
	Type      string     `json:"type" gorm:"column:type;type:char(1);comment:反馈类型 A~H"`
	Kind      string     `json:"kind" gorm:"column:kind;type:char(16);comment:feedback-反馈 complaint-投诉"`
	Content   string     `json:"content" gorm:"column:content;type:text;comment:反馈内容"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);comment:1-处理中 2-已处理"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (NovelFeedback) TableName() string {
	return "app_novel_feedback"
}
