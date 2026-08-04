package config

import (
	"errors"
	"sync"
	"testing"
)

// fakeHandler 模拟热更新处理器：记录 Build/Apply 调用与失败
type fakeHandler struct {
	mu       sync.Mutex
	builds   int
	applies  int
	fail     bool
	panicOn  bool
	staged   string
	applied  string
	lastJSON string
}

func (f *fakeHandler) Build() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.panicOn {
		panic("fake build panic")
	}
	if f.fail {
		return errors.New("fake build error")
	}
	f.builds++
	f.staged = "staged-component"
	return nil
}

func (f *fakeHandler) Apply() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.applies++
	f.applied = f.staged
	f.staged = ""
}

func (f *fakeHandler) reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.builds, f.applies = 0, 0
	f.fail, f.panicOn = false, false
	f.staged, f.applied = "", ""
}

// newTestSettings 构造测试用 Settings（指向全局配置对象，与 Setup 保持一致）
func newTestSettings(handlers ...ReloadHandler) *Settings {
	// bitxx/logger 对空 level 会 log.Fatalf（os.Exit），测试环境必须给合法等级
	LoggerConfig.Level = "info"
	return &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Logger:      LoggerConfig,
			Auth:        AuthConfig,
			Database:    DatabaseConfig,
			Databases:   &DatabasesConfig,
			Gen:         GenConfig,
		Cache:       CacheConfig,
		Queue:       QueueConfig,
		Locker:      LockerConfig,
		RateLimiter: RateLimiterConfig,
		LoginLock:   LoginLockConfig,
	},
		handlers: handlers,
	}
}

func TestApplyBuildThenApplyOrder(t *testing.T) {
	h1, h2 := &fakeHandler{}, &fakeHandler{}
	s := newTestSettings(h1, h2)

	s.Init()

	if h1.builds != 1 || h2.builds != 1 {
		t.Fatalf("expect 1 build each, got %d/%d", h1.builds, h2.builds)
	}
	if h1.applies != 1 || h2.applies != 1 {
		t.Fatalf("expect 1 apply each, got %d/%d", h1.applies, h2.applies)
	}
	if len(s.appliedRaw) == 0 {
		t.Fatal("expect appliedRaw snapshot after successful init")
	}
}

func TestReloadFailureRollback(t *testing.T) {
	oldExpired := CacheConfig.Expired
	defer func() { CacheConfig.Expired = oldExpired }()

	CacheConfig.Expired = 300
	h1, h2 := &fakeHandler{}, &fakeHandler{}
	s := newTestSettings(h1, h2)
	s.Init()

	// 模拟框架扫描到新配置（原地改写全局配置对象）
	CacheConfig.Expired = 999
	h2.fail = true

	s.OnChange()

	if h2.applies != 1 {
		t.Fatalf("expect h2 only applied once (init), got %d", h2.applies)
	}
	if h1.applies != 1 {
		t.Fatalf("h1 must not be applied when h2 build fails, got %d applies", h1.applies)
	}
	if CacheConfig.Expired != 300 {
		t.Fatalf("config must roll back to last applied value, got %d", CacheConfig.Expired)
	}
	if h1.builds != 2 {
		t.Fatalf("h1 must be built again on reload, got %d builds", h1.builds)
	}
	if h2.builds != 1 {
		t.Fatalf("h2 build must fail without counting (fail checked first), got %d builds", h2.builds)
	}
}

func TestReloadSuccessAppliesAll(t *testing.T) {
	h1, h2 := &fakeHandler{}, &fakeHandler{}
	s := newTestSettings(h1, h2)
	s.Init()
	h1.reset()
	h2.reset()

	s.OnChange()

	if h1.applies != 1 || h2.applies != 1 {
		t.Fatalf("expect 1 apply each after successful reload, got %d/%d", h1.applies, h2.applies)
	}
}

func TestReloadPanicRollback(t *testing.T) {
	oldExpired := CacheConfig.Expired
	defer func() { CacheConfig.Expired = oldExpired }()

	CacheConfig.Expired = 300
	h1, h2 := &fakeHandler{}, &fakeHandler{}
	s := newTestSettings(h1, h2)
	s.Init()

	CacheConfig.Expired = 999
	h2.panicOn = true

	s.OnChange() // 不应 panic

	if CacheConfig.Expired != 300 {
		t.Fatalf("config must roll back on panic, got %d", CacheConfig.Expired)
	}
	if h1.applies != 1 {
		t.Fatalf("h1 must not be applied when another handler panics during build")
	}
}

func TestReloadSerialized(t *testing.T) {
	h := &fakeHandler{}
	s := newTestSettings(h)
	s.Init()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.OnChange()
		}()
	}
	wg.Wait()

	// 每次 OnChange：1 次 Build + 1 次 Apply（成功路径），互不交错
	if h.builds != 11 || h.applies != 11 {
		t.Fatalf("expect serialized builds/applies (1 init + 10 reloads), got %d/%d", h.builds, h.applies)
	}
}

