package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerBookReviewRouter)
	routerCheckRole = append(routerCheckRole, registerBookReviewWriteRouter)
}

// registerBookReviewRouter 注册书评浏览路由（无需登录）
func registerBookReviewRouter(v1 *gin.RouterGroup) {
	api := apis.BookReview{}
	v1.GET("/app/novel/book-review/page", api.GetPage)
}

// registerBookReviewWriteRouter 注册书评写路由（需登录）
func registerBookReviewWriteRouter(v1 *gin.RouterGroup) {
	api := apis.BookReview{}
	r := v1.Group("/app/novel/book-review").Use(middleware.Auth())
	{
		r.POST("", api.Insert)
		r.DELETE("", api.Delete)
	}
}
