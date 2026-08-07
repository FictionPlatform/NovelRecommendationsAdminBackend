package main

import (
	"go-admin/core/cmd"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate go run app/admin/sys/models/parseapi/gen_api_desc.go
//go:generate swag init -g main.go -o docs
//go:generate swag init -g main.go -d ./gen/webapidoc,./ -t 读者认证,读者资料,书友关注,我的读者资料,我的书架,系统通知,小说书库,小说书评,小说长文,意见反馈 -o docs/webapi --instanceName webapi

// @title Go-admin 后台管理系统
// @version 2.0.0
// @description 基于Gin的后台管理系统接口文档
// @host localhost:8888
// @BasePath /admin-api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
