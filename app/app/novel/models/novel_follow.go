package models

import "time"

type NovelFollow struct {
	Id           int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId       int64      `json:"userId" gorm:"column:user_id;type:int;comment:关注者用户编号"`
	FollowUserId int64      `json:"followUserId" gorm:"column:follow_user_id;type:int;comment:被关注用户编号"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (NovelFollow) TableName() string {
	return "app_novel_follow"
}
