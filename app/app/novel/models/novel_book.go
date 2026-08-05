package models

import (
	"time"
)

type NovelBook struct {
	Id           int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Title        string     `json:"title" gorm:"column:title;type:varchar(100);comment:书名"`
	Author       string     `json:"author" gorm:"column:author;type:varchar(64);comment:作者"`
	Cover        string     `json:"cover" gorm:"column:cover;type:varchar(500);comment:封面URL或base64"`
	Rating       float64    `json:"rating" gorm:"column:rating;type:decimal(3,1);comment:综合评分"`
	ReviewCount  int        `json:"reviewCount" gorm:"column:review_count;type:int;comment:书评数"`
	SerialStatus string     `json:"serialStatus" gorm:"column:serial_status;type:char(1);comment:连载状态(1-连载中 2-已完结)"`
	Category     string     `json:"category" gorm:"column:category;type:char(2);comment:分类字典app_novel_category"`
	Tags         string     `json:"-" gorm:"column:tags;type:json;comment:标签数组"`
	TagList      []string   `json:"tags" gorm:"-"`
	Slogan       string     `json:"slogan" gorm:"column:slogan;type:varchar(30);comment:一句话推荐语"`
	Description  string     `json:"description" gorm:"column:description;type:varchar(500);comment:简介"`
	Clicks       int        `json:"clicks" gorm:"column:clicks;type:int;comment:点击数"`
	PublishDate  *time.Time `json:"publishDate" gorm:"column:publish_date;type:date;comment:发布日期"`
	WordCount    int        `json:"wordCount" gorm:"column:word_count;type:int;comment:字数"`
	Chapters     int        `json:"chapters" gorm:"column:chapters;type:int;comment:章节数"`
	IsFeatured   int        `json:"isFeatured" gorm:"column:is_featured;type:tinyint;comment:首页精选 0|1"`
	ReadUrl      string     `json:"readUrl" gorm:"column:read_url;type:varchar(500);comment:阅读链接"`
	Status       string     `json:"status" gorm:"column:status;type:char(1);comment:上架状态(1-上架 2-下架)"`
	CreateBy     int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy     int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`

	// 扩展
	Reviews       []NovelBookReview `json:"reviews" gorm:"-"`
	IsCollected   bool              `json:"isCollected" gorm:"-"`
	CategoryLabel string            `json:"categoryLabel" gorm:"-"`
	StatusLabel   string            `json:"statusLabel" gorm:"-"`
}

func (NovelBook) TableName() string {
	return "app_novel_book"
}
