# novel 模块修复记录（第二轮）

- **时间**：2026-08-06
- **范围**：`docs/novel模块代码审查报告.md` 所列 15 项（P1×4 / P2×2 / P3×8）的代码修复
- **验证**：`go build ./...` ✅、`go vet ./app/app/novel/...` ✅、`gofmt` 无差异 ✅

---

# novel 模块变更记录（第四轮：后台读者/帖子管理）

- **时间**：2026-08-07
- **范围**：后台读者管理（禁用账户、限时禁发帖）+ 帖子管理（禁止访问/恢复）；配套前端页面与菜单种子
- **验证**：`go build ./...` ✅、`go generate ./...`（双 spec 重生成）✅、前端 `npx tsc --noEmit` ✅
- **相关前端文档**：`admin/docs/02-views-and-apis.md` §2.4 / §3.4（页面与 API 清单）

---

## 一、修改文件清单

| 文件 | 修改内容 |
|---|---|
| `models/novel_reader_profile.go` | 新增 `ban_post_until`（禁止发帖截止时间，空=未禁言）、`ban_reason`（禁言原因） |
| `service/dto/novel_admin.go` | 新增：`NovelUserQueryReq`/`NovelUserItem`（嵌入 profile + left join app_user 取 userName/status/mobile）/`NovelUserStatusReq`/`NovelBanPostReq`/`NovelPostAdminQueryReq`/`NovelPostAdminStatusReq` |
| `service/novel_admin.go` | 新增：`NovelUser`（GetPage/ChangeStatus/BanPost）与 `NovelPostAdmin`（GetPage/ChangeStatus） |
| `apis/admin/novel_admin.go` | 追加 `NovelUser`、`NovelPostAdmin` handler |
| `router/novel_admin_router.go` | 注册 `/app/novel/user`（GET、PUT `/:id/status`、PUT `/:id/ban-post`）、`/app/novel/post`（GET、PUT `/:id/status`），均 `Auth+AuthCheckRole` |
| `service/novel_post.go` | `Insert` 增加禁言校验：profile 存在且 `ban_post_until > now` 时拒绝发帖（41043） |
| `config/base/lang/app_novel.go` | 新增错误码 41043~41046 |
| `app_mysql.sql` / `app_pgsql.sql` | ①`app_novel_reader_profile` 建表语句补两列（PG 含 `ADD COLUMN IF NOT EXISTS` 迁移语句）；②菜单种子 id 135~149 |

