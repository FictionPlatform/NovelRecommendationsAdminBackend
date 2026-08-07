# `app/app/novel` 代码审查报告（新增模块）

- **审查对象**：`app/app/novel/` 小说推荐平台后端模块（新增代码，首轮完整审查）
- **审查基准**：`docs/小说推荐平台需求文档.md`（§2.4 业务规则 / §4.1–4.5 后端方案）、`docs/LLM使用说明.md`（go-admin 开发规范）、既有 `app/app/user` 五层样板
- **审查范围**：`models/`(8) · `service/dto/`(6) · `service/`(7) · `apis/`(5) · `router/`(6) · `config/base/lang/app_novel.go` · `app_mysql.sql` / `app_pgsql.sql`（8 表 DDL + 4 组字典种子）
- **验证**：`go build ./...` ✅、`go vet ./app/app/novel/...` ✅、`go generate ./...` ✅、`gofmt` ✅
- **结论**：整体结构划分、业务规则落库（评分联动/赞踩互斥/级联删除/字数计算）与前后端契约对齐度高。**首轮存在 4 处"分页失效"类严重问题（P1）** 与若干兼容性/健壮性问题，已在本轮全部修复并经编译/vet 验证；见「修复核对记录」。

---

## 〇、修复核对记录（第二轮）

| # | 严重度 | 处理 | 状态 |
|---|:---:|---|---|
| 1-1~1-3 | P1 | 三个 `GetPage` 追加 `cDto.Paginate(c.GetPageSize(), c.GetPageIndex())`，保持 `.Limit(-1).Offset(-1).Count()` 尾链 | ✅ |
| 1-4 | P1 | `GetHome` 帖子通过 `Pagination{PageSize:5}` 显式限 5 条（latest/hot 各 5） | ✅ |
| 2-1 | P2 | 关键字搜索按驱动分支：PG 用 `CAST(tags AS TEXT)`，MySQL/SQLite 用 `tags like` | ✅ |
| 2-2 | P2 | `AddComment` 改事务：`FOR UPDATE` 锁帖校验存在+`status=正常`，评论与 `comment_count+1` 同事务 | ✅ |
| 2-3 | P2 | **复核为「非缺陷」**：`gen_api_desc` 扫描所有含 `apis` 目录（已生成本模块 handler 描述键），`SaveSysApi` 启动时同步**全部** gin 路由（含 `/web-api/v1/app/**`，`apiType=app`）。仅剩「后台菜单绑定角色权限」为运行期操作，无需改代码（见设计说明 5） | ✅ 复核 |
| 3-1 | P3 | 书评 `Insert` 参数错误由 `DataQueryCode` 改回 `ParamErrCode` | ✅ |
| 3-2 | P3 | 读者资料邮箱改为 `c.Email != data.EmailValue` 判定，支持置空 | ✅ |
| 3-3 | P3 | 新增长度前置校验（书名100/作者64/推荐语30/封面500/简介500/帖题200/评论与书评1000/昵称64/简介500/邮箱128），新增错误码 41018-41027 友好文案 | ✅ |
| 3-4 | P3 | 楼中楼回复校验父评论存在且 `parent.PostId == post.Id`，新增码 41017 | ✅ |
| 3-5 | P3 | 书架/读者资料自动创建并发唯一冲突接 `dberr.IsDuplicateKey`：书架转 `NovelShelfAlreadyExistCode`，资料重查幂等返回 | ✅ |
| 3-6 | P3 | 详情点击当日去重：缓存 `novelClick:book:{id}:{userId}` 24h，命中不计数；匿名不计数 | ✅ |
| 3-7 | P3 | 删除死码 `41004/41011/41015` 与 DTO `NovelBookHomeReq/NovelShelfGetReq/NovelPostCommentGetReq/NovelProfileGetReq`；`41014` 转正复用 | ✅ |
| 3-8 | P3 | 复核为「与基线一致」：`e.Orm = baseOrm.Begin()` 交换写法与 `app/app/user/service/user.go:231` 完全一致，为仓库既有约定，保持现状 | ✅ 复核 |

