package router

import (
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerHomeRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerBookReadRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerCategoryRouter)
	routerCheckRole = append(routerCheckRole, registerBookWriteRouter)
}

// registerCategoryRouter 注册小说分类路由（读者端浏览，公开；带 token 可选注入身份）
func registerCategoryRouter(v1 *gin.RouterGroup) {
	api := apis.NovelCategory{}
	r := v1.Group("/app/novel/category").Use(middleware.AuthOptional())
	{
		r.GET("/list", api.List)
	}
}

// registerHomeRouter 注册首页聚合（无需登录）
func registerHomeRouter(v1 *gin.RouterGroup) {
	api := apis.Book{}
	v1.GET("/app/novel/home", api.Home)
}

// registerBookReadRouter 注册小说书库浏览路由（公开；带 token 可选注入身份以支持 isCollected 等）
func registerBookReadRouter(v1 *gin.RouterGroup) {
	api := apis.Book{}
	r := v1.Group("/app/novel/book").Use(middleware.AuthOptional())
	{
		r.GET("", api.GetPage)
		r.GET("/rank", api.Rank)
		r.GET("/:id", api.Get)
	}
}

// registerBookWriteRouter 注册小说书库写路由（需登录）
func registerBookWriteRouter(v1 *gin.RouterGroup) {
	api := apis.Book{}
	manager := v1.Group("/app/novel/book").Use(middleware.Auth())
	{
		manager.POST("", api.Insert)
		manager.PUT("/:id", api.Update)
		manager.DELETE("", api.Delete)
	}
}
