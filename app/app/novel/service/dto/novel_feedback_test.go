package dto

import "testing"

// TestFeedbackTypeLabelMap 反馈类型文案映射（A~H，与前端表单一致）
func TestFeedbackTypeLabelMap(t *testing.T) {
	cases := []struct {
		code string
		want string
	}{
		{"A", "产品功能建议"},
		{"B", "UI样式建议"},
		{"C", "主界面功能问题"},
		{"D", "视频窗口问题"},
		{"E", "程序报错/性能问题"},
		{"F", "侵犯版权举报"},
		{"G", "侮辱诽谤举报"},
		{"H", "其他"},
	}
	for _, c := range cases {
		t.Run(c.code, func(t *testing.T) {
			if got := FeedbackTypeLabelMap[c.code]; got != c.want {
				t.Errorf("FeedbackTypeLabelMap[%q]=%q, want %q", c.code, got, c.want)
			}
		})
	}
	// 未知类型的文案应为空串（不 panic）
	if _, ok := FeedbackTypeLabelMap["Z"]; ok {
		t.Error("unknown type Z should not be present")
	}
}

// TestNovelFeedbackItemFields NovelFeedbackItem 关键字段契约（与前端 DTO 对齐）
func TestNovelFeedbackItemFields(t *testing.T) {
	item := NovelFeedbackItem{
		Id:        1,
		UserId:    7,
		UserName:  "读友小书狂",
		Type:      "A",
		TypeLabel: FeedbackTypeLabelMap["A"],
		Kind:      "feedback",
		Content:   "希望支持深色模式自动切换",
		Status:    "1",
		CreatedAt: "2026-08-09 10:00:00",
	}
	if item.Id != 1 || item.UserId != 7 || item.UserName != "读友小书狂" {
		t.Errorf("basic fields mismatch: %+v", item)
	}
	if item.TypeLabel != "产品功能建议" {
		t.Errorf("TypeLabel=%q, want 产品功能建议", item.TypeLabel)
	}
	if item.Kind != "feedback" || item.Status != "1" {
		t.Errorf("kind/status mismatch: %+v", item)
	}
}