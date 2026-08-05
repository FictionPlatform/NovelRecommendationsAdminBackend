package service

import (
	"encoding/json"
	"errors"
	"math"
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
	"gorm.io/gorm"
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

	db := e.Orm.Model(&data).Scopes(cDto.MakeCondition(c.GetNeedSearch()))

	// 关键字（书名/作者/标签，不区分大小写）
	if c.Keyword != "" {
		kw := strings.ToLower(c.Keyword)
		db = db.Where("LOWER(title) like ? or LOWER(author) like ? or LOWER(tags) like ?", "%"+kw+"%", "%"+kw+"%", "%"+kw+"%")
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
	// 点击 +1（防刷新刷量后续可按 IP/会话去重）
	_ = e.Orm.Model(&models.NovelBook{}).Where("id = ?", data.Id).Update("clicks", gorm.Expr("clicks + 1")).Error
	data.Clicks++

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
	data.Rating = 5.0
	data.ReviewCount = 0
	data.SerialStatus = c.SerialStatus
	if data.SerialStatus == "" {
		data.SerialStatus = "1"
	}
	data.Category = c.Category
	if len(c.Tags) > 0 {
		tagJSON, _ := json.Marshal(c.Tags)
		data.Tags = string(tagJSON)
	}
	data.Slogan = c.Slogan
	data.Description = c.Description
	data.Clicks = 0
	if c.PublishDate != nil && *c.PublishDate != "" {
		if t, err := time.Parse("2006-01-02", *c.PublishDate); err == nil {
			data.PublishDate = &t
		}
	}
	data.WordCount = c.WordCount
	data.Chapters = c.Chapters
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
		count, respCode, err := e.countByTitle(c.Title)
		if err != nil {
			return false, respCode, err
		}
		if count > 0 {
			return false, baseLang.NovelBookTitleExistCode, lang.MsgErr(baseLang.NovelBookTitleExistCode, e.Lang)
		}
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

// Delete app-删除小说书库（级联清理书评/书架，帖子关联置空）
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

	// 最新/热门帖子
	postService := NewNovelPostService(&e.Service)
	latestPosts, _, respCode, err := postService.GetPage(&dto.NovelPostQueryReq{Sort: "latest"}, 0)
	if err != nil {
		return nil, respCode, err
	}
	hotPosts, _, respCode, err := postService.GetPage(&dto.NovelPostQueryReq{Sort: "hot"}, 0)
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

// fillBookExt app-内部方法，填充标签/字典标签/收藏状态
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
	// 字典标签
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i := range list {
		list[i].CategoryLabel = dictService.GetLabel("app_novel_category", list[i].Category)
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
