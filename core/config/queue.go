package config

import (
	"github.com/redis/go-redis/v9"
	"go-admin/core/utils/storage"
	"go-admin/core/utils/storage/queue"
	"go-admin/core/utils/storage/queue/redisqueue"
	"time"
)

type Queue struct {
	Redis  *QueueRedis
	Memory *QueueMemory
}

type QueueRedis struct {
	RedisConnectOptions
	Producer *redisqueue.ProducerOptions
	Consumer *redisqueue.ConsumerOptions
}

type QueueMemory struct {
	PoolSize uint
}

var QueueConfig = new(Queue)

// Empty 空设置
func (e Queue) Empty() bool {
	return e.Memory == nil && e.Redis == nil
}

// Setup 启用顺序 redis > 其他 > memory
// 注意：不得修改配置对象 e 本身（如把时长字段乘以 time.Second 或把 RedisClient 注入回
// Producer/Consumer），否则配置热更新重复 Setup 会累积污染（时长平方级增长、客户端残留）。
func (e Queue) Setup() (storage.AdapterQueue, error) {
	if e.Redis != nil {
		// 独立客户端：queue 使用自身 Redis 配置（原实现复用首个已建客户端，自身配置被静默忽略）
		options, err := e.Redis.RedisConnectOptions.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		client := StageRedisClient("queue", redis.NewClient(options))
		producer := e.Redis.Producer
		if producer == nil {
			producer = &redisqueue.ProducerOptions{}
		}
		consumer := e.Redis.Consumer
		if consumer == nil {
			consumer = &redisqueue.ConsumerOptions{}
		}
		return queue.NewRedis(
			&redisqueue.ProducerOptions{
				StreamMaxLength:      producer.StreamMaxLength,
				ApproximateMaxLength: producer.ApproximateMaxLength,
				RedisClient:          client,
			},
			&redisqueue.ConsumerOptions{
				Name:              consumer.Name,
				GroupName:         consumer.GroupName,
				VisibilityTimeout: time.Duration(consumer.VisibilityTimeout) * time.Second,
				BlockingTimeout:   time.Duration(consumer.BlockingTimeout) * time.Second,
				ReclaimInterval:   time.Duration(consumer.ReclaimInterval) * time.Second,
				BufferSize:        consumer.BufferSize,
				Concurrency:       consumer.Concurrency,
				RedisClient:       client,
			},
		)
	}
	poolSize := uint(0)
	if e.Memory != nil {
		poolSize = e.Memory.PoolSize
	}
	return queue.NewMemory(poolSize), nil
}
