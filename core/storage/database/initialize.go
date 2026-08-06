package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/core/config"
	"go-admin/core/config/database"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
	"go-admin/core/utils/textutils"
	"os"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Handler 数据库热更新处理器：
// Build 构建全部数据库连接（暂存，不触碰线上），全部构建成功后才 Apply 原子切换；
// 配置未变化时复用现有连接，避免无关配置变更导致无谓断连。
type Handler struct {
	staged      map[string]*gorm.DB
	pendingJSON string
	appliedJSON string
}

func NewHandler() *Handler {
	return &Handler{}
}

// Build 构建新数据库连接（失败返回 error，线上连接不受影响）
func (h *Handler) Build() error {
	// resolveEnv 会原地改写 Source（${ENV} 占位符替换为实际值），先解析再比较，保证与上次应用后的配置可比
	for _, c := range config.DatabasesConfig {
		c.Source = resolveEnv(c.Source)
	}
	curr, err := json.Marshal(config.DatabasesConfig)
	if err != nil {
		return err
	}
	if h.appliedJSON == string(curr) {
		return nil
	}
	dbs := make(map[string]*gorm.DB, len(config.DatabasesConfig))
	for k := range config.DatabasesConfig {
		db, err := setupSimpleDatabase(k, config.DatabasesConfig[k])
		if err != nil {
			return err
		}
		dbs[k] = db
	}
	h.staged = dbs
	h.pendingJSON = string(curr)
	return nil
}

// Apply 切换为新连接。
// 注意：不关闭旧连接池——casbin 等组件仍持有旧 *gorm.DB 引用，关闭会导致权限校验失效；
// 旧连接仅在进程退出时回收。
func (h *Handler) Apply() {
	if h.pendingJSON == "" {
		return
	}
	for k, db := range h.staged {
		runtime.RuntimeConfig.SetDb(k, db)
	}
	h.staged = nil
	h.appliedJSON = h.pendingJSON
	h.pendingJSON = ""
}

func setupSimpleDatabase(host string, c *config.Database) (*gorm.DB, error) {
	if err := validateDSNPassword(c.Source); err != nil {
		return nil, err
	}
	log.Infof("%s => %s", host, textutils.Green(maskDSN(c.Source)))
	registers := make([]database.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		sources := make([]string, len(c.Registers[i].Sources))
		for j := range c.Registers[i].Sources {
			sources[j] = resolveEnv(c.Registers[i].Sources[j])
		}
		replicas := make([]string, len(c.Registers[i].Replicas))
		for j := range c.Registers[i].Replicas {
			replicas[j] = resolveEnv(c.Registers[i].Replicas[j])
		}
		registers[i] = database.NewResolverConfigure(
			sources,
			replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := database.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		return nil, fmt.Errorf("%s connect error: %s", c.Driver, err.Error())
	}
	log.Info(textutils.Green(c.Driver + " connect success !"))

	return db, nil
}

// envVarRegex 仅匹配 ${VAR} 形式的占位符（不展开裸 $VAR，避免破坏密码中的 $ 字符）
var envVarRegex = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)\}`)

// resolveEnv 将 DSN 中的 ${ENV_VAR} 占位符替换为环境变量值，未定义的变量替换为空串
func resolveEnv(source string) string {
	return envVarRegex.ReplaceAllStringFunc(source, func(m string) string {
		return os.Getenv(envVarRegex.FindStringSubmatch(m)[1])
	})
}

// dsnPasswordKVRegex 匹配 kv 风格 DSN 的密码：host=x password=y
var dsnPasswordKVRegex = regexp.MustCompile(`(?i)(\bpassword=)(\S+)`)

// dsnPasswordURLRegex 匹配 URL 风格 DSN 的密码：postgres://user:pass@host
var dsnPasswordURLRegex = regexp.MustCompile(`(?i)://([^:/\s]+):([^@\s/]+)@`)

// extractDSNPassword 提取 DSN 中的数据库密码（kv 风格或 URL 风格），未配置返回空串
func extractDSNPassword(source string) string {
	if m := dsnPasswordKVRegex.FindStringSubmatch(source); m != nil {
		return m[2]
	}
	if m := dsnPasswordURLRegex.FindStringSubmatch(source); m != nil {
		return m[2]
	}
	return ""
}

// maskDSN 脱敏 DSN：密码统一替换为 ***，防止明文口令写入日志
func maskDSN(source string) string {
	masked := dsnPasswordKVRegex.ReplaceAllString(source, `${1}***`)
	return dsnPasswordURLRegex.ReplaceAllString(masked, `://${1}:***@`)
}

// weakPasswords 常见弱口令黑名单
var weakPasswords = map[string]bool{
	"123456": true, "12345678": true, "123456789": true, "1234567890": true,
	"111111": true, "000000": true, "123123": true, "666666": true, "888888": true,
	"root": true, "admin": true, "administrator": true, "password": true,
	"pass": true, "passw0rd": true, "postgres": true, "qwerty": true, "abc123": true,
}

// validateDSNPassword 校验数据库口令：禁止空密码与弱口令，防止配置裸奔
func validateDSNPassword(source string) error {
	pw := extractDSNPassword(source)
	if pw == "" {
		return errors.New("database.source has no password; empty database password is prohibited, inject it via environment variable (e.g. password=${DB_PASSWORD})")
	}
	if len(pw) < 8 || weakPasswords[strings.ToLower(pw)] {
		return fmt.Errorf("database.source uses a weak password (less than 8 chars or common weak password); configure a strong password or inject via environment variable (e.g. password=${DB_PASSWORD})")
	}
	return nil
}
