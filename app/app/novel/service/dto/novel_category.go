package dto

import (
	"go-admin/app/app/novel/models"
	"go-admin/core/dto"
)

// NovelCategoryQueryReq 分类分页查询
type NovelCategoryQueryReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name" search:"type:contains;column:name;table:app_novel_category" comment:"分类名称"`
	Status         string `form:"status" search:"type:exact;column:status;table:app_novel_category" comment:"状态 1-正常 2-停用"`
}

func (m *NovelCategoryQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelCategoryInsertReq 新增分类请求
type NovelCategoryInsertReq struct {
	Name     string `json:"name" comment:"分类名称"`
	Sort     int    `json:"sort" comment:"排序（越小越靠前）"`
	Status   string `json:"status" comment:"状态 1-正常 2-停用"`
	CreateBy int64  `json:"-" comment:"创建者"`
}

// NovelCategoryUpdateReq 更新分类请求
type NovelCategoryUpdateReq struct {
	Id       int64  `json:"-" uri:"id" comment:"分类编号"`
	Name     string `json:"name" comment:"分类名称"`
	Sort     int    `json:"sort" comment:"排序（越小越靠前）"`
	Status   string `json:"status" comment:"状态 1-正常 2-停用"`
	UpdateBy int64  `json:"-" comment:"更新者"`
}

// NovelCategoryDeleteReq 删除分类请求
type NovelCategoryDeleteReq struct {
	Ids []int64 `json:"ids" comment:"分类编号集合"`
}

// NovelCategoryGetReq 分类详情请求
type NovelCategoryGetReq struct {
	Id int64 `uri:"id"`
}

// NovelCategoryItem 分类列表项（含标签数）
type NovelCategoryItem struct {
	models.NovelCategory
	TagCount int64 `json:"tagCount" gorm:"-" comment:"标签数量"`
}

// NovelTagQueryReq 标签分页查询
type NovelTagQueryReq struct {
	dto.Pagination `search:"-"`
	CategoryId     int64  `form:"categoryId" search:"type:exact;column:category_id;table:app_novel_tag" comment:"所属分类id"`
	Name           string `form:"name" search:"type:contains;column:name;table:app_novel_tag" comment:"标签名称"`
	Status         string `form:"status" search:"type:exact;column:status;table:app_novel_tag" comment:"状态 1-正常 2-停用"`
}

func (m *NovelTagQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelTagInsertReq 新增标签请求
type NovelTagInsertReq struct {
	CategoryId int64  `json:"categoryId" comment:"所属分类id"`
	Name       string `json:"name" comment:"标签名称"`
	Sort       int    `json:"sort" comment:"排序（越小越靠前）"`
	Status     string `json:"status" comment:"状态 1-正常 2-停用"`
	CreateBy   int64  `json:"-" comment:"创建者"`
}

// NovelTagUpdateReq 更新标签请求
type NovelTagUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"标签编号"`
	CategoryId int64  `json:"categoryId" comment:"所属分类id"`
	Name       string `json:"name" comment:"标签名称"`
	Sort       int    `json:"sort" comment:"排序（越小越靠前）"`
	Status     string `json:"status" comment:"状态 1-正常 2-停用"`
	UpdateBy   int64  `json:"-" comment:"更新者"`
}

// NovelTagDeleteReq 删除标签请求
type NovelTagDeleteReq struct {
	Ids []int64 `json:"ids" comment:"标签编号集合"`
}

// NovelTagGetReq 标签详情请求
type NovelTagGetReq struct {
	Id int64 `uri:"id"`
}

// NovelTagItem 标签列表项（含分类名）
type NovelTagItem struct {
	models.NovelTag
}
