package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
	"go-admin/core/config"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/log"
)

var blacklistIPs = make([]*net.IPNet, 0)

// RateLimiter 基于 ulule/limiter 的 IP 访问速率限制中间件
func RateLimiter() gin.HandlerFunc {
	if !config.RateLimiterConfig.Enabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	rate := limiter.Rate{
		Period: time.Duration(config.RateLimiterConfig.Period) * time.Second,
		Limit:  config.RateLimiterConfig.Limit,
	}
	if rate.Period <= 0 || rate.Limit <= 0 {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	var store limiter.Store
	var err error
	if config.RateLimiterConfig.Redis != nil && config.GetRedisClient() != nil {
		store, err = redisStore.NewStore(config.GetRedisClient())
		if err != nil {
			log.Fatalf("初始化限流 redis store 失败: %v", err)
		}
	} else {
		store = memory.NewStore()
	}

	instance := limiter.New(store, rate)
	return ginmiddleware.NewMiddleware(instance,
		ginmiddleware.WithKeyGetter(func(c *gin.Context) string {
			return iputils.GetClientIP(c)
		}),
		ginmiddleware.WithLimitReachedHandler(func(c *gin.Context) {
			c.Header("Retry-After", "60")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "请求过于频繁，请稍后再试",
			})
		}),
		ginmiddleware.WithErrorHandler(func(c *gin.Context, err error) {
			log.Errorf("limiter error: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "限流服务异常",
			})
		}),
	)
}

// IPBlacklist IP 黑名单中间件，支持 IP 与 CIDR 网段
func IPBlacklist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := iputils.GetClientIP(c)
		if ip == "" {
			c.Next()
			return
		}
		clientIP := net.ParseIP(ip)
		for _, blackIP := range blacklistIPs {
			if blackIP.Contains(clientIP) {
				log.Warnf("IP %s 命中黑名单，已拒绝访问 %s", ip, c.Request.URL.Path)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"code": http.StatusForbidden,
					"msg":  "禁止访问",
				})
				return
			}
		}
		c.Next()
	}
}

// LoadBlacklist 从配置加载 IP 黑名单
func LoadBlacklist(ips []string) {
	blacklistIPs = make([]*net.IPNet, 0)
	for _, v := range ips {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		// 支持 CIDR 网段: 1.2.3.0/24
		if strings.Contains(v, "/") {
			_, ipNet, err := net.ParseCIDR(v)
			if err != nil {
				log.Errorf("解析黑名单网段 %s 失败: %v", v, err)
				continue
			}
			blacklistIPs = append(blacklistIPs, ipNet)
			continue
		}
		ip := net.ParseIP(v)
		if ip == nil {
			log.Errorf("解析黑名单 IP %s 失败", v)
			continue
		}
		maskBits := 32
		if ip.To4() == nil {
			maskBits = 128
		}
		blacklistIPs = append(blacklistIPs, &net.IPNet{IP: ip, Mask: net.CIDRMask(maskBits, maskBits)})
	}
}
