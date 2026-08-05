package dto

import (
	"go-admin/core/dto"
)

type NovelShelfQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64 `form:"-" search:"-" comment:"当前登录用户(限本人)"`
	CurrUserId     int64 `form:"-" search:"-" comment:"当前登录用户"`
}

func (m *NovelShelfQueryReq) GetNeedSearch() interface{} {
	return *m
}

type NovelShelfInsertReq struct {
	BookId     int64 `json:"bookId" comment:"书籍编号"`
	CurrUserId int64 `json:"-" comment:"当前登录用户"`
}

type NovelShelfGetReq struct {
	Id int64 `form:"id"`
}

type NovelShelfDeleteReq struct {
	Id         int64 `uri:"id"`
	CurrUserId int64 `form:"-" search:"-" comment:"当前登录用户"`
}
