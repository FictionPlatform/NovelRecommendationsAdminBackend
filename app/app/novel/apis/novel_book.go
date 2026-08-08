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

type Book struct {
	api.Api
}

// GetPage app-分页查询小说书库
// @Summary 分页查询小说书库
// @Description 分页查询小说书库（分类/连载状态/字数区间/关键字/排序）
// @Tags 小说书库
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param category query string false "分类"
// @Param serialStatus query string false "连载状态"
// @Param status query string false "上架状态"
// @Param isFeatured query int false "首页精选"
// @Param wordCountMin query int false "字数下限"
// @Param wordCountMax query int false "字数上限"
// @Param keyword query string false "关键字（书名/作者/标签）"
// @Param sort query string false "排序 rating|click|latest"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book [get]
func (e Book) GetPage(c *gin.Context) {
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
	uid, _, _ := auth.Auth.GetUserId(c)
	req.CurrUserId = uid
	list, count, respCode, err := s.GetPage(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get app-查询小说详情
// @Summary 查询小说详情
// @Description 查询小说详情（含最新书评，点击+1）
// @Tags 小说书库
// @Accept json
// @Produce json
// @Param id path int true "书籍编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book/{id} [get]
func (e Book) Get(c *gin.Context) {
	req := dto.NovelBookGetReq{}
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
	result, respCode, err := s.Get(req.Id, uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert app-新增小说
// @Summary 新增小说
// @Description 新增小说（内容管理，权限 novel:book:add）
// @Tags 小说书库
// @Accept json
// @Produce json
// @Param body body dto.NovelBookInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book [post]
func (e Book) Insert(c *gin.Context) {
	req := dto.NovelBookInsertReq{}
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
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update app-更新小说
// @Summary 更新小说
// @Description 更新小说（内容管理，权限 novel:book:edit）
// @Tags 小说书库
// @Accept json
// @Produce json
// @Param id path int true "书籍编号"
// @Param body body dto.NovelBookUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book/{id} [put]
func (e Book) Update(c *gin.Context) {
	req := dto.NovelBookUpdateReq{}
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
	b, respCode, err := s.Update(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(baseLang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete app-删除小说
// @Summary 删除小说
// @Description 删除小说（级联书评/书架，话题关联置空；权限 novel:book:del）
// @Tags 小说书库
// @Accept json
// @Produce json
// @Param body body dto.NovelBookDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book [delete]
func (e Book) Delete(c *gin.Context) {
	req := dto.NovelBookDeleteReq{}
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
	respCode, err := s.Delete(req.Ids)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Rank app-查询小说榜单
// @Summary 查询小说榜单
// @Description type=rating 综合评分榜 type=click 点击热度榜
// @Tags 小说书库
// @Accept json
// @Produce json
// @Param type query string false "rating|click"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/book/rank [get]
func (e Book) Rank(c *gin.Context) {
	req := dto.NovelBookRankReq{}
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
	list, respCode, err := s.Rank(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(list, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Home app-查询首页聚合数据
// @Summary 首页聚合：精选+热门TOP+最新/热门话题（无需登录）
// @Description 首页聚合
// @Tags 小说书库
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/home [get]
func (e Book) Home(c *gin.Context) {
	s := service.NovelBook{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, err.Error())
		return
	}
	result, respCode, err := s.GetHome()
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
