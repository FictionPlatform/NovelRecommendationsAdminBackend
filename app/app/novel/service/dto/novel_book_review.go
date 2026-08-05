package dto

import (
	"go-admin/core/dto"
)

type NovelReviewQueryReq struct {
	dto.Pagination `search:"-"`
	BookId         int64  `form:"bookId" search:"type:exact;column:book_id;table:app_novel_book_review" comment:"书籍编号"`
	Filter         string `form:"filter" search:"-" comment:"all-全部 five-5星 hot-热门"`
	CurrUserId     int64  `form:"-" search:"-" comment:"当前登录用户"`
}

func (m *NovelReviewQueryReq) GetNeedSearch() interface{} {
	return *m
}

type NovelReviewInsertReq struct {
	BookId     int64  `json:"bookId" comment:"书籍编号"`
	Rating     int    `json:"rating" comment:"评分 1~5"`
	Content    string `json:"content" comment:"书评内容"`
	CurrUserId int64  `json:"-" comment:"当前登录用户"`
}

type NovelReviewDeleteReq struct {
	Ids []int64 `json:"ids"`
}
