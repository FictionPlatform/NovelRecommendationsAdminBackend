package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerHomeRouter)
	routerCheckRole = append(routerCheckRole, registerBookRouter)
}

// registerHomeRouter 注册首页聚合（无需登录）
func registerHomeRouter(v1 *gin.RouterGroup) {
	api := apis.Book{}
	v1.GET("/app/novel/home", api.Home)
}

// registerBookRouter 注册小说书库路由
func registerBookRouter(v1 *gin.RouterGroup) {
	api := apis.Book{}
	// 读者浏览（仅需登录）
	reader := v1.Group("/app/novel/book").Use(middleware.Auth())
	{
		reader.GET("", api.GetPage)
		reader.GET("/rank", api.Rank)
		reader.GET("/:id", api.Get)
	}
	// 内容管理（仅需登录，所有注册读者可发布书籍）
	manager := v1.Group("/app/novel/book").Use(middleware.Auth())
	{
		manager.POST("", api.Insert)
		manager.PUT("/:id", api.Update)
		manager.DELETE("", api.Delete)
	}
}
