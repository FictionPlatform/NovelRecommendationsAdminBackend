# go-admin 后台管理系统 —— LLM 使用说明

> 本文档专为 LLM / AI 编程助手编写，用于在本仓库中快速理解架构、遵守开发规范、正确完成任务。
> 阅读对象：任何在本仓库中执行「改代码、加功能、查问题」任务的 AI 助手。

---

## 1. 项目概述

- **模块名**: `go-admin`（Gin 风格 go-admin 后台管理系统）
- **定位**: 基于 Golang 的后台管理系统，前后端分离（本仓库仅含**后端**，`web/` 前端目录不在本仓库内）
- **接口前缀**: 管理后台 API 挂在 `/admin-api/v1` 下（常量 `core/global/constant.go` 中 `RouteRootPath = "/admin-api"`）；小说推荐平台等对外业务模块挂在独立的 `/web-api/v1` 下（模块内自定义前缀，见 `docs/小说推荐平台需求文档.md` §4.1）
- **默认端口**: 8888（`config/settings.yml` 中 `settings.application.port`）
- **默认账户**: admin/123456（顶级账户）、test/123456（受限账户）
- **支持数据库**: MySQL 8.x 与 PostgreSQL 16.x（脚本见根目录 `app_mysql.sql` / `app_pgsql.sql`，共 28 张表）
- **多语言**: 仅中英文，默认中文

## 2. 技术栈

| 领域 | 技术 |
|---|---|
| Web 框架 | gin-gonic/gin v1.12 |
| ORM | gorm.io/gorm v1.31 + dbresolver（读写分离） |
| 鉴权 | appleboy/gin-jwt/v3（JWT）、casbin/casbin/v2（RBAC 接口权限） |
| 配置 | bitxx/load-config（YAML + 热更新） |
| 日志 | bitxx/logger（default/zap/logrus） |
| 缓存/队列/锁 | 内存实现为主，Redis（go-redis/redislock/redisqueue）可选 |
| WAF | coraza/brandy（OWASP CRS 规则集，`rulesets/` 目录） |
| 其他 | excelize（导出）、bluemonday（HTML 消毒）、aliyun-oss-go-sdk、gopsutil（监控）、gorilla/websocket、OpenTelemetry、swaggo（swagger，仅 dev） |
| Go 版本 | 1.26.5 |

## 3. 快速启动

```shell
# 1. 依赖
go mod tidy

# 2. 生成接口描述文件（必须执行！会重新生成 app/admin/sys/models/sys_api_gen_desc.go）
go generate

# 3. 编译
go build -a -o go-admin-api main.go

# 4. 运行（子命令为 server）
./go-admin-api server -c=config/settings.yml

# 测试
go test ./...
```

**重要**：`go generate`（触发 `app/admin/sys/models/parseapi/gen_api_desc.go`，扫描各 `apis/` 目录的注释生成接口常量映射）在每次**打包前和运行前**都必须执行，否则新增接口无法在后台「接口数据同步」中被录入。

## 4. 目录结构（后端）

