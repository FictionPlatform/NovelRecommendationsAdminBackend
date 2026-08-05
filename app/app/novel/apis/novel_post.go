package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/service"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"
)

type Post struct {
	api.Api
}

// GetPage app-分页查询长文帖子
// @Summary 分页查询长文帖子
// @Description sort=latest 最新 hot 热门；topicTag 话题标签筛选
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param topicTag query string false "话题标签"
// @Param sort query string false "latest|hot"
// @Param refBookId query int false "关联小说"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post/page [get]
func (e Post) GetPage(c *gin.Context) {
	req := dto.NovelPostQueryReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, _, _ := auth.Auth.GetUserId(c)
	list, count, respCode, err := s.GetPage(&req, uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get app-查询帖子详情
// @Summary 查询帖子详情
// @Description 帖子详情（含楼中楼评论与当前用户互动状态）
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param id path int true "帖子编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post/{id} [get]
func (e Post) Get(c *gin.Context) {
	req := dto.NovelPostGetReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	result, respCode, err := s.Get(req.Id, uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert app-发布长文帖子
// @Summary 发布长文帖子
// @Description 正文 5000~10000 字
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param body body dto.NovelPostInsertReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post [post]
func (e Post) Insert(c *gin.Context) {
	req := dto.NovelPostInsertReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete app-删除我的帖子
// @Summary 删除我的帖子
// @Description 仅本人可删除（级联清理互动与评论）
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param body body dto.NovelPostDeleteReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post [delete]
func (e Post) Delete(c *gin.Context) {
	req := dto.NovelPostDeleteReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	respCode, err := s.Delete(req.Ids, uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Interact app-帖子互动
// @Summary 帖子点赞/踩/收藏
// @Description type=like|dislike|collect action=add|cancel（赞踩互斥）
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param id path int true "帖子编号"
// @Param body body dto.NovelPostInteractReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post/{id}/interact [post]
func (e Post) Interact(c *gin.Context) {
	req := dto.NovelPostInteractReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	respCode, err := s.Interact(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// AddComment app-新增帖子评论/回复
// @Summary 新增帖子评论/回复
// @Description 支持楼中楼（parentId）与 @昵称（replyToUser）
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param id path int true "帖子编号"
// @Param body body dto.NovelPostCommentInsertReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post/{id}/comment [post]
func (e Post) AddComment(c *gin.Context) {
	req := dto.NovelPostCommentInsertReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.AddComment(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// DeleteComment app-删除评论
// @Summary 删除评论（仅本人）
// @Tags 小说长文
// @Accept json
// @Produce json
// @Param id path int true "评论编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post-comment/{id} [delete]
func (e Post) DeleteComment(c *gin.Context) {
	req := dto.NovelPostCommentDeleteReq{}
	s := service.NovelPost{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	respCode, err := s.DeleteComment(req.Id, uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
