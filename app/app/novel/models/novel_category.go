package models

import "time"

// NovelCategory 小说一级分类（分类.md 定义，仅一级）
type NovelCategory struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name      string     `json:"name" gorm:"column:name;type:varchar(32);uniqueIndex;comment:分类名称"`
	Sort      int        `json:"sort" gorm:"column:sort;type:int;comment:排序（越小越靠前）"`
	Status    string     `json:"status" gorm:"column:status;type:char(1);comment:状态(1-正常 2-停用)"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`

	// 扩展
	TagCount int64       `json:"tagCount" gorm:"-" comment:"标签数量"`
	Tags     []NovelTag  `json:"tags,omitempty" gorm:"-" comment:"该分类下标签列表"`
}

func (NovelCategory) TableName() string {
	return "app_novel_category"
}
