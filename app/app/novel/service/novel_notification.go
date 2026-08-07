package service

import (
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
)

// 通知常量
const (
	NovelNotifyUnread       = "0" // 未读
	NovelNotifyRead         = "1" // 已读
	NovelNotifySourceSystem = "system"
	NovelNotifySourceNotice = "notice"
)

type NovelNotification struct {
	service.Service
}

// NewNovelNotificationService app-实例化系统通知服务
func NewNovelNotificationService(s *service.Service) *NovelNotification {
	var srv = new(NovelNotification)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// Add app-内部方法，向用户追加一条系统通知
func (e *NovelNotification) Add(userId int64, title, content, source string, noticeId *int64) error {
	if userId <= 0 {
		return nil
	}
	now := time.Now()
	data := models.NovelNotification{
		UserId:    userId,
		Title:     title,
		Content:   content,
		IsRead:    NovelNotifyUnread,
		Source:    source,
		NoticeId:  noticeId,
		CreatedAt: &now,
	}
	return e.Orm.Create(&data).Error
}

// Seed app- 注册时播种 3 条种子公告（欢迎/新功能/社区规范）
func (e *NovelNotification) Seed(userId int64) {
	if userId <= 0 {
		return
	}
	seeds := []struct {
		Title   string
		Content string
	}{
		{"欢迎来到墨读", "欢迎加入墨读·小说引力场！在这里你可以浏览书库、点评书籍、发布长文、结识书友。祝你阅读愉快！"},
		{"新功能上线", "书友关注、书架收藏、楼中楼评论现已上线。快去书库发现你的下一本心头好吧！"},
		{"社区规范", "请文明发言，尊重每一位书友。抄袭、辱骂、广告等违规内容将被处理，共建友好阅读社区。"},
	}
	for _, s := range seeds {
		if err := e.Add(userId, s.Title, s.Content, NovelNotifySourceSystem, nil); err != nil {
			e.Log.Errorf("novel notification seed error:%s", err.Error())
		}
	}
}

// GetPage app-分页查询我的系统通知
func (e *NovelNotification) GetPage(c *dto.NovelNotificationQueryReq) ([]models.NovelNotification, int64, int, error) {
	if c.CurrUserId <= 0 {
		return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.NovelNotification
	var list []models.NovelNotification
	var count int64
	err := e.Orm.Model(&data).Where("user_id = ?", c.CurrUserId).
		Order("created_at desc, id desc").
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// UnreadCount app-查询我的未读通知数
func (e *NovelNotification) UnreadCount(userId int64) (int64, int, error) {
	if userId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var count int64
	err := e.Orm.Model(&models.NovelNotification{}).
		Where("user_id = ? and is_read = ?", userId, NovelNotifyUnread).Count(&count).Error
	if err != nil {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return count, baseLang.SuccessCode, nil
}

// Read app-标记已读（ids 为空表示全部标为已读）
func (e *NovelNotification) Read(c *dto.NovelNotificationReadReq) (int, error) {
	if c.CurrUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var affectRows = e.Orm.Model(&models.NovelNotification{}).Where("user_id = ?", c.CurrUserId)
	if len(c.Ids) > 0 {
		affectRows = affectRows.Where("id in ?", c.Ids)
	}
	err := affectRows.Update("is_read", NovelNotifyRead).Error
	if err != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
