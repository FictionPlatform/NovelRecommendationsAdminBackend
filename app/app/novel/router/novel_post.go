package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerPostRouter)
	routerCheckRole = append(routerCheckRole, registerPostWriteRouter)
}

// registerPostRouter 注册长文话题浏览路由（无需登录；带 token 时可选注入身份以支持 mine=1/isCollected=1 与互动状态）
func registerPostRouter(v1 *gin.RouterGroup) {
	api := apis.Post{}
	r := v1.Group("/app/novel").Use(middleware.AuthOptional())
	{
		r.GET("/post/page", api.GetPage)
		r.GET("/post/:id", api.Get)
	}
}

// registerPostWriteRouter 注册长文话题写路由（需登录）
func registerPostWriteRouter(v1 *gin.RouterGroup) {
	api := apis.Post{}
	r := v1.Group("/app/novel").Use(middleware.Auth())
	{
		r.POST("/post", api.Insert)
		r.DELETE("/post", api.Delete)
		r.POST("/post/:id/interact", api.Interact)
		r.POST("/post/:id/comment", api.AddComment)
		r.GET("/post-comment/mine", api.MyComments)
		r.DELETE("/post-comment/:id", api.DeleteComment)
	}
}