```
adminserver/
├── main.go                      # 入口，含 go:generate 指令与 swagger 注解
├── app/                         # ★ 全部业务代码
│   ├── init.go                  # AllRouter() 汇总 admin/app/plugins 三块路由（勿动）
│   ├── admin/sys/               # 系统基础服务（表前缀 admin_sys_*）
│   │   ├── models/              #   GORM 模型 + parseapi/gen_api_desc.go（接口扫描工具）
│   │   ├── apis/                #   Gin 处理器（薄层）
│   │   ├── router/              #   路由注册（init() 自动挂载）
│   │   └── service/             #   业务逻辑 + service/dto/ 请求体
│   ├── app/                     # 个性化业务层（表前缀 app_*），目前仅 user 子业务
│   │   ├── common/dto/join.go   #   跨表联查 DTO
│   │   └── user/                #   用户、等级、账变、区号、配置、操作日志
│   └── plugins/                 # 插件层（表前缀 plugins_*）
│       ├── content/             #   CMS：文章/分类/公告
│       ├── filemgr/             #   App 安装包管理（含 OSS 上传）
│       └── msg/                 #   短信/邮件验证码记录（只读）
├── core/                        # ★ 框架层（一般不改）
│   ├── cmd/                     #   cobra 根命令 + server 子命令（启动流程）
│   ├── config/                  #   配置结构体 + 热更新机制
│   ├── runtime/                 #   运行时服务定位器 RuntimeConfig
│   ├── middleware/              #   全局中间件链 + auth(jwt/casbin) 子包
│   ├── casbin/                  #   RBAC 适配器
│   ├── dto/                     #   分页/搜索/统一响应/API 助手
│   ├── lang/                    #   i18n
│   ├── storage/                 #   热更新初始化器（db/cache/queue/locker）
│   ├── utils/                   #   工具包（encrypt/excel/oss/cache/queue/tree/...）
│   ├── global/                  #   常量（RouteRootPath 等）
│   └── ws/                      #   WebSocket 管理器
├── config/
│   ├── settings.yml             # ★ 运行配置
│   ├── base/constant/           #   业务常量 + 代码生成模板名映射
│   ├── base/lang/               #   中文消息码（按模块分文件）
│   └── lang/en.csv              #   英文翻译
├── static/template/             # ★ 代码生成模板（Go + React 共 9 个）
├── rulesets/                    # WAF 规则（default.conf + OWASP CRS）
├── docs/                        # swagger + 代码审查报告 + 本文档
└── files/                       # 运行时文件存储（上传/日志），配置 fileRootPath
```

## 5. 核心架构要点

### 5.1 启动流程（`core/cmd/api/server.go`）
1. 加载配置 `config.Setup(...)`，注册 db/cache/queue/locker 热更新 Handler（失败即 panic，fail-fast）
2. 为每个库初始化 Casbin enforcer，存入 `runtime.RuntimeConfig`
3. 注册登录/操作日志的异步消费者（内存队列 `queue.Register`）
4. `auth.InitAuth()` 初始化 JWT（校验 secret 强度）、`lang.InitLang()` 加载翻译
5. `AppRouters = append(AppRouters, app.AllRouter()...)` 聚合全部路由闭包
6. `initRouter()`：设置可信代理（留空则不信任何代理）、RequestId、swagger（仅 dev）、`middleware.InitMiddleware(r)` 全局链，然后逐个执行路由闭包
7. 优雅退出（5s 超时 + 队列 Shutdown）

### 5.2 运行时（`core/runtime/`）
全局单例 `runtime.RuntimeConfig`（接口 `Runtime`），线程安全 map 存储：多库 `dbs`（按请求 Host 取库，`"*"` 兜底）、casbin enforcer、gin engine、缓存/队列/锁适配器。业务代码获取 DB 一律用 `ginutils.GetOrm(c)` 或 `runtime.RuntimeConfig.GetDbByKey(c.Request.Host)`。

### 5.3 配置热更新
所有基础设施组件实现 `Build()`（预构建）/ `Apply()`（原子替换）接口，配置文件变更自动热更，失败自动回滚（JSON 快照）。

### 5.4 多租户
按请求 Host 区分租户：DB、Casbin、缓存/队列 key 均带 host 前缀（`"app:" + prefix`、`__host`）。单库场景不受影响。

### 5.5 全局中间件链（`core/middleware/init.go`，顺序敏感）
`Recovery → WithContextDb → LoggerToFile(敏感字段脱敏) → CustomError(panic 契约) → KeepAlive → Options(CORS 白名单) → Secure → WAF(Coraza) → IPBlacklist → RateLimiter → Trace(OTel)`

## 6. 配置说明（config/settings.yml）

