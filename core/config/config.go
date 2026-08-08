package config

import (
	"encoding/json"
	"fmt"
	loadconfig "github.com/bitxx/load-config"
	"github.com/bitxx/load-config/source"
	"go-admin/core/utils/log"
	"sync"
)

var (
	_cfg *Settings
)

// ReloadHandler 热更新处理器：Build 构建新组件（不得修改线上状态），
// 全部组件 Build 成功后才统一 Apply 原子切换；任一 Build 失败则本次热更新整体放弃。
type ReloadHandler interface {
	// Build 基于当前配置构建新组件到暂存区；失败返回 error，线上状态不受影响
	Build() error
	// Apply 将暂存区组件切换为线上状态（仅赋值操作，不应失败）
	Apply()
}

// Settings 兼容原先的配置结构
type Settings struct {
	Settings Config `yaml:"settings"`
	handlers []ReloadHandler

	// reloadMu 串行化热更新，防止并发 OnChange 交错重建组件
	reloadMu sync.Mutex
	// appliedRaw 最近一次成功生效的配置快照（JSON），热更新失败时用于回滚
	appliedRaw []byte
	// appliedLoggerJSON 最近一次已生效的日志配置（JSON），日志未变化时跳过重建
	appliedLoggerJSON string
}

// Init 首次初始化：失败直接 panic，快速暴露配置错误（fail-fast）
func (e *Settings) Init() {
	if err := e.apply(); err != nil {
		panic(fmt.Sprintf("config init error: %s", err.Error()))
	}
	log.Warn("!!! config init")
}

// OnChange 配置热更新（load-config 框架扫描新配置后回调）：
// 1. 串行化执行，避免并发重建；
// 2. 先构建全部新组件（暂存），任一失败即放弃本次更新并回滚配置快照，线上组件保持旧版；
// 3. 全部构建成功后才原子切换，避免在途请求拿到半初始化状态。
func (e *Settings) OnChange() {
	e.reloadMu.Lock()
	defer e.reloadMu.Unlock()

	defer func() {
		if r := recover(); r != nil {
			e.rollback()
			log.Errorf("!!! config reload panic: %v, rolled back to last applied config", r)
		}
	}()

	if err := e.apply(); err != nil {
		e.rollback()
		log.Errorf("!!! config reload failed: %s, keep last applied config", err.Error())
		return
	}
	log.Warn("!!! config change and reload")
}

// apply 构建并切换全部组件
func (e *Settings) apply() error {
	// 0. 校验日志配置：bitxx/logger 对非法配置直接 log.Fatalf（os.Exit），
	// 必须在校验通过后才允许 Setup，保证热更新失败可回滚而不是杀进程
	if err := e.Settings.Logger.Validate(); err != nil {
		return fmt.Errorf("logger config invalid: %w", err)
	}
	// 日志与多库兜底：幂等、低风险，先于组件重建执行
	// 日志必须先于组件构建初始化：组件 Build 过程中会调用 log 输出（如数据库连接日志），
	// 若此时日志尚未初始化会触发 nil pointer panic
	e.Settings.multiDatabase()
	e.setupLoggerIfChanged()

	// 1. 构建全部组件（暂存），任一失败则本次更新整体放弃，线上状态不变
	for _, h := range e.handlers {
		if err := h.Build(); err != nil {
			return fmt.Errorf("config build error: %w", err)
		}
	}

	// 2. 生效快照在切换前生成：Apply 是不可回退的提交动作，快照失败必须发生在切换之前
	raw, err := json.Marshal(e.Settings)
	if err != nil {
		return fmt.Errorf("marshal applied config error: %w", err)
	}

	// 3. 提交：统一切换组件（仅赋值操作）
	for _, h := range e.handlers {
		h.Apply()
	}
	e.appliedRaw = raw
	return nil
}

