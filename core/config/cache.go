package config

import (
	"go-admin/core/utils/storage"
	"go-admin/core/utils/storage/cache"
)

type Cache struct {
	Expired int
	Redis   *RedisConnectOptions
}

// CacheConfig cache配置
var CacheConfig = new(Cache)

// Setup 构造cache 顺序 redis > 其他 > memory
func (e Cache) Setup() (storage.AdapterCache, error) {
	if e.Redis != nil {
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		// 独立客户端：cache 使用自身 Redis 配置（原实现复用首个已建客户端，其余组件的 Redis 配置被静默忽略）
		r, err := cache.NewRedis(nil, options)
		if err != nil {
			return nil, err
		}
		StageRedisClient("cache", r.GetClient())
		_redis = r.GetClient()
		return r, nil
	}
	return cache.NewMemory(), nil
}
