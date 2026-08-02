# Go-Admin 后台管理系统 代码审查报告

- **审查日期**: 2026-08-02
- **审查范围**: 全项目（240 个 Go 文件 + 配置 + SQL 脚本 + 规则集 + 文档）
- **审查方式**: 静态代码审查（只读，未修改任何代码）
- **技术栈**: Gin + GORM + Casbin + JWT + Redis + PostgreSQL/MySQL
- **结论**: 项目结构清晰、CRUD 模式统一、搜索框架参数化无 SQL 注入，但存在**多项严重安全漏洞**（默认密钥、任意文件写入、明文口令存储、越权）和**若干功能缺陷**，建议按优先级修复。

---

## 问题统计

| 严重程度 | 数量 | 关键问题 |
|---------|------|---------|
| 严重 (Critical) | 14 | 默认密钥可伪造 token、代码生成任意文件写入、app 用户密码明文存储、数据权限中间件未挂载、请求体被截断导致接口失效等 |
| 高 (High) | 16 | 权限回收失效、越权提权、公式注入、pageSize 无上限、并发竞态等 |
| 中 (Medium) | 19 | 登录枚举、验证码弱、CORS 全开、热更新不生效、日志脱敏缺失等 |
| 低 (Low) | 14 | 死代码、文档过时、冗余字段、格式错误等 |
| **合计** | **63** | |

---

## 一、严重 (Critical)

### C1. JWT/AES 密钥硬编码入库且共用同一密钥
- **位置**: `config/settings.yml:42`（`secret: admin-api-20231019-jason`，已被 git 跟踪）、`core/middleware/auth/jwtauth/jwtauth.go:70-94`、`core/utils/encrypt/aes.go:10`
- **问题**:
  1. HS256 使用仓库内公开密钥，任何拿到仓库的人可离线伪造任意用户（含 admin）的 token；
  2. 该密钥同时用作 AES 数据加密密钥（settings.yml 注释自述"包括aes密钥共用"），可解密库中全部手机号/邮箱密文；
  3. 密钥已进入 git 历史，仅修改当前文件无法清除。
- **建议**: 密钥从环境变量/密钥管理服务注入；启动时校验拒绝默认值；JWT 签名密钥与数据加密密钥分离；git filter-repo 清理历史并轮换密钥。

### C2. app_user 登录密码/提现密码明文存储（设计缺陷，尚未被触发）
- **位置**: `app_mysql.sql:1374-1375`、`app_pgsql.sql:1610-1611`（`pwd`/`pay_pwd` 明文 `varchar(100)` 列）、`app/app/user/models/user.go:18-19`（字段带 `json:"pwd"/"payPwd"` 标签）、`app/app/user/service/dto/user.go:42-43`（`pwdOrder`/`payPwdOrder` 排序字段）
- **问题**: ① 种子数据中 `pwd`/`pay_pwd` 均为空（`ref_code` 列的 `akIiWm` 等为推荐码，非密码），无明文密码数据；② 但表结构为明文列且 service 层全项目无任何 bcrypt/scrypt 哈希逻辑——一旦 app 端注册/登录功能启用写入密码，写进去即为明文；③ `Pwd`/`PayPwd` 带 `json` 标签，字段一旦有值，任何接口响应都会泄露口令；④ DTO 允许按密码列排序，可探测字段存在性。
- **建议**: 改用 bcrypt/scrypt 哈希（项目已有 `core/utils/encrypt/security.go` 可复用）；模型字段 `json:"-"`；删除按密码列排序的字段；存量数据迁移脚本。

### C3. 代码生成接口可任意文件写入（路径穿越 → 潜在 RCE）
- **位置**: `app/admin/sys/router/sys_gen_table.go:15`（路由组仅 `middleware.Auth()` 无角色校验）、`app/admin/sys/service/sys_gen_table.go:539-571`（GenCode 用 `os.Create` 写文件）、`core/utils/fileutils/file.go:113-123`
- **问题**: 任意登录用户可：① 创建 gen_table 记录 → ② 把 `packageName`/`businessName`/`moduleName` 改为 `../../..` → ③ 触发 GenCode 任意写文件（webshell、覆盖源码）。同路由组的 `/db-tables` 还泄露全部表结构。
- **建议**: 路由补 `middleware.AuthCheckRole()`；路径字段白名单校验（拒绝 `..`、`/`、绝对路径）；`filepath.Clean` 后校验收敛在 `./app/` 与 `config.GenConfig.FrontPath` 内。

### C4. 数据权限中间件只注册未挂载（系统性越权）
- **位置**: `core/middleware/init.go:54`（仅 `SetMiddleware` 存储）、`core/middleware/permission.go:28-91`（`PermissionAction` 无任何路由使用）、`config/settings.yml:16`（`enableDP: false`）
- **问题**: 任何拥有某路由 casbin 权限的管理员可跨部门查看/修改全部 app 用户、账户流水、操作日志、内容、验证码记录；`Get/:id`、`PUT/:id`、`DELETE` 均无归属校验（IDOR）。
- **建议**: 路由组挂载 `PermissionAction()`；缺省策略改 fail-closed；统一 `create_by` 语义（当前存 app 用户 ID，与 admin 体系不一致）后启用 enableDP。

### C5. 日志中间件请求体截断导致小体积 POST/PUT 接口失效
- **位置**: `core/middleware/logger.go:29-42`
- **问题**: 用 `bufio.NewWriter` + `io.Copy` 拷贝请求体但**从未调用 `wt.Flush()`**，body < 4096 字节只进缓冲区不落盘 → `io.ReadAll(bf)` 读到空内容 → `c.Request.Body` 被替换为空 body → 登录等所有小请求体接口业务层拿到空数据（登录表现为"用户不存在"）。同时 `io.ReadAll` 无大小上限，是内存 DoS 向量。
- **建议**: 删除 bufio 包装直接 `io.Copy(bf, c.Request.Body)`，或 `defer wt.Flush()`；对 body 加大小上限（如 `io.LimitReader` 1MB）。

