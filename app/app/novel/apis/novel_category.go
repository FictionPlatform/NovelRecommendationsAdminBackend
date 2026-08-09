package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/service"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
)

type NovelCategory struct {
	api.Api
}

// List app-查询分类列表（含标签）
// @Summary 查询分类列表（含标签）
// @Description 返回启用中的小说一级分类及每个分类下的标签，供书库筛选与发书表单使用
// @Tags 小说分类
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/category/list [get]
func (e NovelCategory) List(c *gin.Context) {
	s := service.NovelCategoryReader{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	list, respCode, err := s.List()
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(list, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
