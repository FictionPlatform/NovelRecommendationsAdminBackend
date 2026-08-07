package service

import (
	"errors"
	"math"
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
	"gorm.io/gorm/clause"
)

// 互动/排序类型常量
const (
	NovelInteractLike    = "like"
	NovelInteractDislike = "dislike"
	NovelInteractCollect = "collect"
	NovelActionAdd       = "add"
	NovelActionCancel    = "cancel"

	NovelLikeType    = "1" // app_novel_post_like.type 点赞
	NovelDislikeType = "2" // app_novel_post_like.type 踩
)

// maxZero 计数下限 0
func maxZero(n int) int {
	if n < 0 {
		return 0
	}
	return n
}

type NovelPost struct {
	service.Service
}

// NewNovelPostService app-实例化小说长文帖子服务
func NewNovelPostService(s *service.Service) *NovelPost {
	var srv = new(NovelPost)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-分页查询小说长文帖子
func (e *NovelPost) GetPage(c *dto.NovelPostQueryReq, currUserId int64) ([]models.NovelPost, int64, int, error) {
	var data models.NovelPost
	var list []models.NovelPost
	var count int64

	db := e.Orm.Model(&data).Where("status = ?", global.SysStatusOk)
	// 只看我的帖子
	if c.Mine == 1 {
		if currUserId <= 0 {
			return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
		}
		db = db.Where("user_id = ?", currUserId)
	}
	// 只看我收藏的帖子
	if c.IsCollected == 1 {
		if currUserId <= 0 {
			return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
		}
		db = db.Where("id in (select post_id from app_novel_post_collect where user_id = ?)", currUserId)
	}
	db = db.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
	)
	// 排序：latest-最新 hot-热门（热门=点赞+评论数）
	switch c.Sort {
	case "hot":
		db = db.Order("(likes + comment_count) desc, created_at desc")
	default:
		db = db.Order("created_at desc, id desc")
	}
	err := db.Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	e.fillPostExt(list, currUserId)
	return list, count, baseLang.SuccessCode, nil
}

// Get app-查询帖子详情（含评论与我的互动状态）
func (e *NovelPost) Get(id int64, currUserId int64) (*models.NovelPost, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.NovelPost{}
	err := e.Orm.Where("status = ?", global.SysStatusOk).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.NovelPostNotExistCode, lang.MsgErr(baseLang.NovelPostNotExistCode, e.Lang)
	}

	tmp := []models.NovelPost{*data}
	e.fillPostExt(tmp, currUserId)
	*data = tmp[0]

	// 评论（楼中楼）
	var comments []models.NovelPostComment
	err = e.Orm.Where("post_id = ?", data.Id).Order("created_at asc").Find(&comments).Error
	if err == nil {
		commentMap := map[int64]*models.NovelPostComment{}
		for i := range comments {
			commentMap[comments[i].Id] = &comments[i]
		}
		root := make([]models.NovelPostComment, 0)
		for i := range comments {
			cm := &comments[i]
			if cm.ParentId != nil && *cm.ParentId > 0 {
				if parent, ok := commentMap[*cm.ParentId]; ok {
					parent.Replies = append(parent.Replies, *cm)
					continue
				}
			}
			root = append(root, *cm)
		}
		data.Comments = root
	}
	return data, baseLang.SuccessCode, nil
}