### C6. 任意文件上传 + 上传目录无鉴权静态服务（存储型 XSS）
- **位置**: `app/plugins/filemgr/apis/filemgr_app.go:174-206`、`service/filemgr_app.go:261-276`、`app/admin/sys/router/router.go:80`（`r.Static(FileRootPath, ...)` 免鉴权）、`app/admin/sys/apis/sys_user.go:327-348`（头像上传）
- **问题**: 上传无扩展名/MIME/魔数/大小校验，可上传 `.html/.svg/.js` → 管理后台**同源**存储型 XSS 窃取 JWT；静态目录免鉴权且开启目录列举，可批量下载全部文件；头像上传 `form, _ := c.MultipartForm()` 忽略错误 → nil 指针 panic；无大小限制 → 磁盘/内存 DoS。
- **建议**: 扩展名+MIME 白名单 + `http.MaxBytesReader` 限大小 + 落盘后验魔数；文件名服务端生成；静态目录仅暴露必要子目录并加 `Content-Disposition: attachment`；上传接口加鉴权。

### C7. 客户端 IP 可完全伪造 → 限流/黑名单绕过 + 定向 DoS
- **位置**: `core/utils/iputils/ip.go:63-80`、`core/middleware/limiter.go:52-93`、`core/cmd/api/server.go`（gin 未调用 `SetTrustedProxies`，默认信任所有代理）
- **问题**: `GetClientIP` 最终值完全由攻击者控制的 `X-Forwarded-For`/`X-real-ip` 决定：① 轮换 XFF 头即可无限绕过 IP 限流与黑名单；② 可伪造受害者 IP 耗尽对方配额实现定向 DoS；③ 设备指纹与登录日志 IP 全部失真。
- **建议**: 显式 `engine.SetTrustedProxies([]string{"<nginx/网关IP>"})`；`GetClientIP` 直接使用配置好可信代理后的 `c.ClientIP()`，删除手工拼接 XFF 逻辑。

### C8. 内存队列消费者自死锁 + Append 无界 goroutine
- **位置**: `core/utils/storage/queue/memory.go:37-61, 80-89`
- **问题**: ① 消费循环出错时 `out <- message` 向自己正在读取的同一 channel 发送 → **永久自死锁**，登录日志/操作日志流从此停止；② `Append` 每消息启动一个 goroutine 阻塞发送，队列满时 goroutine 无限堆积 → 内存耗尽。
- **建议**: 失败消息带重试上限或丢弃+告警；`Append` 改非阻塞 `select`/带超时发送。

### C9. redis 队列消费者 Errors 通道无人消费 → 队列停摆
- **位置**: `core/utils/storage/queue/redisqueue/consumer.go:139,188,268,354,391`、`core/utils/storage/queue/redis.go:61-77`
- **问题**: 注释要求"必须有监听者"，但项目没有任何地方读取 `Errors`；consumer 出错时 `c.Errors <- ...` 无缓冲通道永久阻塞该 worker → 队列整体停摆。
- **建议**: 在 `queue/redis.go` 的 `Run()` 中启动 goroutine 消费 Errors 并打日志。

### C10. 操作日志明文存储密码与 Token
- **位置**: `core/middleware/logger.go:88-89,106-109`（`operParam`/`jsonResult` 原样入库）、`app/admin/sys/apis/sys_user.go:505-541`
- **问题**: `POST /v1/login` 的 `{username, password}` 明文和响应中的 JWT token 原样写入 `admin_sys_oper_log`；文件日志 trace 级全量打印。日志库泄露 = 密码/token 批量泄露。
- **建议**: 日志入库前脱敏（正则过滤 password/token/captcha 字段）；对登录等敏感接口不记录 body。

### C11. 用户禁用/删除/重置密码后旧 JWT 仍然有效
- **位置**: `app/admin/sys/service/sys_user.go:345-427`、`core/middleware/auth/jwtauth/jwtauth.go:504-605`
- **问题**: `authCheck` 不查询用户当前状态与存在性：用户被禁用、删除或密码被重置后，已签发 token 在有效期内（最长 2h+刷新窗口）仍可访问全部接口。管理手段无法回收被盗 token。
- **建议**: `authCheck` 增加用户状态/存在性校验（缓存+DB 兜底）；禁用/删除/重置密码后调用 revokeAllTokens。

### C12. WS 管理器 map 并发读写 + send on closed channel
- **位置**: `core/ws/ws.go:107-151,155-195`（Start 加锁写 Group，三个 Send 服务无锁读）、`:131`（UnRegister `close(mClient.Message)`）
- **问题**: ① Send 系列与 Start 并发读写 `manager.Group` → 运行时 "concurrent map read and map write" panic；② UnRegister 关闭通道的同时 Send 可能正在发送 → "send on closed channel" panic。当前 `WsClient` 未注册路由属死代码，一旦启用即触发。
- **建议**: 读写统一加锁或对 Group 做快照拷贝再发送；关闭通道前先删除 map 条目并加锁保护发送路径。

### C13. 数据库 DSN 密码明文打印日志 + 配置弱口令
- **位置**: `core/storage/database/initialize.go:25`（`log.Infof("%s => %s", host, c.Source)` 打印含密码的完整 DSN）、`config/settings.yml:55`（`password=123456`、`sslmode=disable`、postgres 超级用户）
- **问题**: 启动日志明文输出数据库口令；配置以弱密码+禁用 TLS 提交仓库，生产沿用即裸奔。
- **建议**: 日志仅打印 host/库名；DSN 走环境变量注入，创建最小权限账号，启用 `sslmode=require`。

