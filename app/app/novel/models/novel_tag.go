package models

import "time"

// NovelTag 小说标签（挂在分类下，一级分类 + 多标签）
type NovelTag struct {
	Id         int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	CategoryId int64      `json:"categoryId" gorm:"column:category_id;type:bigint;index;comment:所属分类id"`
	Category   string     `json:"category" gorm:"-" comment:"分类名称(联查填充)"`
	Name       string     `json:"name" gorm:"column:name;type:varchar(32);comment:标签名称"`
	Sort       int        `json:"sort" gorm:"column:sort;type:int;comment:排序（越小越靠前）"`
	Status     string     `json:"status" gorm:"column:status;type:char(1);comment:状态(1-正常 2-停用)"`
	CreateBy   int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy   int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (NovelTag) TableName() string {
	return "app_novel_tag"
}