// Insert app-发布长文帖子（内容 5000~10000 字）
func (e *NovelPost) Insert(c *dto.NovelPostInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if strings.TrimSpace(c.Title) == "" {
		return 0, baseLang.NovelPostTitleEmptyCode, lang.MsgErr(baseLang.NovelPostTitleEmptyCode, e.Lang)
	}
	if len([]rune(c.Title)) > 200 {
		return 0, baseLang.NovelPostTitleTooLongCode, lang.MsgErr(baseLang.NovelPostTitleTooLongCode, e.Lang)
	}
	contentLen := len([]rune(c.Content))
	if contentLen < 5000 || contentLen > 10000 {
		return 0, baseLang.NovelPostContentLenCode, lang.MsgErr(baseLang.NovelPostContentLenCode, e.Lang)
	}
	if c.RefBookId != nil && *c.RefBookId > 0 {
		book := &models.NovelBook{}
		err := e.Orm.First(book, *c.RefBookId).Error
		if err != nil {
			return 0, baseLang.NovelPostRefBookNotExistCode, lang.MsgErr(baseLang.NovelPostRefBookNotExistCode, e.Lang)
		}
	}

	userName, userAvatar, respCode, err := getUserSnapshot(&e.Service, c.CurrUserId)
	if err != nil {
		return 0, respCode, err
	}

	// 禁言校验：管理员设置 ban_post_until 未到期时拒绝发帖
	profile := &models.NovelReaderProfile{}
	perr := e.Orm.Where("user_id = ?", c.CurrUserId).First(profile).Error
	if perr == nil && profile.BanPostUntil != nil && profile.BanPostUntil.After(time.Now()) {
		return 0, baseLang.NovelPostBannedCode, lang.MsgErr(baseLang.NovelPostBannedCode, e.Lang)
	}
	if perr != nil && !errors.Is(perr, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, perr)
	}

	now := time.Now()
	data := models.NovelPost{}
	data.UserId = c.CurrUserId
	data.UserName = userName
	data.UserAvatar = userAvatar
	data.Title = strings.TrimSpace(c.Title)
	data.Content = c.Content
	// summary = 前 150 字
	runes := []rune(c.Content)
	if contentLen > 150 {
		data.Summary = string(runes[:150]) + "..."
	} else {
		data.Summary = string(runes)
	}
	data.WordCount = contentLen
	// read_time = ceil(字符数/400)
	data.ReadTime = int(math.Ceil(float64(contentLen) / 400))
	if data.ReadTime < 1 {
		data.ReadTime = 1
	}
	data.TopicTag = c.TopicTag
	data.RefBookId = c.RefBookId
	data.Status = global.SysStatusOk
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Delete app-删除帖子（本人，级联清理互动与评论）
func (e *NovelPost) Delete(ids []int64, currUserId int64) (int, error) {
	if len(ids) == 0 || currUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 本人校验
	var list []models.NovelPost
	err := e.Orm.Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	for _, p := range list {
		if p.UserId != currUserId {
			return baseLang.NovelNoPermissionCode, lang.MsgErr(baseLang.NovelNoPermissionCode, e.Lang)
		}
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

	// 级联清理互动与评论
	txErr = e.Orm.Where("post_id in (?)", ids).Delete(&models.NovelPostLike{}).Error
	if txErr != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, txErr)
	}
	txErr = e.Orm.Where("post_id in (?)", ids).Delete(&models.NovelPostCollect{}).Error
	if txErr != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, txErr)
	}
	txErr = e.Orm.Where("post_id in (?)", ids).Delete(&models.NovelPostComment{}).Error
	if txErr != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, txErr)
	}
	txErr = e.Orm.Where("id in (?)", ids).Delete(&models.NovelPost{}).Error
	if txErr != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, txErr)
	}
	committed = true
	return baseLang.SuccessCode, nil
}