### C14. 优雅关闭不完整：SIGTERM 无法触发、队列不关闭
- **位置**: `core/cmd/api/server.go:134-135`（仅 `signal.Notify(quit, os.Interrupt)`）、`core/utils/storage/queue/redisqueue/signals.go:12-24`
- **问题**: ① 只监听 SIGINT，SIGTERM（docker stop/k8s 默认）到来不执行 `srv.Shutdown`；② consumer 内部重复注册信号处理且需二次信号才 `os.Exit(1)` → 容器停止超时悬挂；③ 关闭流程从不调用队列 `Shutdown()` → goroutine 泄漏。
- **建议**: 同时监听 `syscall.SIGTERM`；统一信号管理；关闭时依次 Shutdown 队列、DB、Redis。

---

## 二、高 (High)

### H1. casbin 权限回收不完整（删除菜单/API 时残留策略）
- **位置**: `app/admin/sys/service/sys_menu.go:395-443`（菜单 Delete 不清理 casbin）、`app/admin/sys/service/sys_api.go:137-161`、`app/admin/sys/models/sys_api.go:105-124`（Sync 不清理）
- **问题**: 删除菜单/按钮/API、同步删除失效接口时均不删除 `admin_sys_casbin_rule` 对应策略 → 已删除接口的权限持续有效（越权残留）。
- **建议**: 删除前查出关联角色的 (path, method) 策略一并删除，再 `mycasbin.LoadPolicy`。

### H2. 菜单更新 casbin 同步在循环内 return（只更新第一个角色）
- **位置**: `app/admin/sys/service/sys_menu.go:366-392`
- **问题**: `for _, role := range data.SysRole { ... return roleService.UpdateCasbin(...) }` 只处理第一个角色，其余角色权限残留或缺失。
- **建议**: 收集所有角色统一更新后再返回。

### H3. 创建/更新用户可指定任意角色（垂直越权提权）
- **位置**: `app/admin/sys/service/sys_user.go:112-213,280-283`、`dto/sys_user.go:62`（roleId 用户可控）
- **问题**: 有 `sys-user:add/edit` 权限的普通管理员可创建「admin 角色」用户或把数据权限范围内的用户提权为 admin。
- **建议**: 校验 RoleId 存在且未停用；禁止分配 admin 角色（除非操作者即 admin）；校验 DeptId 在操作者数据权限范围内。

### H4. DataScope4（本部门及以下）SQL 引用不存在的列
- **位置**: `core/middleware/permission.go:83`（`"admin_sys_dept.dept_path like ?"`）、`app/admin/sys/models/sys_dept.go:8`（实际列名为 `parent_ids`）
- **问题**: 该数据权限作用域用户的所有列表查询报 `Unknown column`，功能完全失效。
- **建议**: 改为 `parent_ids like ?`（参数 `%/deptId/%`）。

### H5. UpdateDataScope 无权限校验可修改任意角色数据权限（提权）
- **位置**: `app/admin/sys/service/sys_role.go:401-463`
- **问题**: 直接按 id 取角色修改 `data_scope`，无范围校验、无 admin 保护 → 可将任意角色改为"全部数据"。
- **建议**: 增加角色管辖范围校验；禁止修改 admin 角色；修改后强制相关用户重新登录。

### H6. 角色 key 重命名导致权限丢失 / 禁用角色不生效
- **位置**: `app/admin/sys/service/sys_role.go:263-305`（Update 允许改 role_key 但 UpdateCasbin 用旧 key）、`:466-487`（UpdateStatus 无 admin 保护、不 LoadPolicy）、`jwtauth.go:127-170`（AuthCheckRole 不看角色状态）
- **问题**: ① 重命名 role_key 后策略挂旧 key，新 key 用户全 403；② 禁用角色后其用户 token 与 casbin 策略均不变 → 禁用无效；③ admin 角色可被停用。
- **建议**: role_key 禁止修改；停用时 revokeAllTokens 该角色下全部用户并 LoadPolicy；补 admin 保护。

### H7. 角色降权后旧 token 在缓存过期窗口内恢复生效
- **位置**: `core/middleware/auth/jwtauth/jwtauth.go:518-526`
- **问题**: 角色变更写入 `JwtRolePrefix` 缓存（TTL=7 天），缓存过期后旧 token 里的角色声明重新生效；正常用户还会被锁 7 天（可用性缺陷）。
- **建议**: 用"用户角色版本号/更新时间"比对替代可过期缓存；关键操作实时查库校验角色。

### H8. 设备管理并发竞态：authCheck 写设备列表不加锁
- **位置**: `core/middleware/auth/jwtauth/jwtauth.go:557-600`（无锁读改写）对比 `:396-460`（recordDevice 有锁）
- **问题**: 并发请求可互相覆盖设备列表、突破 maxDeviceCount 上限；Redis 场景读改写非原子。
- **建议**: 设备追加统一走 recordDevice；Redis 场景用 Lua 脚本原子化。

### H9. 头像上传 nil 指针 panic + 无大小限制
- **位置**: `app/admin/sys/apis/sys_user.go:327-348`
- **问题**: 非 multipart 请求时 `form` 为 nil → `form.File["avatar"]` panic（全局无 Recovery，连接被重置）；多文件覆盖写同一路径；无大小限制。
- **建议**: 检查 MultipartForm 错误与文件为空；`http.MaxBytesReader` 限大小；校验图片类型（image 解码）。

