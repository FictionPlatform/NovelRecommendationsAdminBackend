package iputils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"go-admin/core/config"
)

// TestMain 初始化日志组件（默认 Level 为空会触发 log.Fatalf）
func TestMain(m *testing.M) {
	config.LoggerConfig.Level = "info"
	config.LoggerConfig.Setup()
	os.Exit(m.Run())
}

func resetLocationState() {
	locationCache.mu.Lock()
	locationCache.items = make(map[string]locationCacheEntry)
	locationCache.mu.Unlock()
}

func TestGetLocationNoKey(t *testing.T) {
	resetLocationState()
	var calls int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&calls, 1)
		_, _ = w.Write([]byte(`{"country":"中国"}`))
	}))
	defer srv.Close()

	oldBase := locationAPIBase
	locationAPIBase = srv.URL
	defer func() { locationAPIBase = oldBase }()

	if got := GetLocation("8.8.8.8", ""); got != "" {
		t.Fatalf("empty key: expect empty, got %q", got)
	}
	if got := GetLocation("8.8.8.8", "---"); got != "" {
		t.Fatalf("placeholder key: expect empty, got %q", got)
	}
	if atomic.LoadInt64(&calls) != 0 {
		t.Fatalf("expect no http call without key, got %d", calls)
	}
}

func TestGetLocationCache(t *testing.T) {
	resetLocationState()
	var calls int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&calls, 1)
		_, _ = w.Write([]byte(`{"country":"中国","province":"北京市","city":"北京市","district":"朝阳区","isp":"电信"}`))
	}))
	defer srv.Close()

	oldBase := locationAPIBase
	locationAPIBase = srv.URL
	defer func() { locationAPIBase = oldBase }()

	// 同一 IP 两次只外呼一次
	got1 := GetLocation("1.2.3.4", "key")
	got2 := GetLocation("1.2.3.4", "key")
	if got1 != "中国-北京市-北京市-朝阳区-电信" {
		t.Fatalf("unexpected location %q", got1)
	}
	if got2 != got1 {
		t.Fatalf("cached value mismatch: %q vs %q", got1, got2)
	}
	if atomic.LoadInt64(&calls) != 1 {
		t.Fatalf("expect 1 http call for same ip, got %d", calls)
	}

	// 不同 IP 再次外呼
	if got := GetLocation("5.6.7.8", "key"); got != got1 {
		t.Fatalf("unexpected location %q", got)
	}
	if atomic.LoadInt64(&calls) != 2 {
		t.Fatalf("expect 2 http calls for 2 ips, got %d", calls)
	}

	// 内网/空 IP 不缓存不外呼
	if got := GetLocation("127.0.0.1", "key"); got != "inner ip" {
		t.Fatalf("inner ip: expect 'inner ip', got %q", got)
	}
	if got := GetLocation("", "key"); got != "inner ip" {
		t.Fatalf("empty ip: expect 'inner ip', got %q", got)
	}
	if atomic.LoadInt64(&calls) != 2 {
		t.Fatalf("inner/empty ip must not call http, got %d", calls)
	}
}

type failRoundTripper struct{ calls int64 }

func (f *failRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	f.calls++
	return nil, errors.New("network unreachable")
}

func TestGetLocationFailureCached(t *testing.T) {
	resetLocationState()
	oldClient := locationClient
	rt := &failRoundTripper{}
	locationClient = &http.Client{Transport: rt, Timeout: locationCallTimeout}
	defer func() { locationClient = oldClient }()

	if got := GetLocation("9.9.9.9", "key"); got != "unknown ip" {
		t.Fatalf("expect 'unknown ip', got %q", got)
	}
	// 失败结果同样缓存，避免每请求重试等待超时
	if got := GetLocation("9.9.9.9", "key"); got != "unknown ip" {
		t.Fatalf("cached failure: expect 'unknown ip', got %q", got)
	}
	if rt.calls != 1 {
		t.Fatalf("expect 1 http attempt (cached after fail), got %d", rt.calls)
	}
}

func TestGetLocationTimeout(t *testing.T) {
	resetLocationState()
	// 服务端故意拖慢响应（3s > 客户端 2s 超时），验证外呼有超时上限
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		_, _ = w.Write([]byte(`{"status":"1","country":"中国"}`))
	}))
	defer srv.Close()

	oldBase := locationAPIBase
	locationAPIBase = srv.URL
	defer func() { locationAPIBase = oldBase }()

	start := time.Now()
	if got := GetLocation("6.6.6.6", "key"); got != "unknown ip" {
		t.Fatalf("timeout: expect 'unknown ip', got %q", got)
	}
	if elapsed := time.Since(start); elapsed > 2500*time.Millisecond {
		t.Fatalf("expect call bounded by 2s timeout, took %v", elapsed)
	}
}

func TestGetLocationBusinessError(t *testing.T) {
	resetLocationState()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":"0","info":"INVALID_USER_KEY"}`))
	}))
	defer srv.Close()

	oldBase := locationAPIBase
	locationAPIBase = srv.URL
	defer func() { locationAPIBase = oldBase }()

	if got := GetLocation("4.4.4.4", "bad-key"); got != "unknown ip" {
		t.Fatalf("business error: expect 'unknown ip', got %q", got)
	}
}
