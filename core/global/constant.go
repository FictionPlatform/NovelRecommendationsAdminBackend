package global

/*
 * 需要和字典匹配
 */

const (
	RouteRootPath = "/admin-api"
	ModelName     = "go-admin"
	LoginLog      = "login_log_queue"
	OperateLog    = "operate_log_queue"
	TrafficKey    = "X-Request-Id"
	LoggerKey     = "_go-admin-logger-request"

	// SysStatusOk 通用-正常
	SysStatusOk      = "1"
	SysStatusNotOk   = "2"
	SysStatusBanned  = "2" // app 读者-禁言（可登录可读，禁止写操作）
	SysStatusCancelled = "3" // app 读者-注销（终态，禁止登录与一切请求）
)

const (
	DBDriverMysql    = "mysql"
	DBDriverPostgres = "postgres"
)

const (
	ModeDev  string = "dev"  //开发模式
	ModeTest string = "test" //测试模式
	ModeProd string = "prod" //生产模式
)