### H10. 导出接口公式注入（CSV/Excel Injection）
- **位置**: `app/admin/sys/service/sys_oper_log.go:94-100`、`sys_login_log.go:95-102`、`sys_config.go:272-273`、`sys_api.go:181-183`、`sys_post.go:277-278`、各 dict 导出、`app/plugins/content/service/content_article.go:238` 等
- **问题**: UA、URL、配置值等用户可控字段以 `=`/`+`/`-`/`@` 开头写入单元格，管理员打开导出文件触发公式执行（钓鱼/窃密）。
- **建议**: 封装单元格清洗函数，对 `=+-@` 开头内容前置 `'`。

### H11. 导出/列表 pageSize 无上限（资源耗尽 DoS）
- **位置**: `core/dto/search.go:43-51`（Paginate 不限制 pageSize）、所有 GetPage API
- **问题**: 登录用户传 `pageSize=1000000` 单请求拉全表；Export 在 `GetWithKeyInt` 失败返回 `-1` 时**不 return** 继续执行 → 全表导出（`sys_opera_log.go:150-153` 等 7 处）。
- **建议**: Paginate 内 clamp（1~100）；Export 出错立即 return，maxSize 下限保护。

### H12. 内存缓存 Increase/Decrease 用 RLock 做读改写 + 过期条目不清理
- **位置**: `core/utils/storage/cache/memory.go:153-183,203-226`
- **问题**: ① 并发计数丢失更新（数据竞态）；② 过期条目仅惰性删除，验证码等随机 key 永不清理 → 内存无限增长。
- **建议**: 读改写改写锁/原子 CAS；增加后台定时清理协程。

### H13. httpclient 关闭 TLS 校验 + 未受检类型断言
- **位置**: `core/utils/httpclient/httpclient.go:183`（`InsecureSkipVerify: true`）、`:103,172`（`h.Body.(map[string]string)` 未检查 ok）
- **问题**: 当前无调用方（死代码），一旦启用即中间人风险 + panic 路径。
- **建议**: 删除 InsecureSkipVerify；类型断言加 ok；若保留增加 URL 白名单防 SSRF。

### H14. OSS 上传实现缺陷
- **位置**: `core/utils/ossutils/oss.go:74-119`
- **问题**: ① `os.Open` 错误未检查 → nil panic；② 分片上传失败不 `AbortMultipartUpload` → 孤儿分片累积；③ 强制 `oss.ACLPublicRead` 公开读，配合签名下载方案自相矛盾。
- **建议**: 检查 err；失败路径 Abort；ACL 改 Private 统一走签名 URL。

### H15. 登录安全：用户名枚举 + admin 用户名绕过禁用状态
- **位置**: `app/admin/sys/service/sys_user.go:545-562`
- **问题**: ① 用户不存在/密码错误返回不同错误码 → 用户名枚举；② `if login.Username == constant.RoleKeyAdmin` 按用户名判断，被禁用的 admin 账号仍可登录。
- **建议**: 统一错误信息；禁用状态按用户 ID 判断；增加账号级失败锁定。

### H16. 核心业务表严重缺失索引（性能/DoS 风险）
- **位置**: `app_mysql.sql:1364-1390`（app_user 的 parent_id/mobile/email/ref_code 无索引无唯一约束）、`:1405-1421`（app_user_account_log 的 **user_id 无索引**）、`:837-854,1190-1208`（日志表 user_id/created_at 无索引）、`:291-308`（dict_data 的 dict_type 无索引）、`:1330-1350`（admin_sys_user 的 username 无唯一索引）
- **问题**: 数据量增长后分页/关联查询全表扫描；`username` 可重复使登录 `First()` 行为不确定；账变记录按用户查询是主路径却无索引。
- **建议**: `username` 加 UNIQUE；`app_user(parent_id)`、`app_user_account_log(user_id)`、`admin_sys_login_log(user_id, created_at)`、`admin_sys_oper_log(user_id, created_at)`、`admin_sys_dict_data(dict_type)` 建索引；pgsql 同步。

---

## 三、中 (Medium)

### M1. 全局无 gin.Recovery，CustomError 对未知 panic 重抛
- **位置**: `core/middleware/customerror.go:13-50`（default 分支 `panic(err)`）、`core/middleware/init.go`
- **问题**: 业务 panic（如 H9、登录 Role nil）无兜底 → 连接被静默断开、无错误日志/堆栈。
- **建议**: 注册 `gin.Recovery()`；CustomError 默认分支记录堆栈并返回 500 JSON。

### M2. 登录接口 Role nil 指针 panic
- **位置**: `app/admin/sys/apis/sys_user.go:538`（`userResp.Role.RoleKey`，Preload 失败/角色被删时 Role 为 nil）
- **建议**: 判空返回业务错误。

### M3. 验证码强度弱 + 登录爆破防护不足
- **位置**: `core/utils/captchautils/captcha.go:36-43`（4 位数字）、`app/admin/sys/apis/sys_user.go:523-528`（**仅非 dev 模式校验验证码**，而默认配置即 dev）
- **问题**: 4 位验证码 10000 组合可穷举；dev 模式默认跳过验证码；无账号失败锁定。
- **建议**: 验证码校验不区分环境；提升复杂度；按用户名+IP 双重限流；失败锁定。

### M4. 随机数使用 math/rand（可预测）
- **位置**: `core/utils/strutils/strutils.go:60-64`（`rand.New(rand.NewSource(time.Now().UnixNano()))` 生成验证码/邀请码）、`core/utils/encrypt/security.go:15-41`、`core/utils/idgen/idgen.go:12-14`
- **问题**: 时间种子可预测，配合枚举可撞验证码/邀请码。
- **建议**: 改用 `crypto/rand` 或 `math/rand/v2`。

