package iputils

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go-admin/core/utils/log"
)

const (
	// locationCallTimeout 高德 API 请求超时，防止上游不可达时拖垮请求
	locationCallTimeout = 2 * time.Second
	// locationCacheTTL 成功归属地缓存时长：同一 IP 周期内最多外呼一次
	locationCacheTTL = time.Hour
	// locationFailCacheTTL 失败结果缓存时长：上游故障期间避免每个请求都等待超时
	locationFailCacheTTL = time.Minute
	// locationCacheMaxSize 归属地缓存最大条目数，防内存膨胀
	locationCacheMaxSize = 10000
	// locationAmpKeyPlaceholder settings.yml 中 ampKey 的默认占位符
	locationAmpKeyPlaceholder = "---"
)

// locationClient 带超时的 HTTP 客户端
var locationClient = &http.Client{Timeout: locationCallTimeout}

// locationAPIBase 高德 IP 定位接口地址（测试时可覆盖）
var locationAPIBase = "https://restapi.amap.com/v5/ip"

// locationCache IP 归属地内存缓存
var locationCache = newLocationCache()

type locationCacheEntry struct {
	value     string
	expiredAt time.Time
}

type locationCacheT struct {
	mu    sync.Mutex
	items map[string]locationCacheEntry
}

func newLocationCache() *locationCacheT {
	return &locationCacheT{items: make(map[string]locationCacheEntry)}
}

func (l *locationCacheT) get(key string) (string, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	it, ok := l.items[key]
	if !ok {
		return "", false
	}
	if time.Now().After(it.expiredAt) {
		delete(l.items, key)
		return "", false
	}
	return it.value, true
}

func (l *locationCacheT) set(key, value string, ttl time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.items) >= locationCacheMaxSize {
		now := time.Now()
		for k, it := range l.items {
			if now.After(it.expiredAt) {
				delete(l.items, k)
			}
		}
		// 清理后仍满，放弃缓存本次结果
		if len(l.items) >= locationCacheMaxSize {
			return
		}
	}
	l.items[key] = locationCacheEntry{value: value, expiredAt: time.Now().Add(ttl)}
}

// GetLocation 获取外网ip地址
// 带超时（2s）+ IP 级缓存：上游不可达时最多阻塞 2s，且同一 IP 周期内只外呼一次
func GetLocation(ip, key string) string {
	if ip == "" || ip == "127.0.0.1" || ip == "localhost" {
		return "inner ip"
	}
	// 未配置 key（含占位符）时跳过外呼，直接返回空
	if key == "" || key == locationAmpKeyPlaceholder {
		return ""
	}
	if v, ok := locationCache.get(ip); ok {
		return v
	}
	u := locationAPIBase + "?ip=" + url.QueryEscape(ip) + "&type=4&key=" + url.QueryEscape(key)
	resp, err := locationClient.Get(u)
	if err != nil {
		log.Errorf("restapi.amap.com failed: %s", err)
		locationCache.set(ip, "unknown ip", locationFailCacheTTL)
		return "unknown ip"
	}
	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read restapi.amap.com body failed: %s", err)
		locationCache.set(ip, "unknown ip", locationFailCacheTTL)
		return "unknown ip"
	}
	m := make(map[string]string)
	if err = json.Unmarshal(s, &m); err != nil {
		log.Errorf("Umarshal failed: %s", err)
		locationCache.set(ip, "unknown ip", locationFailCacheTTL)
		return "unknown ip"
	}
	// 高德返回业务错误（status != 1，如 key 无效/配额耗尽）时按失败处理，避免空串定位被长缓存
	if m["status"] != "" && m["status"] != "1" {
		log.Errorf("restapi.amap.com business error: status=%s info=%s", m["status"], m["info"])
		locationCache.set(ip, "unknown ip", locationFailCacheTTL)
		return "unknown ip"
	}
	loc := m["country"] + "-" + m["province"] + "-" + m["city"] + "-" + m["district"] + "-" + m["isp"]
	locationCache.set(ip, loc, locationCacheTTL)
	return loc
}

// GetLocaHost 获取局域网ip地址
func GetLocaHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Errorf("net.Interfaces failed, err: %s", err)
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}

// GetClientIP 获取客户端真实 IP。
// 依赖 gin 的可信代理配置（SetTrustedProxies）：
// - 未配置可信代理时返回直连对端 IP（RemoteAddr），客户端伪造的 X-Forwarded-For 无效；
// - 配置了可信代理（如 nginx）时，返回代理链中最左侧的真实客户端 IP。
func GetClientIP(c *gin.Context) string {
	return c.ClientIP()
}
