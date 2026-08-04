package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/core/runtime"
	"net/http"
)

func WithContextDb(c *gin.Context) {
	db := runtime.RuntimeConfig.GetDbByKey(c.Request.Host)
	if db == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "数据库连接获取失败",
		})
		return
	}
	c.Set("db", db.WithContext(c))
	c.Next()
}
