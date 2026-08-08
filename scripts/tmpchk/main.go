package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(postgres.Open("host=127.0.0.1 port=5432 user=postgres dbname=public password=nVZypeJxyuXZ4J4J sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	var rows []struct {
		ID       int64
		UserName string
		Status   string
	}
	db.Raw("SELECT id, user_name, status FROM app_user ORDER BY id").Scan(&rows)
	for _, r := range rows {
		fmt.Printf("id=%d status=%s name=%q\n", r.ID, r.Status, r.UserName)
	}
}
