package models

import (
	"time"
)

type NovelPostComment struct {
	Id               int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	PostId           int64      `json:"postId" gorm:"column:post_id;type:int;comment:帖子编号"`
	UserId           int64      `json:"userId" gorm:"column:user_id;type:int;comment:评论用户编号"`
	UserName         string     `json:"userName" gorm:"column:user_name;type:varchar(64);comment:昵称快照"`
	UserAvatar       string     `json:"userAvatar" gorm:"column:user_avatar;type:varchar(500);comment:头像快照"`
	ParentId         *int64     `json:"parentId" gorm:"column:parent_id;type:int;comment:楼中楼父评论编号"`
	ReplyToUser      string     `json:"replyToUser" gorm:"column:reply_to_user;type:varchar(64);comment:被回复人昵称"`
	ReplyToCommentId *int64     `json:"replyToCommentId" gorm:"column:reply_to_comment_id;type:int;comment:被回复评论编号"`
	Content          string     `json:"content" gorm:"column:content;type:varchar(1000);comment:评论内容"`
	Likes            int        `json:"likes" gorm:"column:likes;type:int;comment:点赞数"`
	CreatedAt        *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`

	// 扩展
	Replies   []NovelPostComment `json:"replies,omitempty" gorm:"-"`
	PostTitle string             `json:"postTitle" gorm:"-"`
}

func (NovelPostComment) TableName() string {
	return "app_novel_post_comment"
}