| 段 | 关键项 | 说明 |
|---|---|---|
| `settings.application` | mode/host/port, `enableDP`(数据权限开关), `fileRootPath`(文件根，不能以`./`或`/`开头), `isSingleLogin`(单点登录), `trustedProxies`(可信代理，留空防 XFF 伪造), `corsOrigins`(跨域白名单), `ampKey`(高德定位 key) | |
| `settings.logger` | path/stdout/level/enabledDB/type/cap | |
| `settings.auth` | `timeout`(7200s), `maxRefresh`(7天), `secret`(JWT 密钥，启动强校验), `secretAes`(16/24/32 字节 AES 密钥，用于手机号/邮箱加密), `enableDeviceCheck`(设备数限制), `enableBlacklist`(JWT 黑名单), `maxDeviceCount`(2) | |
| `settings.database` | driver(mysql/postgres)，`source` 支持 `${DB_PASSWORD}` 环境变量注入；禁止空密码/弱口令 | |
| `settings.gen` | `frontPath: ./web/src`（代码生成前端输出路径，前端目录不在本仓库） | |
| `settings.cache` | `expired`(300s 默认 TTL)；redis 注释掉即用内存缓存 | |
| `settings.queue` | memory.poolSize；redis 可选 | |
| `settings.rateLimiter` | enabled/limit(100 次)/period(60s)/blacklist(IP+CIDR)；配 redis 则分布式计数 | |

## 7. ★ 业务开发规范（新增/修改业务时必读）

### 7.1 分层目录约定（每个业务模块固定五层，代码生成器按此产出）
```
app/<app|admin|plugins>/<业务名>/
├── models/<表名>.go          # GORM 模型（TableName() 方法；密码字段 bcrypt BeforeCreate/BeforeUpdate；敏感字段用 encrypt.AesEncrypt）
├── service/<表名>.go         # 业务逻辑（GetPage/Get/Insert/Update/Delete/Export）
├── service/dto/<表名>.go     # 请求体（QueryReq 内嵌 dto.Pagination，字段带 search:"..." 标签）
├── apis/<表名>.go            # 薄 handler（MakeContext/MakeOrm/Bind/MakeService 链式调用）
└── router/<表名>.go          # init() 中调用 register<类名>Router，把函数 append 进 routerNoCheckRole / routerCheckRole
```

### 7.2 路由注册机制
- 每个模块 `router.go` 声明两个包级切片：
  - `routerNoCheckRole []func(*gin.RouterGroup)` —— 无需登录（如 captcha、login）
  - `routerCheckRole []func(v1 *gin.RouterGroup)` —— 需登录 + 角色校验
- 各业务 router 文件在 `init()` 里往这两个切片 append 注册函数
- 模块 `router.go` 的 `InitRouter()` 从 `runtime.RuntimeConfig.GetEngine()` 取 engine，建 `v1 := r.Group(global.RouteRootPath + "/v1")`，执行所有已注册函数
- 业务分组建议：`v1.Group("/<包名>/<业务名>/<模块名>", middleware.Auth(), middleware.AuthCheckRole())`

### 7.3 命名与注释规范（严格遵守）
1. **数据库表前缀**：系统 `admin_sys_*`、主业务 `app_*`、插件 `plugins_*`，后接「子业务_模块」，如 `app_user_level`
2. **接口方法注释**（接口同步依赖，**必须**）：`// 方法名 功能说明`，例如 `// GetPage 分页查询用户`
3. **按钮权限命名**：`业务名:包名:作用`，如 `app:user_level:add` / `del` / `edit` / `query`，需与前端页面 `permission` 属性一致
4. **菜单路由路径**：必须满足「目录/菜单」层级逻辑
5. 新业务模块接入：把生成的 `router.go.bk` 去掉 `.bk`，然后在 `app/init.go` 对应层 `init.go` 中 `append` 该模块的 `InitRouter`

### 7.4 DTO 搜索标签（核心能力）
`core/dto/search` 通过反射解析 `search:"type:...;column:...;table:...;on:...;join:..."` 标签自动生成查询条件。支持类型：`exact/iexact、contains/icontains、gt/gte/lt/lte、startswith/endswith、in、isnull、order`，以及 `left/inner` 联表（`on:a:b`）。零值字段自动跳过。分页上限 100（导出可经 `PageSizeLimit` 提升，硬上限 10000）。

### 7.5 统一响应与错误
- 响应体：`core/dto/response` 的 `OK/Error/PageOK/Download`；`core/dto/api` 提供 `Api` 助手（OK/PageOK/Error/DownloadExcel/Zip）
- 业务错误：抛 `CustomError#<httpCode>#<msg>`（中间件 `customererror.go` 捕获），或直接 `c.Error`；**不要** panic 未处理
- 消息码：新增错误消息在 `config/base/lang/` 对应文件注册（数字码），用 `lang.MsgByCode` / `MsgErr` 取文案

