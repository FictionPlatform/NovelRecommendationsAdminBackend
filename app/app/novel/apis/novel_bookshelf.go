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

type BookShelf struct {
	api.Api
}

// GetPage app-分页查询我的书架
// @Summary 分页查询我的书架
// @Description 当前登录用户书架
// @Tags 我的书架
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/bookshelf [get]
func (e BookShelf) GetPage(c *gin.Context) {
	req := dto.NovelShelfQueryReq{}
	s := service.NovelBookshelf{}
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
	list, count, respCode, err := s.GetPage(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert app-加入书架
// @Summary 加入书架
// @Description 重复加入幂等返回
// @Tags 我的书架
// @Accept json
// @Produce json
// @Param body body dto.NovelShelfInsertReq "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/bookshelf [post]
func (e BookShelf) Insert(c *gin.Context) {
	req := dto.NovelShelfInsertReq{}
	s := service.NovelBookshelf{}
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

// Delete app-移除书架
// @Summary 移除书架
// @Description 仅本人可移除
// @Tags 我的书架
// @Accept json
// @Produce json
// @Param id path int true "书架记录编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/bookshelf/{id} [delete]
func (e BookShelf) Delete(c *gin.Context) {
	req := dto.NovelShelfDeleteReq{}
	s := service.NovelBookshelf{}
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
	respCode, err := s.Delete(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
