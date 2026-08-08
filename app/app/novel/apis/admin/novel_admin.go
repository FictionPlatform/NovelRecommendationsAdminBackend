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

type Notice struct {
	api.Api
}

// GetPage 后台分页查询系统公告
// @Summary 后台分页查询系统公告
// @Tags 小说平台公告
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param keyword query string false "标题关键字"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/notice [get]
func (e Notice) GetPage(c *gin.Context) {
	req := dto.NovelNoticeQueryReq{}
	s := service.NovelNotice{}
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

// Insert 后台发布系统公告（广播到全部读者）
// @Summary 后台发布系统公告
// @Description 发布后广播到全部注册读者，读者端生成系统通知
// @Tags 小说平台公告
// @Accept json
// @Produce json
// @Param body body dto.NovelNoticeInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/notice [post]
func (e Notice) Insert(c *gin.Context) {
	req := dto.NovelNoticeInsertReq{}
	s := service.NovelNotice{}
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

// Delete 后台删除系统公告（级联删除读者通知）
// @Summary 后台删除系统公告
// @Description 删除公告并级联删除广播生成的读者通知
// @Tags 小说平台公告
// @Accept json
// @Produce json
// @Param body body dto.NovelNoticeDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/notice [delete]
func (e Notice) Delete(c *gin.Context) {
	req := dto.NovelNoticeDeleteReq{}
	s := service.NovelNotice{}
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

type FeedbackAdmin struct {
	api.Api
}

// GetPage 后台分页查询反馈/投诉
// @Summary 后台分页查询反馈/投诉
// @Tags 小说平台反馈
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param kind query string false "feedback|complaint"
// @Param type query string false "反馈类型 A~H"
// @Param keyword query string false "内容关键字"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/feedback [get]
func (e FeedbackAdmin) GetPage(c *gin.Context) {
	req := dto.NovelFeedbackPageReq{}
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
	list, count, respCode, err := s.GetAdminPage(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Delete 后台删除反馈/投诉
// @Summary 后台删除反馈/投诉
// @Tags 小说平台反馈
// @Accept json
// @Produce json
// @Param body body dto.NovelFeedbackDeleteReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/feedback [delete]
func (e FeedbackAdmin) Delete(c *gin.Context) {
	req := dto.NovelFeedbackDeleteReq{}
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
	respCode, err := s.Delete(req.Ids)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

type NovelUser struct {
	api.Api
}

// GetPage 后台分页查询读者
// @Summary 后台分页查询读者
// @Tags 小说平台读者
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param keyword query string false "用户名关键字"
// @Param status query string false "状态 1-正常 2-异常"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/user [get]
func (e NovelUser) GetPage(c *gin.Context) {
	req := dto.NovelUserQueryReq{}
	s := service.NovelUser{}
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

// ChangeStatus 后台启用/禁用读者账户
// @Summary 后台启用/禁用读者账户
// @Description 禁用（status=2）后该读者全部接口实时拒绝访问（token 立即失效）
// @Tags 小说平台读者
// @Accept json
// @Produce json
// @Param id path int true "读者用户编号"
// @Param body body dto.NovelUserStatusReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/user/{id}/status [put]
func (e NovelUser) ChangeStatus(c *gin.Context) {
	req := dto.NovelUserStatusReq{}
	s := service.NovelUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	respCode, err := s.ChangeStatus(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// BanPost 后台禁止/解除禁止读者发帖
// @Summary 后台禁止/解除禁止读者发帖
// @Description banUntil 为空表示解除禁止；禁止期间发帖接口返回 41043
// @Tags 小说平台读者
// @Accept json
// @Produce json
// @Param id path int true "读者用户编号"
// @Param body body dto.NovelBanPostReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/user/{id}/ban-post [put]
func (e NovelUser) BanPost(c *gin.Context) {
	req := dto.NovelBanPostReq{}
	s := service.NovelUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	respCode, err := s.BanPost(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

type NovelPostAdmin struct {
	api.Api
}

// GetPage 后台分页查询话题
// @Summary 后台分页查询话题
// @Tags 小说平台话题
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param keyword query string false "标题关键字"
// @Param status query string false "状态 1-正常 2-禁止访问"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post [get]
func (e NovelPostAdmin) GetPage(c *gin.Context) {
	req := dto.NovelPostAdminQueryReq{}
	s := service.NovelPostAdmin{}
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

// ChangeStatus 后台禁止访问/恢复话题
// @Summary 后台禁止访问/恢复话题
// @Description 禁止（status=2）后读者端列表与详情均不可见
// @Tags 小说平台话题
// @Accept json
// @Produce json
// @Param id path int true "话题编号"
// @Param body body dto.NovelPostAdminStatusReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/novel/post/{id}/status [put]
func (e NovelPostAdmin) ChangeStatus(c *gin.Context) {
	req := dto.NovelPostAdminStatusReq{}
	s := service.NovelPostAdmin{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	respCode, err := s.ChangeStatus(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}
