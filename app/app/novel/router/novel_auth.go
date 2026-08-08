package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerAuthRouter)
	routerCheckRole = append(routerCheckRole, registerAuthCancelRouter)
}

// registerAuthRouter 注册读者注册/登录路由（无需登录）
func registerAuthRouter(v1 *gin.RouterGroup) {
	api := apis.Auth{}
	r := v1.Group("/app/novel/user")
	{
		r.POST("/login", api.Login)
		r.POST("/register", api.Register)
	}
}

// registerAuthCancelRouter 注册读者主动注销路由（需登录）
func registerAuthCancelRouter(v1 *gin.RouterGroup) {
	api := apis.Auth{}
	r := v1.Group("/app/novel/user").Use(middleware.Auth())
	{
		r.POST("/cancel", api.CancelAccount)
	}
}
