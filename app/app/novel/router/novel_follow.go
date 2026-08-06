package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFollowRouter)
}

// registerFollowRouter 注册书友关注路由
func registerFollowRouter(v1 *gin.RouterGroup) {
	api := apis.Follow{}
	r := v1.Group("/app/novel").Use(middleware.Auth())
	{
		r.POST("/follow", api.Follow)
		r.DELETE("/follow/:targetUserId", api.Unfollow)
		r.GET("/following", api.Following)
		r.GET("/profile/:userId", api.Profile)
	}
}
