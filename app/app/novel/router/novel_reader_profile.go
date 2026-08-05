package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerReaderProfileRouter)
}

// registerReaderProfileRouter 注册读者资料路由
func registerReaderProfileRouter(v1 *gin.RouterGroup) {
	api := apis.ReaderProfile{}
	r := v1.Group("/app/novel/reader-profile").Use(middleware.Auth())
	{
		r.GET("", api.Get)
		r.PUT("", api.Update)
	}
}
