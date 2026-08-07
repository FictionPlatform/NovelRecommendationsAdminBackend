package dto

import "go-admin/core/dto"

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
