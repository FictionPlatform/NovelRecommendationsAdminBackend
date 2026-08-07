package service

import (
	"fmt"
	"strings"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
)

// 反馈/投诉类型常量
const (
	NovelFeedbackKind     = "feedback"  // 意见反馈
	NovelComplaintKind    = "complaint" // 投诉
	NovelFeedbackStatusDo = "1"         // 处理中
)

type NovelFeedback struct {
	service.Service
}

// NewNovelFeedbackService app-实例化意见反馈服务
func NewNovelFeedbackService(s *service.Service) *NovelFeedback {
	var srv = new(NovelFeedback)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// Insert app-提交反馈/投诉（成功后自动生成系统通知）
func (e *NovelFeedback) Insert(c *dto.NovelFeedbackInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Kind == "" {
		c.Kind = NovelFeedbackKind
	}
	if c.Kind != NovelFeedbackKind && c.Kind != NovelComplaintKind {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if _, ok := dto.FeedbackTypeLabelMap[c.Type]; !ok {
		return 0, baseLang.NovelFeedbackTypeErrCode, lang.MsgErr(baseLang.NovelFeedbackTypeErrCode, e.Lang)
	}
	if strings.TrimSpace(c.Content) == "" {
		return 0, baseLang.NovelFeedbackEmptyCode, lang.MsgErr(baseLang.NovelFeedbackEmptyCode, e.Lang)
	}
	if len([]rune(c.Content)) > 500 {
		return 0, baseLang.NovelFeedbackTooLongCode, lang.MsgErr(baseLang.NovelFeedbackTooLongCode, e.Lang)
	}

	userName, _, respCode, err := getUserSnapshot(&e.Service, c.CurrUserId)
	if err != nil {
		return 0, respCode, err
	}
	now := time.Now()
	data := models.NovelFeedback{
		UserId:    c.CurrUserId,
		UserName:  userName,
		Type:      c.Type,
		Kind:      c.Kind,
		Content:   c.Content,
		Status:    NovelFeedbackStatusDo,
		CreatedAt: &now,
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}

	// 生成系统通知（失败仅记日志，不阻断反馈提交）
	title := "反馈已提交"
	content := fmt.Sprintf("你的反馈（%s）已收到，我们会尽快处理。", dto.FeedbackTypeLabelMap[c.Type])
	source := "feedback"
	if c.Kind == NovelComplaintKind {
		title = "投诉已提交"
		content = fmt.Sprintf("你的投诉（%s）已收到，我们会尽快核实处理。", dto.FeedbackTypeLabelMap[c.Type])
		source = "complaint"
	}
	if nErr := NewNovelNotificationService(&e.Service).Add(c.CurrUserId, title, content, source, nil); nErr != nil {
		e.Log.Errorf("novel feedback notify error:%s", nErr.Error())
	}
	return data.Id, baseLang.SuccessCode, nil
}

// GetPage app-分页查询我的反馈/投诉
func (e *NovelFeedback) GetPage(c *dto.NovelFeedbackQueryReq) ([]dto.NovelFeedbackItem, int64, int, error) {
	if c.CurrUserId <= 0 {
		return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.NovelFeedback
	var list []models.NovelFeedback
	var count int64
	db := e.Orm.Model(&data).Where("user_id = ?", c.CurrUserId)
	if c.Kind != "" {
		db = db.Where("kind = ?", c.Kind)
	}
	err := db.Order("created_at desc, id desc").
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	items := make([]dto.NovelFeedbackItem, 0, len(list))
	for _, f := range list {
		items = append(items, e.toItem(f))
	}
	return items, count, baseLang.SuccessCode, nil
}

// GetAdminPage app-后台分页查询反馈/投诉
func (e *NovelFeedback) GetAdminPage(c *dto.NovelFeedbackPageReq) ([]dto.NovelFeedbackItem, int64, int, error) {
	var data models.NovelFeedback
	var list []models.NovelFeedback
	var count int64
	err := e.Orm.Model(&data).
		Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("created_at desc, id desc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	items := make([]dto.NovelFeedbackItem, 0, len(list))
	for _, f := range list {
		items = append(items, e.toItem(f))
	}
	return items, count, baseLang.SuccessCode, nil
}

// Delete app-后台删除反馈/投诉
func (e *NovelFeedback) Delete(ids []int64) (int, error) {
	if len(ids) == 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Where("id in ?", ids).Delete(&models.NovelFeedback{}).Error; err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// toItem 转换列表项（补类型文案）
func (e *NovelFeedback) toItem(f models.NovelFeedback) dto.NovelFeedbackItem {
	return dto.NovelFeedbackItem{
		Id:        f.Id,
		UserId:    f.UserId,
		UserName:  f.UserName,
		Type:      f.Type,
		TypeLabel: dto.FeedbackTypeLabelMap[f.Type],
		Kind:      f.Kind,
		Content:   f.Content,
		Status:    f.Status,
		CreatedAt: f.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
