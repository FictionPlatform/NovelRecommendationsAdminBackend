package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go-admin/core/config"
	"go-admin/core/global"
	"go-admin/core/runtime"
	"go-admin/core/ws"
	"mime"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup), 0)
)

// InitRouter 初始化路由
func InitRouter() {
	var r *gin.Engine
	h := runtime.RuntimeConfig.GetEngine()
	if h == nil {
		panic("not found engine...")
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		panic("not found engine...")
	}
	InitSysRouter(r)

	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r)

}

// noCheckRoleRouter 无需认证的路由示例
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group(global.RouteRootPath + "/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// checkRoleRouter 需要认证的路由示例
func checkRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group(global.RouteRootPath + "/v1")

	for _, f := range routerCheckRole {
		f(v1)
	}
}

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	// 上传文件静态服务：禁止目录列举、防路径穿越、危险类型强制下载
	r.GET(global.RouteRootPath+"/"+config.ApplicationConfig.FileRootPath+"/*filepath", func(c *gin.Context) {
		serveFileNoList(c, config.ApplicationConfig.FileRootPath)
	})
	// static 静态目录：同样禁止目录列举（gin.Static 基于 http.FileServer 可列目录）
	r.GET("/static/*filepath", func(c *gin.Context) {
		serveFileNoList(c, "./static")
	})
}

// dangerousExt 可作为脚本执行的危险扩展名，强制以附件形式下载
var dangerousExt = map[string]bool{
	".html": true, ".htm": true, ".svg": true, ".xhtml": true, ".mhtml": true,
	".xml": true, ".js": true, ".css": true, ".json": true, ".txt": true,
}

// serveFileNoList 目录静态服务：禁止目录列举、防路径穿越、危险类型强制下载
func serveFileNoList(c *gin.Context, root string) {
	cleanRoot := filepath.Clean(root)
	filePath := c.Param("filepath")
	// 禁止目录列举
	if filePath == "" || filePath == "/" || strings.HasSuffix(filePath, "/") {
		c.Status(http.StatusNotFound)
		return
	}
	fullPath := filepath.Join(cleanRoot, filepath.FromSlash(filePath))
	// 防路径穿越
	if fullPath != cleanRoot && !strings.HasPrefix(fullPath, cleanRoot+string(filepath.Separator)) {
		c.Status(http.StatusNotFound)
		return
	}
	info, err := os.Stat(fullPath)
	if err != nil || info.IsDir() {
		c.Status(http.StatusNotFound)
		return
	}
	// 危险类型强制以附件下载，不在浏览器中执行
	if dangerousExt[strings.ToLower(filepath.Ext(fullPath))] {
		c.Header("Content-Disposition", "attachment")
	}
	c.File(fullPath)
}
