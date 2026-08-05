package models

import (
	"time"
)

type NovelBookReview struct {
	Id         int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	BookId     int64      `json:"bookId" gorm:"column:book_id;type:int;comment:书籍编号"`
	UserId     int64      `json:"userId" gorm:"column:user_id;type:int;comment:读者用户编号"`
	UserName   string     `json:"userName" gorm:"column:user_name;type:varchar(64);comment:昵称快照"`
	UserAvatar string     `json:"userAvatar" gorm:"column:user_avatar;type:varchar(500);comment:头像快照"`
	Rating     int        `json:"rating" gorm:"column:rating;type:tinyint;comment:评分1-5"`
	Content    string     `json:"content" gorm:"column:content;type:varchar(1000);comment:书评内容"`
	Likes      int        `json:"likes" gorm:"column:likes;type:int;comment:点赞数"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`

	// 扩展
	Book *NovelBook `json:"book" gorm:"-"`
}

func (NovelBookReview) TableName() string {
	return "app_novel_book_review"
}