### M5. CORS 全开 + 无 X-Frame-Options + JWT 支持 query 传参
- **位置**: `core/middleware/header.go:26-47`（`ACAO: *`，X-Frame-Options 注释掉）、`jwtauth.go:109`（`TokenLookup: "header: Authorization, query: token, cookie: jwt"`）
- **问题**: ① 管理后台可被 iframe 嵌入（点击劫持）；② token 经 `?token=` 传递会出现在访问日志、Referer、浏览器历史。
- **建议**: ACAO 改域名白名单；启用 `X-Frame-Options: DENY`；去掉 query 传 token。

### M6. Swagger 文档无条件暴露 + 静态目录开启目录列举
- **位置**: `core/cmd/api/server.go:174`（`/swagger/*any` 无认证）、`app/admin/sys/router/router.go:80`（gin Static 基于 http.FileServer 可列目录）
- **建议**: Swagger 仅 dev/test 注册；自定义 handler 禁目录列表。

### M7. 多库/多实例路由与 Redis 客户端共享失效
- **位置**: `core/runtime/application.go:51-78`（存在 `"*"` 键时无条件返回 `"*"` 实例，按 Host 分库永不生效）、`core/config/option_redis.go:20-25` + `cache.go:27-28` + `locker.go` + `queue.go:39-47`（cache/locker/queue 各自配置不同 Redis 时被静默忽略）
- **建议**: 先精确匹配 key 再回退 `"*"`；各组件持有独立客户端或校验配置一致性。

### M8. Runtime 容器 map 无锁返回引用（读竞态）
- **位置**: `core/runtime/application.go:44-48,66-68,127-129,183-187`
- **问题**: 加锁后直接返回内部 map 引用，调用方锁外迭代，配置热更新并发 Set 时数据竞态。
- **建议**: 返回深拷贝或仅提供原子读取接口。

### M9. 配置热更新无保护重建全部组件
- **位置**: `core/config/config.go:36-40`（`init()` 重跑 Logger/multiDatabase/callbacks，无锁无回滚）
- **问题**: 文件变更时同步重建 DB/缓存/队列/锁，重建期间在途请求可能拿到半初始化状态。
- **建议**: 热更新与初始化分离，配置校验通过后原子切换。

### M10. limiter 配置热更新不生效 + 黑名单全局切片并发隐患
- **位置**: `core/middleware/limiter.go:19,52-93,96-124`、`core/middleware/init.go:46`（仅启动时 `LoadBlacklist` 一次）
- **问题**: ① 配置框架的 OnChange 热重载不含限流/黑名单，运行期改配置完全不生效需重启；② `blacklistIPs` 全局可变切片，`LoadBlacklist` 为导出函数，一旦运行期重载即与每请求遍历产生数据竞态；③ Redis store 初始化失败 `log.Fatalf` 直接崩进程（应降级内存）；④ `Retry-After: 60` 硬编码与配置 Period 不符。
- **建议**: `atomic.Pointer`/RWMutex + "先构建后替换"；注册到 OnChange 回调实现热更新；失败降级内存 store；Retry-After 用实际 Period。

### M11. 登录失败无审计日志 + GetLocation 同步外呼无超时
- **位置**: `app/admin/sys/apis/sys_user.go:530-540`、`service/sys_user.go:638`、`core/utils/iputils/ip.go:14-37`
- **问题**: ① 登录失败不落库，无法审计暴力破解；② 每次登录同步调用高德 API（`http.Get` 无超时），上游不可达时拖垮登录。
- **建议**: 失败也记录；`http.Client{Timeout: 2s}` + 异步化/缓存。

### M12. 队列配置缺省校验缺失（nil 解引用）
- **位置**: `core/config/queue.go:36-49`（`e.Redis.Consumer.ReclaimInterval` 等未判 nil）
- **建议**: 初始化时判空给默认值或报错。

### M13. 动态搜索条件反射解析存在 panic 路径
- **位置**: `core/dto/search/query.go:37,95`（无 search tag 的基础类型字段递归 `Field(i)` panic；`IsNil` 对非可空类型 panic）
- **建议**: 递归前判断 Kind；IsNil 前判断可空类型；或捕获 panic 返回空条件。

### M14. i18n CSV 缺行长度校验
- **位置**: `core/lang/i18n.go:66-69`（`row[0]`/`row[1]` 直接下标访问，空行/单列即 panic；初始化失败仅 Warn 静默降级）
- **建议**: 访问前 `len(row) >= 2`；失败升级为 error。

### M15. 验证码明文存储并回显管理端
- **位置**: `app/plugins/msg/models/msg_code.go:10`、`app/plugins/msg/apis/msg_code.go:30-80`（plugins_msg_code.code 明文，列表/详情直接返回）
- **建议**: 仅存哈希/密文；列表详情不返回 code 字段。

### M16. "先查后插"非原子并发竞态（重复数据）
- **位置**: `app/app/user/service/user.go:275-289,326-335`、`user_level.go:117-126`、`user_country_code.go:115-135`、`content_category.go:111-119`、`content_announcement.go:157-168`
- **问题**: 手机号/邮箱/邀请码/名称全部先 Count 再 Create，无唯一索引兜底 → 并发下重复数据。
- **建议**: 唯一索引 + 冲突错误处理；邀请码冲突循环重试。

### M17. UserConf.Update 可通过请求体禁用任意用户（逻辑型 IDOR）
- **位置**: `app/app/user/service/user_conf.go:206-226`（`c.UserId` 取自请求体，未校验与路径 `:id` 记录一致）
- **建议**: `c.UserId` 服务端从 `data.UserId` 取值并强制一致。

### M18. 系统监控接口错误全部忽略 → nil 解引用 panic + 敏感信息
- **位置**: `app/admin/sys/apis/sys_monitor.go:45-112`（`dis.Usage`/`mem.VirtualMemory` 失败返回 nil 后直接取字段 panic；未走 MakeContext 导致 e.Lang 为空；暴露主机名/本机 IP）
- **建议**: 判 err；接口限 admin 权限。

