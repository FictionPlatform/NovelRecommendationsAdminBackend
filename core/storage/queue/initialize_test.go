package queue

import (
	"go-admin/core/config"
	"testing"
)

func TestHandlerChangeDetection(t *testing.T) {
	oldQueue := config.QueueConfig
	defer func() { config.QueueConfig = oldQueue }()

	// 默认配置为空（未配置队列）：首次构建走"移除"路径，Apply 后无线上队列
	h := NewHandler()
	if err := h.Build(); err != nil {
		t.Fatalf("first build error: %v", err)
	}
	if h.pendingJSON == "" {
		t.Fatal("expect pending change on first build")
	}
	if h.staged != nil {
		t.Fatal("expect staged nil when queue config empty")
	}
	h.Apply()
	if h.appliedJSON == "" {
		t.Fatal("expect appliedJSON recorded after apply")
	}

	// 配置未变化：再次构建应跳过
	if err := h.Build(); err != nil {
		t.Fatalf("build error: %v", err)
	}
	if h.pendingJSON != "" {
		t.Fatal("expect no pending change when config unchanged")
	}

	// 配置变化（memory 队列）：必须产生待应用变更
	config.QueueConfig = &config.Queue{
		Memory: &config.QueueMemory{PoolSize: 100},
	}
	if err := h.Build(); err != nil {
		t.Fatalf("build error: %v", err)
	}
	if h.pendingJSON == "" {
		t.Fatal("expect pending change when config changed")
	}
	if h.staged == nil {
		t.Fatal("expect staged memory queue when queue config set")
	}
	h.Apply()
	if h.appliedJSON == "" {
		t.Fatal("expect appliedJSON updated after apply")
	}
}

func TestHandlerRemoval(t *testing.T) {
	oldQueue := config.QueueConfig
	defer func() { config.QueueConfig = oldQueue }()

	config.QueueConfig = &config.Queue{
		Memory: &config.QueueMemory{PoolSize: 100},
	}
	h := NewHandler()
	if err := h.Build(); err != nil {
		t.Fatalf("build error: %v", err)
	}
	h.Apply()

	// 清空队列配置：Apply 后 staged 为 nil（移除路径）
	config.QueueConfig = &config.Queue{}
	if err := h.Build(); err != nil {
		t.Fatalf("build error: %v", err)
	}
	if h.pendingJSON == "" {
		t.Fatal("expect pending change when queue config cleared")
	}
	h.Apply()
	if h.staged != nil {
		t.Fatal("expect no staged queue after removal apply")
	}
}
