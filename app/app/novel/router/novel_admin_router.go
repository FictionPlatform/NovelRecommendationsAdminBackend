package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis/admin"
	"go-admin/core/middleware"
)

var adminRouterCheckRole = make([]func(v1 *gin.RouterGroup), 0)

func init() {
	adminRouterCheckRole = append(adminRouterCheckRole, registerNoticeRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerFeedbackAdminRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelUserRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelPostAdminRouter)
}

// registerNoticeRouter 注册后台系统公告路由（/admin-api/v1/app/novel/notice/**）
func registerNoticeRouter(v1 *gin.RouterGroup) {
	api := admin.Notice{}
	r := v1.Group("/app/novel/notice").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.POST("", api.Insert)
		r.DELETE("", api.Delete)
	}
}

// registerFeedbackAdminRouter 注册后台反馈/投诉管理路由（/admin-api/v1/app/novel/feedback）
func registerFeedbackAdminRouter(v1 *gin.RouterGroup) {
	api := admin.FeedbackAdmin{}
	r := v1.Group("/app/novel/feedback").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.DELETE("", api.Delete)
	}
}

// registerNovelUserRouter 注册后台读者管理路由（/admin-api/v1/app/novel/user/**）
func registerNovelUserRouter(v1 *gin.RouterGroup) {
	api := admin.NovelUser{}
	r := v1.Group("/app/novel/user").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.PUT("/:id/status", api.ChangeStatus)
		r.PUT("/:id/ban-post", api.BanPost)
	}
}

// registerNovelPostAdminRouter 注册后台帖子管理路由（/admin-api/v1/app/novel/post）
func registerNovelPostAdminRouter(v1 *gin.RouterGroup) {
	api := admin.NovelPostAdmin{}
	r := v1.Group("/app/novel/post").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.PUT("/:id/status", api.ChangeStatus)
	}
}
