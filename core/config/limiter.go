package config

type RateLimiter struct {
	Enabled   bool
	Limit     int64
	Period    int64
	Redis     *RedisConnectOptions
	Blacklist []string
}

var RateLimiterConfig = new(RateLimiter)
