package dto

import (
	"go-admin/app/app/novel/models"
	"go-admin/core/dto"
)

type NovelBookQueryReq struct {
	dto.Pagination `search:"-"`
	Category       string `form:"category" search:"type:exact;column:category;table:app_novel_book" comment:"分类"`
	SerialStatus   string `form:"serialStatus" search:"type:exact;column:serial_status;table:app_novel_book" comment:"连载状态 1-连载中 2-已完结"`
	Status         string `form:"status" search:"type:exact;column:status;table:app_novel_book" comment:"上架状态 1-上架 2-下架"`
	IsFeatured     int    `form:"isFeatured" search:"type:exact;column:is_featured;table:app_novel_book" comment:"首页精选 0|1"`
	WordCountMin   int    `form:"wordCountMin" search:"type:gte;column:word_count;table:app_novel_book" comment:"字数下限"`
	WordCountMax   int    `form:"wordCountMax" search:"type:lte;column:word_count;table:app_novel_book" comment:"字数上限"`
	Keyword        string `form:"keyword" search:"-" comment:"关键字(书名/作者/标签不区分大小写)"`
	Sort           string `form:"sort" search:"-" comment:"排序 rating|click|latest"`
	AllStatus      bool   `form:"allStatus" search:"-" comment:"查询全部上架状态（管理端）"`
	CategoryLabel  string `form:"categoryLabel" search:"-" comment:"分类标签"`
	StatusLabel    string `form:"statusLabel" search:"-" comment:"状态标签"`
	CurrUserId     int64  `form:"-" search:"-" comment:"当前登录用户"`
}

func (m *NovelBookQueryReq) GetNeedSearch() interface{} {
	return *m
}

type NovelBookInsertReq struct {
	Title        string   `json:"title" comment:"书名"`
	Author       string   `json:"author" comment:"作者"`
	Cover        string   `json:"cover" comment:"封面URL或base64"`
	SerialStatus string   `json:"serialStatus" comment:"连载状态 1-连载中 2-已完结"`
	Category     string   `json:"category" comment:"分类字典值"`
	Tags         []string `json:"tags" comment:"标签数组"`
	Slogan       string   `json:"slogan" comment:"一句话推荐语"`
	Description  string   `json:"description" comment:"简介"`
	PublishDate  *string  `json:"publishDate" comment:"发布日期 YYYY-MM-DD"`
	WordCount    int      `json:"wordCount" comment:"字数"`
	Chapters     int      `json:"chapters" comment:"章节数"`
	IsFeatured   int      `json:"isFeatured" comment:"首页精选 0|1"`
	ReadUrl      string   `json:"readUrl" comment:"阅读链接"`
	Status       string   `json:"status" comment:"上架状态 1-上架 2-下架"`
	CurrUserId   int64    `json:"-" comment:"当前登录用户"`
}

type NovelBookUpdateReq struct {
	Id           int64    `json:"-" uri:"id" comment:"书籍编号"`
	Title        string   `json:"title" comment:"书名"`
	Author       string   `json:"author" comment:"作者"`
	Cover        string   `json:"cover" comment:"封面URL或base64"`
	SerialStatus string   `json:"serialStatus" comment:"连载状态"`
	Category     string   `json:"category" comment:"分类字典值"`
	Tags         []string `json:"tags" comment:"标签数组"`
	Slogan       string   `json:"slogan" comment:"一句话推荐语"`
	Description  string   `json:"description" comment:"简介"`
	PublishDate  *string  `json:"publishDate" comment:"发布日期 YYYY-MM-DD"`
	WordCount    int      `json:"wordCount" comment:"字数"`
	Chapters     int      `json:"chapters" comment:"章节数"`
	IsFeatured   int      `json:"isFeatured" comment:"首页精选 0|1"`
	ReadUrl      string   `json:"readUrl" comment:"阅读链接"`
	Status       string   `json:"status" comment:"上架状态 1-上架 2-下架"`
	CurrUserId   int64    `json:"-" comment:"当前登录用户"`
}

type NovelBookGetReq struct {
	Id int64 `uri:"id"`
}

type NovelBookDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type NovelBookRankReq struct {
	Type       string `form:"type" comment:"rating-综合评分 click-点击热度"`
	CurrUserId int64  `form:"-" comment:"当前登录用户"`
}

// NovelHomeResp 首页聚合响应
type NovelHomeResp struct {
	Featured    []models.NovelBook `json:"featured" comment:"精选轮播"`
	HotBooks    []models.NovelBook `json:"hotBooks" comment:"热门TOP5"`
	LatestPosts []models.NovelPost `json:"latestPosts" comment:"最新帖子"`
	HotPosts    []models.NovelPost `json:"hotPosts" comment:"热门帖子"`
}
