package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/runtime"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NovelBook struct {
	service.Service
}

// NewNovelBookService app-实例化小说书库服务
func NewNovelBookService(s *service.Service) *NovelBook {
	var srv = new(NovelBook)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-分页查询小说书库
func (e *NovelBook) GetPage(c *dto.NovelBookQueryReq) ([]models.NovelBook, int64, int, error) {
	var data models.NovelBook
	var list []models.NovelBook
	var count int64

	// 读者端默认只展示上架书籍，管理端传 allStatus=true 查询全部
	if !c.AllStatus && c.Status == "" {
		c.Status = global.SysStatusOk
	}

	db := e.Orm.Model(&data).Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
	)

	// 关键字（书名/作者/标签，不区分大小写）；tags 为 JSON，PG 需 CAST(TEXT)
	if c.Keyword != "" {
		kw := strings.ToLower(c.Keyword)
		if e.Orm.Dialector.Name() == "postgres" {
			db = db.Where("LOWER(title) like ? or LOWER(author) like ? or CAST(tags AS TEXT) like ?", "%"+kw+"%", "%"+kw+"%", "%"+kw+"%")
		} else {
			db = db.Where("LOWER(title) like ? or LOWER(author) like ? or tags like ?", "%"+kw+"%", "%"+kw+"%", "%"+kw+"%")
		}
	}

	// 标签筛选：tags JSON 中包含该标签（精确匹配 JSON 字符串元素）
	if c.Tag != "" {
		tagJSON := `"` + c.Tag + `"`
		if e.Orm.Dialector.Name() == "postgres" {
			db = db.Where("CAST(tags AS TEXT) like ?", "%"+tagJSON+"%")
		} else {
			db = db.Where("tags like ?", "%"+tagJSON+"%")
		}
	}

	// 排序：latest-最新发布 rating-高分优先 click-热门热度
	switch c.Sort {
	case "rating":
		db = db.Order("rating desc, review_count desc")
	case "click":
		db = db.Order("clicks desc")
	case "latest":
		db = db.Order("publish_date desc, id desc")
	default:
		db = db.Order("created_at desc")
	}

	err := db.Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	e.fillBookExt(list, c.CurrUserId)
	return list, count, baseLang.SuccessCode, nil
}

// Get app-查询小说书库详情（点击+1）
func (e *NovelBook) Get(id int64, currUserId int64) (*models.NovelBook, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelBook{}
	err := e.Orm.First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.NovelBookNotExistCode, lang.MsgErr(baseLang.NovelBookNotExistCode, e.Lang)
	}
	// 点击 +1（当日去重防刷；匿名无标识用户不计数）
	if currUserId > 0 {
		cache := runtime.RuntimeConfig.GetCacheAdapter()
		deduped := false
		if cache != nil {
			v, _ := cache.Get(NovelClickCachePrefix, fmt.Sprintf("book:%d:%d", data.Id, currUserId))
			deduped = v != ""
		}
		if !deduped {
			if cache == nil || cache.Set(NovelClickCachePrefix, fmt.Sprintf("book:%d:%d", data.Id, currUserId), "1", 86400) == nil {
				_ = e.Orm.Model(&models.NovelBook{}).Where("id = ?", data.Id).Update("clicks", gorm.Expr("clicks + 1")).Error
				data.Clicks++
			}
		}
	}

	tmp := []models.NovelBook{*data}
	e.fillBookExt(tmp, currUserId)
	*data = tmp[0]

	// 附带最新书评
	var reviews []models.NovelBookReview
	err = e.Orm.Where("book_id = ?", data.Id).Order("created_at desc").Limit(10).Find(&reviews).Error
	if err == nil {
		data.Reviews = reviews
	}
	return data, baseLang.SuccessCode, nil
}

