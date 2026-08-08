package dto

import (
	"go-admin/app/app/novel/models"
	"go-admin/core/dto"
)

type NovelFollowInsertReq struct {
	TargetUserId int64 `json:"targetUserId" comment:"被关注用户编号"`
	CurrUserId   int64 `json:"-" comment:"当前登录用户"`
}

type NovelFollowDeleteReq struct {
	TargetUserId int64 `uri:"targetUserId" comment:"被关注用户编号"`
	CurrUserId   int64 `form:"-" search:"-" comment:"当前登录用户"`
}

type NovelFollowQueryReq struct {
	dto.Pagination `search:"-"`
	CurrUserId     int64 `form:"-" search:"-" comment:"当前登录用户（我关注的人）"`
}

func (m *NovelFollowQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelFollowingItem 关注列表项（目标用户画像）
type NovelFollowingItem struct {
	UserId     int64  `json:"userId" comment:"目标用户编号"`
	Name       string `json:"name" comment:"昵称"`
	Avatar     string `json:"avatar" comment:"头像"`
	Bio        string `json:"bio" comment:"简介"`
	Status     string `json:"status" comment:"状态 1-正常 2-禁言 3-注销"`
	IsFollowed bool   `json:"isFollowed" comment:"是否已关注"`
}

// NovelProfileReq 书友名片查询
type NovelProfileReq struct {
	UserId     int64 `uri:"userId" comment:"目标用户编号"`
	CurrUserId int64 `form:"-" search:"-" comment:"当前登录用户"`
}

// NovelProfileResp 书友名片聚合响应
type NovelProfileResp struct {
	UserId           int64                     `json:"userId"`
	UniqueId         *int64                    `json:"uniqueId,omitempty" comment:"唯一9位数字ID（仅本人可见）"`
	Name             string                    `json:"name"`
	Avatar           string                    `json:"avatar"`
	Bio              string                    `json:"bio"`
	Status           string                    `json:"status"`
	IsSelf           bool                      `json:"isSelf"`
	IsFollowed       bool                      `json:"isFollowed"`
	PostCount        int64                     `json:"postCount"`
	Posts            []models.NovelPost        `json:"posts"`
	ReviewCount      int64                     `json:"reviewCount"`
	Reviews          []models.NovelBookReview  `json:"reviews"`
	CommentCount     int64                     `json:"commentCount"`
	Comments         []models.NovelPostComment `json:"comments"`
	CollectPostCount int64                     `json:"collectPostCount"`
	CollectPosts     []models.NovelPost        `json:"collectPosts"`
	BookCount        int64                     `json:"bookCount"`
	Books            []models.NovelBookshelf   `json:"books"`
	FollowCount      int64                     `json:"followCount"`
	FanCount         int64                     `json:"fanCount"`
}
