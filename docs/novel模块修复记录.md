# novel 模块修复记录（第二轮）

- **时间**：2026-08-06
- **范围**：`docs/novel模块代码审查报告.md` 所列 15 项（P1×4 / P2×2 / P3×8）的代码修复
- **验证**：`go build ./...` ✅、`go vet ./app/app/novel/...` ✅、`gofmt` 无差异 ✅

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