func TestInitFailFast(t *testing.T) {
	h := &fakeHandler{fail: true}
	s := newTestSettings(h)
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expect Init to panic on build failure")
		}
	}()
	s.Init()
}

func TestLoggerValidate(t *testing.T) {
	if err := (Logger{Level: "trace", Path: ""}).Validate(); err != nil {
		t.Fatalf("valid level must pass, got %v", err)
	}
	if err := (Logger{Level: "bogus"}).Validate(); err == nil {
		t.Fatal("invalid level must fail validation")
	}
	if err := (Logger{Level: "info", Path: t.TempDir()}).Validate(); err != nil {
		t.Fatalf("creatable path must pass, got %v", err)
	}
	if err := (Logger{Level: "info", Path: "Z:/no-such-dir-xxxx/logs"}).Validate(); err == nil {
		t.Fatal("uncreatable path must fail validation")
	}
}

func TestReloadScanUpdatesGlobalsInPlace(t *testing.T) {
	oldExpired := CacheConfig.Expired
	defer func() { CacheConfig.Expired = oldExpired }()
	CacheConfig.Expired = 300

	h := &fakeHandler{}
	s := newTestSettings(h)
	s.Init()
	h.reset()

	yml := "settings:\n  cache:\n    expired: 777\n"
	if err := s.scan([]byte(yml)); err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	if CacheConfig.Expired != 777 {
		t.Fatalf("reload must update global config in place, got %d", CacheConfig.Expired)
	}
	// 指针身份保持：e.Settings 与全局配置仍为同一对象
	if s.Settings.Cache != CacheConfig {
		t.Fatal("pointer identity must be preserved after reload scan")
	}
	if h.applies != 1 {
		t.Fatalf("expect 1 apply after successful reload, got %d", h.applies)
	}
}

func TestReloadScanInvalidKeepsOld(t *testing.T) {
	oldExpired := CacheConfig.Expired
	defer func() { CacheConfig.Expired = oldExpired }()
	CacheConfig.Expired = 300

	h := &fakeHandler{}
	s := newTestSettings(h)
	s.Init()
	h.reset()

	if err := s.scan([]byte("settings:\n  cache: [broken yaml")); err == nil {
		t.Fatal("expect parse error for invalid yaml")
	}
	if CacheConfig.Expired != 300 {
		t.Fatalf("invalid reload must keep last applied config, got %d", CacheConfig.Expired)
	}
	if h.applies != 0 {
		t.Fatalf("invalid reload must not apply any handler, got %d applies", h.applies)
	}
}

func TestReloadScanNilFieldRelink(t *testing.T) {
	oldExpired := CacheConfig.Expired
	defer func() { CacheConfig.Expired = oldExpired }()
	CacheConfig.Expired = 300

	h := &fakeHandler{}
	s := newTestSettings(h)
	s.Init()
	h.reset()

	// 新配置删除 cache 段：字段被置 nil 后 relinkGlobals 应重新指回全局对象
	yml := "settings:\n  application:\n    name: test\n"
	if err := s.scan([]byte(yml)); err != nil {
		t.Fatalf("scan failed: %v", err)
	}
	if s.Settings.Cache != CacheConfig {
		t.Fatal("nil field must be relinked back to global config object")
	}
	if CacheConfig.Expired != 300 {
		t.Fatalf("cache config untouched by unrelated reload, got %d", CacheConfig.Expired)
	}
	if h.applies != 1 {
		t.Fatalf("expect 1 apply after successful reload, got %d", h.applies)
	}
}

func TestReloadScanUpdatesLoginLockInPlace(t *testing.T) {
	oldMax := LoginLockConfig.MaxFailsByAccount
	defer func() { LoginLockConfig.MaxFailsByAccount = oldMax }()
	LoginLockConfig.MaxFailsByAccount = 5

	h := &fakeHandler{}
	s := newTestSettings(h)
	s.Init()
	h.reset()

	yml := "settings:\n  loginLock:\n    enabled: false\n    maxFailsByAccount: 9\n"
	if err := s.scan([]byte(yml)); err != nil {
		t.Fatalf("scan failed: %v", err)
	}
	if LoginLockConfig.Enabled {
		t.Fatal("enabled must be false after reload")
	}
	if LoginLockConfig.MaxFailsByAccount != 9 {
		t.Fatalf("reload must update loginLock in place, got %d", LoginLockConfig.MaxFailsByAccount)
	}
	if s.Settings.LoginLock != LoginLockConfig {
		t.Fatal("pointer identity must be preserved for loginLock after reload scan")
	}
}
