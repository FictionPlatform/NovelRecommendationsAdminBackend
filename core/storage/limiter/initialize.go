package limiter

import (
	"encoding/json"
	"go-admin/core/config"
	"go-admin/core/middleware"
	"net"
)

// Handler 限流器/IP黑名单热更新处理器：
// Build 基于当前配置预构建新实例（暂存，不触碰线上），成功后才 Apply 原子替换；
// 配置未变化时跳过重建。
type Handler struct {
	stagedBlacklist *[]*net.IPNet
	stagedState     *middleware.LimiterState
	pendingJSON     string
	appliedJSON     string
}

func NewHandler() *Handler {
	return &Handler{}
}

// Build 构建新限流实例与黑名单；配置与上次生效一致时直接跳过
func (h *Handler) Build() error {
	curr, err := json.Marshal(config.RateLimiterConfig)
	if err != nil {
		return err
	}
	if h.appliedJSON == string(curr) {
		return nil
	}
	h.pendingJSON = string(curr)
	h.stagedBlacklist = middleware.BuildBlacklist(config.RateLimiterConfig.Blacklist)
	h.stagedState = middleware.BuildRateLimiterState()
	return nil
}

// Apply 原子替换线上限流器与黑名单（仅赋值操作，不应失败）
func (h *Handler) Apply() {
	if h.pendingJSON == "" {
		return
	}
	middleware.ApplyBlacklist(h.stagedBlacklist)
	middleware.ApplyRateLimiterState(h.stagedState)
	// 切换成功后关闭被替换的旧 Redis 客户端
	config.CommitRedisClients()
	h.stagedBlacklist = nil
	h.stagedState = nil
	h.appliedJSON = h.pendingJSON
	h.pendingJSON = ""
}
