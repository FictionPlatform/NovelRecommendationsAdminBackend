package service

import (
	"encoding/json"
	"errors"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/utils/dberr"
	"gorm.io/gorm"
)

type NovelFollow struct {
	service.Service
}

// NewNovelFollowService app-实例化书友关注服务
func NewNovelFollowService(s *service.Service) *NovelFollow {
	var srv = new(NovelFollow)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// Follow app-关注书友（幂等，唯一索引兜底）
func (e *NovelFollow) Follow(c *dto.NovelFollowInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.TargetUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.CurrUserId == c.TargetUserId {
		return 0, baseLang.NovelFollowSelfCode, lang.MsgErr(baseLang.NovelFollowSelfCode, e.Lang)
	}
	// 目标用户存在性校验
	u := &userModels.User{}
	err := e.Orm.First(u, c.TargetUserId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	// 已关注 → 幂等返回既有记录
	old := &models.NovelFollow{}
	err = e.Orm.Where("user_id = ? and follow_user_id = ?", c.CurrUserId, c.TargetUserId).First(old).Error
	if err == nil {
		return old.Id, baseLang.NovelFollowAlreadyExistCode, lang.MsgErr(baseLang.NovelFollowAlreadyExistCode, e.Lang)
	}
	now := time.Now()
	data := models.NovelFollow{UserId: c.CurrUserId, FollowUserId: c.TargetUserId, CreatedAt: &now}
	err = e.Orm.Create(&data).Error
	if err != nil {
		if dberr.IsDuplicateKey(err) {
			return 0, baseLang.NovelFollowAlreadyExistCode, lang.MsgErr(baseLang.NovelFollowAlreadyExistCode, e.Lang)
		}
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Unfollow app-取消关注
func (e *NovelFollow) Unfollow(c *dto.NovelFollowDeleteReq) (int, error) {
	if c.CurrUserId <= 0 || c.TargetUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.CurrUserId == c.TargetUserId {
		return baseLang.NovelFollowSelfCode, lang.MsgErr(baseLang.NovelFollowSelfCode, e.Lang)
	}
	res := e.Orm.Where("user_id = ? and follow_user_id = ?", c.CurrUserId, c.TargetUserId).Delete(&models.NovelFollow{})
	if res.Error != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, res.Error)
	}
	if res.RowsAffected == 0 {
		return baseLang.NovelFollowNotExistCode, lang.MsgErr(baseLang.NovelFollowNotExistCode, e.Lang)
	}
	return baseLang.SuccessCode, nil
}

// GetFollowingPage app-分页查询我的关注列表
func (e *NovelFollow) GetFollowingPage(c *dto.NovelFollowQueryReq) ([]dto.NovelFollowingItem, int64, int, error) {
	if c.CurrUserId <= 0 {
		return nil, 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var data models.NovelFollow
	var list []models.NovelFollow
	var count int64
	err := e.Orm.Model(&data).Where("user_id = ?", c.CurrUserId).
		Order("created_at desc, id desc").
		Scopes(cDto.Paginate(c.GetPageSize(), c.GetPageIndex())).
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	items := make([]dto.NovelFollowingItem, 0, len(list))
	for _, f := range list {
		name, avatar, bio := getReaderUserInfo(&e.Service, f.FollowUserId)
		items = append(items, dto.NovelFollowingItem{
			UserId:     f.FollowUserId,
			Name:       name,
			Avatar:     avatar,
			Bio:        bio,
			IsFollowed: true,
		})
	}
	return items, count, baseLang.SuccessCode, nil
}

// GetProfile app-书友名片聚合（帖子/书评/评论/收藏/关注）
func (e *NovelFollow) GetProfile(c *dto.NovelProfileReq) (*dto.NovelProfileResp, int, error) {
	if c.UserId <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	// 目标用户存在性校验
	u := &userModels.User{}
	if err := e.Orm.First(u, c.UserId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, baseLang.NovelUserNotExistCode, lang.MsgErr(baseLang.NovelUserNotExistCode, e.Lang)
		}
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	name, avatar, bio := getReaderUserInfo(&e.Service, c.UserId)
	resp := &dto.NovelProfileResp{
		UserId: c.UserId,
		Name:   name,
		Avatar: avatar,
		Bio:    bio,
	}
	resp.IsSelf = c.UserId == c.CurrUserId
	// 当前用户对目标用户的关注状态
	if c.CurrUserId > 0 && !resp.IsSelf {
		var cnt int64
		_ = e.Orm.Model(&models.NovelFollow{}).
			Where("user_id = ? and follow_user_id = ?", c.CurrUserId, c.UserId).Count(&cnt).Error
		resp.IsFollowed = cnt > 0
	}

	// 帖子（最新 5 条 + 总数）
	postSvc := NewNovelPostService(&e.Service)
	_ = e.Orm.Where("user_id = ? and status = ?", c.UserId, global.SysStatusOk).
		Order("created_at desc, id desc").Limit(5).Find(&resp.Posts).Error
	_ = e.Orm.Model(&models.NovelPost{}).
		Where("user_id = ? and status = ?", c.UserId, global.SysStatusOk).Count(&resp.PostCount).Error
	postSvc.fillPostExt(resp.Posts, c.CurrUserId)

	// 书评（最新 5 条 + 总数）
	_ = e.Orm.Where("user_id = ?", c.UserId).
		Order("created_at desc, id desc").Limit(5).Find(&resp.Reviews).Error
	_ = e.Orm.Model(&models.NovelBookReview{}).
		Where("user_id = ?", c.UserId).Count(&resp.ReviewCount).Error
	e.fillReviewBooks(resp.Reviews)

	// 评论（最新 5 条 + 总数，含帖子标题）
	_ = e.Orm.Where("user_id = ?", c.UserId).
		Order("created_at desc, id desc").Limit(5).Find(&resp.Comments).Error
	_ = e.Orm.Model(&models.NovelPostComment{}).
		Where("user_id = ?", c.UserId).Count(&resp.CommentCount).Error
	postSvc.fillCommentExt(resp.Comments)

	// 收藏的帖子（最新 5 条 + 总数）
	_ = e.Orm.Where("id in (select post_id from app_novel_post_collect where user_id = ?) and status = ?", c.UserId, global.SysStatusOk).
		Order("created_at desc, id desc").Limit(5).Find(&resp.CollectPosts).Error
	_ = e.Orm.Model(&models.NovelPostCollect{}).
		Where("user_id = ?", c.UserId).Count(&resp.CollectPostCount).Error
	postSvc.fillPostExt(resp.CollectPosts, c.CurrUserId)

	// 书架书籍（最新 5 条 + 总数）
	_ = e.Orm.Where("user_id = ?", c.UserId).
		Order("created_at desc, id desc").Limit(5).Find(&resp.Books).Error
	_ = e.Orm.Model(&models.NovelBookshelf{}).
		Where("user_id = ?", c.UserId).Count(&resp.BookCount).Error
	fillShelfBooks(&e.Service, resp.Books)

	// 关注/粉丝数
	_ = e.Orm.Model(&models.NovelFollow{}).Where("user_id = ?", c.UserId).Count(&resp.FollowCount).Error
	_ = e.Orm.Model(&models.NovelFollow{}).Where("follow_user_id = ?", c.UserId).Count(&resp.FanCount).Error

	return resp, baseLang.SuccessCode, nil
}

// getReaderUserInfo app-内部方法，读取用户公开画像（优先读者资料，回退 app_user 快照）
func getReaderUserInfo(e *service.Service, userId int64) (string, string, string) {
	userName, userAvatar, _, _ := getUserSnapshot(e, userId)
	name, avatar, bio := userName, userAvatar, ""
	profile := &models.NovelReaderProfile{}
	pErr := e.Orm.Where("user_id = ?", userId).First(profile).Error
	if pErr == nil {
		if profile.Nickname != "" {
			name = profile.Nickname
		}
		if profile.Avatar != "" {
			avatar = profile.Avatar
		}
		bio = profile.Bio
	}
	return name, avatar, bio
}

// fillReviewBooks app-内部方法，批量填充书评关联书籍
func (e *NovelFollow) fillReviewBooks(list []models.NovelBookReview) {
	if len(list) == 0 {
		return
	}
	bookIds := make([]int64, 0, len(list))
	for _, r := range list {
		bookIds = append(bookIds, r.BookId)
	}
	var books []models.NovelBook
	err := e.Orm.Where("id in (?)", bookIds).Find(&books).Error
	if err != nil {
		return
	}
	bookMap := map[int64]models.NovelBook{}
	for _, b := range books {
		if b.Tags != "" {
			_ = json.Unmarshal([]byte(b.Tags), &b.TagList)
		}
		bookMap[b.Id] = b
	}
	for i := range list {
		if b, ok := bookMap[list[i].BookId]; ok {
			list[i].Book = &b
		}
	}
}
