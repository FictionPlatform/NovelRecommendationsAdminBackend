package service

import (
	"fmt"
	"os"
	"testing"
	"time"

	"go-admin/app/app/novel/models"
	"go-admin/app/app/novel/service/dto"
	userModels "go-admin/app/app/user/models"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// novelTestDB 集成测试专用连接（独立测试库 novel_test，不影响 public 生产库）
var novelTestDB *gorm.DB

// novelTestDSN 与 config/settings.yml database.source 一致（仅库名不同）
const novelTestDSN = "host=127.0.0.1 port=5432 user=postgres dbname=novel_test password=nVZypeJxyuXZ4J4J sslmode=disable TimeZone=Asia/Shanghai"

func TestMain(m *testing.M) {
	var err error
	novelTestDB, err = gorm.Open(postgres.Open(novelTestDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "setup: open novel_test db failed: %v\n", err)
		os.Exit(1)
	}
	sqlDB, _ := novelTestDB.DB()
	if err := sqlDB.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "setup: ping novel_test db failed: %v\n", err)
		os.Exit(1)
	}

	// 建表：PG 兼容 DDL（模型 tag 沿用 MySQL 类型名，PG 无 tinyint/datetime/type:json 直建，故用显式 DDL）
	ddl := []string{
		`CREATE TABLE IF NOT EXISTS app_user (
			id BIGSERIAL PRIMARY KEY, level_id BIGINT NOT NULL DEFAULT 1,
			user_name varchar(100) NOT NULL DEFAULT '', true_name varchar(100) NOT NULL DEFAULT '',
			money numeric(30,18) NOT NULL DEFAULT 0, email varchar(300), mobile_title varchar(255) NOT NULL DEFAULT '+86',
			mobile varchar(100), avatar varchar(1000), pay_pwd varchar(100) NOT NULL DEFAULT '', pwd varchar(100) NOT NULL DEFAULT '',
			ref_code varchar(255), parent_id BIGINT NOT NULL DEFAULT 0, parent_ids varchar(1000) NOT NULL DEFAULT '',
			tree_sort BIGINT NOT NULL DEFAULT 0, tree_sorts varchar(1000) NOT NULL DEFAULT '0',
			tree_leaf varchar(1) NOT NULL DEFAULT '0', tree_level BIGINT NOT NULL DEFAULT 0,
			status varchar(1) NOT NULL DEFAULT '1', unique_id BIGINT, remark varchar(255),
			create_by BIGINT NOT NULL DEFAULT 0, update_by BIGINT NOT NULL DEFAULT 0,
			created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE UNIQUE INDEX IF NOT EXISTS uniq_it_user_unique_id ON app_user (unique_id)`,
		`CREATE TABLE IF NOT EXISTS app_novel_category (
			id BIGSERIAL PRIMARY KEY, name varchar(32) NOT NULL, sort int NOT NULL DEFAULT 0,
			status varchar(1) NOT NULL DEFAULT '1', create_by BIGINT NOT NULL DEFAULT 0, update_by BIGINT NOT NULL DEFAULT 0,
			created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE UNIQUE INDEX IF NOT EXISTS uniq_it_category_name ON app_novel_category (name)`,
		`CREATE TABLE IF NOT EXISTS app_novel_book (
			id BIGSERIAL PRIMARY KEY, title varchar(100) NOT NULL, author varchar(64) NOT NULL,
			cover varchar(500) NOT NULL DEFAULT '', rating numeric(3,1) NOT NULL DEFAULT 0,
			review_count int NOT NULL DEFAULT 0, serial_status varchar(1) NOT NULL DEFAULT '1',
			category bigint NOT NULL DEFAULT 0, tags jsonb NOT NULL DEFAULT '[]',
			slogan varchar(30) NOT NULL DEFAULT '', description varchar(500) NOT NULL DEFAULT '',
			clicks int NOT NULL DEFAULT 0, publish_date date, word_count int NOT NULL DEFAULT 0,
			chapters int NOT NULL DEFAULT 0, is_featured smallint NOT NULL DEFAULT 0,
			read_url varchar(500) NOT NULL DEFAULT '', status varchar(1) NOT NULL DEFAULT '1',
			create_by BIGINT NOT NULL DEFAULT 0, update_by BIGINT NOT NULL DEFAULT 0,
			created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS app_novel_book_review (
			id BIGSERIAL PRIMARY KEY, book_id BIGINT NOT NULL, user_id BIGINT NOT NULL,
			user_name varchar(64) NOT NULL DEFAULT '', user_avatar varchar(500) NOT NULL DEFAULT '',
			rating int NOT NULL DEFAULT 5, content text NOT NULL DEFAULT '', likes int NOT NULL DEFAULT 0,
			created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_it_review_book ON app_novel_book_review (book_id)`,
		`CREATE TABLE IF NOT EXISTS app_novel_bookshelf (
			id BIGSERIAL PRIMARY KEY, user_id BIGINT NOT NULL, book_id BIGINT NOT NULL,
			sort_no int NOT NULL DEFAULT 0, created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE UNIQUE INDEX IF NOT EXISTS uniq_it_shelf ON app_novel_bookshelf (user_id, book_id)`,
		`CREATE TABLE IF NOT EXISTS app_novel_reader_profile (
			id BIGSERIAL PRIMARY KEY, user_id BIGINT NOT NULL, nickname varchar(64) NOT NULL DEFAULT '',
			avatar varchar(500) NOT NULL DEFAULT '', bio varchar(500) NOT NULL DEFAULT '',
			email varchar(128) NOT NULL DEFAULT '', preferred_categories jsonb NOT NULL DEFAULT '[]',
			notify_comment smallint NOT NULL DEFAULT 1, notify_book_update smallint NOT NULL DEFAULT 1,
			ban_post_until timestamptz, ban_reason varchar(255) NOT NULL DEFAULT '',
			created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE UNIQUE INDEX IF NOT EXISTS uniq_it_profile_user ON app_novel_reader_profile (user_id)`,
	}
	for _, q := range ddl {
		if err := novelTestDB.Exec(q).Error; err != nil {
			fmt.Fprintf(os.Stderr, "setup: ddl failed: %v\n%s\n", err, q)
			os.Exit(1)
		}
	}

	code := m.Run()

	// 清理（保留库，清空测试数据）
	novelTestDB.Exec(`TRUNCATE TABLE app_novel_book_review, app_novel_bookshelf, app_novel_reader_profile, app_novel_book, app_novel_category, app_user RESTART IDENTITY CASCADE`)
	sqlDB.Close()
	os.Exit(code)
}

