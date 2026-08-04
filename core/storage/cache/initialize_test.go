package cache

import (
	"go-admin/core/config"
	"testing"
)

func TestHandlerChangeDetection(t *testing.T) {
	oldExpired := config.CacheConfig.Expired
	defer func() { config.CacheConfig.Expired = oldExpired }()

	h := NewHandler()

	// 首次构建：必须产生待应用变更
	if err := h.Build(); err != nil {
		t.Fatalf("first build error: %v", err)
	}
	if h.pendingJSON == "" {
		t.Fatal("expect pending change on first build")
	}
	h.Apply()
	if h.appliedJSON == "" {
		t.Fatal("expect appliedJSON recorded after apply")
	}
	if h.pendingJSON != "" {
		t.Fatal("expect pending cleared after apply")
	}

	// 配置未变化：再次构建应跳过（不产生新 adapter）
	if err := h.Build(); err != nil {
		t.Fatalf("build error: %v", err)
	}
	if h.pendingJSON != "" {
		t.Fatal("expect no pending change when config unchanged")
	}

	// 配置变化：必须产生待应用变更
	config.CacheConfig.Expired = oldExpired + 1
	if err := h.Build(); err != nil {
		t.Fatalf("build error: %v", err)
	}
	if h.pendingJSON == "" {
		t.Fatal("expect pending change when config changed")
	}
	h.Apply()
	if h.appliedJSON == "" {
		t.Fatal("expect appliedJSON updated after apply")
	}
}