### M19. 导出/列表按用户字段搜索必报 SQL 错（同表双 JOIN）
- **位置**: `app/app/user/service/user_conf.go:52`、`user_account_log.go:52`（`Joins("User")`）与 `app/app/common/dto/join.go:3-8`（`inner join app_user` 标签）叠加 → PostgreSQL duplicate table 报错 500。
- **建议**: 二选一，联表统一走 Preload。

---

## 四、低 (Low)

| # | 位置 | 问题 |
|---|------|------|
| L1 | `core/utils/encrypt/aes.go:10,26-50` | AES 采用 **ECB 模式**（无 IV，同明文同密文可碰撞比对）、填充剥离逻辑只判断字节值 ≤16 不校验 PKCS7 一致性（可能误剥数据）、`hex.DecodeString` 错误被忽略；多处调用方忽略加密错误静默存空 |
| L2 | `core/middleware/auth/jwtauth/jwtauth.go:382` | `strings.SplitN(message, "_", 1)` n=1 不切分，错误码解析逻辑永远不生效，客户端收到带 `401_` 前缀的消息 |
| L3 | `core/utils/fileutils/file.go:31-35` | `IsFileExist` 用 `os.IsExist(err)`，Stat 成功时恒返回 false（逻辑反转） |
| L4 | `core/utils/fileutils/file.go:77-98,213-219` | `GetType` 打开文件后不 Close（FD 泄漏）；`GetFileSize` Walk 出错时 f 为 nil → panic |
| L5 | `core/storage/database/initialize.go:50` | `fmt.Sprintf(c.Driver+" connect error :", err)` 无占位符 → 输出 `%!(EXTRA ...)` |
| L6 | `core/middleware/logger.go:64-66` | `st.(int)` 类型断言无 ok 检查，非 int 即 panic |
| L7 | `core/utils/iputils/ip.go:67` | `strings.Contains(ip, "127.0.0.1")` 子串判断脆弱（`10.127.0.0.1` 误判） |
| L8 | `core/runtime/application.go` 返回 map 引用；`core/runtime/cache.go:21` `wxTokenStoreKey` 未使用 | 见 M8；死代码 |
| L9 | `core/middleware/auth/jwtauth/jwtauth.go:234-283` | 多处缓存 Set 错误被 `_ =` 忽略（黑名单写入失败静默） |
| L10 | `app/admin/sys/service/sys_user.go:359` | **UpdateStatus 用 `u.Avatar != c.Status` 比较**（笔误，应为 `u.Status`），状态更新条件完全失效 |
| L11 | `app/admin/sys/service/sys_user.go:387` | ResetPwd 的 `u.Password != c.Password` 用 bcrypt 哈希与明文比较恒为 true，"未变更跳过"逻辑永不生效 |
| L12 | `app/admin/sys/service/sys_user.go:655-658` | 登录日志记录成功（err==nil）后仍走 `Log.Errorf`，每次登录打一条假 Error 日志 |
| L13 | `app/admin/sys/service/sys_post.go:236-247` | 删除岗位引用检查用 `userReq.RoleId = id`（应为 PostId），引用检查失效 |
| L14 | `app/admin/sys/service/sys_dept.go:160-248` | 部门移动节点无循环检测（可成环）；删除不检查用户/角色引用（悬空） |
| L15 | `core/dto/api/api.go:39-43` | `api.Lang` 从未赋值（getAcceptLanguage 定义了未调用），多语言错误消息失效 |
| L16 | `core/dto/api/binding.go` | `setBinding` 从未被调用，每请求重复反射解析 DTO（性能） |
| L17 | `core/middleware/trace.go:36` | `ctx.Set("traceSpan", span)` 无处消费，与 OTel 标准做法重复 |
| L18 | `core/ws/ws.go:274-281,300-302` | `CheckOrigin` 恒 true 且无鉴权（CSWSH 隐患）；WsClient 内 `time.Sleep(15s)` + 硬编码路径 `tmp/logs/job/db-20200820.log`（必然不存在） |
| L19 | `core/casbin/mycasbin.go:45,49-56` | `EnableLog(true)` 每请求 Enforce 打日志（性能）；`LoadPolicy` 失败返回 nil 被所有调用点忽略（权限变更静默失效） |
| L20 | `core/casbin/adapter.go:248-268` | casbin 表唯一索引创建非原子，多实例同时启动可能 panic |
| L21 | `core/middleware/auth/jwtauth/jwtauth.go:37-59` | deviceLockManager 按 userID 建 mutex 永不删除，内存缓慢增长 |
| L22 | `core/dto/response/return.go:13-20` | 业务错误码 ≤600 直接映射为 HTTP 状态码，语义混淆 |
| L23 | `core/utils/captchautils/store.go:17` | 注释引用 `SetCustomStore` 与实际 `SetStore` 不一致 |
| L24 | `app/admin/sys/service/sys_user.go:387`、`app/admin/sys/service/sys_gen_table.go:126-131,580-587` | gen 事务内子 service 用独立连接（事务形同虚设）；`e.Orm = e.Orm.Begin()` 写回共享成员属脆弱写法（app 模块同款） |
| L25 | `app/plugins/filemgr/service/filemgr_app.go:164-170,241-253` | 下载地址拼接无协议校验（可存 `javascript:` 链接）；删除先删 OSS 后删 DB（DB 失败则文件已丢） |
| L26 | SQL 脚本 | ① `admin_sys_role_dept` 的 role_id/dept_id 为 smallint 与主表 int 不一致且无外键索引；② `SET FOREIGN_KEY_CHECKS=0` 全局关闭外键；③ dict 表字段名用保留字 `default`；④ admin_sys_user 的 `salt` 列永不使用（bcrypt 自带盐）；⑤ 种子数据含真实数据哈希（手机号 MD5 极弱可还原） |
| L27 | `go.mod` | WAF 依赖 `wprimadi/brandy v1.0.1` 小众低维护；`aliyun-oss-go-sdk v3.0.2+incompatible` 过旧（2020 年） |
| L28 | `.gitignore:21` | **go.sum 被忽略未提交**，破坏构建可复现性与供应链完整性（Go 官方要求必须提交） |
| L29 | `config/settings.yml:6,29` | 默认 `mode: dev` + `level: trace`，生产误用泄露 SQL/参数/堆栈 |
| L30 | `README.md:24,76,99,111` | 引用不存在的 logo.png；称"不提供 swagger"与事实不符；Go 版本写 1.25.4（实际 1.26.5）；目录写 `conf/`（实际 `config/`） |
| L31 | `docs/` | swagger 生成物未纳入 git（易过期无法回滚） |
| L32 | `rulesets/default.conf:8,176` | WAF 直接开启拦截模式（未经 DetectionOnly 试运行，CRS 对中文业务误报可能阻断功能）；`SecDataDir /tmp/` 共享目录不安全；`SecRequestBodyLimit` 12.5MB 偏大；multipart 严格校验可能拦截正常上传 |

