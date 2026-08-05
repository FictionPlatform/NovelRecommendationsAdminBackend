package router

import (
	"go-admin/core/runtime"

	"github.com/gin-gonic/gin"
)

// NovelRouteRootPath 小说推荐平台独立前缀（与管理后台 /admin-api 分离，互不干扰）
const NovelRouteRootPath = "/web-api"

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

// InitRouter 初始化小说平台路由
func InitRouter() {
	var r *gin.Engine
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		panic("not found engine...")
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		panic("not support other engine")
	}

	// 无需登录（首页聚合）
	noCheckRoleRouter(r)
	// 需登录（读者操作/内容管理）
	checkRoleRouter(r)
}

// noCheckRoleRouter 无需登录路由
func noCheckRoleRouter(r *gin.Engine) {
	v1 := r.Group(NovelRouteRootPath + "/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// checkRoleRouter 需登录路由
func checkRoleRouter(r *gin.Engine) {
	v1 := r.Group(NovelRouteRootPath + "/v1")
	for _, f := range routerCheckRole {
		f(v1)
	}
}
