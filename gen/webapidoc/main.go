// 墨读·小说引力场 独立 web-api swagger spec 扫描入口。
// 仅用于 swag init 生成独立文档，不注册任何路由；生成的 spec 经 docs/webapi 提供。
// 生成命令（searchDir 限定该目录 + import 链，避免混入后台管理接口）：
//
//	swag init -g main.go -d ./gen/webapidoc -o docs/webapi --instanceName webapi
//
// 与主 spec（main.go，@BasePath /admin-api/v1）区分，读者接口文档路径与真实 /web-api/v1 前缀一致。
package main

import (
	_ "go-admin/app/app/novel/apis"
)

// @title 墨读·小说引力场 读者端 API
// @version 2.0.0
// @description 小说推荐平台读者端接口（/web-api/v1/app/novel/**，读公开、写需登录）
// @host localhost:8888
// @BasePath /web-api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
}
