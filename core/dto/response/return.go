package response

import (
	"go-admin/core/utils/strutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Default = &response{}

// Error 失败数据处理
// L22：仅当业务码落在合法 4xx/5xx 区间（AuthErr=401/ForbitErr=403/ServerErr=500）时
// 才作为 HTTP 状态码透传，其余业务码（≥888）一律返回 400，避免 2xx/3xx 等
// 非错误码语义泄露到 HTTP 层
func Error(c *gin.Context, code int, msg string) {
	httpCode := http.StatusBadRequest
	if code >= http.StatusBadRequest && code <= 599 {
		httpCode = code
	}
	ErrorByHttpCode(c, httpCode, code, msg)

}

// ErrorByHttpCode 自定义httpCode
func ErrorByHttpCode(c *gin.Context, httpCode, code int, msg string) {
	res := Default.Clone()
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(strutils.GenerateMsgIDFromContext(c))
	res.SetCode(code)
	res.SetSuccess(false)
	if httpCode > 600 {
		httpCode = http.StatusBadRequest
	}
	c.AbortWithStatusJSON(httpCode, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, code int, msg string) {
	res := Default.Clone()
	res.SetData(data)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetSuccess(true)
	res.SetTraceID(strutils.GenerateMsgIDFromContext(c))
	res.SetCode(code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func Download(c *gin.Context, data []byte, filename, contentType string) {
	c.Header("Content-Disposition", "attachment;filename="+filename)
	c.Header("Access-Control-Expose-Headers", "Content-Disposition") //允许获取懂到指定header
	c.Header("Content-AppType", contentType)
	c.Data(http.StatusOK, contentType, data)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result, extend interface{}, count int64, pageIndex, pageSize, code int, msg string) {
	var res page
	res.List = result
	res.Extend = extend
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, code, msg)
}
