package config

import (
	"github.com/redis/go-redis/v9"
	"go-admin/core/utils/storage"
	"go-admin/core/utils/storage/locker"
)

var LockerConfig = new(Locker)

type Locker struct {
	Redis *RedisConnectOptions
}

// Empty 空设置
func (e Locker) Empty() bool {
	return e.Redis == nil
}

// Setup 启用顺序 redis > 其他 > memory
func (e Locker) Setup() (storage.AdapterLocker, error) {
	if e.Redis != nil {
		// 独立客户端：locker 使用自身 Redis 配置（原实现复用首个已建客户端，自身配置被静默忽略）
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		client := StageRedisClient("locker", redis.NewClient(options))
		return locker.NewRedis(client), nil
	}
	return nil, nil
}
