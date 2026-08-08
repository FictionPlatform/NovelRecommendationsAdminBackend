package dto

import (
	"go-admin/core/dto"
)

type NovelPostQueryReq struct {
	dto.Pagination `search:"-"`
	TopicTag       string `form:"topicTag" search:"type:exact;column:topic_tag;table:app_novel_post" comment:"话题标签"`
	Sort           string `form:"sort" search:"-" comment:"latest-最新 hot-热门"`
	RefBookId      int64  `form:"refBookId" search:"type:exact;column:ref_book_id;table:app_novel_post" comment:"关联小说编号"`
	Mine           int    `form:"mine" search:"-" comment:"1-只看我的话题"`
	IsCollected    int    `form:"isCollected" search:"-" comment:"1-只看我收藏的话题"`
	CurrUserId     int64  `form:"-" search:"-" comment:"当前登录用户"`
}

func (m *NovelPostQueryReq) GetNeedSearch() interface{} {
	return *m
}

type NovelPostGetReq struct {
	Id         int64 `uri:"id"`
	CurrUserId int64 `form:"-" comment:"当前登录用户"`
}

type NovelPostInsertReq struct {
	Title      string `json:"title" comment:"长文标题"`
	Content    string `json:"content" comment:"正文 5000~10000字"`
	TopicTag   string `json:"topicTag" comment:"话题标签字典值"`
	RefBookId  *int64 `json:"refBookId" comment:"关联小说编号(可选)"`
	CurrUserId int64  `json:"-" comment:"当前登录用户"`
}

type NovelPostDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type NovelPostInteractReq struct {
	Id         int64  `json:"-" uri:"id" comment:"话题编号"`
	Type       string `json:"type" comment:"like-点赞 dislike-踩 collect-收藏"`
	Action     string `json:"action" comment:"add-增加 cancel-取消"`
	CurrUserId int64  `json:"-" comment:"当前登录用户"`
}

type NovelPostCommentInsertReq struct {
	PostId           int64  `json:"-" uri:"id" comment:"话题编号"`
	Content          string `json:"content" comment:"评论内容"`
	ParentId         *int64 `json:"parentId" comment:"楼中楼父评论编号(可选)"`
	ReplyToUser      string `json:"replyToUser" comment:"被回复人昵称"`
	ReplyToCommentId *int64 `json:"replyToCommentId" comment:"被回复评论编号"`
	CurrUserId       int64  `json:"-" comment:"当前登录用户"`
}

type NovelPostCommentDeleteReq struct {
	Id         int64 `uri:"id"`
	CurrUserId int64 `json:"-" comment:"当前登录用户"`
}

type NovelCommentQueryReq struct {
	dto.Pagination `search:"-"`
	CurrUserId     int64 `form:"-" search:"-" comment:"当前登录用户（限本人评论）"`
}
