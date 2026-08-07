package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFeedbackRouter)
}

// registerFeedbackRouter 注册意见反馈/投诉路由（需登录）
func registerFeedbackRouter(v1 *gin.RouterGroup) {
	api := apis.Feedback{}
	r := v1.Group("/app/novel").Use(middleware.Auth())
	{
		r.POST("/feedback", api.Insert)
		r.GET("/feedback/page", api.GetPage)
	}
}