### 7.6 代码生成器（后台「系统管理 → 代码生成」）
流程：导入库表（`sys_gen_table`）→ 编辑字段规则（`sys_gen_column`）→ 预览 / 生成到磁盘 / 下载 zip / 生成菜单（`/admin/sys/sys-table/gen/*`，仅 admin 角色）。
模板：`static/template/` 下 9 个 `.template`（model/dto/service/apis/router/business_router + react.api.ts/react.formmodal.tsx/react.view.tsx），模板名与输出路径映射见 `config/base/constant/admin_sys.go` 的 `TemplatInfo`。
**安全注意**：生成路径有白名单校验（`^[a-zA-Z0-9_-]+$`），改动模板时勿破坏该校验。

## 8. 权限体系

1. **JWT 登录**（`core/middleware/auth/jwtauth/jwtauth.go`，734 行）：Bearer token；`authCheck` 每请求实时校验 用户/角色状态（禁用、删除、改密即失效）、单点登录、JWT 黑名单（jti）、设备指纹（`dev_fp`，UserAgent+IP+Accept 摘要，设备数上限 `maxDeviceCount`）
2. **Casbin 接口权限**（`core/middleware/auth.go` 的 `AuthCheckRole`）：`admin` 角色全通过；否则 `e.Enforce(roleKey, path, method)`；`CasbinExclude` 列表（登录/登出/captcha/profile 等）豁免
3. **菜单/按钮绑定**：`admin_sys_api` 表 + 「接口数据同步」（`sys_api.Sync()`，读取 `sys_api_gen_desc.go` 的 `ApiDescMap`）→ 菜单管理绑定接口 → 角色管理分配菜单
4. **数据权限**（`core/middleware/permission.go`）：5 级数据范围（1 全部 / 2 自定义 / 3 本部门 / 4 本部门及子部门 / 5 本人），由 `settings.application.enableDP` 控制（默认关闭）
5. **AdminOnly**：仅 `admin` 角色可访问（代码生成接口使用）

## 9. 国际化（i18n）

- 中文消息码注册表：`config/base/lang/*.go`（`lang.MsgInfo` map，按模块分文件：base/admin_sys/app/plugins_*）
- 英文翻译：`config/lang/en.csv`（目前仅 3 条，其余回退中文）
- 获取当前语言：`ParseAcceptLanguage` / `GetAcceptLanguage(c)`（默认 zh-CN）
- 支持语言列表如需扩展：改 `core/lang/lang.go`

## 10. 日志与异步

- 请求日志：中间件 `LoggerToFile` 写文件（`settings.logger.path`），敏感字段（password/token/captcha 等）自动脱敏为 `***`，请求体缓冲上限 1MB
- 登录/操作日志：写入内存队列（`global.LoginLog` / `global.OperateLog`）→ 消费者 `models.SaveLoginLog/SaveOperLog` 落库，**不要**在 handler 里同步写库
- WebSocket：`core/ws` 管理器（`SendAll/SendGroup/SendOne`），配合 `fileutils.FileMonitoringById` 可做文件监控推送

## 11. LLM 常见任务指引

### 11.1 新增一个业务模块（标准做法）
1. 建表（遵循 `7.3` 命名），更新 SQL 脚本或使用后台代码生成导入
2. **首选**：后台「代码生成」导入表 → 配置字段 → 生成到磁盘 → `router.go.bk` 改名 → 在模块 `init.go` append 路由 → `go generate` → 编译
3. **手工编写**：严格按 7.1 五层结构写 models → dto → service → apis → router，每个方法加 `// 方法名 功能说明` 注释
4. 如需新错误文案：在 `config/base/lang/` 对应文件注册消息码
5. 验证：`go build ./...` + `go vet ./...` + `go test ./...`

### 11.2 修改既有功能
- 入口先找 `app/<层>/<业务>/apis/` 对应文件确认接口路径，再下钻 service/models
- 涉及鉴权/权限改动：检查 `core/middleware/auth/`、`CasbinExclude` 列表
- 涉及配置：改 `config/settings.yml` 时**同步**确认 `core/config/` 对应结构体与热更新 Handler