// Interact app-帖子互动（点赞/踩/收藏，事务内互斥）
func (e *NovelPost) Interact(c *dto.NovelPostInteractReq) (int, error) {
	if c.CurrUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Type != NovelInteractLike && c.Type != NovelInteractDislike && c.Type != NovelInteractCollect {
		return baseLang.NovelActionTypeErrCode, lang.MsgErr(baseLang.NovelActionTypeErrCode, e.Lang)
	}
	if c.Action != NovelActionAdd && c.Action != NovelActionCancel {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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

	// 锁定帖子行，保证计数原子
	post := &models.NovelPost{}
	txErr = e.Orm.Clauses(clause.Locking{Strength: "UPDATE"}).First(post, c.Id).Error
	if txErr != nil {
		if errors.Is(txErr, gorm.ErrRecordNotFound) {
			txErr = nil
			return baseLang.NovelPostNotExistCode, lang.MsgErr(baseLang.NovelPostNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, txErr)
	}
	if post.Status != global.SysStatusOk {
		return baseLang.NovelPostNotExistCode, lang.MsgErr(baseLang.NovelPostNotExistCode, e.Lang)
	}

	switch c.Type {
	case NovelInteractCollect:
		// 收藏（唯一索引 (user_id, post_id) 兜底）
		var old models.NovelPostCollect
		err := e.Orm.Where("user_id = ? and post_id = ?", c.CurrUserId, c.Id).First(&old).Error
		hasOld := true
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				hasOld = false
			} else {
				txErr = err
				return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
			}
		}
		if c.Action == NovelActionAdd {
			if hasOld {
				return baseLang.DataNotUpdateCode, lang.MsgErr(baseLang.DataNotUpdateCode, e.Lang)
			}
			now := time.Now()
			if err := e.Orm.Create(&models.NovelPostCollect{PostId: c.Id, UserId: c.CurrUserId, CreatedAt: &now}).Error; err != nil {
				txErr = err
				return baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
			}
			post.Collections++
		} else {
			if !hasOld {
				return baseLang.DataNotUpdateCode, lang.MsgErr(baseLang.DataNotUpdateCode, e.Lang)
			}
			if err := e.Orm.Where("user_id = ? and post_id = ?", c.CurrUserId, c.Id).Delete(&models.NovelPostCollect{}).Error; err != nil {
				txErr = err
				return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
			}
			post.Collections = maxZero(post.Collections - 1)
		}

	case NovelInteractLike, NovelInteractDislike:
		likeType := NovelLikeType
		if c.Type == NovelInteractDislike {
			likeType = NovelDislikeType
		}
		// 一条 (user_id, post_id) 唯一记录（UNIQUE 兜底），赞/踩互斥
		old := &models.NovelPostLike{}
		err := e.Orm.Where("user_id = ? and post_id = ?", c.CurrUserId, c.Id).First(old).Error
		hasOld := true
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				hasOld = false
				old.Type = ""
			} else {
				txErr = err
				return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
			}
		}

		if c.Action == NovelActionAdd {
			if hasOld && old.Type == likeType {
				// 幂等：已存在同类型互动
				return baseLang.DataNotUpdateCode, lang.MsgErr(baseLang.DataNotUpdateCode, e.Lang)
			}
			// 互斥：先撤销旧记录并回滚计数
			if hasOld && old.Type != likeType {
				if err := e.Orm.Where("user_id = ? and post_id = ?", c.CurrUserId, c.Id).Delete(&models.NovelPostLike{}).Error; err != nil {
					txErr = err
					return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
				}
				if old.Type == NovelLikeType {
					post.Likes = maxZero(post.Likes - 1)
				} else {
					post.Dislikes = maxZero(post.Dislikes - 1)
				}
			}
			if err := e.Orm.Create(&models.NovelPostLike{PostId: c.Id, UserId: c.CurrUserId, Type: likeType}).Error; err != nil {
				txErr = err
				return baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
			}
			if likeType == NovelLikeType {
				post.Likes++
			} else {
				post.Dislikes++
			}
		} else {
			// cancel：仅当存在同类型记录时撤销
			if !hasOld || old.Type != likeType {
				return baseLang.DataNotUpdateCode, lang.MsgErr(baseLang.DataNotUpdateCode, e.Lang)
			}
			if err := e.Orm.Where("user_id = ? and post_id = ? and type = ?", c.CurrUserId, c.Id, likeType).Delete(&models.NovelPostLike{}).Error; err != nil {
				txErr = err
				return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
			}
			if old.Type == NovelLikeType {
				post.Likes = maxZero(post.Likes - 1)
			} else {
				post.Dislikes = maxZero(post.Dislikes - 1)
			}
		}
	}

	// 回写计数（事务内行锁，避免并发覆盖）
	txErr = e.Orm.Model(&models.NovelPost{}).Where("id = ?", post.Id).
		Updates(map[string]interface{}{
			"likes":       post.Likes,
			"dislikes":    post.Dislikes,
			"collections": post.Collections,
			"updated_at":  time.Now(),
		}).Error
	if txErr != nil {
		return baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, txErr)
	}
	committed = true
	return baseLang.SuccessCode, nil
}

