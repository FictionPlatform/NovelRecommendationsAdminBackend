package dberr

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func TestIsDuplicateKey(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"plain error", errors.New("something broke"), false},
		{"gorm duplicated", gorm.ErrDuplicatedKey, true},
		{"gorm wrapped", errors.Join(errors.New("context"), gorm.ErrDuplicatedKey), true},
		{"mysql 1062", &mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'x' for key 'uniq_app_user_ref_code'"}, true},
		{"mysql wrapped", errors.Join(errors.New("context"), &mysql.MySQLError{Number: 1062}), true},
		{"mysql other", &mysql.MySQLError{Number: 1146}, false},
		{"pg 23505", &pgconn.PgError{Code: "23505", ConstraintName: "uniq_app_user_ref_code"}, true},
		{"pg wrapped", errors.Join(errors.New("context"), &pgconn.PgError{Code: "23505"}), true},
		{"pg other", &pgconn.PgError{Code: "22P02"}, false},
		{"msg duplicate entry", errors.New("Error 1062: Duplicate entry '1' for key 'PRIMARY'"), true},
		{"msg unique constraint", errors.New("ERROR: duplicate key value violates unique constraint \"uniq_app_user_email\""), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDuplicateKey(tt.err); got != tt.want {
				t.Errorf("IsDuplicateKey(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestIsDuplicateIndex(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"plain error", errors.New("something broke"), false},
		{"mysql 1061", &mysql.MySQLError{Number: 1061, Message: "Duplicate key name 'idx_admin_sys_casbin_rule'"}, true},
		{"mysql wrapped", errors.Join(errors.New("context"), &mysql.MySQLError{Number: 1061}), true},
		{"mysql other", &mysql.MySQLError{Number: 1062}, false},
		{"pg 42P07", &pgconn.PgError{Code: "42P07", Message: "relation \"idx_admin_sys_casbin_rule\" already exists"}, true},
		{"pg wrapped", errors.Join(errors.New("context"), &pgconn.PgError{Code: "42P07"}), true},
		{"pg other", &pgconn.PgError{Code: "23505"}, false},
		{"msg duplicate key name", errors.New("Error 1061: Duplicate key name 'idx_x'"), true},
		{"msg already exists", errors.New("ERROR: relation \"idx_x\" already exists"), true},
		{"msg unrelated", errors.New("Error 1146: Table 'x' doesn't exist"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDuplicateIndex(tt.err); got != tt.want {
				t.Errorf("IsDuplicateIndex(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}
