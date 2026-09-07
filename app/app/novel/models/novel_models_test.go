package models

import (
	"testing"

	"gorm.io/gorm"
)

// TestNovelReaderProfileBeforeCreate 空 JSON 列兜底为 []
func TestNovelReaderProfileBeforeCreate(t *testing.T) {
	p := &NovelReaderProfile{}
	if err := p.BeforeCreate(&gorm.DB{}); err != nil {
		t.Fatalf("BeforeCreate error: %v", err)
	}
	if p.PreferredCategories != "[]" {
		t.Errorf("empty PreferredCategories should become [], got %q", p.PreferredCategories)
	}
}

// TestNovelReaderProfileBeforeCreateNonEmpty 已有 JSON 不覆盖
func TestNovelReaderProfileBeforeCreateNonEmpty(t *testing.T) {
	p := &NovelReaderProfile{PreferredCategories: `["玄幻","科幻"]`}
	if err := p.BeforeCreate(&gorm.DB{}); err != nil {
		t.Fatalf("BeforeCreate error: %v", err)
	}
	if p.PreferredCategories != `["玄幻","科幻"]` {
		t.Errorf("non-empty PreferredCategories should stay, got %q", p.PreferredCategories)
	}
}

// TestNovelReaderProfileBeforeUpdate 更新路径同样兜底（map update 误写空串场景）
func TestNovelReaderProfileBeforeUpdate(t *testing.T) {
	p := &NovelReaderProfile{}
	if err := p.BeforeUpdate(&gorm.DB{}); err != nil {
		t.Fatalf("BeforeUpdate error: %v", err)
	}
	if p.PreferredCategories != "[]" {
		t.Errorf("empty PreferredCategories should become [], got %q", p.PreferredCategories)
	}
}

// TestNovelModelsTableName 表名与文档约定一致
func TestNovelModelsTableName(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"NovelBook", (NovelBook{}).TableName(), "app_novel_book"},
		{"NovelPost", (NovelPost{}).TableName(), "app_novel_post"},
		{"NovelCategory", (NovelCategory{}).TableName(), "app_novel_category"},
		{"NovelTag", (NovelTag{}).TableName(), "app_novel_tag"},
		{"NovelReaderProfile", (NovelReaderProfile{}).TableName(), "app_novel_reader_profile"},
		{"NovelFollow", (NovelFollow{}).TableName(), "app_novel_follow"},
		{"NovelFeedback", (NovelFeedback{}).TableName(), "app_novel_feedback"},
		{"NovelNotification", (NovelNotification{}).TableName(), "app_novel_notification"},
		{"NovelNotice", (NovelNotice{}).TableName(), "app_novel_notice"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("TableName()=%q, want %q", c.got, c.want)
			}
		})
	}
}