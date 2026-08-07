package dto

import "go-admin/core/dto"

// NovelFeedbackInsertReq 提交反馈/投诉请求
type NovelFeedbackInsertReq struct {
	Type       string `json:"type" comment:"反馈类型 A~H"`
	Content    string `json:"content" comment:"反馈内容"`
	Kind       string `json:"kind" comment:"feedback-反馈 complaint-投诉"`
	CurrUserId int64  `json:"-" comment:"当前登录用户"`
}

// NovelFeedbackQueryReq 我的反馈/投诉分页查询
type NovelFeedbackQueryReq struct {
	dto.Pagination `search:"-"`
	Kind           string `form:"kind" search:"-" comment:"feedback|complaint 为空查全部"`
	CurrUserId     int64  `form:"-" search:"-" comment:"当前登录用户"`
}

func (m *NovelFeedbackQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelFeedbackPageReq 后台反馈管理分页查询
type NovelFeedbackPageReq struct {
	dto.Pagination `search:"-"`
	Kind           string `form:"kind" search:"type:exact;column:kind;table:app_novel_feedback" comment:"feedback|complaint"`
	Type           string `form:"type" search:"type:exact;column:type;table:app_novel_feedback" comment:"反馈类型 A~H"`
	Keyword        string `form:"keyword" search:"type:contains;column:content;table:app_novel_feedback" comment:"内容关键字"`
}

func (m *NovelFeedbackPageReq) GetNeedSearch() interface{} {
	return *m
}

// NovelFeedbackDeleteReq 后台删除反馈请求
type NovelFeedbackDeleteReq struct {
	Ids []int64 `json:"ids" comment:"反馈编号集合"`
}

// NovelFeedbackItem 反馈列表项（含类型文案）
type NovelFeedbackItem struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	UserName  string `json:"userName"`
	Type      string `json:"type"`
	TypeLabel string `json:"typeLabel"`
	Kind      string `json:"kind"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

// FeedbackTypeLabelMap 反馈类型文案映射（与前端 A~H 一致）
var FeedbackTypeLabelMap = map[string]string{
	"A": "产品功能建议",
	"B": "UI样式建议",
	"C": "主界面功能问题",
	"D": "视频窗口问题",
	"E": "程序报错/性能问题",
	"F": "侵犯版权举报",
	"G": "侮辱诽谤举报",
	"H": "其他",
}
