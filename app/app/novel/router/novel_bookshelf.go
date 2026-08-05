package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerBookshelfRouter)
}

// registerBookshelfRouter 注册我的书架路由
func registerBookshelfRouter(v1 *gin.RouterGroup) {
	api := apis.BookShelf{}
	r := v1.Group("/app/novel/bookshelf").Use(middleware.Auth())
	{
		r.GET("", api.GetPage)
		r.POST("", api.Insert)
		r.DELETE("/:id", api.Delete)
	}
}