---

## 五、已核查确认安全/无明显问题

- **SQL 注入**: 搜索框架 `core/dto/search/query.go:98-102` 的 where 全部参数化、order 方向受白名单（desc/asc）限制、列名来自 DTO struct tag（受信任），未发现可利用注入点。
- **密码存储（admin 侧）**: `app/admin/sys/models/sys_user.go:38-62` BeforeCreate/BeforeUpdate 均触发 bcrypt，密码为 bcrypt 哈希且 `json:"-"` 不泄露。
- **验证码一次性**: 登录时 `clear=true` 删除，不可重放。
- **JWT 流程**: 登录 → 验证码 → bcrypt 校验 → claims 注入 → 签发，单点登录/黑名单/设备数限制均有实现。
- **static/ 与 files/ 目录**: 内容干净，无密钥/日志/上传文件被提交。
- **SQL 脚本**: 纯静态 DDL/DML，无动态拼接用户输入。

---

## 六、修复优先级建议

**第一优先级（安全，立即处理）**
1. C1 默认密钥硬编码 → 环境变量注入 + 强制校验 + 密钥分离
2. C2 app 用户密码明文存储 → bcrypt 哈希 + DTO 脱敏
3. C3 代码生成任意文件写入 → 权限收敛 + 路径白名单
4. C5 请求体截断（功能故障，登录不可用）→ 修 bufio Flush
5. C7 客户端 IP 可伪造 → SetTrustedProxies
6. C6 文件上传链路 → 类型/大小校验 + 静态目录鉴权
7. C4 数据权限挂载 → 全局 PermissionAction

**第二优先级（高可用/权限正确性）**
8. C8/C9 队列死锁/停摆 → 修复消息回投与 Errors 消费
9. C11/C14 失效控制与优雅关闭 → revokeAllTokens + SIGTERM
10. H1-H8 权限回收/提权/数据权限故障
11. C10 操作日志脱敏
12. H16 SQL 索引补齐

**第三优先级（治理与规范）**
13. M 级全部（验证码、CORS、随机数、热更新、日志脱敏）
14. L 级清理死代码、修文档、提交 go.sum
15. 建立 CI（lint + 单元测试 + 依赖漏洞扫描 govulncheck）与部署规范（Dockerfile、密钥管理、防火墙策略）

---

*报告生成方式: 静态代码审查（Gin 中间件链、casbin 权限流、数据权限、队列、缓存、配置加载、SQL 脚本等交叉验证），部分行号可能随代码变动略有偏移。*


# 补充

部署提醒: 如果你在 nginx 后面部署,请:
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
并配置 trustedProxies: [nginx的IP]。若 nginx 不覆盖 XFF 头,客户端仍可注入伪造值,所以 nginx 必须使用 $proxy_add_x_forwarded_for(它会追加客户端真实 IP,并丢弃客户端伪造部分)。
默认行为变化: 不配置 trustedProxies 时,直连访问的 ClientIP() = TCP 对端 IP,不可伪造;登录日志的 IP 定位(GetLocation)与之前一致,只是 IP 来源更可靠。

# 修复进度

