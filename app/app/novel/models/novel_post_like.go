package models

import "time"

type NovelPostLike struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	PostId    int64      `json:"postId" gorm:"column:post_id;type:int;comment:帖子编号"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`
	Type      string     `json:"type" gorm:"column:type;type:char(1);comment:1-点赞 2-踩"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (NovelPostLike) TableName() string {
	return "app_novel_post_like"
}