## 二、后台管理接口清单（base `/admin-api/v1`，权限见菜单种子）

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/app/novel/user` | 读者分页：query `keyword`（用户名包含搜索，映射 app_user.user_name）、`status`（1-正常 2-禁用） |
| PUT | `/app/novel/user/{id}/status` | 启用/禁用账户：body `{status: "1"\|"2"}`；禁用后 `jwtauth` 实时校验拒绝该读者全部请求 |
| PUT | `/app/novel/user/{id}/ban-post` | 禁止发帖：body `{banUntil: RFC3339 或 null, reason≤255}`；资料不存在自动创建；`banUntil` 为空即解除 |
| GET | `/app/novel/post` | 帖子分页（含全部状态）：query `keyword`（标题包含搜索）、`status` |
| PUT | `/app/novel/post/{id}/status` | 禁止访问/恢复：body `{status: "1"\|"2"}`；禁止后读者端列表与详情均不可见（读者端仅查 status=正常） |

## 三、错误码变更清单（`config/base/lang/app_novel.go`）

| 码 | 常量 | 文案 |
|---|---|---|
| 41043 | `NovelPostBannedCode` | 您已被禁止发帖，请于解禁时间后再试 |
| 41044 | `NovelUserStatusErrCode` | 用户状态无效 |
| 41045 | `NovelBanTimeErrCode` | 禁言截止时间必须晚于当前时间 |
| 41046 | `NovelPostStatusErrCode` | 帖子状态无效 |

## 四、菜单种子（id 135~149，MySQL/PostgreSQL 已同步）

| id | 名称 | 类型 | 路径/权限 |
|---|---|---|---|
| 135 | 小说平台 | 目录 | `/app/novel` |
| 136 | 读者管理 | 菜单 | `/app/novel/novel-user` |
| 137 | 读者查询 | 按钮 | `app:novel-user:query` |
| 138 | 启禁用账户 | 按钮 | `app:novel-user:status` |
| 139 | 禁/解禁发帖 | 按钮 | `app:novel-user:ban-post` |
| 140 | 帖子管理 | 菜单 | `/app/novel/novel-post` |
| 141 | 帖子查询 | 按钮 | `app:novel-post:query` |
| 142 | 帖子状态 | 按钮 | `app:novel-post:status` |
| 143 | 反馈/投诉管理 | 菜单 | `/app/novel/novel-feedback` |
| 144 | 反馈查询 | 按钮 | `app:novel-feedback:query` |
| 145 | 反馈删除 | 按钮 | `app:novel-feedback:del` |
| 146 | 公告管理 | 菜单 | `/app/novel/novel-notice` |
| 147 | 公告查询 | 按钮 | `app:novel-notice:query` |
| 148 | 公告发布 | 按钮 | `app:novel-notice:add` |
| 149 | 公告删除 | 按钮 | `app:novel-notice:del` |

## 五、验证结果

- `go build ./...` ✅
- `go generate ./...`：主 spec 与 webapi spec 重生成，新增 `NovelUserStatusReq`/`NovelBanPostReq`/`NovelPostAdminStatusReq` DTO 已收录 ✅
- 前端 `npx tsc --noEmit` ✅（页面 `novel-user`/`novel-post` + API 封装，文档见 `admin/docs/02-views-and-apis.md`）

> 注意：存量库需执行建表语句同目录的 `ALTER TABLE` 迁移语句（MySQL 追加两列；PG 用 `ADD COLUMN IF NOT EXISTS`），并给角色分配 136~149 菜单。

---

# novel 模块变更记录（第三轮）

- **时间**：2026-08-07
- **范围**：① 公开读接口免登录改造；② 意见反馈/投诉 + 系统通知 + 后台公告广播 + 后台反馈管理；③ 独立 web-api swagger spec 落地
- **验证**：`go build ./...` ✅、`go vet ./app/app/novel/... ./gen/... ./core/cmd/api/...` ✅、`gofmt` 无差异 ✅、`go generate ./...`（含双 spec 生成）✅

---

## 一、修改文件清单

| 文件 | 修改内容 |
|---|---|
| `router/novel_post.go` | `GET /post/page`、`GET /post/:id` 移入 `routerNoCheckRole`（公开）；写操作留在 `routerCheckRole` |
| `router/novel_book_review.go` | `GET /book-review/page` 移入 `routerNoCheckRole`（公开） |
| `apis/novel_post.go` | 读 handler `uid, _, _ := auth.GetUserId(c)`，匿名可读 |
| `models/novel_feedback.go` | 新增（表 `app_novel_feedback`） |
| `models/novel_notification.go` | 新增（表 `app_novel_notification`，`user_id` 索引 + `source+notice_id` 复合索引） |
| `models/novel_notice.go` | 新增（表 `app_novel_notice`） |
| `service/dto/novel_feedback.go` | 新增（`NovelFeedbackInsertReq` 等） |
| `service/dto/novel_notification.go` | 新增（`NovelNotificationReadReq` 等） |
| `service/dto/novel_notice.go` | 新增（`NovelNoticeInsertReq` 等） |
| `service/novel_feedback.go` | 新增（提交/分页/删除，提交成功自动生成系统通知） |
| `service/novel_notification.go` | 新增（分页/未读数/标记已读/种子播种/广播写入） |
| `service/novel_notice.go` | 新增（发布/分页/删除，事务广播每批 500 条，级联删通知） |
| `apis/novel_feedback.go` | 新增（读者端 Feedback handler） |
| `apis/novel_notification.go` | 新增（读者端 Notification handler） |
| `apis/novel_admin.go` → `apis/admin/novel_admin.go` | 后台 `Notice`/`FeedbackAdmin` handler 移入 `apis/admin` 子包（包结构调整，供双 spec 区分） |
| `router/novel_feedback.go` | 新增（读者端 /web-api 路由） |
| `router/novel_notification.go` | 新增（读者端 /web-api 路由） |
| `router/novel_admin_router.go` | 新增（后台 /admin-api 路由：`/app/novel/notice`、`/app/novel/feedback`，`Auth+AuthCheckRole`） |
| `service/novel_auth.go` | 注册成功播种 3 条种子系统通知（欢迎/新功能/社区规范，`source=system`） |
| `config/base/lang/app_novel.go` | 新增错误码 41036~41042 |
| `app_mysql.sql` / `app_pgsql.sql` | 新增三张表 DDL + id 4~7 种子读者各 3 条种子通知 |
| `gen/webapidoc/main.go` | 新增：读者端独立 swagger spec 扫描入口（`@BasePath /web-api/v1`） |
| `core/cmd/api/server.go` | dev 模式注册 `/webapi/swagger/doc.json`（`swag.ReadDoc("webapi")`）+ `/webapi/swagger/*any` |
| `main.go` | go:generate 新增第二条 `swag init`（webapi spec） |
| `docs/` | 重生成主 spec；新增 `docs/webapi/` 独立 spec；需求文档/审查报告/LLM 说明同步 |

---

## 二、错误码变更清单（`config/base/lang/app_novel.go`）

| 码 | 常量 | 文案 |
|---|---|---|
| 41036 | `NovelFeedbackEmptyCode` | 反馈/投诉内容不能为空 |
| 41037 | `NovelFeedbackTooLongCode` | 反馈/投诉内容不能超过 500 字 |
| 41038 | `NovelFeedbackTypeErrCode` | 反馈类型不合法 |
| 41039 | `NovelNoticeTitleEmptyCode` | 公告标题不能为空 |
| 41040 | `NovelNoticeContentEmptyCode` | 公告内容不能为空 |
| 41041 | `NovelNoticeTitleTooLongCode` | 公告标题不能超过 100 字 |
| 41042 | `NovelNoticeContentTooLongCode` | 公告内容不能超过 2000 字 |

> 注：系统通知（`app_novel_notification`）未新增独立错误码，复用通用码（如 `DataQueryCode`/`ParamErrCode`）。

---

## 三、swagger 双 spec 说明

- **主 spec**（`docs/`，`@BasePath /admin-api/v1`，dev 访问 `/swagger/index.html`）：管理后台全量接口，含后台公告/反馈管理。
- **读者侧独立 spec**（`docs/webapi/`，`@BasePath /web-api/v1`，dev 访问 `/webapi/swagger/index.html`）：仅读者接口 27 条路径（登录/注册/书库/书评/长文/书架/关注/反馈/通知/资料），通过 `-t` 按读者侧 tags 白名单过滤，排除后台管理接口。
- 生成：`go generate`（`main.go` 中两条 `swag init`；webapi 侧 `-g main.go -d ./gen/webapidoc,./ -t 读者认证,读者资料,书友关注,我的读者资料,我的书架,系统通知,小说书库,小说书评,小说长文,意见反馈 -o docs/webapi --instanceName webapi`）。

---

## 四、验证结果

- `go build ./...` ✅
- `go vet ./app/app/novel/... ./gen/... ./core/cmd/api/...` ✅
- `gofmt -l`（novel + gen + core/cmd/api + main.go）无输出 ✅
- `go generate ./...` 复跑：主 spec 28 条 novel 路径（含后台管理）、webapi spec 27 条纯读者路径、零 admin/user/plugins 混入 ✅
- 新 handler（`Feedback`/`Notification`/`Notice`/`FeedbackAdmin`）已通过 `go generate` 收录进 `sys_api_gen_desc.go` 接口同步 ✅

---

## 一、修改文件清单

| 文件 | 修改内容 |
|---|---|
| `service/novel_book.go` | ①`GetPage` 补 `cDto.Paginate`；②关键字搜索拆驱动分支（PG `CAST(tags AS TEXT)`）；③`GetHome` 帖子 `Pagination{PageSize:5}`；④`Get` 点击当日去重；⑤`Insert/Update` 长度前置校验 |
| `service/novel_post.go` | ①`GetPage` 补 `cDto.Paginate`；②`Insert` 帖题长 200 校验；③`AddComment` 重构为事务（锁帖校验 + 父评论同帖校验 + 内容长度） |
| `service/novel_book_review.go` | ①`GetPage` 补 `cDto.Paginate`；②`Insert` 参数错误码 `DataQueryCode→ParamErrCode`；③书评长 1000 校验 |
| `service/novel_bookshelf.go` | `Insert` 并发唯一冲突转 `NovelShelfAlreadyExistCode`（幂等） |
| `service/novel_reader_profile.go` | ①`Get` 自动创建唯一冲突幂等重查；②`Update` 邮箱支持置空 + 昵称/简介/邮箱长度校验 |
| `service/novel_common.go` | 新增点击去重缓存前缀常量 `NovelClickCachePrefix` |
| `service/dto/novel_book.go` | 删除死代码 `NovelBookHomeReq` |
| `service/dto/novel_bookshelf.go` | 删除死代码 `NovelShelfGetReq` |
| `service/dto/novel_post.go` | 删除死代码 `NovelPostCommentGetReq` |
| `service/dto/novel_reader_profile.go` | 删除死代码 `NovelProfileGetReq` |
| `config/base/lang/app_novel.go` | 删除死码 41004/41011/41015，`41014` 转正复用；新增 41017-41028──见下方清单 |

---

## 二、逐项处理明细

### P1 严重

1. **书库列表分页失效**：`novel_book.go GetPage` 挂上 `Scopes(MakeCondition, Paginate(GetPageSize,GetPageIndex))`。
2. **帖子列表分页失效**：`novel_post.go GetPage` 同上。
3. **书评圈列表分页失效**：`novel_book_review.go GetPage` 同上。
4. **首页聚合全表返回**：`GetHome` 用 `cDto.Pagination{PageSize:5}` 限定首页 latest/hot 帖子各 5 条（`Pagination` 定义于 `core/dto`，别名为 `cDto`）。

### P2 应修复

- **2-1 PG 兼容**：关键字 `like` 中受 `tags`（JSON）影响，改为按驱动分支：
  - `postgres` → `CAST(tags AS TEXT) like`
  - 其它（MySQL/SQLite）→ `tags like`
- **2-2 帖子补充校验**：`AddComment` 重构为事务——
  `FOR UPDATE` 锁帖 → 校验存在且 `status=正常`（防孤儿评论与 `comment_count` 空增）→ 建评论 → `comment_count+1` 同事务提交。
- **2-3 复核「非缺陷」**：`gen_api_desc` 扫描所有含 `apis` 目录，已生成 `go-admin.Book.GetPage-fm` 等描述键；`SaveSysApi`（`app/admin/sys/models/sys_api.go`）启动时同步全部 gin 路由（含 `/web-api/v1/app/novel/**`，`apiType=app`）。仅剩后台角色绑定为运行期操作。

### P3 一般

- **3-1** 书评 `Insert` 对 `CurrUserId<=0` 改回 `ParamErrCode`（原误用 `DataQueryCode`）。
- **3-2** 读者资料邮箱置空：判定改 `c.Email != data.EmailValue`，`c.Email==""` 时写入空串。
- **3-3** 新增长度前置校验（校验 `len([]rune(...))`）：
  - 书名 100、作者 64、推荐语 30、封面/简介 500（`novel_book.go` Insert/Update）
  - 帖子标题 200（`novel_post.go` Insert）
  - 评论与书评内容 1000
  - 昵称 64、简介 500、邮箱 128（`novel_reader_profile.go` Update）
- **3-4** 楼中楼回复校验父评论存在且 `parent.PostId == post.Id`（新增码 41017）。
- **3-5** 唯一冲突兜底：`bookshelf.Insert` 接 `core/utils/dberr.IsDuplicateKey` → 转 41014 幂等返回既有记录；`reader_profile.Get` 自动创建并发 → 重查幂等返回。
- **3-6** 详情点击当日去重：缓存 `novelClick:book:{bookId}:{userId}` TTL 24h，命中不计数；匿名用户不计数；无缓存适配器时回退原计数行为。
- **3-7** 删除死代码：
  - 错误码 `NovelReviewNotExistCode(41004)`、`NovelProfileNotExistCode(41011)`、`NovelCommentDeleteErrCode(41015)`；`NovelShelfAlreadyExistCode(41014)` 转正复用。
  - DTO `NovelBookHomeReq`、`NovelShelfGetReq`、`NovelPostCommentGetReq`、`NovelProfileGetReq`。
- **3-8 复核「与基线一致」**：`e.Orm = baseOrm.Begin()` 交换写法与 `app/app/user/service/user.go:231` 一致，系仓库约定，保持现状。

---

## 三、错误码变更清单（`config/base/lang/app_novel.go`）

| 码 | 常量 | 文案 | 处置 |
|---|---|---|---|
| 41004 | `NovelReviewNotExistCode` | 书评不存在 | 删除 |
| 41011 | `NovelProfileNotExistCode` | 读者资料不存在 | 删除 |
| 41015 | `NovelCommentDeleteErrCode` | 评论删除失败 | 删除 |
| 41014 | `NovelShelfAlreadyExistCode` | 该书已在书架上 | 保留并转正复用（3-5） |
| 41017 | `NovelParentNotSamePostCode` | 父评论不属于该帖子 | 新增（3-4） |
| 41018 | `NovelBookTitleTooLongCode` | 书名不能超过 100 字 | 新增（3-3） |
| 41019 | `NovelBookAuthorTooLongCode` | 作者名不能超过 64 字 | 新增 |
| 41020 | `NovelBookSloganTooLongCode` | 推荐语不能超过 30 字 | 新增 |
| 41021 | `NovelBookCoverTooLongCode` | 封面地址不能超过 500 字 | 新增 |
| 41022 | `NovelBookDescTooLongCode` | 简介不能超过 500 字 | 新增 |
| 41023 | `NovelPostTitleTooLongCode` | 帖子标题不能超过 200 字 | 新增 |
| 41024 | `NovelContentTooLongCode` | 评论/书评内容不能超过 1000 字 | 新增 |
| 41025 | `NovelNicknameTooLongCode` | 昵称不能超过 64 字 | 新增 |
| 41026 | `NovelBioTooLongCode` | 简介不能超过 500 字 | 新增 |
| 41027 | `NovelEmailTooLongCode` | 邮箱不能超过 128 字 | 新增 |
| 41028 | `NovelPostCommentEmptyCode` | 评论内容不得为空 | 新增 |

---

## 四、验证结果

- `go build ./...` ✅
- `go vet ./app/app/novel/...` ✅
- `gofmt -l`（novel + lang）无输出 ✅
- `go run app/admin/sys/models/parseapi/gen_api_desc.go` 重新生成 `sys_api_gen_desc.go`，已含 `go-admin.Book.GetPage-fm` / `go-admin.Post.AddComment-fm` / `go-admin.Post.Interact-fm` / `go-admin.BookShelf.*` / `go-admin.ReaderProfile.*` 等 novel handler 描述键 ✅