package reglimit

import (
	"strconv"
	"time"

	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage"
)

// 缓存 key 前缀与维度分段
const (
	prefix = "reglimit"
	// 计数键：ip:{ip}，TTL 至当日 24:00 自然衰减，跨天自动清零
	keyIP = "ip:"
)

// Check 检查 IP 当日注册次数是否已达上限。返回是否已达上限；缓存/配置不可用时视为未达上限。
func Check(ip string) bool {
	cfg := config.RegLimitConfig
	if cfg == nil || !cfg.Enabled || cfg.MaxRegPerIP <= 0 {
		return false
	}
	c := getCache()
	if c == nil {
		return false
	}
	if ip == "" {
		return false
	}
	v, err := c.Get(prefix, keyIP+ip)
	if err != nil {
		return false
	}
	n, e := strconv.Atoi(v)
	return e == nil && n >= cfg.MaxRegPerIP
}

// Record 记录一次注册成功：IP 当日计数 +1（TTL 刷新至当日 24:00）。
// 计数读写非原子（缓存适配器无 INCR），并发窗口内可能多放行少量注册，可接受。
func Record(ip string) {
	cfg := config.RegLimitConfig
	if cfg == nil || !cfg.Enabled || cfg.MaxRegPerIP <= 0 {
		return
	}
	c := getCache()
	if c == nil || ip == "" {
		return
	}
	key := keyIP + ip
	count := 1
	if v, err := c.Get(prefix, key); err == nil {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			count = n + 1
		}
	}
	_ = c.Set(prefix, key, count, ttlToMidnight())
}

// ttlToMidnight 返回距当日 24:00 的剩余秒数
func ttlToMidnight() int {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1)
	seconds := int(midnight.Sub(now).Seconds())
	if seconds < 1 {
		seconds = 1
	}
	return seconds
}

func getCache() storage.AdapterCache {
	return runtime.RuntimeConfig.GetCacheAdapter()
}
