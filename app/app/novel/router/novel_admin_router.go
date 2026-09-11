package router

import (
	"go-admin/app/app/novel/apis/admin"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

var adminRouterCheckRole = make([]func(v1 *gin.RouterGroup), 0)

func init() {
	adminRouterCheckRole = append(adminRouterCheckRole, registerNoticeRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerFeedbackAdminRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelUserRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelPostAdminRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelCategoryAdminRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelTagAdminRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNovelBookAdminRouter)
	adminRouterCheckRole = append(adminRouterCheckRole, registerNotificationAdminRouter)
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

// registerNovelPostAdminRouter 注册后台话题管理路由（/admin-api/v1/app/novel/post）
func registerNovelPostAdminRouter(v1 *gin.RouterGroup) {
	api := admin.NovelPostAdmin{}
	r := v1.Group("/app/novel/post").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.PUT("/:id/status", api.ChangeStatus)
	}
}

// registerNovelCategoryAdminRouter 注册后台分类管理路由（/admin-api/v1/app/novel/category）
func registerNovelCategoryAdminRouter(v1 *gin.RouterGroup) {
	api := admin.NovelCategoryAdmin{}
	r := v1.Group("/app/novel/category").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

// registerNovelTagAdminRouter 注册后台标签管理路由（/admin-api/v1/app/novel/tag）
func registerNovelTagAdminRouter(v1 *gin.RouterGroup) {
	api := admin.NovelTagAdmin{}
	r := v1.Group("/app/novel/tag").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

// registerNovelBookAdminRouter 注册后台书籍管理路由（/admin-api/v1/app/novel/book）
func registerNovelBookAdminRouter(v1 *gin.RouterGroup) {
	api := admin.NovelBookAdmin{}
	r := v1.Group("/app/novel/book").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.POST("/merge", api.Merge)
	}
}

// registerNotificationAdminRouter 注册后台通知管理路由（/admin-api/v1/app/novel/notification）
func registerNotificationAdminRouter(v1 *gin.RouterGroup) {
	api := admin.NotificationAdmin{}
	r := v1.Group("/app/novel/notification").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.POST("", api.Send)
		r.DELETE("", api.Delete)
	}
}