// AddComment app-新增评论/回复（@昵称，事务内校验帖子并回写计数）
func (e *NovelPost) AddComment(c *dto.NovelPostCommentInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.PostId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if strings.TrimSpace(c.Content) == "" {
		return 0, baseLang.NovelPostCommentEmptyCode, lang.MsgErr(baseLang.NovelPostCommentEmptyCode, e.Lang)
	}
	if len([]rune(c.Content)) > 1000 {
		return 0, baseLang.NovelContentTooLongCode, lang.MsgErr(baseLang.NovelContentTooLongCode, e.Lang)
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

	// 锁定帖子行，校验存在且状态正常（防孤儿评论 + comment_count 空增）
	post := &models.NovelPost{}
	txErr = e.Orm.Clauses(clause.Locking{Strength: "UPDATE"}).First(post, c.PostId).Error
	if txErr != nil {
		if errors.Is(txErr, gorm.ErrRecordNotFound) {
			txErr = nil
			return 0, baseLang.NovelPostNotExistCode, lang.MsgErr(baseLang.NovelPostNotExistCode, e.Lang)
		}
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, txErr)
	}
	if post.Status != global.SysStatusOk {
		return 0, baseLang.NovelPostNotExistCode, lang.MsgErr(baseLang.NovelPostNotExistCode, e.Lang)
	}

	// 父评论存在性 + 必须属于同一帖子（楼中楼结构一致）
	if c.ParentId != nil && *c.ParentId > 0 {
		parent := &models.NovelPostComment{}
		err := e.Orm.First(parent, *c.ParentId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, baseLang.NovelPostCommentNotExistCode, lang.MsgErr(baseLang.NovelPostCommentNotExistCode, e.Lang)
			}
			txErr = err
			return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
		}
		if parent.PostId != post.Id {
			return 0, baseLang.NovelParentNotSamePostCode, lang.MsgErr(baseLang.NovelParentNotSamePostCode, e.Lang)
		}
	}

	userName, userAvatar, respCode, err := getUserSnapshot(&e.Service, c.CurrUserId)
	if err != nil {
		return 0, respCode, err
	}

	now := time.Now()
	data := models.NovelPostComment{}
	data.PostId = c.PostId
	data.UserId = c.CurrUserId
	data.UserName = userName
	data.UserAvatar = userAvatar
	data.ParentId = c.ParentId
	data.ReplyToUser = c.ReplyToUser
	data.ReplyToCommentId = c.ReplyToCommentId
	data.Content = c.Content
	data.CreatedAt = &now
	if err := e.Orm.Create(&data).Error; err != nil {
		txErr = err
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	// 评论数 +1（事务内与帖子行锁联动）
	if err := e.Orm.Model(&models.NovelPost{}).Where("id = ?", c.PostId).
		Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
		txErr = err
		return 0, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	committed = true
	return data.Id, baseLang.SuccessCode, nil
}