> 说明：构造首页限条数需用 `cDto.Pagination`（`Pagination` 定义于 `core/dto`，非 novel 的 service dto）。`GetPageIndex/GetPageSize` 默认 10、上限 100 的既有行为不变；点击去重缓存键 `novelClick:book:{bookId}:{userId}`，TTL 24h，匿名用户不计数。

---

## 一、问题清单

### P1 严重（需修复，影响功能正确性）

| # | 严重度 | 问题 | 位置 |
|---|:---:|---|---|
| 1-1 | P1 | **书库列表分页失效（全表返回）**：`GetPage` 未挂 `cDto.Paginate(c.GetPageSize(), c.GetPageIndex())`。框架的 `MakeCondition` 不会自动分页（`dto.Pagination` 被标 `search:"-"`），故 `pageIndex/pageSize` 形同虚设，`.Find(&list).Limit(-1).Offset(-1).Count()` 下 list 无任何 limit。对照 `app/app/user/service/user.go:67-72`（正确写法）。 | `service/novel_book.go:44-64` |
| 1-2 | **P1** | 帖子列表分页失效，同 1-1。 | `service/novel_post.go:58-66` |
| 1-3 | **P1** | 书评圈列表分页失效，同 1-1。 | `service/novel_book_review.go:35-44` |
| 1-4 | **P1** | 首页聚合 `GetHome` 拉取**全表帖子**注入最新/热门（依赖 1-2 的错误），每次首页请求全表扫描 + 大响应；修复 1-2 后仍应显式 `Limit(5~10)`。 | `service/novel_book.go:327-338` |

> **修复建议（1-1~1-3）**：在 `db := e.Orm.Model(&data)…` 后追加 `Scopes(cDto.MakeCondition(c.GetNeedSearch()), cDto.Paginate(c.GetPageSize(), c.GetPageIndex()))`，保持 `.Find(&list).Limit(-1).Offset(-1).Count(&count)` 尾链不变（与 user 模块一致）。`1-4` 增加 `.Limit(5)` 后再由 service 填充扩展。

### P2 应修复（兼容性 / 健壮性 / 权限）

| # | 严重度 | 问题 | 位置 |
|---|:---:|---|---|
| 2-1 | P2 | 关键字搜索 `LOWER(tags) like` 对 **PostgreSQL 不兼容**（JSON 类型不支持 `LOWER()`，MySQL 可隐式转字符串），PG 环境直接报错。建议 `LOWER(CAST(tags AS CHAR))`/`CAST(tags AS TEXT)`，或标签拆独立列分句。 | `service/novel_book.go:49` |
| 2-2 | P2 | `AddComment` **未校验目标帖子存在/状态**：被删帖或下架帖 id 仍可写入评论 → 孤儿评论 + `comment_count` 空增（静默）。应在事务内先锁定/校验 post 存在且 `status=正常`。 | `service/novel_post.go:377-414` |
| 2-3 | P2 | **内容管理权限集成缺口**：书架路由中 `POST/PUT/DELETE /app/novel/book` 挂了 `AuthCheckRole`，但 `go generate` 的接口同步仅扫描 `/admin-api/v1`，`sys_api` 不会收录任何 `/web-api/v1/app/novel/**` → 后台「接口管理→同步」无法给「编辑/读者」角色写入 Casbin 策略，**非 admin 角色将一律 403**。建议：扩展 `gen_api_desc` 扫描 web-api；或在 SQL 中手工补 `admin_sys_api`+casbin 数据；或短期切换为 `AdminOnly()`。 | `router/novel_book.go:17-28`、`core/…/gen_api_desc*` |

### P3 一般（健壮性 / 一致性 / 清理）

