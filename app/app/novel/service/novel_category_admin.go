package service

import (
	"errors"
	"strings"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"

	"gorm.io/gorm"
)

// NovelCategoryReader 读者端分类查询服务（分类列表含标签）
type NovelCategoryReader struct {
	service.Service
}

// NewNovelCategoryReaderService app-实例化分类查询服务
func NewNovelCategoryReaderService(s *service.Service) *NovelCategoryReader {
	var srv = new(NovelCategoryReader)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// List 读者端-查询启用中的全部分类（含各分类下启用中的标签）
func (e *NovelCategoryReader) List() ([]models.NovelCategory, int, error) {
	var cats []models.NovelCategory
	if err := e.Orm.Where("status = ?", global.SysStatusOk).Order("sort asc, id asc").Find(&cats).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if len(cats) == 0 {
		return cats, baseLang.SuccessCode, nil
	}
	catIds := make([]int64, 0, len(cats))
	for _, c := range cats {
		catIds = append(catIds, c.Id)
	}
	var tags []models.NovelTag
	if err := e.Orm.Where("category_id in (?) and status = ?", catIds, global.SysStatusOk).
		Order("sort asc, id asc").Find(&tags).Error; err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	tagMap := map[int64][]models.NovelTag{}
	for _, t := range tags {
		tagMap[t.CategoryId] = append(tagMap[t.CategoryId], t)
	}
	for i := range cats {
		cats[i].Tags = tagMap[cats[i].Id]
	}
	return cats, baseLang.SuccessCode, nil
}

// NovelCategoryAdmin 后台分类管理服务
type NovelCategoryAdmin struct {
	service.Service
}

// NewNovelCategoryAdminService app-实例化分类管理服务
func NewNovelCategoryAdminService(s *service.Service) *NovelCategoryAdmin {
	var srv = new(NovelCategoryAdmin)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 后台-分页查询分类（含标签数）
func (e *NovelCategoryAdmin) GetPage(c *dto.NovelCategoryQueryReq) ([]dto.NovelCategoryItem, int64, int, error) {
	var data models.NovelCategory
	var list []dto.NovelCategoryItem
	var count int64
	err := e.Orm.Model(&data).
		Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("sort asc, id asc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 填充标签数
	if len(list) > 0 {
		ids := make([]int64, 0, len(list))
		for _, it := range list {
			ids = append(ids, it.Id)
		}
		type cntRow struct {
			CategoryId int64
			Cnt        int64
		}
		var rows []cntRow
		if err := e.Orm.Model(&models.NovelTag{}).
			Select("category_id, count(*) as cnt").
			Where("category_id in (?)", ids).
			Group("category_id").Scan(&rows).Error; err == nil {
			m := map[int64]int64{}
			for _, r := range rows {
				m[r.CategoryId] = r.Cnt
			}
			for i := range list {
				list[i].TagCount = m[list[i].Id]
			}
		}
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get 后台-分类详情
func (e *NovelCategoryAdmin) Get(id int64) (*models.NovelCategory, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelCategory{}
	err := e.Orm.First(data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, baseLang.NovelCategoryNotExistCode, lang.MsgErr(baseLang.NovelCategoryNotExistCode, e.Lang)
		}
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return data, baseLang.SuccessCode, nil
}

// Insert 后台-新增分类
func (e *NovelCategoryAdmin) Insert(c *dto.NovelCategoryInsertReq) (int64, int, error) {
	if strings.TrimSpace(c.Name) == "" {
		return 0, baseLang.NovelCategoryNameEmptyCode, lang.MsgErr(baseLang.NovelCategoryNameEmptyCode, e.Lang)
	}
	if len([]rune(c.Name)) > 32 {
		return 0, baseLang.NovelCategoryNameTooLongCode, lang.MsgErr(baseLang.NovelCategoryNameTooLongCode, e.Lang)
	}
	// 同名校验
	var cnt int64
	if err := e.Orm.Model(&models.NovelCategory{}).Where("name = ?", strings.TrimSpace(c.Name)).Count(&cnt).Error; err != nil {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if cnt > 0 {
		return 0, baseLang.NovelCategoryNameExistCode, lang.MsgErr(baseLang.NovelCategoryNameExistCode, e.Lang)
	}
	now := time.Now()
	data := models.NovelCategory{
		Name:      strings.TrimSpace(c.Name),
		Sort:      c.Sort,
		Status:    c.Status,
		CreateBy:  c.CreateBy,
		UpdateBy:  c.CreateBy,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if data.Status == "" {
		data.Status = global.SysStatusOk
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update 后台-更新分类
func (e *NovelCategoryAdmin) Update(c *dto.NovelCategoryUpdateReq) (bool, int, error) {
	if c.Id <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelCategory{}
	err := e.Orm.First(data, c.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, baseLang.NovelCategoryNotExistCode, lang.MsgErr(baseLang.NovelCategoryNotExistCode, e.Lang)
		}
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	updates := map[string]interface{}{}
	if strings.TrimSpace(c.Name) != "" && data.Name != strings.TrimSpace(c.Name) {
		if len([]rune(c.Name)) > 32 {
			return false, baseLang.NovelCategoryNameTooLongCode, lang.MsgErr(baseLang.NovelCategoryNameTooLongCode, e.Lang)
		}
		var cnt int64
		if err := e.Orm.Model(&models.NovelCategory{}).
			Where("name = ? and id <> ?", strings.TrimSpace(c.Name), c.Id).Count(&cnt).Error; err != nil {
			return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if cnt > 0 {
			return false, baseLang.NovelCategoryNameExistCode, lang.MsgErr(baseLang.NovelCategoryNameExistCode, e.Lang)
		}
		updates["name"] = strings.TrimSpace(c.Name)
	}
	updates["sort"] = c.Sort
	if c.Status != "" {
		updates["status"] = c.Status
	}
	updates["update_by"] = c.UpdateBy
	updates["updated_at"] = time.Now()
	err = e.Orm.Model(&models.NovelCategory{}).Where("id = ?", c.Id).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// Delete 后台-删除分类（存在标签或书籍引用时禁止删除）
func (e *NovelCategoryAdmin) Delete(ids []int64) (int, error) {
	if len(ids) == 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 存在标签则禁止删除
	var tagCnt int64
	if err := e.Orm.Model(&models.NovelTag{}).Where("category_id in (?)", ids).Count(&tagCnt).Error; err != nil {
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if tagCnt > 0 {
		return baseLang.NovelCategoryHasTagCode, lang.MsgErr(baseLang.NovelCategoryHasTagCode, e.Lang)
	}
	// 存在书籍引用则禁止删除
	var bookCnt int64
	if err := e.Orm.Model(&models.NovelBook{}).Where("category in (?)", ids).Count(&bookCnt).Error; err != nil {
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if bookCnt > 0 {
		return baseLang.NovelCategoryHasBookCode, lang.MsgErr(baseLang.NovelCategoryHasBookCode, e.Lang)
	}
	if err := e.Orm.Where("id in (?)", ids).Delete(&models.NovelCategory{}).Error; err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// NovelTagAdmin 后台标签管理服务
type NovelTagAdmin struct {
	service.Service
}

// NewNovelTagAdminService app-实例化标签管理服务
func NewNovelTagAdminService(s *service.Service) *NovelTagAdmin {
	var srv = new(NovelTagAdmin)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 后台-分页查询标签（联查分类名）
func (e *NovelTagAdmin) GetPage(c *dto.NovelTagQueryReq) ([]dto.NovelTagItem, int64, int, error) {
	var data models.NovelTag
	var list []dto.NovelTagItem
	var count int64
	err := e.Orm.Model(&data).
		Select("app_novel_tag.*, app_novel_category.name as category").
		Joins("left join app_novel_category on app_novel_category.id = app_novel_tag.category_id").
		Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Order("app_novel_tag.sort asc, app_novel_tag.id asc").
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get 后台-标签详情
func (e *NovelTagAdmin) Get(id int64) (*models.NovelTag, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelTag{}
	err := e.Orm.First(data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, baseLang.NovelTagNotExistCode, lang.MsgErr(baseLang.NovelTagNotExistCode, e.Lang)
		}
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return data, baseLang.SuccessCode, nil
}

// Insert 后台-新增标签
func (e *NovelTagAdmin) Insert(c *dto.NovelTagInsertReq) (int64, int, error) {
	if c.CategoryId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if strings.TrimSpace(c.Name) == "" {
		return 0, baseLang.NovelTagNameEmptyCode, lang.MsgErr(baseLang.NovelTagNameEmptyCode, e.Lang)
	}
	if len([]rune(c.Name)) > 32 {
		return 0, baseLang.NovelTagNameTooLongCode, lang.MsgErr(baseLang.NovelTagNameTooLongCode, e.Lang)
	}
	// 分类存在性
	var catCnt int64
	if err := e.Orm.Model(&models.NovelCategory{}).Where("id = ?", c.CategoryId).Count(&catCnt).Error; err != nil {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if catCnt == 0 {
		return 0, baseLang.NovelCategoryNotExistCode, lang.MsgErr(baseLang.NovelCategoryNotExistCode, e.Lang)
	}
	// 分类内同名校验
	var cnt int64
	if err := e.Orm.Model(&models.NovelTag{}).
		Where("category_id = ? and name = ?", c.CategoryId, strings.TrimSpace(c.Name)).Count(&cnt).Error; err != nil {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if cnt > 0 {
		return 0, baseLang.NovelTagNameExistCode, lang.MsgErr(baseLang.NovelTagNameExistCode, e.Lang)
	}
	now := time.Now()
	data := models.NovelTag{
		CategoryId: c.CategoryId,
		Name:       strings.TrimSpace(c.Name),
		Sort:       c.Sort,
		Status:     c.Status,
		CreateBy:   c.CreateBy,
		UpdateBy:   c.CreateBy,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	}
	if data.Status == "" {
		data.Status = global.SysStatusOk
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update 后台-更新标签
func (e *NovelTagAdmin) Update(c *dto.NovelTagUpdateReq) (bool, int, error) {
	if c.Id <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelTag{}
	err := e.Orm.First(data, c.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, baseLang.NovelTagNotExistCode, lang.MsgErr(baseLang.NovelTagNotExistCode, e.Lang)
		}
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	updates := map[string]interface{}{}
	if c.CategoryId > 0 && data.CategoryId != c.CategoryId {
		var catCnt int64
		if err := e.Orm.Model(&models.NovelCategory{}).Where("id = ?", c.CategoryId).Count(&catCnt).Error; err != nil {
			return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if catCnt == 0 {
			return false, baseLang.NovelCategoryNotExistCode, lang.MsgErr(baseLang.NovelCategoryNotExistCode, e.Lang)
		}
		updates["category_id"] = c.CategoryId
	}
	if strings.TrimSpace(c.Name) != "" && data.Name != strings.TrimSpace(c.Name) {
		if len([]rune(c.Name)) > 32 {
			return false, baseLang.NovelTagNameTooLongCode, lang.MsgErr(baseLang.NovelTagNameTooLongCode, e.Lang)
		}
		catId := c.CategoryId
		if catId <= 0 {
			catId = data.CategoryId
		}
		var cnt int64
		if err := e.Orm.Model(&models.NovelTag{}).
			Where("category_id = ? and name = ? and id <> ?", catId, strings.TrimSpace(c.Name), c.Id).Count(&cnt).Error; err != nil {
			return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if cnt > 0 {
			return false, baseLang.NovelTagNameExistCode, lang.MsgErr(baseLang.NovelTagNameExistCode, e.Lang)
		}
		updates["name"] = strings.TrimSpace(c.Name)
	}
	updates["sort"] = c.Sort
	if c.Status != "" {
		updates["status"] = c.Status
	}
	updates["update_by"] = c.UpdateBy
	updates["updated_at"] = time.Now()
	err = e.Orm.Model(&models.NovelTag{}).Where("id = ?", c.Id).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// Delete 后台-删除标签
func (e *NovelTagAdmin) Delete(ids []int64) (int, error) {
	if len(ids) == 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Where("id in (?)", ids).Delete(&models.NovelTag{}).Error; err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}
