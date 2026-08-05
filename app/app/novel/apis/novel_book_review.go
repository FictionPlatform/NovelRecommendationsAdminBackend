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

type BookReview struct {
	api.Api
}

// GetPage app-书评圈分页查询
// @Summary 书评圈分页查询
// @Description filter=all 全部 five 5星 hot 热门
// @Tags 小说书评
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param bookId query int false "书籍编号"
// @Param filter query string false "all|five|hot"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book-review/page [get]
func (e BookReview) GetPage(c *gin.Context) {
	req := dto.NovelReviewQueryReq{}
	s := service.NovelBookReview{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	list, count, respCode, err := s.GetPage(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert app-新增书评（评分联动）
// @Summary 新增书评
// @Description 新增书评并联动书籍评分（事务聚合）
// @Tags 小说书评
// @Accept json
// @Produce json
// @Param body body dto.NovelReviewInsertReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book-review [post]
func (e BookReview) Insert(c *gin.Context) {
	req := dto.NovelReviewInsertReq{}
	s := service.NovelBookReview{}
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

// Delete app-删除我的书评
// @Summary 删除我的书评
// @Description 仅本人可删除
// @Tags 小说书评
// @Accept json
// @Produce json
// @Param body body dto.NovelReviewDeleteReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book-review [delete]
func (e BookReview) Delete(c *gin.Context) {
	req := dto.NovelReviewDeleteReq{}
	s := service.NovelBookReview{}
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
