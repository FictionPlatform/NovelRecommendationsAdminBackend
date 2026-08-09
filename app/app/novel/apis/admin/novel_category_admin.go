package admin

import (
	"go-admin/app/app/novel/service"
	"go-admin/app/app/novel/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"

	"github.com/gin-gonic/gin"
)

type NovelCategoryAdmin struct {
	api.Api
}

// GetPage 后台分页查询分类
// @Summary 后台分页查询分类
// @Tags 小说平台分类
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param name query string false "分类名称"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/category [get]
func (e NovelCategoryAdmin) GetPage(c *gin.Context) {
	req := dto.NovelCategoryQueryReq{}
	s := service.NovelCategoryAdmin{}
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

// Get 后台分类详情
// @Summary 后台分类详情
// @Tags 小说平台分类
// @Accept json
// @Produce json
// @Param id path int true "分类编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/category/{id} [get]
func (e NovelCategoryAdmin) Get(c *gin.Context) {
	req := dto.NovelCategoryGetReq{}
	s := service.NovelCategoryAdmin{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	data, respCode, err := s.Get(req.Id)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(data, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert 后台新增分类
// @Summary 后台新增分类
// @Tags 小说平台分类
// @Accept json
// @Produce json
// @Param body body dto.NovelCategoryInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/category [post]
func (e NovelCategoryAdmin) Insert(c *gin.Context) {
	req := dto.NovelCategoryInsertReq{}
	s := service.NovelCategoryAdmin{}
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
	req.CreateBy = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update 后台更新分类
// @Summary 后台更新分类
// @Tags 小说平台分类
// @Accept json
// @Produce json
// @Param id path int true "分类编号"
// @Param body body dto.NovelCategoryUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/category/{id} [put]
func (e NovelCategoryAdmin) Update(c *gin.Context) {
	req := dto.NovelCategoryUpdateReq{}
	s := service.NovelCategoryAdmin{}
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
	req.UpdateBy = uid
	b, respCode, err := s.Update(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(b, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete 后台删除分类
// @Summary 后台删除分类
// @Description 存在标签或书籍引用时禁止删除
// @Tags 小说平台分类
// @Accept json
// @Produce json
// @Param body body dto.NovelCategoryDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/category [delete]
func (e NovelCategoryAdmin) Delete(c *gin.Context) {
	req := dto.NovelCategoryDeleteReq{}
	s := service.NovelCategoryAdmin{}
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

type NovelTagAdmin struct {
	api.Api
}

// GetPage 后台分页查询标签
// @Summary 后台分页查询标签
// @Tags 小说平台标签
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param categoryId query int false "所属分类id"
// @Param name query string false "标签名称"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/tag [get]
func (e NovelTagAdmin) GetPage(c *gin.Context) {
	req := dto.NovelTagQueryReq{}
	s := service.NovelTagAdmin{}
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

// Get 后台标签详情
// @Summary 后台标签详情
// @Tags 小说平台标签
// @Accept json
// @Produce json
// @Param id path int true "标签编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/tag/{id} [get]
func (e NovelTagAdmin) Get(c *gin.Context) {
	req := dto.NovelTagGetReq{}
	s := service.NovelTagAdmin{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	data, respCode, err := s.Get(req.Id)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(data, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert 后台新增标签
// @Summary 后台新增标签
// @Tags 小说平台标签
// @Accept json
// @Produce json
// @Param body body dto.NovelTagInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/tag [post]
func (e NovelTagAdmin) Insert(c *gin.Context) {
	req := dto.NovelTagInsertReq{}
	s := service.NovelTagAdmin{}
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
	req.CreateBy = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update 后台更新标签
// @Summary 后台更新标签
// @Tags 小说平台标签
// @Accept json
// @Produce json
// @Param id path int true "标签编号"
// @Param body body dto.NovelTagUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/tag/{id} [put]
func (e NovelTagAdmin) Update(c *gin.Context) {
	req := dto.NovelTagUpdateReq{}
	s := service.NovelTagAdmin{}
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
	req.UpdateBy = uid
	b, respCode, err := s.Update(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(b, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete 后台删除标签
// @Summary 后台删除标签
// @Tags 小说平台标签
// @Accept json
// @Produce json
// @Param body body dto.NovelTagDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/tag [delete]
func (e NovelTagAdmin) Delete(c *gin.Context) {
	req := dto.NovelTagDeleteReq{}
	s := service.NovelTagAdmin{}
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
