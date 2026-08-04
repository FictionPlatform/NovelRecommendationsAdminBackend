package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/redis/go-redis/v9"
	"go-admin/core/utils/log"
	"os"
	"sync"
)

var (
	// _redis 兼容旧接口的主客户端（默认取 cache 组件的客户端）
	_redis *redis.Client

	// clientMu 保护 clients/deprecated
	clientMu sync.Mutex
	// clients 各组件（cache/locker/queue/limiter）独立的 Redis 客户端。
	// 组件必须使用自身配置的独立客户端——共享同一客户端会导致其余组件的
	// Redis 配置（不同 addr/db/密码）被静默忽略（原实现：首个已建客户端被全员复用）。
	clients = make(map[string]*redis.Client)
	// deprecated 已被替换、等待 Apply 成功切换后关闭的旧客户端
	deprecated []*redis.Client
)

// GetRedisClient 获取兼容入口的主客户端（外部显式设置优先，缺省为 cache 组件的客户端）
func GetRedisClient() *redis.Client {
	clientMu.Lock()
	defer clientMu.Unlock()
	if c := clients["external"]; c != nil {
		return c
	}
	return clients["cache"]
}

// SetRedisClient 设置外部主客户端，替换旧客户端（旧客户端立即关闭）
func SetRedisClient(c *redis.Client) {
	clientMu.Lock()
	defer clientMu.Unlock()
	if old := clients["external"]; old != nil && old != c {
		_ = old.Shutdown(context.Background())
	}
	clients["external"] = c
}

// StageRedisClient 登记组件（cache/locker/queue/limiter）的 Redis 客户端：
// 替换时旧客户端挂起，等待 Apply 成功切换后统一关闭（构建阶段不能关停仍被线上组件使用的客户端）。
func StageRedisClient(id string, c *redis.Client) *redis.Client {
	clientMu.Lock()
	defer clientMu.Unlock()
	if old, ok := clients[id]; ok && old != c {
		deprecated = append(deprecated, old)
	}
	clients[id] = c
	return c
}

// CommitRedisClients Apply 成功切换组件后关闭被替换的旧客户端（各组件 Handler 的 Apply 调用）
func CommitRedisClients() {
	clientMu.Lock()
	defer clientMu.Unlock()
	for _, c := range deprecated {
		_ = c.Shutdown(context.Background())
	}
	deprecated = nil
}

// CloseAllRedisClients 关闭全部登记客户端（进程退出时调用）
func CloseAllRedisClients() {
	clientMu.Lock()
	defer clientMu.Unlock()
	for id, c := range clients {
		_ = c.Shutdown(context.Background())
		delete(clients, id)
	}
	deprecated = nil
	_redis = nil
}

type RedisConnectOptions struct {
	Network    string `yaml:"network" json:"network"`
	Addr       string `yaml:"addr" json:"addr"`
	Username   string `yaml:"username" json:"username"`
	Password   string `yaml:"password" json:"password"`
	DB         int    `yaml:"db" json:"db"`
	PoolSize   int    `yaml:"poolSize" json:"poolSize"`
	Tls        *Tls   `yaml:"tls" json:"tls"`
	MaxRetries int    `yaml:"maxRetries" json:"maxRetries"`
}

type Tls struct {
	Cert string `yaml:"cert" json:"cert"`
	Key  string `yaml:"key" json:"key"`
	Ca   string `yaml:"ca" json:"ca"`
}

func (e RedisConnectOptions) GetRedisOptions() (*redis.Options, error) {
	r := &redis.Options{
		Network:    e.Network,
		Addr:       e.Addr,
		Username:   e.Username,
		Password:   e.Password,
		DB:         e.DB,
		MaxRetries: e.MaxRetries,
		PoolSize:   e.PoolSize,
	}
	var err error
	r.TLSConfig, err = getTLS(e.Tls)
	return r, err
}

func getTLS(c *Tls) (*tls.Config, error) {
	if c != nil && c.Cert != "" {
		// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
		cert, err := tls.LoadX509KeyPair(c.Cert, c.Key)
		if err != nil {
			log.Errorf("tls.LoadX509KeyPair err: %v\n", err)
			return nil, err
		}
		// 创建一个新的、空的 CertPool，并尝试解析 PEM 编码的证书，解析成功会将其加到 CertPool 中
		certPool := x509.NewCertPool()
		ca, err := os.ReadFile(c.Ca)
		if err != nil {
			log.Errorf("ioutil.ReadFile err: %v\n", err)
			return nil, err
		}

		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			log.Error("certPool.AppendCertsFromPEM err")
			return nil, err
		}
		return &tls.Config{
			// 设置证书链，允许包含一个或多个
			Certificates: []tls.Certificate{cert},
			// 要求必须校验客户端的证书
			ClientAuth: tls.RequireAndVerifyClientCert,
			// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
			ClientCAs: certPool,
		}, nil
	}
	return nil, nil
}
