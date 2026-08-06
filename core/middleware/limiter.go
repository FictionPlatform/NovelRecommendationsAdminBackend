package middleware

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/log"
)

// blacklistIPs 当前生效的 IP 黑名单网段列表。
// 通过 atomic.Pointer 整体替换（先构建后替换），避免热更新重载与每请求遍历之间的数据竞态。
var blacklistIPs atomic.Pointer[[]*net.IPNet]

// LimiterState 限流器运行态快照；配置热更新时整体重建后原子替换，运行期改动即时生效
type LimiterState struct {
	enabled bool
	limiter *limiter.Limiter
	period  time.Duration
}

var currentLimiter atomic.Pointer[LimiterState]

var (
	// limiterStoreMu 串行化限流底层 store 的创建
	limiterStoreMu sync.Mutex
	// limiterStore 复用的限流计数存储（内存/Redis 只创建一次）。
	// ulule 内存 store 每次创建都会启动一个清理 goroutine（仅 GC finalizer 才停止），
	// 热更新重建 limiter 时复用 store 可避免 goroutine 累积，同时保留存量限流计数
	// （被限流的 IP 不会因配置重载而立刻清零，防止通过触发重载绕过限流）。
	limiterStore limiter.Store
)

// getLimiterStore 返回复用的限流计数存储；Redis 可用时优先 Redis，初始化失败降级内存
func getLimiterStore() limiter.Store {
	limiterStoreMu.Lock()
	defer limiterStoreMu.Unlock()
	if limiterStore != nil {
		return limiterStore
	}
	// 独立客户端：限流器使用自身 rateLimiter.redis 配置（原实现依赖 GetRedisClient，
	// 当 cache 未配 Redis 时限流的 Redis 配置被静默忽略）
	if config.RateLimiterConfig.Redis != nil {
		options, err := config.RateLimiterConfig.Redis.GetRedisOptions()
		if err != nil {
			log.Errorf("限流 redis 配置解析失败，降级为内存 store: %v", err)
		} else {
			client := config.StageRedisClient("limiter", redis.NewClient(options))
			store, err := redisStore.NewStore(client)
			if err != nil {
				log.Errorf("初始化限流 redis store 失败，降级为内存 store: %v", err)
			} else {
				limiterStore = store
				return limiterStore
			}
		}
	}
	limiterStore = memory.NewStore()
	return limiterStore
}

// RateLimiter 基于 ulule/limiter 的 IP 访问速率限制中间件。
// 每次请求从原子指针读取最新限流实例，支持配置热更新即时生效，无需重启。
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		st := currentLimiter.Load()
		if st == nil || !st.enabled || st.limiter == nil {
			c.Next()
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()

		lctx, err := st.limiter.Get(ctx, iputils.GetClientIP(c))
		if err != nil {
			log.Errorf("limiter error: %v", err)
			response.ErrorByHttpCode(c, http.StatusInternalServerError, baseLang.RateLimitServerErrCode,
				lang.MsgByCode(baseLang.RateLimitServerErrCode, lang.GetAcceptLanguage(c)))
			return
		}
		if lctx.Reached {
			// Retry-After 使用配置的实际周期（秒），与限流窗口保持一致
			c.Header("Retry-After", strconv.FormatInt(int64(st.period.Seconds()), 10))
			response.ErrorByHttpCode(c, http.StatusTooManyRequests, baseLang.RateLimitErrCode,
				lang.MsgByCode(baseLang.RateLimitErrCode, lang.GetAcceptLanguage(c)))
			return
		}
		c.Next()
	}
}

// IPBlacklist IP 黑名单中间件，支持 IP 与 CIDR 网段
func IPBlacklist() gin.HandlerFunc {
	return func(c *gin.Context) {
		list := blacklistIPs.Load()
		if list == nil || len(*list) == 0 {
			c.Next()
			return
		}
		ip := iputils.GetClientIP(c)
		if ip == "" {
			c.Next()
			return
		}
		clientIP := net.ParseIP(ip)
		for _, blackIP := range *list {
			if blackIP.Contains(clientIP) {
				log.Warnf("IP %s 命中黑名单，已拒绝访问 %s", ip, c.Request.URL.Path)
				response.ErrorByHttpCode(c, http.StatusForbidden, baseLang.IpBlacklistCode,
					lang.MsgByCode(baseLang.IpBlacklistCode, lang.GetAcceptLanguage(c)))
				return
			}
		}
		c.Next()
	}
}

// BuildBlacklist 基于配置解析 IP 黑名单（仅构建暂存，不改变线上状态）
func BuildBlacklist(ips []string) *[]*net.IPNet {
	parsed := make([]*net.IPNet, 0, len(ips))
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
			parsed = append(parsed, ipNet)
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
		parsed = append(parsed, &net.IPNet{IP: ip, Mask: net.CIDRMask(maskBits, maskBits)})
	}
	return &parsed
}

// ApplyBlacklist 原子替换生效中的黑名单列表
func ApplyBlacklist(list *[]*net.IPNet) {
	blacklistIPs.Store(list)
}

// LoadBlacklist 从配置加载 IP 黑名单（先构建后原子替换，与请求遍历无竞态）
func LoadBlacklist(ips []string) {
	ApplyBlacklist(BuildBlacklist(ips))
}

// BuildRateLimiterState 基于当前配置构建限流器实例（仅构建暂存，不改变线上状态）。
// 底层 store 复用全局实例（见 getLimiterStore），Redis store 初始化失败时降级为内存 store。
func BuildRateLimiterState() *LimiterState {
	cfg := config.RateLimiterConfig
	st := &LimiterState{enabled: cfg.Enabled}

	rate := limiter.Rate{
		Period: time.Duration(cfg.Period) * time.Second,
		Limit:  cfg.Limit,
	}
	if !cfg.Enabled || rate.Period <= 0 || rate.Limit <= 0 {
		st.enabled = false
		return st
	}
	st.period = rate.Period
	st.limiter = limiter.New(getLimiterStore(), rate)
	return st
}

// ApplyRateLimiterState 原子替换生效中的限流器实例
func ApplyRateLimiterState(st *LimiterState) {
	currentLimiter.Store(st)
}
