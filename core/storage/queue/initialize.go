/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package queue

import (
	"encoding/json"
	"fmt"
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage"
)

// Handler 队列热更新处理器：
// Build 构建新队列（暂存，不触碰线上），成功后才 Apply 切换；
// 配置未变化时复用现有队列，避免热更新无关配置导致消费流中断。
type Handler struct {
	staged      storage.AdapterQueue
	pendingJSON string
	appliedJSON string
}

func NewHandler() *Handler {
	return &Handler{}
}

// Build 构建新队列；队列配置被清空时暂存移除标记（Apply 时关停旧队列）
func (h *Handler) Build() error {
	curr, err := json.Marshal(config.QueueConfig)
	if err != nil {
		return err
	}
	if h.appliedJSON == string(curr) {
		return nil
	}
	h.pendingJSON = string(curr)
	h.staged = nil
	if config.QueueConfig.Empty() {
		return nil
	}
	q, err := config.QueueConfig.Setup()
	if err != nil {
		h.pendingJSON = ""
		return fmt.Errorf("queue setup error: %s", err.Error())
	}
	h.staged = q
	return nil
}

// Apply 关停旧队列后切换为新队列（staged 为 nil 表示移除队列配置）
func (h *Handler) Apply() {
	if h.pendingJSON == "" {
		return
	}
	if old := runtime.RuntimeConfig.GetQueueAdapter(); old != nil {
		old.Shutdown()
	}
	if h.staged != nil {
		runtime.RuntimeConfig.SetQueueAdapter(h.staged)
		go h.staged.Run()
	} else {
		runtime.RuntimeConfig.SetQueueAdapter(nil)
	}
	// 切换成功后关闭被替换的旧 Redis 客户端
	config.CommitRedisClients()
	h.staged = nil
	h.appliedJSON = h.pendingJSON
	h.pendingJSON = ""
}
