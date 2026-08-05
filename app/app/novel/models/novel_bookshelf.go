package models

import "time"

type NovelBookshelf struct {
	Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	UserId    int64      `json:"userId" gorm:"column:user_id;type:int;comment:用户编号"`
	BookId    int64      `json:"bookId" gorm:"column:book_id;type:int;comment:书籍编号"`
	SortNo    int        `json:"sortNo" gorm:"column:sort_no;type:int;comment:排序号"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`

	// 扩展
	Book *NovelBook `json:"book" gorm:"-"`
}

func (NovelBookshelf) TableName() string {
	return "app_novel_bookshelf"
}
