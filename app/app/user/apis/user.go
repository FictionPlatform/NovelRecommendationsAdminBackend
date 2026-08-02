package apis

import (
	"github.com/gin-gonic/gin"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/app/user/service"
	"go-admin/app/app/user/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/dateutils"
	"time"
)

type User struct {
	api.Api
}

// GetPage app-获取用户管理分页列表
// @Summary 获取用户管理分页列表
// @Description 获取用户管理分页列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param pageIndex query int false "页码"
// @Param pageSize query int false "每页条数"
// @Param levelId query int false "用户等级编号"
// @Param userName query string false "用户昵称"
// @Param trueName query string false "真实姓名"
// @Param email query string false "电子邮箱"
// @Param mobileTitle query string false "国家区号"
// @Param mobile query string false "手机号码"
// @Param parentId query int false "父级编号"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user [get]
func (e User) GetPage(c *gin.Context) {
	req := dto.UserQueryReq{}
	s := service.User{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	req.ShowInfo = false
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	result, _, _ := s.GetSummaries(&req, p)
	e.PageOK(list, result, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Get app-获取用户管理详情
// @Summary 获取用户管理详情
// @Description 获取用户管理详情
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户编号"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user/{id} [get]
func (e User) Get(c *gin.Context) {
	req := dto.UserGetReq{}
	s := service.User{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Insert app-新增用户管理
// @Summary 新增用户管理
// @Description 新增用户管理
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param body body dto.UserInsertReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user [post]
func (e User) Insert(c *gin.Context) {
	req := dto.UserInsertReq{}
	s := service.User{}
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
	respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(baseLang.SuccessCode, e.Lang))
}

// Update app-更新用户管理
// @Summary 更新用户管理
// @Description 更新用户管理
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户编号"
// @Param body body dto.UserUpdateReq true "请求参数"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user/{id} [put]
func (e User) Update(c *gin.Context) {
	req := dto.UserUpdateReq{}
	s := service.User{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.Update(&req, p)
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

// Export app-导出用户管理
// @Summary 导出用户管理
// @Description 导出用户管理
// @Tags 用户管理
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param levelId query int false "用户等级编号"
// @Param userName query string false "用户昵称"
// @Param status query string false "状态"
// @Security Bearer
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求失败"
// @Router /app/user/user/export [get]
func (e User) Export(c *gin.Context) {
	req := dto.UserQueryReq{}
	s := service.User{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(baseLang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, baseLang.DataDecodeCode, baseLang.DataDecodeLogCode, err).Error())
		return
	}

	sysConfService := adminService.NewSysConfigService(&s.Service)
	maxSize, respCode, err := sysConfService.GetWithKeyInt("admin_sys_max_export_size")
	if err != nil || maxSize <= 0 {
		//配置缺失/非法时使用兜底值，避免导出数据量异常
		maxSize = constant.DefaultExportMaxSize
	}
	if maxSize > constant.ExportMaxSizeLimit {
		maxSize = constant.ExportMaxSizeLimit
	}
	p := middleware.GetPermissionFromContext(c)
	req.PageIndex = 1
	req.PageSize = maxSize
	req.PageSizeLimit = maxSize
	req.ShowInfo = true
	list, _, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	data, _ := s.Export(list)
	fileName := "user_" + dateutils.ConvertToStr(time.Now(), 3) + ".xlsx"
	e.DownloadExcel(fileName, data)
}
