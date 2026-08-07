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

type Feedback struct {
	api.Api
}

// Insert app-提交意见反馈/投诉
// @Summary 提交意见反馈/投诉
// @Description type=A~H；kind=feedback 反馈 / complaint 投诉；提交成功自动生成系统通知
// @Tags 意见反馈
// @Accept json
// @Produce json
// @Param body body dto.NovelFeedbackInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/feedback [post]
func (e Feedback) Insert(c *gin.Context) {
	req := dto.NovelFeedbackInsertReq{}
	s := service.NovelFeedback{}
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

// GetPage app-分页查询我的反馈/投诉
// @Summary 分页查询我的反馈/投诉
// @Description kind=feedback 我的反馈 / complaint 我的投诉，为空查全部
// @Tags 意见反馈
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param kind query string false "feedback|complaint"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/feedback/page [get]
func (e Feedback) GetPage(c *gin.Context) {
	req := dto.NovelFeedbackQueryReq{}
	s := service.NovelFeedback{}
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
