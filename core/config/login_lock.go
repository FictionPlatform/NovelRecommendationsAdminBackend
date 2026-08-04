package config

// LoginLock 登录爆破防护配置：账号/IP 双维度失败计数，达到阈值后临时锁定。
// 字段均带 yaml tag（热更新重扫使用 yaml.v3 精确匹配），初始加载走 encoding/json（字段名大小写不敏感，不受影响）。
type LoginLock struct {
	Enabled            bool `yaml:"enabled"`
	MaxFailsByAccount  int  `yaml:"maxFailsByAccount"`
	AccountLockMinutes int  `yaml:"accountLockMinutes"`
	MaxFailsByIP       int  `yaml:"maxFailsByIP"`
	IPLockMinutes      int  `yaml:"ipLockMinutes"`
}

// LoginLockConfig 登录爆破防护配置，缺省开启：账号 5 次失败锁 15 分钟、IP 20 次失败锁 60 分钟
var LoginLockConfig = &LoginLock{
	Enabled:            true,
	MaxFailsByAccount:  5,
	AccountLockMinutes: 15,
	MaxFailsByIP:       20,
	IPLockMinutes:      60,
}
