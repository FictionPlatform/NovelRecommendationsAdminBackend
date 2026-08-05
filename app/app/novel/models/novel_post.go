package models

import (
	"time"
)

type NovelPost struct {
	Id           int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId       int64      `json:"userId" gorm:"column:user_id;type:int;comment:作者用户编号"`
	UserName     string     `json:"userName" gorm:"column:user_name;type:varchar(64);comment:昵称快照"`
	UserAvatar   string     `json:"userAvatar" gorm:"column:user_avatar;type:varchar(500);comment:头像快照"`
	Title        string     `json:"title" gorm:"column:title;type:varchar(200);comment:长文标题"`
	Content      string     `json:"content" gorm:"column:content;type:text;comment:正文"`
	Summary      string     `json:"summary" gorm:"column:summary;type:varchar(500);comment:摘要"`
	WordCount    int        `json:"wordCount" gorm:"column:word_count;type:int;comment:字符数"`
	ReadTime     int        `json:"readTime" gorm:"column:read_time;type:int;comment:阅读分钟"`
	TopicTag     string     `json:"topicTag" gorm:"column:topic_tag;type:char(2);comment:话题标签字典app_novel_topic"`
	RefBookId    *int64     `json:"refBookId" gorm:"column:ref_book_id;type:int;comment:关联小说编号"`
	Likes        int        `json:"likes" gorm:"column:likes;type:int;comment:点赞数"`
	Dislikes     int        `json:"dislikes" gorm:"column:dislikes;type:int;comment:踩数"`
	Collections  int        `json:"collections" gorm:"column:collections;type:int;comment:收藏数"`
	CommentCount int        `json:"commentCount" gorm:"column:comment_count;type:int;comment:评论数"`
	Status       string     `json:"status" gorm:"column:status;type:char(1);comment:1-正常 2-删除 3-审核中"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`

	// 扩展
	RefBookTitle string             `json:"refBookTitle" gorm:"-"`
	IsLiked      bool               `json:"isLiked" gorm:"-"`
	IsDisliked   bool               `json:"isDisliked" gorm:"-"`
	IsCollected  bool               `json:"isCollected" gorm:"-"`
	Comments     []NovelPostComment `json:"comments" gorm:"-"`
}

func (NovelPost) TableName() string {
	return "app_novel_post"
}
