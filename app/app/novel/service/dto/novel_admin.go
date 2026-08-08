package dto

import (
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/core/dto"
)

// NovelUserQueryReq 后台读者分页查询
type NovelUserQueryReq struct {
	dto.Pagination `search:"-"`
	Keyword        string `form:"keyword" search:"type:contains;column:user_name;table:app_user" comment:"用户名关键字"`
	Status         string `form:"status" search:"type:exact;column:status;table:app_user" comment:"状态 1-正常 2-禁言 3-注销"`
}

func (m *NovelUserQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelUserItem 后台读者列表项（读者资料 + app_user 扩展字段）
type NovelUserItem struct {
	models.NovelReaderProfile
	UserName string `json:"userName" gorm:"column:user_name" comment:"用户名"`
	Status   string `json:"status" gorm:"column:status" comment:"状态 1-正常 2-禁言 3-注销"`
	Mobile   string `json:"mobile" gorm:"column:mobile" comment:"手机号码"`
}

// NovelUserStatusReq 读者状态变更请求（1-正常 2-禁言 3-注销）
type NovelUserStatusReq struct {
	Id       int64      `json:"-" uri:"id" comment:"读者用户编号"`
	Status   string     `json:"status" comment:"状态 1-正常 2-禁言 3-注销"`
	BanUntil *time.Time `json:"banUntil" comment:"禁言截止时间(空=永久禁言，仅 status=2 生效)"`
	Reason   string     `json:"reason" comment:"禁言原因(仅 status=2 生效)"`
}

// NovelBanPostReq 禁止/解除禁止读者发帖请求
type NovelBanPostReq struct {
	Id       int64      `json:"-" uri:"id" comment:"读者用户编号"`
	BanUntil *time.Time `json:"banUntil" comment:"禁止发帖截止时间(空=解除)"`
	Reason   string     `json:"reason" comment:"禁言原因"`
}

// NovelPostAdminQueryReq 后台话题分页查询
type NovelPostAdminQueryReq struct {
	dto.Pagination `search:"-"`
	Keyword        string `form:"keyword" search:"type:contains;column:title;table:app_novel_post" comment:"标题关键字"`
	Status         string `form:"status" search:"type:exact;column:status;table:app_novel_post" comment:"状态 1-正常 2-禁止访问"`
}

func (m *NovelPostAdminQueryReq) GetNeedSearch() interface{} {
	return *m
}

// NovelPostAdminStatusReq 后台话题状态变更请求
type NovelPostAdminStatusReq struct {
	Id     int64  `json:"-" uri:"id" comment:"话题编号"`
	Status string `json:"status" comment:"状态 1-正常 2-禁止访问"`
}
