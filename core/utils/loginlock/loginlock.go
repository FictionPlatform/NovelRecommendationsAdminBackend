package loginlock

import (
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage"
	"strconv"
)

// 缓存 key 前缀与维度分段
const (
	prefix = "loginlock"
	// fail 计数键：fail:user:{account} / fail:ip:{ip}，TTL 与对应锁定时长一致，自然衰减
	// lock 锁定键：lock:user:{account} / lock:ip:{ip}，存在即锁定
	keyUser = "user:"
	keyIP   = "ip:"
)

// Check 检查账号/IP 是否处于锁定状态。返回是否锁定；缓存/配置不可用时视为未锁定。
func Check(account, ip string) (locked bool) {
	if !enabled() {
		return false
	}
	c := getCache()
	if c == nil {
		return false
	}
	if account != "" {
		if v, err := c.Get(prefix, keyUser+"lock:"+account); err == nil && v != "" {
			return true
		}
	}
	if ip != "" {
		if v, err := c.Get(prefix, keyIP+"lock:"+ip); err == nil && v != "" {
			return true
		}
	}
	return false
}

// RecordFail 记录一次登录失败：账号/IP 计数 +1，达到阈值时转为锁定并清零计数。
// 计数读写非原子（缓存适配器无 INCR），并发窗口内可能多放行少量尝试，可接受。
func RecordFail(account, ip string) {
	cfg := config.LoginLockConfig
	if cfg == nil || !cfg.Enabled {
		return
	}
	c := getCache()
	if c == nil {
		return
	}
	if account != "" && cfg.MaxFailsByAccount > 0 {
		incr(c, keyUser, account, cfg.MaxFailsByAccount, cfg.AccountLockMinutes)
	}
	if ip != "" && cfg.MaxFailsByIP > 0 {
		incr(c, keyIP, ip, cfg.MaxFailsByIP, cfg.IPLockMinutes)
	}
}

// Clear 登录成功后清除账号/IP 的失败计数与锁定状态
func Clear(account, ip string) {
	if !enabled() {
		return
	}
	c := getCache()
	if c == nil {
		return
	}
	if account != "" {
		_ = c.Del(prefix, keyUser+"fail:"+account)
		_ = c.Del(prefix, keyUser+"lock:"+account)
	}
	if ip != "" {
		_ = c.Del(prefix, keyIP+"fail:"+ip)
		_ = c.Del(prefix, keyIP+"lock:"+ip)
	}
}

// incr 维度计数：计数达到 maxFails 则写锁定键（TTL=lockMinutes），否则计数 +1（TTL=lockMinutes）
func incr(c storage.AdapterCache, dim, id string, maxFails, lockMinutes int) {
	lockTTL := lockMinutes * 60
	if lockTTL <= 0 {
		lockTTL = 60
	}
	failKey := dim + "fail:" + id
	count := 1
	if v, err := c.Get(prefix, failKey); err == nil {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			count = n + 1
		}
	}
	if count >= maxFails {
		_ = c.Set(prefix, dim+"lock:"+id, "1", lockTTL)
		_ = c.Del(prefix, failKey)
		return
	}
	_ = c.Set(prefix, failKey, count, lockTTL)
}

func enabled() bool {
	return config.LoginLockConfig != nil && config.LoginLockConfig.Enabled
}

func getCache() storage.AdapterCache {
	return runtime.RuntimeConfig.GetCacheAdapter()
}
