package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPostRouter)
}

// registerPostRouter 注册长文帖子路由
func registerPostRouter(v1 *gin.RouterGroup) {
	api := apis.Post{}
	r := v1.Group("/app/novel").Use(middleware.Auth())
	{
		r.GET("/post/page", api.GetPage)
		r.GET("/post/:id", api.Get)
		r.POST("/post", api.Insert)
		r.DELETE("/post", api.Delete)
		r.POST("/post/:id/interact", api.Interact)
		r.POST("/post/:id/comment", api.AddComment)
		r.DELETE("/post-comment/:id", api.DeleteComment)
	}
}
