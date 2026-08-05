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
)

type NovelBookshelf struct {
	service.Service
}

// NewNovelBookshelfService app-实例化我的书架服务
func NewNovelBookshelfService(s *service.Service) *NovelBookshelf {
	var srv = new(NovelBookshelf)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-分页查询我的书架
func (e *NovelBookshelf) GetPage(c *dto.NovelShelfQueryReq) ([]models.NovelBookshelf, int64, int, error) {
	if c.CurrUserId <= 0 {
		return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.NovelBookshelf
	var list []models.NovelBookshelf
	var count int64
	err := e.Orm.Model(&data).Where("user_id = ?", c.CurrUserId).
		Order("sort_no asc, id desc").
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	fillShelfBooks(&e.Service, list)
	return list, count, baseLang.SuccessCode, nil
}

// Insert app-新增书架（upsert 语义）
func (e *NovelBookshelf) Insert(c *dto.NovelShelfInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.BookId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 书籍存在性
	book := &models.NovelBook{}
	err := e.Orm.First(book, c.BookId).Error
	if err != nil {
		return 0, baseLang.NovelBookNotExistCode, lang.MsgErr(baseLang.NovelBookNotExistCode, e.Lang)
	}
	// 已存在则幂等返回
	old := &models.NovelBookshelf{}
	err = e.Orm.Where("user_id = ? and book_id = ?", c.CurrUserId, c.BookId).First(old).Error
	if err == nil {
		return old.Id, baseLang.SuccessCode, nil
	}
	now := time.Now()
	data := models.NovelBookshelf{}
	data.UserId = c.CurrUserId
	data.BookId = c.BookId
	data.SortNo = 0
	data.CreatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Delete app-移除书架（本人）
func (e *NovelBookshelf) Delete(c *dto.NovelShelfDeleteReq) (int, error) {
	if c.CurrUserId <= 0 || c.Id <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	shelf := &models.NovelBookshelf{}
	err := e.Orm.First(shelf, c.Id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return baseLang.NovelShelfNotExistCode, lang.MsgErr(baseLang.NovelShelfNotExistCode, e.Lang)
	}
	if shelf.UserId != c.CurrUserId {
		return baseLang.NovelNoPermissionCode, lang.MsgErr(baseLang.NovelNoPermissionCode, e.Lang)
	}
	err = e.Orm.Where("id = ?", shelf.Id).Delete(&models.NovelBookshelf{}).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
