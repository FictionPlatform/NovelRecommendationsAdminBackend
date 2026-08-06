package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/runtime"
)

func WithContextDb(c *gin.Context) {
	db := runtime.RuntimeConfig.GetDbByKey(c.Request.Host)
	if db == nil {
		response.ErrorByHttpCode(c, http.StatusInternalServerError, baseLang.DbConnFailCode,
			lang.MsgByCode(baseLang.DbConnFailCode, lang.GetAcceptLanguage(c)))
		return
	}
	c.Set("db", db.WithContext(c))
	c.Next()
}