// Insert app-新增小说书库
func (e *NovelBook) Insert(c *dto.NovelBookInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if strings.TrimSpace(c.Title) == "" {
		return 0, baseLang.NovelBookTitleEmptyCode, lang.MsgErr(baseLang.NovelBookTitleEmptyCode, e.Lang)
	}
	if strings.TrimSpace(c.Category) == "" {
		return 0, baseLang.NovelBookCategoryEmptyCode, lang.MsgErr(baseLang.NovelBookCategoryEmptyCode, e.Lang)
	}
	if len([]rune(c.Title)) > 100 {
		return 0, baseLang.NovelBookTitleTooLongCode, lang.MsgErr(baseLang.NovelBookTitleTooLongCode, e.Lang)
	}
	if len([]rune(c.Author)) > 64 {
		return 0, baseLang.NovelBookAuthorTooLongCode, lang.MsgErr(baseLang.NovelBookAuthorTooLongCode, e.Lang)
	}
	if len([]rune(c.Slogan)) > 30 {
		return 0, baseLang.NovelBookSloganTooLongCode, lang.MsgErr(baseLang.NovelBookSloganTooLongCode, e.Lang)
	}
	if len([]rune(c.Cover)) > 500 {
		return 0, baseLang.NovelBookCoverTooLongCode, lang.MsgErr(baseLang.NovelBookCoverTooLongCode, e.Lang)
	}
	if len([]rune(c.Description)) > 500 {
		return 0, baseLang.NovelBookDescTooLongCode, lang.MsgErr(baseLang.NovelBookDescTooLongCode, e.Lang)
	}
	// 书名唯一校验
	count, respCode, err := e.countByTitle(c.Title)
	if err != nil {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.NovelBookTitleExistCode, lang.MsgErr(baseLang.NovelBookTitleExistCode, e.Lang)
	}

	now := time.Now()
	data := models.NovelBook{}
	data.Title = strings.TrimSpace(c.Title)
	data.Author = strings.TrimSpace(c.Author)
	data.Cover = c.Cover
	// 新书默认值对齐需求文档 §3.2：评分 9.5 / 书评数 1 / 点击 1200 / 章节 30 / 当天发布 / 默认 3 标签（首个取所选分类）
	data.Rating = 9.5
	data.ReviewCount = 1
	data.SerialStatus = c.SerialStatus
	if data.SerialStatus == "" {
		data.SerialStatus = "1"
	}
	data.Category = c.Category
	// 分类存在性校验 + 取分类名（默认标签首个用分类名）
	categoryName := ""
	if id, err := strconv.ParseInt(c.Category, 10, 64); err == nil && id > 0 {
		var cat models.NovelCategory
		if err := e.Orm.First(&cat, id).Error; err == nil {
			categoryName = cat.Name
		} else {
			return 0, baseLang.NovelCategoryNotExistCode, lang.MsgErr(baseLang.NovelCategoryNotExistCode, e.Lang)
		}
	} else {
		return 0, baseLang.NovelBookCategoryEmptyCode, lang.MsgErr(baseLang.NovelBookCategoryEmptyCode, e.Lang)
	}
	if len(c.Tags) > 0 {
		tagJSON, _ := json.Marshal(c.Tags)
		data.Tags = string(tagJSON)
	} else {
		tagJSON, _ := json.Marshal([]string{"#" + categoryName, "#新书热推", "#精彩必读"})
		data.Tags = string(tagJSON)
	}
	data.Slogan = c.Slogan
	data.Description = c.Description
	data.Clicks = 1200
	if c.PublishDate != nil && *c.PublishDate != "" {
		if t, err := time.Parse("2006-01-02", *c.PublishDate); err == nil {
			data.PublishDate = &t
		}
	} else {
		if t, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02")); err == nil {
			data.PublishDate = &t
		}
	}
	data.WordCount = c.WordCount
	data.Chapters = c.Chapters
	if data.Chapters <= 0 {
		data.Chapters = 30
	}
	data.IsFeatured = c.IsFeatured
	data.ReadUrl = c.ReadUrl
	data.Status = c.Status
	if data.Status == "" {
		data.Status = global.SysStatusOk
	}
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update app-更新小说书库
func (e *NovelBook) Update(c *dto.NovelBookUpdateReq) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelBook{}
	err := e.Orm.First(data, c.Id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, baseLang.NovelBookNotExistCode, lang.MsgErr(baseLang.NovelBookNotExistCode, e.Lang)
	}
	// 改名时校验唯一
	if strings.TrimSpace(c.Title) != "" && c.Title != data.Title {
		if len([]rune(c.Title)) > 100 {
			return false, baseLang.NovelBookTitleTooLongCode, lang.MsgErr(baseLang.NovelBookTitleTooLongCode, e.Lang)
		}
		count, respCode, err := e.countByTitle(c.Title)
		if err != nil {
			return false, respCode, err
		}
		if count > 0 {
			return false, baseLang.NovelBookTitleExistCode, lang.MsgErr(baseLang.NovelBookTitleExistCode, e.Lang)
		}
	}
	if c.Author != "" && len([]rune(c.Author)) > 64 {
		return false, baseLang.NovelBookAuthorTooLongCode, lang.MsgErr(baseLang.NovelBookAuthorTooLongCode, e.Lang)
	}
	if c.Slogan != "" && len([]rune(c.Slogan)) > 30 {
		return false, baseLang.NovelBookSloganTooLongCode, lang.MsgErr(baseLang.NovelBookSloganTooLongCode, e.Lang)
	}
	if c.Cover != "" && len([]rune(c.Cover)) > 500 {
		return false, baseLang.NovelBookCoverTooLongCode, lang.MsgErr(baseLang.NovelBookCoverTooLongCode, e.Lang)
	}
	if c.Description != "" && len([]rune(c.Description)) > 500 {
		return false, baseLang.NovelBookDescTooLongCode, lang.MsgErr(baseLang.NovelBookDescTooLongCode, e.Lang)
	}

	updates := map[string]interface{}{}
	if c.Title != "" && data.Title != c.Title {
		updates["title"] = c.Title
	}
	if c.Author != "" && data.Author != c.Author {
		updates["author"] = c.Author
	}
	if c.Cover != "" && data.Cover != c.Cover {
		updates["cover"] = c.Cover
	}
	if c.SerialStatus != "" && data.SerialStatus != c.SerialStatus {
		updates["serial_status"] = c.SerialStatus
	}
	if c.Category != "" && data.Category != c.Category {
		// 分类存在性校验
		if id, err := strconv.ParseInt(c.Category, 10, 64); err == nil && id > 0 {
			var cnt int64
			if err := e.Orm.Model(&models.NovelCategory{}).Where("id = ?", id).Count(&cnt).Error; err != nil {
				return false, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
			}
			if cnt == 0 {
				return false, baseLang.NovelCategoryNotExistCode, lang.MsgErr(baseLang.NovelCategoryNotExistCode, e.Lang)
			}
		} else {
			return false, baseLang.NovelBookCategoryEmptyCode, lang.MsgErr(baseLang.NovelBookCategoryEmptyCode, e.Lang)
		}
		updates["category"] = c.Category
	}
	if c.Tags != nil {
		tagJSON, _ := json.Marshal(c.Tags)
		updates["tags"] = string(tagJSON)
	}
	if c.Slogan != "" && data.Slogan != c.Slogan {
		updates["slogan"] = c.Slogan
	}
	if c.Description != "" && data.Description != c.Description {
		updates["description"] = c.Description
	}
	if c.PublishDate != nil && *c.PublishDate != "" {
		if t, err := time.Parse("2006-01-02", *c.PublishDate); err == nil {
			updates["publish_date"] = t
		}
	}
	if c.WordCount > 0 && data.WordCount != c.WordCount {
		updates["word_count"] = c.WordCount
	}
	if c.Chapters > 0 && data.Chapters != c.Chapters {
		updates["chapters"] = c.Chapters
	}
	if c.IsFeatured != data.IsFeatured {
		updates["is_featured"] = c.IsFeatured
	}
	if c.ReadUrl != "" && data.ReadUrl != c.ReadUrl {
		updates["read_url"] = c.ReadUrl
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
	}
	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.NovelBook{}).Where("id = ?", data.Id).Updates(updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Delete app-删除小说书库（级联清理书评/书架，话题关联置空）
func (e *NovelBook) Delete(ids []int64) (int, error) {
	if len(ids) == 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	baseOrm := e.Orm
	e.Orm = baseOrm.Begin()
	defer func() {
		if err := recover(); err != nil {
			e.Orm.Rollback()
			panic(err)
		}
	}()
	var err error
	committed := false
	defer func() {
		if committed {
			e.Orm.Commit()
		} else {
			e.Orm.Rollback()
		}
		e.Orm = baseOrm
	}()

	err = e.Orm.Where("id in (?)", ids).Delete(&models.NovelBook{}).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	err = e.Orm.Where("book_id in (?)", ids).Delete(&models.NovelBookReview{}).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	err = e.Orm.Where("book_id in (?)", ids).Delete(&models.NovelBookshelf{}).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	err = e.Orm.Model(&models.NovelPost{}).Where("ref_book_id in (?)", ids).Update("ref_book_id", nil).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	committed = true
	return baseLang.SuccessCode, nil
}

// Rank app-查询小说榜单
func (e *NovelBook) Rank(c *dto.NovelBookRankReq) ([]models.NovelBook, int, error) {
	var list []models.NovelBook
	db := e.Orm.Where("status = ?", global.SysStatusOk).Limit(50)
	switch c.Type {
	case "click":
		db = db.Order("clicks desc")
	default:
		db = db.Order("rating desc, review_count desc")
	}
	err := db.Find(&list).Error
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	e.fillBookExt(list, c.CurrUserId)
	return list, baseLang.SuccessCode, nil
}

// GetHome app-查询首页聚合数据
func (e *NovelBook) GetHome() (*dto.NovelHomeResp, int, error) {
	resp := &dto.NovelHomeResp{}

	// 精选轮播
	err := e.Orm.Where("status = ? and is_featured = 1", global.SysStatusOk).Order("id desc").Limit(5).Find(&resp.Featured).Error
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 热门 TOP5
	err = e.Orm.Where("status = ?", global.SysStatusOk).Order("clicks desc").Limit(5).Find(&resp.HotBooks).Error
	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	e.fillBookExt(resp.Featured, 0)
	e.fillBookExt(resp.HotBooks, 0)

	// 最新/热门话题（首页各取 5 条）
	postService := NewNovelPostService(&e.Service)
	latestPosts, _, respCode, err := postService.GetPage(&dto.NovelPostQueryReq{Pagination: cDto.Pagination{PageSize: 5}, Sort: "latest"}, 0)
	if err != nil {
		return nil, respCode, err
	}
	hotPosts, _, respCode, err := postService.GetPage(&dto.NovelPostQueryReq{Pagination: cDto.Pagination{PageSize: 5}, Sort: "hot"}, 0)
	if err != nil {
		return nil, respCode, err
	}
	resp.LatestPosts = latestPosts
	resp.HotPosts = hotPosts
	return resp, baseLang.SuccessCode, nil
}

// countByTitle app-内部方法，按书名统计
func (e *NovelBook) countByTitle(title string) (int64, int, error) {
	var count int64
	err := e.Orm.Model(&models.NovelBook{}).Where("title = ?", title).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return count, baseLang.SuccessCode, nil
}

// fillBookExt app-内部方法，填充标签/分类名/字典标签/收藏状态
func (e *NovelBook) fillBookExt(list []models.NovelBook, currUserId int64) {
	if len(list) == 0 {
		return
	}
	// 标签
	for i := range list {
		if list[i].Tags != "" {
			_ = json.Unmarshal([]byte(list[i].Tags), &list[i].TagList)
		}
	}
	// 分类名联查（app_novel_category）
	catIds := make([]int64, 0, len(list))
	for i := range list {
		if list[i].Category != "" {
			if id, err := strconv.ParseInt(list[i].Category, 10, 64); err == nil {
				catIds = append(catIds, id)
			}
		}
	}
	catNameMap := map[int64]string{}
	if len(catIds) > 0 {
		var cats []models.NovelCategory
		if err := e.Orm.Where("id in (?)", catIds).Find(&cats).Error; err == nil {
			for _, c := range cats {
				catNameMap[c.Id] = c.Name
			}
		}
	}
	// 字典标签（连载状态）
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i := range list {
		if id, err := strconv.ParseInt(list[i].Category, 10, 64); err == nil {
			list[i].CategoryLabel = catNameMap[id]
			list[i].CategoryName = catNameMap[id]
		}
		list[i].StatusLabel = dictService.GetLabel("app_novel_serial_status", list[i].SerialStatus)
	}
	// 收藏状态（当前用户）
	if currUserId <= 0 {
		return
	}
	bookIds := make([]int64, 0, len(list))
	for _, b := range list {
		bookIds = append(bookIds, b.Id)
	}
	var shelves []models.NovelBookshelf
	err := e.Orm.Where("user_id = ? and book_id in (?)", currUserId, bookIds).Find(&shelves).Error
	if err != nil {
		return
	}
	collected := map[int64]bool{}
	for _, s := range shelves {
		collected[s.BookId] = true
	}
	for i := range list {
		list[i].IsCollected = collected[list[i].Id]
	}
}

// avgRating app-内部方法，评分联动聚合（保留 1 位小数）
func avgRating(oldRating float64, oldCount int, newScore int) float64 {
	total := oldRating*float64(oldCount) + float64(newScore)
	avg := total / float64(oldCount+1)
	return math.Round(avg*10) / 10
}

// mergeBizErr 合并业务的错误码包装（事务内返回，事务外解包）
type mergeBizErr struct {
	code int
}

func (e *mergeBizErr) Error() string {
	return fmt.Sprintf("merge biz error code=%d", e.code)
}

// MergeBooks 后台-合并书籍：源书（被合并）→ 目标书（保留）
// 迁移：书评/书架收藏/话题引用；聚合：书评数/评分（加权）/点击；源书下架保留
func (e *NovelBook) MergeBooks(c *dto.NovelBookMergeReq) (int, error) {
	if c.SourceBookId <= 0 || c.TargetBookId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.SourceBookId == c.TargetBookId {
		return baseLang.NovelMergeSameBookCode, lang.MsgErr(baseLang.NovelMergeSameBookCode, e.Lang)
	}
	txErr := e.Orm.Transaction(func(tx *gorm.DB) error {
		// 锁定源书与目标书（FOR UPDATE，防并发聚合错乱）
		var source, target models.NovelBook
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&source, c.SourceBookId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &mergeBizErr{baseLang.NovelMergeSourceNotExistCode}
			}
			return err
		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&target, c.TargetBookId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &mergeBizErr{baseLang.NovelMergeTargetNotExistCode}
			}
			return err
		}
		if target.Status != global.SysStatusOk {
			return &mergeBizErr{baseLang.NovelMergeTargetOfflineCode}
		}

		// 1. 迁移书评：book_id 源→目标（保留全部书评）
		if err := tx.Model(&models.NovelBookReview{}).
			Where("book_id = ?", source.Id).
			Update("book_id", target.Id).Error; err != nil {
			return err
		}

		// 2. 迁移书架：先去重（已收藏目标书的用户，删除其源书收藏），再迁移
		var dupUsers []int64
		if err := tx.Model(&models.NovelBookshelf{}).
			Where("book_id = ?", target.Id).Pluck("user_id", &dupUsers).Error; err != nil {
			return err
		}
		if len(dupUsers) > 0 {
			if err := tx.Where("book_id = ? and user_id in (?)", source.Id, dupUsers).
				Delete(&models.NovelBookshelf{}).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&models.NovelBookshelf{}).
			Where("book_id = ?", source.Id).
			Update("book_id", target.Id).Error; err != nil {
			return err
		}

		// 3. 迁移话题引用：ref_book_id 源→目标
		if err := tx.Model(&models.NovelPost{}).
			Where("ref_book_id = ?", source.Id).
			Update("ref_book_id", target.Id).Error; err != nil {
			return err
		}

		// 4. 聚合目标书：书评数（实际 count）+ 评分（加权平均）+ 点击（累加）
		var reviewCount int64
		if err := tx.Model(&models.NovelBookReview{}).
			Where("book_id = ?", target.Id).Count(&reviewCount).Error; err != nil {
			return err
		}
		newRating := 0.0
		if reviewCount > 0 {
			total := target.Rating*float64(target.ReviewCount) + source.Rating*float64(source.ReviewCount)
			newRating = math.Round(total/float64(reviewCount)*10) / 10
		}
		if err := tx.Model(&models.NovelBook{}).Where("id = ?", target.Id).Updates(map[string]interface{}{
			"review_count": reviewCount,
			"rating":       newRating,
			"clicks":       target.Clicks + source.Clicks,
			"updated_at":   time.Now(),
		}).Error; err != nil {
			return err
		}

		// 5. 源书下架保留（status=2），清空书评数/评分
		if err := tx.Model(&models.NovelBook{}).Where("id = ?", source.Id).Updates(map[string]interface{}{
			"status":       global.SysStatusNotOk,
			"review_count": 0,
			"rating":       0,
			"updated_at":   time.Now(),
		}).Error; err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		var biz *mergeBizErr
		if errors.As(txErr, &biz) {
			return biz.code, lang.MsgErr(biz.code, e.Lang)
		}
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, txErr)
	}
	return baseLang.SuccessCode, nil
}
