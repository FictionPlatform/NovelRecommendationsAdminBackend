package router

import (
	"go-admin/app/app/novel/apis"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
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
		r.GET("/fans", api.Fans)
		r.GET("/profile/:userId", api.Profile)
	}
}
