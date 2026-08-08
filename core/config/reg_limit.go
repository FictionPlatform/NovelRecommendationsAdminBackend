package config

// RegLimit 注册频率限制配置：同一 IP 每天最多注册 maxRegPerIP 次。
// 字段均带 yaml tag（热更新重扫使用 yaml.v3 精确匹配），初始加载走 encoding/json（字段名大小写不敏感，不受影响）。
type RegLimit struct {
	Enabled     bool `yaml:"enabled"`
	MaxRegPerIP int  `yaml:"maxRegPerIP"`
}

// RegLimitConfig 注册频率限制配置，缺省开启：每个 IP 每天最多注册 5 次
var RegLimitConfig = &RegLimit{
	Enabled:     true,
	MaxRegPerIP: 5,
}