## Critical
- C1 ✅ 密钥硬编码：JWT/AES 密钥启动校验（拒绝占位符/旧默认值/共用），配置随机化；`config/` 已停止 git 追踪；DSN 密码环境变量注入 `${DB_PASSWORD}` + 弱口令/空密码启动拒绝 + 日志脱敏
- C2 ✅ app 用户密码：`pwd`/`pay_pwd` 改 bcrypt 哈希（`app/app/user/models/user.go` BeforeCreate/BeforeUpdate + CheckPwd 兼容），字段 `json:"-"`（commit 42784cc）
- C3 ✅ 代码生成任意写：`sys-table` 路由挂 `AdminOnly()`；gen 路径字段白名单正则校验
- C4 ⚠️ 部分：新增 `middleware.AdminOnly()`（仅 admin 角色）并挂载到敏感路由；数据权限 `PermissionAction()` 全量挂载待 create_by 语义统一后启用（enableDP 仍为 false）
- C5 ✅ 日志中间件：删除 bufio（原未 Flush 导致小 body 接口失效），`io.LimitReader` 1MB 上限
- C6 ✅ 上传链路：filemgr_app 扩展名白名单 + 200MB 上限；头像 2MB + 魔数校验 + nil 检查；静态目录自实现路由（禁目录列举、防路径穿越、危险类型强制下载）
- C7 ✅ IP 伪造：显式 `SetTrustedProxies`（未配置不信任任何代理），`GetClientIP` 直接用 `c.ClientIP()`，删除手工拼接 XFF
- C8 ✅ 内存队列：消费失败本地重试 3 次后丢弃（消除向自身 channel 回投的自死锁/忙等）；Append 有界等待入队（100ms 超时丢弃，防 goroutine 堆积）
- C9 ✅ redis 队列：`Run()` 启动 goroutine 消费 `consumer.Errors` 并记日志（防无缓冲通道阻塞停摆）
- C10 ✅ 操作日志脱敏：JSON 递归 + 正则兜底过滤 password/token/captcha 等敏感字段后入库
- C11 ✅ 旧 JWT：`MarkPwdChanged(uid)` 统一失效机制（禁用/删除/重置密码/改密后全部旧 token 立即失效）；`checkUserStatus` 实时查库校验用户存在/状态/角色一致性
- C12 ✅ WS 管理器：Send 系列加锁快照再发送；`safeSend`（client 级互斥 + closed 标志）杜绝 send on closed channel / map 竞态
- C13 ✅ DSN 日志脱敏 + 弱口令校验（见 C1）
- C14 ✅ 优雅关闭：监听 SIGTERM + SIGINT；关闭流程执行队列 `Shutdown()`

## High
- H1 ✅ 删除 API/菜单时回收 casbin 策略（受影响角色重建 + 全局 enforcer 重载）
- H2 ✅ 菜单更新 casbin 循环内 return 修复（全部角色更新后才返回）
- H3 ✅ 创建/更新用户角色校验（`checkRoleAssignable`）+ 部门范围校验（`checkDeptAssignable`）
- H4 ✅ `dept_path` → `parent_ids`（Permission 中间件 + 用户查询 + DTO 标签）
- H5 ✅ UpdateDataScope：admin 角色保护 + 数据范围不得放宽校验 + 变更后 `MarkPwdChanged`
- H6 ✅ role_key 禁止重命名；admin 角色禁止停用；每请求实时校验角色启用状态
- H7 ✅ 角色降权即时生效：删除 JwtRolePrefix 过期缓存，改实时 DB join 比对
- H8 ✅ 设备列表写入加锁（`addDevice` 原子读改写）
- H9 ✅ 头像上传：大小限制 + nil 检查 + 魔数校验
- H10 ✅ 导出公式注入：`excelutils.SafeRow` 统一清洗（`=+-@` 前缀转义）
- H11 ✅ 导出/列表 pageSize：`GetPageSize` 钳制 + `PageSizeLimit` 服务端字段（客户端不可传）；15 个导出接口配置缺失兜底 1000、上限 10000、错误后 return
- H12 ✅ 内存缓存：Increase/Decrease/Expire/HashDel 写锁 + copy-on-write；后台协程 30s 清理过期条目
- H13 ✅ httpclient：删除 InsecureSkipVerify；类型断言加 ok 检查
- H14 ✅ OSS：os.Open 错误检查；分片失败/完成失败 `AbortMultipartUpload`；ACL 改 Private（配合签名 URL）
- H15 ✅ 登录：用户不存在/密码错误统一返回"用户名或密码错误"（防枚举）；admin 账号同受 status 校验（被禁用不可登录）
- H16 ✅ 索引：`admin_sys_user(username)` UNIQUE；`app_user(parent_id/mobile/email/ref_code)`、`app_user_account_log(user_id)`、`admin_sys_login_log(user_id, created_at)`、`admin_sys_oper_log(user_id, created_at)`、`admin_sys_dict_data(dict_type)` 索引（mysql+pgsql 同步）

## Medium（未处理）
M3-M4、M6-M19 未处理（验证码强度、随机数、Swagger 暴露、热更新、M19 同表双 JOIN 报错等）

## Medium
- M1 ✅ 全局 `gin.CustomRecovery` 兜底所有 panic（记录堆栈 + 统一 500 JSON）；`CustomError` 默认分支不再重抛，改日志 + 500 响应
- M2 ✅ `LoginVerify` 增加角色判空（`Role == nil || RoleKey == ""` 返回"该账户尚未分配角色"），消除登录接口 Role nil 指针 panic
- M5 ✅ CORS 改配置白名单 `corsOrigins`（请求 Origin 不在白名单不返回 ACAO，`Vary: Origin`；响应包 Download 的 ACAO `*` 一并移除）；启用 `X-Frame-Options: DENY`；`TokenLookup` 移除 query 传参（仅 header/cookie）
- M6 ✅ Swagger 仅 `mode: dev` 注册（其他环境 404）；`/static` 目录弃用 `gin.Static`（http.FileServer 可列目录），改与上传目录同款 `serveFileNoList`（禁目录列举、防路径穿越、危险类型强制下载）
- M8 ✅ Runtime 容器 map 返回副本：`GetDb`/`GetCasbin`/`GetMiddleware`/`GetHandler` 深拷贝 map（切片一并拷贝）、`GetHandlerPrefix`/`GetRouter` 返回切片副本；`GetDbByKey`/`GetCasbinKey`/`GetMiddlewareKey` 改用 RLock

## 待办
- C4 数据权限全量挂载（依赖 create_by 语义统一）
- M 级全部
- 存量库执行新增索引的 ALTER 语句（新装库直接跑 app_mysql.sql / app_pgsql.sql 即可）
- 清理 git 历史中的旧密钥（git filter-repo）并轮换线上密钥
- settings.yml.back 中硬编码了本地库口令（nVZypeJxyuXZ4J4J），仅限本地开发，勿用于生产