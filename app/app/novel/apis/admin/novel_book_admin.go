package admin

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

type NovelBookAdmin struct {
	api.Api
}

// GetPage 后台分页查询书籍（含下架）
// @Summary 后台分页查询书籍
// @Description 管理端查询全部书籍（含下架），支持关键字/分类/状态筛选
// @Tags 小说平台书籍管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param keyword query string false "关键字（书名/作者/标签）"
// @Param category query string false "分类"
// @Param status query string false "上架状态 1-上架 2-下架"
// @Param serialStatus query string false "连载状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book [get]
func (e NovelBookAdmin) GetPage(c *gin.Context) {
	req := dto.NovelBookQueryReq{}
	s := service.NovelBook{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	req.AllStatus = true
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	list, count, respCode, err := s.GetPage(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Merge 后台合并书籍
// @Summary 后台合并书籍
// @Description 将源书（被合并）的书评/书架收藏/话题引用迁移到目标书（保留），聚合评分/书评数/点击，源书下架保留
// @Tags 小说平台书籍管理
// @Accept json
// @Produce json
// @Param body body dto.NovelBookMergeReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book/merge [post]
func (e NovelBookAdmin) Merge(c *gin.Context) {
	req := dto.NovelBookMergeReq{}
	s := service.NovelBook{}
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
	respCode, err := s.MergeBooks(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