// setupLoggerIfChanged 仅当日志配置变化时重建日志组件。
// 校验已通过，且无关热更新不再重复创建日志实例。
func (e *Settings) setupLoggerIfChanged() {
	raw, err := json.Marshal(e.Settings.Logger)
	if err != nil || string(raw) == e.appliedLoggerJSON {
		return
	}
	e.Settings.Logger.Setup()
	e.appliedLoggerJSON = string(raw)
}

// rollback 热更新失败时，将配置回滚到最近一次成功生效的版本
func (e *Settings) rollback() {
	if len(e.appliedRaw) == 0 {
		return
	}
	if err := json.Unmarshal(e.appliedRaw, &e.Settings); err != nil {
		log.Errorf("config rollback error: %s", err.Error())
		return
	}
	// 回滚可能将配置字段置 nil（如新配置缺失该段），重新指回全局配置对象保证指针身份一致
	e.relinkGlobals()
	// 失败尝试若已重建日志，恢复为生效版本
	e.setupLoggerIfChanged()
	log.Warn("!!! config rolled back to last applied version")
}

// relinkGlobals 将扫描/回滚后为 nil 的配置字段重新指回包级全局配置对象，
// 保证 e.Settings 与全局配置（handlers 读取的 config.xxxConfig）始终共享同一对象。
func (e *Settings) relinkGlobals() {
	if e.Settings.Application == nil {
		e.Settings.Application = ApplicationConfig
	}
	if e.Settings.Logger == nil {
		e.Settings.Logger = LoggerConfig
	}
	if e.Settings.Auth == nil {
		e.Settings.Auth = AuthConfig
	}
	if e.Settings.Database == nil {
		e.Settings.Database = DatabaseConfig
	}
	if e.Settings.Databases == nil {
		e.Settings.Databases = &DatabasesConfig
	}
	if e.Settings.Gen == nil {
		e.Settings.Gen = GenConfig
	}
	if e.Settings.Cache == nil {
		e.Settings.Cache = CacheConfig
	}
	if e.Settings.Queue == nil {
		e.Settings.Queue = QueueConfig
	}
	if e.Settings.Locker == nil {
		e.Settings.Locker = LockerConfig
	}
	if e.Settings.RateLimiter == nil {
		e.Settings.RateLimiter = RateLimiterConfig
	}
	if e.Settings.LoginLock == nil {
		e.Settings.LoginLock = LoginLockConfig
	}
	if e.Settings.RegLimit == nil {
		e.Settings.RegLimit = RegLimitConfig
	}
}

// Config 配置集合
type Config struct {
	Application *Application          `yaml:"application"`
	Logger      *Logger               `yaml:"logger"`
	Auth        *Auth                 `yaml:"auth"`
	Database    *Database             `yaml:"database"`
	Databases   *map[string]*Database `yaml:"databases"`
	Gen         *Gen                  `yaml:"gen"`
	Cache       *Cache                `yaml:"cache"`
	Queue       *Queue                `yaml:"queue"`
	Locker      *Locker               `yaml:"locker"`
	RateLimiter *RateLimiter          `yaml:"rateLimiter"`
	LoginLock   *LoginLock            `yaml:"loginLock"`
	RegLimit    *RegLimit             `yaml:"regLimit"`
}

// 多db改造
func (e *Config) multiDatabase() {
	if e.Database == nil || len(*e.Databases) > 0 {
		return
	}
	*e.Databases = map[string]*Database{
		"*": e.Database,
	}
}

// Setup 载入配置文件
func Setup(s source.Source,
	handlers ...ReloadHandler) {
	_cfg = &Settings{
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
			RegLimit:    RegLimitConfig,
		},
		handlers: handlers,
	}
	var err error
	loadconfig.DefaultConfig, err = loadconfig.NewConfig(
		loadconfig.WithSource(s),
		loadconfig.WithEntity(_cfg),
	)
	if err != nil {
		panic(fmt.Sprintf("New config object fail: %s", err.Error()))
	}
	_cfg.Init()
}
