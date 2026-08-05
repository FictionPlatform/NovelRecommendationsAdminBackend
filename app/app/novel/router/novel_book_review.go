package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerBookReviewRouter)
}

// registerBookReviewRouter 注册书评路由
func registerBookReviewRouter(v1 *gin.RouterGroup) {
	api := apis.BookReview{}
	r := v1.Group("/app/novel/book-review").Use(middleware.Auth())
	{
		r.GET("/page", api.GetPage)
		r.POST("", api.Insert)
		r.DELETE("", api.Delete)
	}
}
