package service

import (
	"errors"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NovelBookReview struct {
	service.Service
}

// NewNovelBookReviewService app-实例化书评服务
func NewNovelBookReviewService(s *service.Service) *NovelBookReview {
	var srv = new(NovelBookReview)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-书评圈分页查询
func (e *NovelBookReview) GetPage(c *dto.NovelReviewQueryReq) ([]models.NovelBookReview, int64, int, error) {
	var data models.NovelBookReview
	var list []models.NovelBookReview
	var count int64

	db := e.Orm.Model(&data)
	// 只看我的书评
	if c.Mine == 1 {
		if c.CurrUserId <= 0 {
			return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
		}
		db = db.Where("user_id = ?", c.CurrUserId)
	}
	db = db.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
	)
	switch c.Filter {
	case "five":
		db = db.Where("rating = 5")
	case "hot":
		db = db.Order("likes desc, created_at desc")
	default:
		db = db.Order("created_at desc")
	}
	err := db.Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	// 批量填充关联书籍
	bookIds := make([]int64, 0, len(list))
	for _, r := range list {
		bookIds = append(bookIds, r.BookId)
	}
	if len(bookIds) > 0 {
		var books []models.NovelBook
		_ = e.Orm.Where("id in (?)", bookIds).Find(&books).Error
		bookMap := map[int64]models.NovelBook{}
		for _, b := range books {
			bookMap[b.Id] = b
		}
		for i := range list {
			if b, ok := bookMap[list[i].BookId]; ok {
				list[i].Book = &b
			}
		}
	}
	return list, count, baseLang.SuccessCode, nil
}

// Insert app-新增书评（评分联动事务）
func (e *NovelBookReview) Insert(c *dto.NovelReviewInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.BookId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Rating < 1 || c.Rating > 5 {
		return 0, baseLang.NovelReviewRatingRangeCode, lang.MsgErr(baseLang.NovelReviewRatingRangeCode, e.Lang)
	}
	if len([]rune(c.Content)) > 1000 {
		return 0, baseLang.NovelContentTooLongCode, lang.MsgErr(baseLang.NovelContentTooLongCode, e.Lang)
	}
	// 写操作权限校验（注销/禁言拦截，禁言到期惰性恢复）
	if respCode, err := CheckReaderWritePermission(&e.Service, c.CurrUserId); err != nil {
		return 0, respCode, err
	}

	userName, userAvatar, respCode, err := getUserSnapshot(&e.Service, c.CurrUserId)
	if err != nil {
		return 0, respCode, err
	}

	baseOrm := e.Orm
	e.Orm = baseOrm.Begin()
	var txErr error
	committed := false
	defer func() {
		if committed {
			e.Orm.Commit()
		} else {
			e.Orm.Rollback()
		}
		e.Orm = baseOrm
	}()

	// 锁定书籍行（FOR UPDATE），保证评分聚合原子
	book := &models.NovelBook{}
	txErr = e.Orm.Clauses(clause.Locking{Strength: "UPDATE"}).First(book, c.BookId).Error
	if txErr != nil {
		if errors.Is(txErr, gorm.ErrRecordNotFound) {
			txErr = nil
			return 0, baseLang.NovelBookNotExistCode, lang.MsgErr(baseLang.NovelBookNotExistCode, e.Lang)
		}
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, txErr)
	}

	now := time.Now()
	review := models.NovelBookReview{}
	review.BookId = c.BookId
	review.UserId = c.CurrUserId
	review.UserName = userName
	review.UserAvatar = userAvatar
	review.Rating = c.Rating
	review.Content = c.Content
	review.Likes = 0
	review.CreatedAt = &now
	txErr = e.Orm.Create(&review).Error
	if txErr != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, txErr)
	}

	// 评分聚合更新 rating=(旧rating×旧数+新评分)/(新数)，review_count+1
	newRating := avgRating(book.Rating, book.ReviewCount, c.Rating)
	txErr = e.Orm.Model(&models.NovelBook{}).Where("id = ?", book.Id).
		Updates(map[string]interface{}{
			"rating":       newRating,
			"review_count": gorm.Expr("review_count + 1"),
			"updated_at":   time.Now(),
		}).Error
	if txErr != nil {
		return 0, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, txErr)
	}
	committed = true
	return review.Id, baseLang.SuccessCode, nil
}

// Delete app-删除我的书评（本人校验）
func (e *NovelBookReview) Delete(ids []int64, currUserId int64) (int, error) {
	if len(ids) == 0 || currUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var list []models.NovelBookReview
	err := e.Orm.Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	for _, r := range list {
		if r.UserId != currUserId {
			return baseLang.NovelNoPermissionCode, lang.MsgErr(baseLang.NovelNoPermissionCode, e.Lang)
		}
	}
	err = e.Orm.Where("id in (?)", ids).Delete(&models.NovelBookReview{}).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