// DeleteComment app-删除评论（本人）
func (e *NovelPost) DeleteComment(id int64, currUserId int64) (int, error) {
	if id <= 0 || currUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	comment := &models.NovelPostComment{}
	err := e.Orm.First(comment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return baseLang.NovelPostCommentNotExistCode, lang.MsgErr(baseLang.NovelPostCommentNotExistCode, e.Lang)
		}
		return baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if comment.UserId != currUserId {
		return baseLang.NovelNoPermissionCode, lang.MsgErr(baseLang.NovelNoPermissionCode, e.Lang)
	}
	if err := e.Orm.Where("id = ?", comment.Id).Delete(&models.NovelPostComment{}).Error; err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	// 评论数 -1（下限 0）
	_ = e.Orm.Model(&models.NovelPost{}).Where("id = ?", comment.PostId).
		Update("comment_count", gorm.Expr("CASE WHEN comment_count > 0 THEN comment_count - 1 ELSE 0 END")).Error
	return baseLang.SuccessCode, nil
}

// GetMyComments app-分页查询我的评论（含原帖标题）
func (e *NovelPost) GetMyComments(c *dto.NovelCommentQueryReq) ([]models.NovelPostComment, int64, int, error) {
	if c.CurrUserId <= 0 {
		return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.NovelPostComment
	var list []models.NovelPostComment
	var count int64
	err := e.Orm.Model(&data).Where("user_id = ?", c.CurrUserId).
		Order("created_at desc, id desc").
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	e.fillCommentExt(list)
	return list, count, baseLang.SuccessCode, nil
}

// fillCommentExt app-内部方法，批量填充评论所属帖子标题
func (e *NovelPost) fillCommentExt(list []models.NovelPostComment) {
	if len(list) == 0 {
		return
	}
	postIds := make([]int64, 0, len(list))
	for _, cm := range list {
		postIds = append(postIds, cm.PostId)
	}
	var posts []models.NovelPost
	err := e.Orm.Where("id in (?)", postIds).Find(&posts).Error
	if err != nil {
		return
	}
	titleMap := map[int64]string{}
	for _, p := range posts {
		titleMap[p.Id] = p.Title
	}
	for i := range list {
		list[i].PostTitle = titleMap[list[i].PostId]
	}
}

// fillPostExt app-内部方法，填充关联书名与当前用户互动状态
func (e *NovelPost) fillPostExt(list []models.NovelPost, currUserId int64) {
	if len(list) == 0 {
		return
	}
	postIds := make([]int64, 0, len(list))
	for _, p := range list {
		postIds = append(postIds, p.Id)
	}
	// 关联书名
	bookIds := make([]int64, 0, len(list))
	for _, p := range list {
		if p.RefBookId != nil && *p.RefBookId > 0 {
			bookIds = append(bookIds, *p.RefBookId)
		}
	}
	if len(bookIds) > 0 {
		var books []models.NovelBook
		_ = e.Orm.Where("id in (?)", bookIds).Find(&books).Error
		titleMap := map[int64]string{}
		for _, b := range books {
			titleMap[b.Id] = b.Title
		}
		for i := range list {
			if list[i].RefBookId != nil {
				list[i].RefBookTitle = titleMap[*list[i].RefBookId]
			}
		}
	}
	if currUserId <= 0 {
		return
	}
	// 我的互动状态
	var likes []models.NovelPostLike
	_ = e.Orm.Where("user_id = ? and post_id in (?)", currUserId, postIds).Find(&likes).Error
	likeState := map[int64]string{}
	for _, l := range likes {
		likeState[l.PostId] = l.Type
	}
	var collects []models.NovelPostCollect
	_ = e.Orm.Where("user_id = ? and post_id in (?)", currUserId, postIds).Find(&collects).Error
	collectMap := map[int64]bool{}
	for _, cl := range collects {
		collectMap[cl.PostId] = true
	}
	for i := range list {
		if t, ok := likeState[list[i].Id]; ok {
			list[i].IsLiked = t == NovelLikeType
			list[i].IsDisliked = t == NovelDislikeType
		}
		list[i].IsCollected = collectMap[list[i].Id]
	}
}
