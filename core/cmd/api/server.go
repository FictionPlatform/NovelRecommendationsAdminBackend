package api

import (
	"context"
	"errors"
	"fmt"
	"go-admin/app"
	mycasbin "go-admin/core/casbin"
	"go-admin/core/config"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"
	"go-admin/core/runtime"
	"go-admin/core/storage/cache"
	"go-admin/core/storage/database"
	limiterSetup "go-admin/core/storage/limiter"
	"go-admin/core/storage/locker"
	queueSetup "go-admin/core/storage/queue"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/log"
	"go-admin/core/utils/storage"
	"go-admin/core/utils/strutils"
	"go-admin/core/utils/textutils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bitxx/load-config/source/file"

	"go-admin/app/admin/sys/models"
	"go-admin/core/global"
	"go-admin/core/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-admin/docs"
)

var (
	configPath string
	StartCmd   *cobra.Command
	// appQueue 全局消息队列实例，优雅关闭时统一释放
	appQueue storage.AdapterQueue
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			//初始化权限校验
			auth.InitAuth()

			//国际化-初始化底层
			if err := lang.InitLang(); err != nil {
				return err
			}

			//国际化-业务

			AppRouters = append(AppRouters, app.AllRouter()...)

			return run()
		},
	}

	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	// 1. 读取配置（各组件 Handler 支持配置热更新：全部构建成功后才原子切换，失败保持旧组件）
	config.Setup(
		file.NewSource(file.WithPath(configPath)),
		database.NewHandler(),
		cache.NewHandler(),
		queueSetup.NewHandler(),
		locker.NewHandler(),
		limiterSetup.NewHandler(),
	)

	// 2.casbin设置
	for host := range config.DatabasesConfig {
		db := runtime.RuntimeConfig.GetDbByKey(host)
		e := mycasbin.Setup(db, "admin_sys_")
		runtime.RuntimeConfig.SetCasbin(host, e)
	}

	// 3. 配置热更新：轮询监听配置文件，变更后受保护重建 DB/缓存/队列/锁/限流（全部构建成功才切换，失败自动回滚）
	go config.Watch(configPath, 3*time.Second)

	// 3. 注册监听函数
	queue := runtime.RuntimeConfig.GetMemoryQueue("")
	appQueue = queue
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperLog)
	go queue.Run()
	log.Info(`starting api server...`)
}

func run() error {
	if config.ApplicationConfig.Mode == global.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: runtime.RuntimeConfig.GetEngine(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverErr := make(chan error, 1)
	go func() {
		// 服务连接，不考虑https，该服务结偶，由专业的转发工具提供，如nginx
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("Server listen error: %v", err)
			serverErr <- err
		}
	}()
	log.Info(textutils.Red(string(global.LogoContent)))
	tip()
	log.Info(textutils.Green("Server run at:"))
	log.Infof("-  Local:   http://localhost:%d/ \r", config.ApplicationConfig.Port)
	log.Infof("-  Network: http://%s:%d/ \r", iputils.GetLocaHost(), config.ApplicationConfig.Port)
	log.Infof("%s Enter Control + C Shutdown Server \r", strutils.GetCurrentTimeStr())

	// 等待中断/终止信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// 同时监听 SIGTERM（docker stop / k8s 默认发送），避免容器停止时悬挂
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case <-quit:
		// 正常关闭流程
		log.Infof("%s Shutdown Server ... \r", strutils.GetCurrentTimeStr())
	case err := <-serverErr:
		// 启动失败，直接返回错误
		log.Errorf("Server failed to start: %v", err)
		return err
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Server shutdown error: %v", err)
		return err // 返回错误给 RunE
	}

	// 关闭消息队列，避免 goroutine 泄漏
	if appQueue != nil {
		appQueue.Shutdown()
	}
	// 关闭全部组件登记的 Redis 客户端
	config.CloseAllRedisClients()
	log.Info("Server exiting")

	return nil
}

func tip() {
	usageStr := `欢迎使用 ` + textutils.Green(config.ApplicationConfig.Name+" "+config.ApplicationConfig.Version) + ` 可以使用 ` + textutils.Red(`--help`) + ` 查看命令`
	log.Infof("%s", usageStr)
}

func initRouter() {
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		h = gin.New()
		runtime.RuntimeConfig.SetEngine(h)
	}
	r, ok := h.(*gin.Engine)
	if !ok {
		panic("not support other engine")
	}
	// 显式设置可信代理：未配置时不信任任何代理，拒绝伪造的 X-Forwarded-For（防 IP 伪造绕过限流/黑名单）
	if len(config.ApplicationConfig.TrustedProxies) > 0 {
		if err := r.SetTrustedProxies(config.ApplicationConfig.TrustedProxies); err != nil {
			panic("SetTrustedProxies error: " + err.Error())
		}
	} else {
		r.SetTrustedProxies(nil)
	}
	//r.Use(middleware.Metrics())
	r.Use(middleware.RequestId()).Use(log.SetRequestLogger)

	// swagger 文档仅开发模式（settings.yml mode: dev）开放，避免生产暴露接口文档
	if config.ApplicationConfig.Mode == global.ModeDev {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	middleware.InitMiddleware(r)
}