### 11.3 排错要点
- 启动失败：先看配置校验（JWT secret、DB 密码强度、WAF 规则加载）
- 接口 401/403：token 状态、角色菜单绑定、Casbin 策略（`admin_sys_casbin_rule` 表）
- 数据查不到：数据权限 scope、`search` 标签联表条件、分页上限 100
- 前端「接口数据同步」为空：未执行 `go generate` 或注释不合规（见 7.3-2）

## 12. 安全红线（勿违反）

已知问题清单详见 `docs/code-review-report.md`（63 项，含 14 个 Critical）。重点：
1. **密钥**：不得硬编码/提交 JWT `secret`、AES `secretAes`、数据库密码（配置支持 `${DB_PASSWORD}` 环境变量注入）；代码评审已发现历史密钥泄露（C1/C2）
2. **代码生成**：`sys_gen_table` 的任意文件写入是 RCE 风险点（C3），生成路径白名单校验不可移除，接口保持 `AdminOnly`
3. **日志**：禁止将明文密码/token 写入日志（脱敏由中间件处理，勿绕过）
4. **上传**：`files/` 静态服务无鉴权（C6），新增上传功能必须做类型白名单 + 大小限制 + 危险扩展名强制下载（参考 filemgr：仅 `.apk/.ipa/.zip`、200MB 上限）
5. **导出**：Excel 导出统一走 `excelutils.SanitizeCell` 防公式注入
6. **HTML**：富文本内容入库前必须 `bluemonday` 消毒（参考 content 插件）
7. **IP 安全**：限流/黑名单依赖真实 IP，`trustedProxies` 配置错误可被 XFF 伪造绕过
8. **异步队列**：消费者错误通道必须被消费（redisqueue 的 Errors 通道），防止死锁/泄漏

## 13. 关键文件索引（改代码前先看这些）

| 目的 | 文件 |
|---|---|
| 入口/编译指令 | `main.go` |
| 启动流程 | `core/cmd/api/server.go` |
| 常量（路由前缀等） | `core/global/constant.go` |
| 运行时服务定位器 | `core/runtime/*.go` |
| 中间件链 | `core/middleware/init.go` |
| JWT 引擎 | `core/middleware/auth/jwtauth/jwtauth.go` |
| 路由豁免清单 | `core/middleware/auth/casbin/settings.go` |
| 统一响应/分页/搜索 | `core/dto/response/`、`core/dto/pagination.go`、`core/dto/search/` |
| 配置结构体 | `core/config/*.go` |
| 代码生成引擎 | `app/admin/sys/service/sys_gen_table.go` |
| 代码生成模板 | `static/template/*.template` |
| 接口扫描工具 | `app/admin/sys/models/parseapi/gen_api_desc.go` |
| 中文消息码 | `config/base/lang/*.go` |
| 业务常量 | `config/base/constant/*.go` |
| 数据库脚本 | `app_mysql.sql` / `app_pgsql.sql` |
| 已知问题 | `docs/code-review-report.md` |
| 接口文档 | `docs/swagger.yaml`（dev 模式可访问 `/swagger/index.html`） |

## 14. 现有业务模块清单（路由前缀均为 /admin-api/v1）

| 模块 | 前缀 | 说明 |
|---|---|---|
| 系统管理 | `/admin/sys/sys-{user,role,post,dept,menu,dict,config,api,oper-log,login-log,monitor,table}` | 基础 RBAC + 日志 + 监控 + 代码生成 |
| app 用户 | `/app/user/{user,user-level,user-conf,user-account-log,user-country-code,user-oper-log}` | 业务用户（AES 加密手机号/邮箱、邀请树）、等级、账变、区号 |
| CMS | `/plugins/content/content-{article,category,announcement}` | 文章/分类/公告 |
| 安装包 | `/plugins/filemgr/filemgr-app` | App 包管理（本地/外链/阿里云 OSS） |
| 消息 | `/plugins/msg/msg-code` | 验证码发送记录（只读） |
