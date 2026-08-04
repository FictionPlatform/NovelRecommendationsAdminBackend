package config

import (
	"fmt"
	"github.com/bitxx/logger/logbase"
	"go-admin/core/utils/log"
	"os"
)

type Logger struct {
	Type      string
	Path      string
	Level     string
	Stdout    string
	EnabledDB bool
	Cap       uint
}

// Validate 校验日志配置是否合法。
// bitxx/logger 的 NewLogger 对非法 level / 目录创建失败会直接 log.Fatalf（os.Exit），
// 热更新路径必须先校验再 Setup，防止一条坏配置杀死整个进程。
func (e Logger) Validate() error {
	if _, err := logbase.GetLevel(e.Level); err != nil {
		return fmt.Errorf("invalid logger level %q: %w", e.Level, err)
	}
	if e.Path != "" {
		if err := os.MkdirAll(e.Path, os.ModePerm); err != nil {
			return fmt.Errorf("create log dir %q error: %w", e.Path, err)
		}
	}
	return nil
}

func (e Logger) Setup() {

	log.Init(log.LoggerConf{
		Type:      e.Type,
		Path:      e.Path,
		Level:     e.Level,
		Stdout:    e.Stdout,
		EnabledDB: e.EnabledDB,
		Cap:       e.Cap,
	})
}

var LoggerConfig = new(Logger)
