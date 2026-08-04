package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go-admin/core/config"
)

// TestMain 初始化日志组件（默认 Level 为空会触发 log.Fatalf）
func TestMain(m *testing.M) {
	config.LoggerConfig.Level = "info"
	config.LoggerConfig.Setup()
	os.Exit(m.Run())
}

// resetLimiterStoreForTest 丢弃共享 store 与生效状态，使测试用例相互独立
func resetLimiterStoreForTest() {
	limiterStoreMu.Lock()
	limiterStore = nil
	limiterStoreMu.Unlock()
	ApplyRateLimiterState(BuildRateLimiterState())
}

func TestIPBlacklist(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(IPBlacklist())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	do := func(ip string) int {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.RemoteAddr = ip
		router.ServeHTTP(w, req)
		return w.Code
	}

	// 正常放行
	LoadBlacklist([]string{"192.168.1.1", "10.0.0.0/8", "bad-cidr"})
	if code := do("8.8.8.8:80"); code != http.StatusOK {
		t.Fatalf("normal ip: expect 200, got %d", code)
	}

	// 精确 IP 命中
	if code := do("192.168.1.1:1234"); code != http.StatusForbidden {
		t.Fatalf("blacklisted ip: expect 403, got %d", code)
	}

	// CIDR 命中
	if code := do("10.1.2.3:5678"); code != http.StatusForbidden {
		t.Fatalf("blacklisted cidr: expect 403, got %d", code)
	}

	// 热更新后新黑名单生效、旧黑名单失效
	LoadBlacklist([]string{"10.0.0.0/8"})
	if code := do("192.168.1.1:1234"); code != http.StatusOK {
		t.Fatalf("after reload, old ip: expect 200, got %d", code)
	}
	if code := do("10.1.2.3:5678"); code != http.StatusForbidden {
		t.Fatalf("after reload, cidr: expect 403, got %d", code)
	}
}

func TestRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	old := config.RateLimiterConfig
	config.RateLimiterConfig = &config.RateLimiter{
		Enabled: true,
		Limit:   2,
		Period:  60,
	}
	defer func() { config.RateLimiterConfig = old }()
	resetLimiterStoreForTest()

	ApplyRateLimiterState(BuildRateLimiterState())

	router := gin.New()
	router.Use(RateLimiter())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	do := func(ip string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.RemoteAddr = ip
		router.ServeHTTP(w, req)
		return w
	}

	for i := 0; i < 2; i++ {
		if w := do("1.2.3.4:80"); w.Code != http.StatusOK {
			t.Fatalf("request %d: expect 200, got %d", i+1, w.Code)
		}
	}

	w := do("1.2.3.4:80")
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("over limit: expect 429, got %d", w.Code)
	}
	// Retry-After 与配置周期一致（60s），而非硬编码
	if got := w.Header().Get("Retry-After"); got != "60" {
		t.Fatalf("Retry-After: expect 60, got %q", got)
	}

	// 其他 IP 不受影响
	if w := do("5.6.7.8:80"); w.Code != http.StatusOK {
		t.Fatalf("other ip: expect 200, got %d", w.Code)
	}
}

func TestRateLimiterDisabled(t *testing.T) {
	gin.SetMode(gin.TestMode)

	old := config.RateLimiterConfig
	config.RateLimiterConfig = &config.RateLimiter{Enabled: false}
	defer func() { config.RateLimiterConfig = old }()
	resetLimiterStoreForTest()

	ApplyRateLimiterState(BuildRateLimiterState())

	router := gin.New()
	router.Use(RateLimiter())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.RemoteAddr = "1.2.3.4:80"
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("disabled limiter: expect 200, got %d", w.Code)
	}
}

func TestRateLimiterPeriodScale(t *testing.T) {
	gin.SetMode(gin.TestMode)

	old := config.RateLimiterConfig
	config.RateLimiterConfig = &config.RateLimiter{
		Enabled: true,
		Limit:   1,
		Period:  1,
	}
	defer func() { config.RateLimiterConfig = old }()
	resetLimiterStoreForTest()

	ApplyRateLimiterState(BuildRateLimiterState())

	router := gin.New()
	router.Use(RateLimiter())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	do := func(ip string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.RemoteAddr = ip
		router.ServeHTTP(w, req)
		return w
	}

	if w := do("9.9.9.9:80"); w.Code != http.StatusOK {
		t.Fatalf("first request: expect 200, got %d", w.Code)
	}
	w := do("9.9.9.9:80")
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("over limit: expect 429, got %d", w.Code)
	}
	if got := w.Header().Get("Retry-After"); got != "1" {
		t.Fatalf("Retry-After: expect 1, got %q", got)
	}

	// 周期滚动后恢复（内存 store 基于 wall clock）
	time.Sleep(1100 * time.Millisecond)
	if w := do("9.9.9.9:80"); w.Code != http.StatusOK {
		t.Fatalf("after window reset: expect 200, got %d", w.Code)
	}
}

func TestRateLimiterStoreReusedAcrossReload(t *testing.T) {
	gin.SetMode(gin.TestMode)

	old := config.RateLimiterConfig
	defer func() { config.RateLimiterConfig = old }()
	resetLimiterStoreForTest()

	router := gin.New()
	router.Use(RateLimiter())
	router.GET("/test", func(c *gin.Context) { c.Status(http.StatusOK) })

	do := func(ip string) int {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.RemoteAddr = ip
		router.ServeHTTP(w, req)
		return w.Code
	}

	config.RateLimiterConfig = &config.RateLimiter{Enabled: true, Limit: 1, Period: 60}
	ApplyRateLimiterState(BuildRateLimiterState())
	if code := do("7.7.7.7:80"); code != http.StatusOK {
		t.Fatalf("first request: expect 200, got %d", code)
	}
	if code := do("7.7.7.7:80"); code != http.StatusTooManyRequests {
		t.Fatalf("over limit: expect 429, got %d", code)
	}

	// 模拟热更新（配置变更后重建限流器）：store 复用，存量计数保留，
	// 被限流 IP 重载后仍受限（直至窗口过期），无法通过触发重载绕过限流
	config.RateLimiterConfig = &config.RateLimiter{Enabled: true, Limit: 1, Period: 30}
	ApplyRateLimiterState(BuildRateLimiterState())
	if code := do("7.7.7.7:80"); code != http.StatusTooManyRequests {
		t.Fatalf("after reload, previously limited ip must stay limited, got %d", code)
	}
	if code := do("8.8.8.8:80"); code != http.StatusOK {
		t.Fatalf("fresh ip after reload: expect 200, got %d", code)
	}

	// 同一 store 实例被复用（内存 store 只创建一次，不随重载泄漏清理 goroutine）
	limiterStoreMu.Lock()
	shared := limiterStore
	limiterStoreMu.Unlock()
	if shared == nil {
		t.Fatal("expect shared limiter store to exist")
	}
}
