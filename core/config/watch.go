package config

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-admin/core/utils/log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Watch 轮询监听配置文件变更（load-config v1.6.4 未内置文件监听，OnChange 不会被框架触发）。
// 内容变化后重新扫描配置并走 OnChange 受保护热更新：全部组件构建成功才原子切换，任一失败自动回滚。
// 解析失败仅记录日志，线上配置与组件保持旧版。阻塞运行，由调用方以 goroutine 启动。
func Watch(path string, interval time.Duration) {
	if interval <= 0 {
		interval = 3 * time.Second
	}
	last := fileDigest(path)
	for {
		time.Sleep(interval)
		cur := fileDigest(path)
		if cur == "" || cur == last {
			continue
		}
		last = cur
		if err := ReloadFromFile(path); err != nil {
			log.Errorf("!!! config watch reload failed: %s", err.Error())
		}
	}
}

// ReloadFromFile 从文件重新扫描配置到全局配置对象（保持指针身份），并触发受保护热更新。
// 解析失败返回 error，线上配置与组件均不受影响。
func ReloadFromFile(path string) error {
	if _cfg == nil {
		return fmt.Errorf("config not setup yet")
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file error: %w", err)
	}
	return _cfg.scan(raw)
}

// scan 将 YAML 配置原地写入配置对象（yaml.v3 对非 nil 指针复用目标对象，全局配置随之更新），
// 然后执行 OnChange 受保护热更新。
func (e *Settings) scan(raw []byte) error {
	// yaml 对 map 只做增改不删除，先清空多库映射，保证被移除的库在本次重载中生效
	*e.Settings.Databases = map[string]*Database{}
	if err := yaml.Unmarshal(raw, e); err != nil {
		return fmt.Errorf("parse config error: %w", err)
	}
	e.relinkGlobals()
	e.OnChange()
	return nil
}

// fileDigest 返回文件内容 SHA-256；读取失败返回空串（文件暂时不可读时跳过本次轮询）
func fileDigest(path string) string {
	raw, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}
