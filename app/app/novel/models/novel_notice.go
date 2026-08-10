package models

import "time"

type NovelNotice struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Title     string     `json:"title" gorm:"column:title;type:varchar(100);comment:公告标题"`
	Content   string     `json:"content" gorm:"column:content;type:text;comment:公告内容"`
	ValidDays int        `json:"validDays" gorm:"column:valid_days;type:int;comment:有效天数(0=永久,过期后登录不再收取)"`
	CreateBy  int64      `json:"createBy" gorm:"column:create_by;type:int;comment:发布管理员编号"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

func (NovelNotice) TableName() string {
	return "app_novel_notice"
}