| # | 严重度 | 问题 | 位置 |
|---|:---:|---|---|
| 3-1 | P3 | 书评 `Insert` 对 `CurrUserId<=0` 误用 `DataQueryCode`（“数据查询失败”），应换 `ParamInvalidCode/DataErr`。 | `service/novel_book_review.go:72-73` |
| 3-2 | P3 | 读者资料邮箱无法**清空**：`if c.Email != ""` 判断使置空失效。 | `service/novel_reader_profile.go:96` |
| 3-3 | P3 | 各字段缺少长度前置校验（标题>200、作者>64、书名>100、评论>1000 等依赖 DB strict 报 500），应输出友好文案。 | `service/novel_post.go:122`、`service/novel_book.go:108-113`、`service/novel_post.go:381` |
| 3-4 | P3 | 楼中楼回复（`AddComment`）未校验「父评论属于同一帖子」，结构与数据一致性弱。 | `service/novel_post.go:385-391` |
| 3-5 | P3 | 书架/读者资料"自动创建"并发下依赖唯一索引兜底，冲突会返回通用 `DataInsertError`，建议捕获唯一冲突（`gorm.ErrDuplicatedKey`）转 `NovelShelfAlreadyExistCode`/幂等成功。 | `service/novel_bookshelf.go:48-74`、`service/novel_reader_profile.go:36-54` |
| 3-6 | P3 | `clicks +1` 无防刷（需求 §4.4 标注"可选"）；列表页刷新即刷量，生产应加用户/IP 当日去重。 | `service/novel_book.go:87` |
| 3-7 | P3 | 死代码：未使用错误码 `NovelShelfAlreadyExistCode / NovelCommentDeleteErrCode / NovelProfileNotExistCode / NovelReviewNotExistCode`；未使用 DTO `NovelBookHomeReq / NovelShelfGetReq / NovelPostCommentGetReq / NovelProfileGetReq`。保留则用于 3-5；否则删除。 | `config/base/lang/app_novel.go`、`service/dto/novel_*.go` |
| 3-8 | P3 | `book_review.Insert` 中"锁定书籍行 → create → 更新聚合"使用事务正确，但 `Interact`/`Insert` 中 `return` 前的 `data` / `txErr` 命名与既有模块异（`e.Orm = baseOrm.Begin()` 后再由 deferred `Commit/Rollback`，写法与 user 模块 L24 的 `tx` 用法不一致），建议统一从上下文读取事务对象，避免多实例共享 `Orm` 的并发隐患。 | `service/novel_post.go:237-248` 等 |

---

## 二、设计说明 / 已知待办（非缺陷，留档）

1. **删书评不回回溯书籍聚合**：`rating/review_count` 在删除书评后保持旧值。此为需求 §2.3 `handleDeleteUserReview`"不回溯聚合"的明确约定，属设计行为；但会导致评分与书评数失真，生产期可评估是否纳入"后台校正任务"。
2. **物理删除 vs 软删除**：帖子/书评/互动采用物理删除（需求 §2.4-6 允许软删除选项）。若需审计还原，可迁移至 `status=2` 软删除。
3. **Token 体系**：读者与 admin 共用同一套 JWT 签发（`app_user` 登录沿用 admin jwt-secret）。需求 §8 建议读者独立签发避免泄漏，当前 admin 侧均有 `AuthCheckRole`+Casbin 兜底（无角色 claim 的读者 token 无法访问 admin 业务接口），风险可控但建议评估独立 JWT。
4. **swagger 文档（已修复）**：原 `@Router` 均用 `/app/novel/*` 且 `@BasePath` 为 `/admin-api/v1`，new 接口落在管理后台 spec 下。现按需求 §8 落地「第二份独立 spec」：`gen/webapidoc/main.go` 作为扫描入口（`swag init -g main.go -d ./gen/webapidoc,./ -t <读者侧tags> -o docs/webapi --instanceName webapi`），`@BasePath /web-api/v1`，仅含读者侧接口；`core/cmd/api/server.go` dev 模式注册 `/webapi/swagger/doc.json`（`swag.ReadDoc("webapi")`）+ `/webapi/swagger/*any`。后台管理接口（公告/反馈管理）移至 `apis/admin` 包，只进主 spec（`/swagger/index.html`）。
5. **接口同步（已复核，非缺陷）**：`go generate` 生成的 `ApiDescMap` 已收录全部 novel handler（`go-admin.Book.GetPage-fm` 等）；`SaveSysApi`（`app/admin/sys/models/sys_api.go`）在服务启动时遍历 **gin 全部路由**（`core/runtime` 的 `GetRouter`），`/web-api/v1/app/novel/**` 会以 `apiType=app` 落入 `admin_sys_api`，无需手工 SQL。遗留仅为运行期操作：在后台「角色管理」给非 admin 角色绑定这些 app 接口（生成 casbin `p` 策略）后即可授权访问。

