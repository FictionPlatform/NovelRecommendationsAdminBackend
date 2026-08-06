package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerAuthRouter)
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
