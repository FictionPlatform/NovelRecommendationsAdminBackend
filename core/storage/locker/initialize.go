/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package locker

import (
	"encoding/json"
	"fmt"
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage"
)

// Handler 分布式锁热更新处理器：
// Build 构建新锁（暂存，不触碰线上），成功后才 Apply 原子切换；
// 配置未变化时复用现有锁。
type Handler struct {
	staged      storage.AdapterLocker
	pendingJSON string
	appliedJSON string
}

func NewHandler() *Handler {
	return &Handler{}
}

// Build 构建新锁；锁配置被清空时暂存移除标记（Apply 时清除线上锁）
func (h *Handler) Build() error {
	curr, err := json.Marshal(config.LockerConfig)
	if err != nil {
		return err
	}
	if h.appliedJSON == string(curr) {
		return nil
	}
	h.pendingJSON = string(curr)
	h.staged = nil
	if config.LockerConfig.Empty() {
		return nil
	}
	l, err := config.LockerConfig.Setup()
	if err != nil {
		h.pendingJSON = ""
		return fmt.Errorf("locker setup error: %s", err.Error())
	}
	h.staged = l
	return nil
}

// Apply 切换为新锁（staged 为 nil 表示移除锁配置）
func (h *Handler) Apply() {
	if h.pendingJSON == "" {
		return
	}
	runtime.RuntimeConfig.SetLockerAdapter(h.staged)
	// 切换成功后关闭被替换的旧 Redis 客户端
	config.CommitRedisClients()
	h.staged = nil
	h.appliedJSON = h.pendingJSON
	h.pendingJSON = ""
}