---

## 三、对照需求复查结论（§2.4 规则）

| 规则 | 实现核对 | 结论 |
|---|---|---|
| §2.4-1 书籍组合筛选（分类/连载/字数区间/关键字） | search 标签 + 关键字 `OR`（PG 已按驱动分支 `CAST(tags AS TEXT)`，见 2-1） | ✅ |
| §2.4-2 排序（rating / clicks / publish_date） | 榜单与列表 order | ✅ |
| §2.4-3 评分联动保留 1 位小数 | `avgRating` 事务 + `FOR UPDATE` | ✅ |
| §2.4-4 赞踩互斥 + 计数下限0 | 唯一索引 + 行锁 + `maxZero` | ✅ |
| §2.4-5 帖子字数/summary/read_time | `5000~10000` 校验、前150字、`ceil(字符/400)` | ✅ |
| §2.4-6 级联删除（书→书评/书架；帖子→互动/评论） | 事务级联实现 | ✅ |

---

## 四、下一步建议

1. ✅ 已修复 1-1~1-4（补 `cDto.Paginate` + `GetHome` 限 5 条）；建议补一条针对列表分页的单元测试（`pageSize/pageIndex` 生效断言）。
2. ✅ 已补 2-2 帖子存在性校验（事务内 `FOR UPDATE` 锁定）。
3. ✅ 联调验证：在 admin 后台给「编辑/读者」角色绑 `/web-api/v1/app/novel/**` 的 `sys_api` 数据（启动后已自动收录），确认 Casbin 通过（2026-08-07 复核：`go generate` 已收录新 handler `Feedback`/`Notification`/`Notice`/`FeedbackAdmin` 描述键）。
4. ✅ 已逐项处理 3-1~3-8 并清理死代码；`go build ./...`、`go vet ./app/app/novel/...` 通过。

---

## 五、2026-08-07 补充变更（免登录 + 反馈/通知/公告 + 双 swagger spec）

| 项 | 变更 |
|---|---|
| 免登录读接口 | `GET /app/novel/post/page`、`GET /app/novel/post/:id`、`GET /app/novel/book-review/page` 由「需登录」改为公开（`routerNoCheckRole`）；读 handler 内 `uid, _, _ := auth.GetUserId(c)` 匿名降级 |
| 意见反馈/投诉 | 新表 `app_novel_feedback`，`POST /app/novel/feedback`（kind=feedback/complaint，type A~H）+ `GET /app/novel/feedback/page`；错误码 41036~41038 |
| 系统通知 | 新表 `app_novel_notification`（user_id 索引 + source+notice_id 复合索引），`GET /notification/page`、`GET /notification/unread-count`、`POST /notification/read`；注册时播种 3 条种子通知；错误码 41039~41041 |
| 后台公告广播 | 新表 `app_novel_notice`，`POST/GET/DELETE /admin-api/v1/app/novel/notice`（发布/分页/删除，事务广播每批 500 条，级联删读者通知）；错误码 41042 |
| 后台反馈管理 | `GET/DELETE /admin-api/v1/app/novel/feedback` |
| 包结构调整 | 后台 handler（`Notice`/`FeedbackAdmin`）移入 `apis/admin` 子包，路由挂 `adminRouterCheckRole`（`Auth+AuthCheckRole`） |
| 双 swagger spec | 读者侧独立 spec（`docs/webapi/`，`@BasePath /web-api/v1`）+ 管理后台主 spec；生成入口 `gen/webapidoc/main.go`（`swag init -d ./gen/webapidoc,./ -t <读者tags>`）；dev 模式 `/webapi/swagger/index.html` |
| 验证 | `go build ./...`、`go vet`、`gofmt`、`go generate ./...` 全通过；webapi spec 27 条纯读者路径、主 spec 28 条 novel 路径（含后台管理接口） |