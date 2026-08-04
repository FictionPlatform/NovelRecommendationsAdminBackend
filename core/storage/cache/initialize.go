/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package cache

import (
	"encoding/json"
	"fmt"
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/captchautils"
	"go-admin/core/utils/storage"
)

// Handler 缓存热更新处理器：
// Build 构建新缓存（暂存，不触碰线上），成功后才 Apply 原子切换；
// 配置未变化时复用现有缓存，避免清空已有会话/验证码数据。
type Handler struct {
	staged      storage.AdapterCache
	pendingJSON string
	appliedJSON string
}

func NewHandler() *Handler {
	return &Handler{}
}

// Build 构建新缓存
func (h *Handler) Build() error {
	curr, err := json.Marshal(config.CacheConfig)
	if err != nil {
		return err
	}
	if h.appliedJSON == string(curr) {
		return nil
	}
	adapter, err := config.CacheConfig.Setup()
	if err != nil {
		return fmt.Errorf("cache setup error: %s", err.Error())
	}
	h.staged = adapter
	h.pendingJSON = string(curr)
	return nil
}

// Apply 切换为新缓存，并同步更新验证码存储
func (h *Handler) Apply() {
	if h.pendingJSON == "" {
		return
	}
	runtime.RuntimeConfig.SetCacheAdapter(h.staged)
	captchautils.SetStore(captchautils.NewCacheStore(h.staged, config.CacheConfig.Expired))
	// 切换成功后关闭被替换的旧 Redis 客户端
	config.CommitRedisClients()
	h.staged = nil
	h.appliedJSON = h.pendingJSON
	h.pendingJSON = ""
}
