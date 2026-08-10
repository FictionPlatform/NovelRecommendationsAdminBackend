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

// PullNotices app-登录时按需收取未消费的未过期公告
// 机制：发公告不再广播；用户登录/查看通知时，将「有效期内 + 该用户尚未消费」的公告生成为其通知
func (e *NovelNotification) PullNotices(userId int64) {
	if userId <= 0 {
		return
	}
	now := time.Now()
	// 未过期的公告：valid_days=0（永久）或 created_at 距今未超过 valid_days 天
	// 跨库兼容：先取全部公告，再在内存中按 valid_days 过滤（公告量小，可接受）
	var notices []models.NovelNotice
	if err := e.Orm.Model(&models.NovelNotice{}).
		Order("created_at desc, id desc").Find(&notices).Error; err != nil {
		e.Log.Errorf("novel notification pull notices query error:%s", err.Error())
		return
	}
	// 过滤未过期
	valid := notices[:0]
	for _, n := range notices {
		if n.ValidDays <= 0 || n.CreatedAt == nil {
			valid = append(valid, n)
			continue
		}
		expireAt := n.CreatedAt.AddDate(0, 0, n.ValidDays)
		if expireAt.After(now) {
			valid = append(valid, n)
		}
	}
	if len(valid) == 0 {
		return
	}
	// 已消费的公告（该用户已有对应通知记录）
	noticeIds := make([]int64, 0, len(valid))
	for _, n := range valid {
		noticeIds = append(noticeIds, n.Id)
	}
	var consumed []int64
	if err := e.Orm.Model(&models.NovelNotification{}).
		Where("user_id = ? and source = ? and notice_id in (?)", userId, NovelNotifySourceNotice, noticeIds).
		Pluck("notice_id", &consumed).Error; err != nil {
		e.Log.Errorf("novel notification pull consumed query error:%s", err.Error())
		return
	}
	consumedMap := map[int64]bool{}
	for _, id := range consumed {
		consumedMap[id] = true
	}
	// 逐条生成未消费公告通知
	for _, n := range valid {
		if consumedMap[n.Id] {
			continue
		}
		noticeId := n.Id
		if err := e.Add(userId, n.Title, n.Content, NovelNotifySourceNotice, &noticeId); err != nil {
			e.Log.Errorf("novel notification pull notice error:%s", err.Error())
		}
	}
}

// GetPage app-分页查询我的系统通知
func (e *NovelNotification) GetPage(c *dto.NovelNotificationQueryReq) ([]models.NovelNotification, int64, int, error) {
	if c.CurrUserId <= 0 {
		return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 登录查看时按需收取未消费的未过期公告
	e.PullNotices(c.CurrUserId)
	var data models.NovelNotification
	var list []models.NovelNotification
	var count int64
	err := e.Orm.Model(&data).Where("user_id = ?", c.CurrUserId).
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("created_at desc, id desc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 过滤已过期的公告通知（source=notice 且对应公告已过期 → 彻底不可见）
	list, count = e.filterExpiredNotice(c.CurrUserId, list, count)
	return list, count, baseLang.SuccessCode, nil
}

// UnreadCount app-查询我的未读通知数
func (e *NovelNotification) UnreadCount(userId int64) (int64, int, error) {
	if userId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 登录查看时按需收取未消费的未过期公告
	e.PullNotices(userId)
	var count int64
	err := e.Orm.Model(&models.NovelNotification{}).
		Where("user_id = ? and is_read = ?", userId, NovelNotifyUnread).Count(&count).Error
	if err != nil {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 过滤已过期的公告通知未读数
	expired := e.expiredNoticeIds()
	if len(expired) > 0 {
		var expiredUnread int64
		err = e.Orm.Model(&models.NovelNotification{}).
			Where("user_id = ? and is_read = ? and source = ? and notice_id in (?)", userId, NovelNotifyUnread, NovelNotifySourceNotice, expired).
			Count(&expiredUnread).Error
		if err == nil {
			count = count - expiredUnread
			if count < 0 {
				count = 0
			}
		}
	}
	return count, baseLang.SuccessCode, nil
}

// expiredNoticeIds app-内部方法，返回已过期的公告 id 集合
func (e *NovelNotification) expiredNoticeIds() []int64 {
	now := time.Now()
	var notices []models.NovelNotice
	if err := e.Orm.Model(&models.NovelNotice{}).
		Where("valid_days > 0").Find(&notices).Error; err != nil {
		return nil
	}
	var expired []int64
	for _, n := range notices {
		if n.CreatedAt == nil {
			continue
		}
		expireAt := n.CreatedAt.AddDate(0, 0, n.ValidDays)
		if !expireAt.After(now) {
			expired = append(expired, n.Id)
		}
	}
	return expired
}

// filterExpiredNotice app-内部方法，从通知列表移除已过期的公告通知
func (e *NovelNotification) filterExpiredNotice(userId int64, list []models.NovelNotification, count int64) ([]models.NovelNotification, int64) {
	expired := e.expiredNoticeIds()
	if len(expired) == 0 {
		return list, count
	}
	expiredMap := map[int64]bool{}
	for _, id := range expired {
		expiredMap[id] = true
	}
	filtered := list[:0]
	for _, n := range list {
		if n.Source == NovelNotifySourceNotice && n.NoticeId != nil && expiredMap[*n.NoticeId] {
			continue
		}
		filtered = append(filtered, n)
	}
	// 修正总数：从分页结果中减去的仅是当前页过期数，总数需另查
	if len(list) != len(filtered) {
		var realCount int64
		if err := e.Orm.Model(&models.NovelNotification{}).
			Where("user_id = ?", userId).
			Where("not (source = ? and notice_id in (?))", NovelNotifySourceNotice, expired).
			Count(&realCount).Error; err == nil {
			count = realCount
		}
	}
	return filtered, count
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
