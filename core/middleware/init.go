package middleware

import (
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"

	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/wprimadi/brandy"
)

const (
	JwtTokenCheck   string = "JwtToken"
	RoleCheck       string = "AuthCheckRole"
	PermissionCheck string = "PermissionAction"
)

func InitMiddleware(r *gin.Engine) {
	// 数据库链接
	r.Use(WithContextDb)
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	// IsKeepAlive is a middleware function that appends headers
	r.Use(KeepAlive)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)

	// 1. 初始化 WAF 引擎并加载规则集
	rulesetPaths := []string{
		"rulesets/default.conf",
		"rulesets/owasp-crs/rules/*.conf",
	}
	waf, err := brandy.InitWaf(rulesetPaths)
	if err != nil {
		log.Fatalf("加载 WAF 规则失败: %v", err)
	}

	// 2. 应用 WAF 中间件
	r.Use(brandy.Waf(waf, ""))
	// IP 黑名单
	LoadBlacklist(config.RateLimiterConfig.Blacklist)
	r.Use(IPBlacklist())
	// IP 访问速率限制
	r.Use(RateLimiter())
	// 链路追踪
	r.Use(Trace())
	runtime.RuntimeConfig.SetMiddleware(JwtTokenCheck, (*jwt.GinJWTMiddleware).MiddlewareFunc)
	runtime.RuntimeConfig.SetMiddleware(RoleCheck, AuthCheckRole())
	runtime.RuntimeConfig.SetMiddleware(PermissionCheck, PermissionAction())
}