// newNovelService 构造带测试 DB 的 service.Service
func newNovelService() service.Service {
	return service.Service{Orm: novelTestDB, Lang: "zh-CN"}
}

// seedUser 创建测试读者并返回 id
func seedUser(t *testing.T, status string, banUntil *time.Time, banReason string) int64 {
	t.Helper()
	now := time.Now()
	u := userModels.User{
		UserName: fmt.Sprintf("it_user_%d", now.UnixNano()),
		LevelId:  1,
		Status:   status,
		TreeLeaf: "1",
	}
	if err := novelTestDB.Create(&u).Error; err != nil {
		t.Fatalf("seed user failed: %v", err)
	}
	p := models.NovelReaderProfile{
		UserId:      u.Id,
		Nickname:    u.UserName,
		BanPostUntil: banUntil,
		BanReason:   banReason,
	}
	if err := novelTestDB.Create(&p).Error; err != nil {
		t.Fatalf("seed profile failed: %v", err)
	}
	return u.Id
}

// seedBook 创建测试书籍并返回 id
func seedBook(t *testing.T, rating float64, reviewCount int) int64 {
	t.Helper()
	b := models.NovelBook{
		Title:        fmt.Sprintf("it_book_%d", time.Now().UnixNano()),
		Author:       "test",
		Category:     "1",
		Tags:         "[]", // 合法 JSON，避免 PG jsonb 空串报错
		Rating:       rating,
		ReviewCount:  reviewCount,
		SerialStatus: "1",
		Status:       "1",
	}
	if err := novelTestDB.Create(&b).Error; err != nil {
		t.Fatalf("seed book failed: %v", err)
	}
	return b.Id
}

