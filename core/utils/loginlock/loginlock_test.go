package loginlock

import (
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/storage/cache"
	"testing"
)

func setup(t *testing.T) {
	t.Helper()
	runtime.RuntimeConfig.SetCacheAdapter(cache.NewMemory())
}

func TestCheckUnlocked(t *testing.T) {
	setup(t)
	if Check("admin", "10.0.0.1") {
		t.Error("fresh account/ip should not be locked")
	}
}

func TestAccountLockAfterMaxFails(t *testing.T) {
	setup(t)
	cfg := &config.LoginLock{Enabled: true, MaxFailsByAccount: 3, AccountLockMinutes: 5, MaxFailsByIP: 10, IPLockMinutes: 5}
	old := config.LoginLockConfig
	config.LoginLockConfig = cfg
	defer func() { config.LoginLockConfig = old }()

	for i := 0; i < 2; i++ {
		RecordFail("admin", "10.0.0.1")
		if Check("admin", "10.0.0.1") {
			t.Fatalf("locked too early at fail %d", i+1)
		}
	}
	RecordFail("admin", "10.0.0.1")
	if !Check("admin", "10.0.0.1") {
		t.Error("account should be locked after max fails")
	}
}

func TestClearAfterSuccess(t *testing.T) {
	setup(t)
	cfg := &config.LoginLock{Enabled: true, MaxFailsByAccount: 3, AccountLockMinutes: 5, MaxFailsByIP: 10, IPLockMinutes: 5}
	old := config.LoginLockConfig
	config.LoginLockConfig = cfg
	defer func() { config.LoginLockConfig = old }()

	RecordFail("admin", "10.0.0.1")
	Clear("admin", "10.0.0.1")
	for i := 0; i < 2; i++ {
		RecordFail("admin", "10.0.0.1")
		if Check("admin", "10.0.0.1") {
			t.Fatalf("clear should reset counters, locked at fail %d", i+1)
		}
	}
	RecordFail("admin", "10.0.0.1")
	if !Check("admin", "10.0.0.1") {
		t.Error("should lock again after 2 more fails")
	}
}

func TestIPLockIndependent(t *testing.T) {
	setup(t)
	cfg := &config.LoginLock{Enabled: true, MaxFailsByAccount: 100, AccountLockMinutes: 5, MaxFailsByIP: 3, IPLockMinutes: 5}
	old := config.LoginLockConfig
	config.LoginLockConfig = cfg
	defer func() { config.LoginLockConfig = old }()

	RecordFail("a", "10.0.0.2")
	RecordFail("b", "10.0.0.2")
	if Check("a", "10.0.0.2") {
		t.Error("account a should not be locked, only ip")
	}
	RecordFail("c", "10.0.0.2")
	if !Check("c", "10.0.0.2") {
		t.Error("ip should be locked after 3 fails")
	}
}

func TestDisabledNoOp(t *testing.T) {
	setup(t)
	cfg := &config.LoginLock{Enabled: false}
	old := config.LoginLockConfig
	config.LoginLockConfig = cfg
	defer func() { config.LoginLockConfig = old }()

	for i := 0; i < 10; i++ {
		RecordFail("admin", "10.0.0.3")
	}
	if Check("admin", "10.0.0.3") {
		t.Error("disabled config must not lock")
	}
}
