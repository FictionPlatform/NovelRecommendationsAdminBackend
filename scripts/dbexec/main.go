package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bitxx/load-config"
	"github.com/bitxx/load-config/source/file"
	"go-admin/core/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// entity 仅读取配置文件中 database 段（复用项目 config.Database 结构，避免配置漂移）
type entity struct {
	Settings struct {
		Database *config.Database `yaml:"database"`
	} `yaml:"settings"`
}

// OnChange 实现 loadconfig.Entity 接口（脚本不订阅热更新，留空即可）
func (e *entity) OnChange() {}

func main() {
	configPath := flag.String("c", "config/settings.yml", "配置文件路径")
	sqlFile := flag.String("file", "app_pgsql.sql", "SQL 文件路径")
	strict := flag.Bool("strict", false, "出错即停止（默认跳过错误继续执行）")
	yes := flag.Bool("yes", false, "跳过确认提示")
	reset := flag.Bool("reset-schema", false, "执行前重建 public schema（仅 postgres，清空全部数据）")
	flag.Parse()

	if *sqlFile == "" {
		fatal("缺少参数 -file")
	}
	sqlBytes, err := os.ReadFile(*sqlFile)
	if err != nil {
		fatal("读取 SQL 文件失败: " + err.Error())
	}

	e := &entity{}
	if err := initConfig(*configPath, e); err != nil {
		fatal("加载配置失败: " + err.Error())
	}
	dbCfg := e.Settings.Database
	if dbCfg == nil || dbCfg.Source == "" {
		fatal("配置中未找到 database.source")
	}
	if *reset {
		if strings.ToLower(dbCfg.Driver) != "postgres" && !strings.Contains(strings.ToLower(dbCfg.Driver), "pg") {
			fatal("-reset-schema 仅支持 postgres")
		}
		if !*yes {
			fmt.Print("-reset-schema 将 DROP SCHEMA public CASCADE 清空整个库。输入 yes 继续: ")
			line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			if strings.TrimSpace(line) != "yes" {
				fmt.Println("已取消")
				return
			}
		}
	} else if strings.Contains(string(sqlBytes), "DROP TABLE") && !*yes {
		fmt.Printf("警告: %s 包含 DROP TABLE 语句，执行将重建并清空已有数据。输入 yes 继续: ", *sqlFile)
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(line) != "yes" {
			fmt.Println("已取消")
			return
		}
	}

	db, err := openDB(dbCfg.Driver, dbCfg.Source)
	if err != nil {
		fatal("连接数据库失败: " + err.Error())
	}
	if *reset {
		if err := db.Exec("DROP SCHEMA public CASCADE").Error; err != nil {
			fatal("DROP SCHEMA 失败: " + err.Error())
		}
		if err := db.Exec("CREATE SCHEMA public").Error; err != nil {
			fatal("CREATE SCHEMA 失败: " + err.Error())
		}
		fmt.Println("public schema 已重建")
	}

	bad, ok := 0, 0
	for i, stmt := range splitSQL(string(sqlBytes)) {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" || strings.HasPrefix(stmt, "--") {
			continue
		}
		if err := db.Exec(stmt).Error; err != nil {
			bad++
			head := stmt
			if n := strings.IndexByte(head, '\n'); n > 0 {
				head = head[:n]
			}
			if len(head) > 120 {
				head = head[:120] + "..."
			}
			fmt.Printf("  [%d] 失败: %v\n      %s\n", i, err, head)
			if *strict {
				fmt.Printf("共执行 %d 条，失败 %d 条\n", ok+bad, bad)
				os.Exit(1)
			}
		} else {
			ok++
		}
	}
	fmt.Printf("执行完成: 成功 %d 条，失败 %d 条\n", ok, bad)
	if bad > 0 {
		os.Exit(1)
	}
}

func initConfig(path string, e *entity) error {
	cfg, err := loadconfig.NewConfig(
		loadconfig.WithSource(file.NewSource(file.WithPath(path))),
		loadconfig.WithEntity(e),
	)
	if err != nil {
		return err
	}
	return cfg.Load()
}

func openDB(driver, source string) (*gorm.DB, error) {
	switch strings.ToLower(driver) {
	case "postgres", "postgresql", "pg":
		return gorm.Open(postgres.Open(source), &gorm.Config{})
	case "mysql", "mariadb":
		return gorm.Open(mysql.Open(source), &gorm.Config{})
	default:
		return nil, fmt.Errorf("不支持的数据库驱动: %s", driver)
	}
}

func fatal(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// splitSQL 按分号切分，感知单/双引号与 -- 行注释（处理字符串内分号与含引号的注释行）
func splitSQL(sql string) []string {
	var out []string
	var cur []rune
	var quote rune
	inComment := false
	for _, r := range sql {
		if inComment {
			if r == '\n' {
				inComment = false
				cur = append(cur, r)
			}
			continue
		}
		if quote == 0 {
			if r == '-' && len(cur) > 0 && cur[len(cur)-1] == '-' {
				inComment = true
				cur = cur[:len(cur)-1]
				continue
			}
			if r == '\'' || r == '"' {
				quote = r
				cur = append(cur, r)
				continue
			}
			if r == ';' {
				out = append(out, string(cur))
				cur = nil
				continue
			}
			cur = append(cur, r)
			continue
		}
		if r == quote {
			quote = 0
		}
		cur = append(cur, r)
	}
	if len(cur) > 0 {
		out = append(out, string(cur))
	}
	return out
}
