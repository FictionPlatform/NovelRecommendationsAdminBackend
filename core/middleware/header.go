package middleware

import (
	"go-admin/core/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// KeepAlive is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func KeepAlive(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// setCorsOrigin 按白名单设置跨域头：仅回显白名单内的 Origin，
// 防止任意站点跨域调用；白名单为空则完全不开放跨域（默认同源）
func setCorsOrigin(c *gin.Context) {
	origin := c.GetHeader("Origin")
	if origin == "" {
		return
	}
	for _, o := range config.ApplicationConfig.CorsOrigins {
		if o == "*" || strings.EqualFold(o, origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			return
		}
	}
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	}
	setCorsOrigin(c)
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-AppType", "application/json")
	c.AbortWithStatus(200)
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(c *gin.Context) {
	setCorsOrigin(c)
	// 禁止页面被 iframe 嵌入，防点击劫持
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-AppType-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy headers
	c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
