package dto

import (
	"go-admin/app/app/novel/models"
	"go-admin/core/dto"
)

// NovelNotificationQueryReq 我的通知分页查询
type NovelNotificationQueryReq struct {
	dto.Pagination `search:"-"`
	CurrUserId     int64 `form:"-" search:"-" comment:"当前登录用户"`
}

func (m *NovelNotificationQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelNotificationReadReq 标记已读请求（ids 为空表示全部标为已读）
type NovelNotificationReadReq struct {
	Ids        []int64 `json:"ids" comment:"通知编号集合，空=全部"`
	CurrUserId int64   `json:"-" comment:"当前登录用户"`
}

// NovelNotificationItem 通知列表项
type NovelNotificationItem struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	IsRead    string `json:"isRead"`
	Source    string `json:"source"`
	CreatedAt string `json:"createdAt"`
}

// NovelAdminNotificationQueryReq 后台通知管理分页查询
type NovelAdminNotificationQueryReq struct {
	dto.Pagination `search:"-"`
	Keyword        string `form:"keyword" search:"-" comment:"标题/内容关键字"`
	Source         string `form:"source" search:"-" comment:"来源 system|notice|admin|feedback|complaint"`
}

func (m *NovelAdminNotificationQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelAdminNotificationItem 后台通知管理列表项（含收件人）
type NovelAdminNotificationItem struct {
	models.NovelNotification
	UserName    string `json:"userName" comment:"收件人昵称"`
	CreatedAtStr string `json:"createdAtStr" comment:"创建时间字符串"`
}

// NovelAdminNotificationSendReq 后台发送通知请求
type NovelAdminNotificationSendReq struct {
	UserId  int64  `json:"userId" comment:"目标读者ID（0=全员发送）"`
	Title   string `json:"title" comment:"通知标题"`
	Content string `json:"content" comment:"通知内容"`
}
