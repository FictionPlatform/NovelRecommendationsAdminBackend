package dto

import (
	"go-admin/app/app/novel/models"
	"go-admin/core/dto"
)

// NovelNoticeInsertReq 发布公告请求
type NovelNoticeInsertReq struct {
	Title     string `json:"title" comment:"公告标题"`
	Content   string `json:"content" comment:"公告内容"`
	ValidDays int    `json:"validDays" comment:"有效天数(0=永久,过期后登录不再收取)"`
	CreateBy  int64  `json:"-" comment:"发布管理员编号"`
}

// NovelNoticeQueryReq 公告分页查询
type NovelNoticeQueryReq struct {
	dto.Pagination `search:"-"`
	Keyword        string `form:"keyword" search:"type:contains;column:title;table:app_novel_notice" comment:"标题关键字"`
}

func (m *NovelNoticeQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelNoticeDeleteReq 删除公告请求
type NovelNoticeDeleteReq struct {
	Ids []int64 `json:"ids" comment:"公告编号集合"`
}

// NovelNoticeItem 公告列表项（含读者触达数）
type NovelNoticeItem struct {
	models.NovelNotice
	RecipientCount int64 `json:"recipientCount" gorm:"-" comment:"触达读者数"`
}