// TestIntegrationBookReviewRatingAggregation 书评提交 → 评分加权聚合 + 书评数+1
func TestIntegrationBookReviewRatingAggregation(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test in short mode")
	}
	uid := seedUser(t, "1", nil, "")
	bookId := seedBook(t, 9.5, 1)

	srv := NewNovelBookReviewService(&service.Service{Orm: novelTestDB, Lang: "zh-CN"})
	id, respCode, err := srv.Insert(&dto.NovelReviewInsertReq{
		BookId:     bookId,
		Rating:     4,
		Content:    "集成测试书评",
		CurrUserId: uid,
	})
	if err != nil {
		t.Fatalf("insert review failed: code=%d err=%v", respCode, err)
	}
	if id <= 0 {
		t.Fatal("expected review id > 0")
	}

	var book models.NovelBook
	if err := novelTestDB.First(&book, bookId).Error; err != nil {
		t.Fatal(err)
	}
	// (9.5*1 + 4)/2 = 6.75 → 6.8
	if book.Rating != 6.8 {
		t.Errorf("expected rating 6.8, got %.2f", book.Rating)
	}
	if book.ReviewCount != 2 {
		t.Errorf("expected review_count 2, got %d", book.ReviewCount)
	}
}

// TestIntegrationBookshelfIdempotent 书架重复加入幂等（返回既有记录，不重复创建）
func TestIntegrationBookshelfIdempotent(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test in short mode")
	}
	uid := seedUser(t, "1", nil, "")
	bookId := seedBook(t, 0, 0)

	srv := NewNovelBookshelfService(&service.Service{Orm: novelTestDB, Lang: "zh-CN"})
	first, respCode, err := srv.Insert(&dto.NovelShelfInsertReq{BookId: bookId, CurrUserId: uid})
	if err != nil {
		t.Fatalf("first insert failed: code=%d err=%v", respCode, err)
	}
	second, respCode, err := srv.Insert(&dto.NovelShelfInsertReq{BookId: bookId, CurrUserId: uid})
	if err != nil {
		t.Fatalf("second insert failed: code=%d err=%v", respCode, err)
	}
	if first == 0 || second != first {
		t.Errorf("expected second insert to return same id, first=%d second=%d", first, second)
	}
	var cnt int64
	novelTestDB.Model(&models.NovelBookshelf{}).Where("book_id = ?", bookId).Count(&cnt)
	if cnt != 1 {
		t.Errorf("expected 1 shelf row, got %d", cnt)
	}
}

// TestIntegrationBannedUserWriteRejected 禁言用户写操作被 41043 拒绝
func TestIntegrationBannedUserWriteRejected(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test in short mode")
	}
	future := time.Now().Add(24 * time.Hour)
	uid := seedUser(t, "2", &future, "it ban")

	srv := NewNovelBookReviewService(&service.Service{Orm: novelTestDB, Lang: "zh-CN"})
	_, respCode, err := srv.Insert(&dto.NovelReviewInsertReq{
		BookId:     999, // 即便是不存在书籍，权限校验也应先行拦截
		Rating:     5,
		Content:    "banned",
		CurrUserId: uid,
	})
	if err == nil {
		t.Fatal("expected error for banned user write")
	}
	if respCode != baseLang.NovelPostBannedCode {
		t.Errorf("expected code 41043(NovelPostBannedCode), got %d", respCode)
	}
}

// TestIntegrationBannedExpiryLazyRestore 禁言到期 → 惰性恢复后写操作放行
func TestIntegrationBannedExpiryLazyRestore(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration test in short mode")
	}
	past := time.Now().Add(-24 * time.Hour)
	uid := seedUser(t, "2", &past, "expired ban")
	bookId := seedBook(t, 0, 0)

	srv := NewNovelBookReviewService(&service.Service{Orm: novelTestDB, Lang: "zh-CN"})
	_, respCode, err := srv.Insert(&dto.NovelReviewInsertReq{
		BookId:     bookId,
		Rating:     4,
		Content:    "after expiry",
		CurrUserId: uid,
	})
	if err != nil {
		t.Fatalf("expected write allowed after ban expiry, code=%d err=%v", respCode, err)
	}
	// 惰性恢复：禁言字段清空，状态 2 → 1
	var u userModels.User
	novelTestDB.First(&u, uid)
	if u.Status != "1" {
		t.Errorf("expected status restored to 1, got %s", u.Status)
	}
	var p models.NovelReaderProfile
	novelTestDB.Where("user_id = ?", uid).First(&p)
	if p.BanPostUntil != nil || p.BanReason != "" {
		t.Errorf("expected ban fields cleared, got until=%v reason=%q", p.BanPostUntil, p.BanReason)
	}
}