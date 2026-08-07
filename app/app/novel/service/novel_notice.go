package service

import (
	"strings"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"gorm.io/gorm"
)

// 公告广播分批大小
const noticeBroadcastBatch = 500

type NovelNotice struct {
	service.Service
}

// NewNovelNoticeService app-实例化系统公告服务
func NewNovelNoticeService(s *service.Service) *NovelNotice {
	var srv = new(NovelNotice)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-后台分页查询公告
func (e *NovelNotice) GetPage(c *dto.NovelNoticeQueryReq) ([]dto.NovelNoticeItem, int64, int, error) {
	var data models.NovelNotice
	var list []models.NovelNotice
	var count int64
	err := e.Orm.Model(&data).
		Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("created_at desc, id desc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	items := make([]dto.NovelNoticeItem, 0, len(list))
	for _, n := range list {
		var recipient int64
		_ = e.Orm.Model(&models.NovelNotification{}).
			Where("source = ? and notice_id = ?", NovelNotifySourceNotice, n.Id).Count(&recipient).Error
		items = append(items, dto.NovelNoticeItem{NovelNotice: n, RecipientCount: recipient})
	}
	return items, count, baseLang.SuccessCode, nil
}

// Insert app-发布公告并广播到全部读者
func (e *NovelNotice) Insert(c *dto.NovelNoticeInsertReq) (int64, int, error) {
	if strings.TrimSpace(c.Title) == "" {
		return 0, baseLang.NovelNoticeTitleEmptyCode, lang.MsgErr(baseLang.NovelNoticeTitleEmptyCode, e.Lang)
	}
	if len([]rune(c.Title)) > 100 {
		return 0, baseLang.NovelNoticeTitleTooLongCode, lang.MsgErr(baseLang.NovelNoticeTitleTooLongCode, e.Lang)
	}
	if strings.TrimSpace(c.Content) == "" {
		return 0, baseLang.NovelNoticeContentEmptyCode, lang.MsgErr(baseLang.NovelNoticeContentEmptyCode, e.Lang)
	}
	if len([]rune(c.Content)) > 2000 {
		return 0, baseLang.NovelNoticeContentTooLongCode, lang.MsgErr(baseLang.NovelNoticeContentTooLongCode, e.Lang)
	}
	now := time.Now()
	notice := models.NovelNotice{
		Title:     c.Title,
		Content:   c.Content,
		CreateBy:  c.CreateBy,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := e.Orm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&notice).Error; err != nil {
			return err
		}
		var userIds []int64
		if err := tx.Model(&models.NovelReaderProfile{}).Pluck("user_id", &userIds).Error; err != nil {
			return err
		}
		if len(userIds) == 0 {
			return nil
		}
		notifications := make([]models.NovelNotification, 0, len(userIds))
		for _, uid := range userIds {
			notifications = append(notifications, models.NovelNotification{
				UserId:    uid,
				Title:     c.Title,
				Content:   c.Content,
				IsRead:    NovelNotifyUnread,
				Source:    NovelNotifySourceNotice,
				NoticeId:  &notice.Id,
				CreatedAt: &now,
			})
		}
		for i := 0; i < len(notifications); i += noticeBroadcastBatch {
			end := i + noticeBroadcastBatch
			if end > len(notifications) {
				end = len(notifications)
			}
			if err := tx.Create(notifications[i:end]).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return notice.Id, baseLang.SuccessCode, nil
}

// Delete app-后台删除公告（级联删除对应读者通知）
func (e *NovelNotice) Delete(ids []int64) (int, error) {
	if len(ids) == 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	err := e.Orm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id in ?", ids).Delete(&models.NovelNotice{}).Error; err != nil {
			return err
		}
		return tx.Where("source = ? and notice_id in ?", NovelNotifySourceNotice, ids).
			Delete(&models.NovelNotification{}).Error
	})
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
