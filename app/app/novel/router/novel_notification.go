package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerNotificationRouter)
}

// registerNotificationRouter 注册系统通知路由（需登录）
func registerNotificationRouter(v1 *gin.RouterGroup) {
	api := apis.Notification{}
	r := v1.Group("/app/novel").Use(middleware.Auth())
	{
		r.GET("/notification/page", api.GetPage)
		r.GET("/notification/unread-count", api.UnreadCount)
		r.POST("/notification/read", api.Read)
	}
}
