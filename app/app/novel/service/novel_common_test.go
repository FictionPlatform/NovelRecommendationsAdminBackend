package service

import (
	"errors"
	"testing"
)

// TestAvgRating 评分联动聚合：保留 1 位小数
func TestAvgRating(t *testing.T) {
	cases := []struct {
		name      string
		oldRating float64
		oldCount  int
		newScore  int
		want      float64
	}{
		{"首次评分", 0, 0, 5, 5},
		{"零评分+多条", 0, 0, 4, 4},
		{"普通加权", 9.5, 1, 8, 8.8}, // (9.5+8)/2=8.75 → round→8.8
		{"聚合后仍一位小数", 8.8, 2, 5, 7.5},
		{"小数位四舍五入下降", 9.1, 1, 6, 7.6}, // (9.1+6)/2=7.55 → math.Round→7.6? 见下
		{"大量书评趋近新分", 9.2, 100, 1, 9.1}, // (9.2*100+1)/101=9.1188 → 9.1
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := avgRating(c.oldRating, c.oldCount, c.newScore)
			if got != c.want {
				t.Errorf("avgRating(%.1f,%d,%d)=%.2f, want %.2f", c.oldRating, c.oldCount, c.newScore, got, c.want)
			}
		})
	}
}

// TestAvgRatingOneDecimal 结果始终保留 1 位小数（乘以10取整后除以10）
func TestAvgRatingOneDecimal(t *testing.T) {
	// 直接校验关键收敛点
	if v := avgRating(4.9, 0, 5); v != 5.0 {
		t.Fatalf("expected 5.0, got %.1f", v)
	}
	// oldCount=0 时无历史权重，结果即新评分
	if v := avgRating(5, 0, 4); v != 4.0 {
		t.Fatalf("expected 4.0 (no historical weight), got %.1f", v)
	}
}

// TestMaxZero 计数下限 0
func TestMaxZero(t *testing.T) {
	if got := maxZero(-1); got != 0 {
		t.Errorf("maxZero(-1)=%d, want 0", got)
	}
	if got := maxZero(0); got != 0 {
		t.Errorf("maxZero(0)=%d, want 0", got)
	}
	if got := maxZero(10); got != 10 {
		t.Errorf("maxZero(10)=%d, want 10", got)
	}
}

// TestIsUniqueIDConflict 唯一 ID 冲突错误判定（PG/MySQL 双驱动）
func TestIsUniqueIDConflict(t *testing.T) {
	if isUniqueIDConflict(nil) {
		t.Error("nil error should not be conflict")
	}
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"PG 唯一约束", errors.New(`ERROR: duplicate key value violates unique constraint "uniq_app_user_unique_id" (SQLSTATE 23505)`), true},
		{"MySQL 唯一约束", errors.New("Error 1062: Duplicate entry '100000102' for key 'uniq_app_user_unique_id'"), true},
		{"仅 23505 无关键字", errors.New("pq: duplicate key value violates unique constraint (SQLSTATE 23505)"), false}, // 无 unique_id 字样，保守判定 false
		{"普通错误", errors.New("something wrong"), false},
		{"含 unique_id 的其他错误", errors.New("column unique_id does not exist"), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isUniqueIDConflict(c.err); got != c.want {
				t.Errorf("isUniqueIDConflict(%q)=%v, want %v", c.err.Error(), got, c.want)
			}
		})
	}
}