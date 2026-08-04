package middleware

import (
	"bytes"
	"encoding/json"
	"go-admin/core/config"
	"go-admin/core/middleware/auth"
	"go-admin/core/runtime"
	"go-admin/core/utils/iputils"
	"go-admin/core/utils/log"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"go-admin/core/global"
)

// maxBodyLogSize 请求体记录上限，超过则不缓存请求体（防止内存耗尽，且不影响大文件上传）
const maxBodyLogSize = 1 << 20

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		rLog := log.GetRequestLogger(c)
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			// 请求体过大时跳过记录，保留原始 body 给业务处理（如文件上传）
			if c.Request.ContentLength > maxBodyLogSize {
				break
			}
			bf := bytes.NewBuffer(nil)
			_, err := io.Copy(bf, io.LimitReader(c.Request.Body, maxBodyLogSize))
			if err != nil {
				rLog.Warnf("copy body error, %s", err.Error())
			}
			rb := bf.Bytes()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}

		c.Next()
		// 结束时间
		endTime := time.Now()
		if c.Request.Method == http.MethodOptions {
			return
		}

		rt, bl := c.Get("result")
		var result = ""
		if bl {
			rb, err := json.Marshal(rt)
			if err != nil {
				rLog.Warnf("json Marshal result error, %s", err.Error())
			} else {
				result = string(rb)
			}
		}

		st, bl := c.Get("status")
		var statusBus = 0
		if bl {
			if v, ok := st.(int); ok {
				statusBus = v
			}
		}

		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := iputils.GetClientIP(c)
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 日志格式
		logData := map[string]interface{}{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}
		rLog.WithFields(logData).Info()

		if c.Request.Method != "OPTIONS" && config.LoggerConfig.EnabledDB && statusCode != 404 {
			// 敏感字段脱敏后再入库，防止密码/token 明文落库
			SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime, maskSensitiveFields(body), maskSensitiveFields(result), statusBus)
		}
	}
}

// isSensitiveKey 判断是否敏感字段
func isSensitiveKey(key string) bool {
	switch strings.ToLower(strings.ReplaceAll(key, "-", "")) {
	case "password", "pwd", "paypwd", "pay_pwd", "oldpassword", "newpassword", "confirmpassword",
		"token", "accesstoken", "refreshtoken", "access_token", "refresh_token", "authorization",
		"secret", "secretkey", "appsecret", "captcha", "verifycode", "verificationcode", "validatecode":
		return true
	}
	return false
}

// maskJSON 递归脱敏 JSON 中的敏感字段值
func maskJSON(obj interface{}) interface{} {
	switch v := obj.(type) {
	case map[string]interface{}:
		for k, val := range v {
			if isSensitiveKey(k) {
				v[k] = "***"
			} else {
				v[k] = maskJSON(val)
			}
		}
		return v
	case []interface{}:
		for i, item := range v {
			v[i] = maskJSON(item)
		}
		return v
	default:
		return obj
	}
}

// maskSensitiveFields 对请求体/响应体脱敏：优先 JSON 结构脱敏，非 JSON 用正则兜底
func maskSensitiveFields(data string) string {
	if data == "" {
		return data
	}
	var obj interface{}
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return maskSensitiveFieldsRegex(data)
	}
	if b, err := json.Marshal(maskJSON(obj)); err == nil {
		return string(b)
	}
	return data
}

// maskSensitiveFieldsRegex 正则兜底：匹配 "key": "value" 形式
var sensitiveFieldsRegex = regexp.MustCompile(`(?i)("(?:password|pwd|payPwd|pay_pwd|oldPassword|newPassword|confirmPassword|token|accessToken|refreshToken|access_token|refresh_token|authorization|secret|secretKey|appSecret|captcha|verifyCode|verificationCode|validateCode)"\s*:\s*")[^"]*(")`)

func maskSensitiveFieldsRegex(data string) string {
	return sensitiveFieldsRegex.ReplaceAllString(data, `${1}***${2}`)
}

// SetDBOperLog 写入操作日志表 fixme 该方法后续即将弃用
func SetDBOperLog(c *gin.Context, clientIP string, statusCode int, reqUri string, reqMethod string, latencyTime time.Duration, body string, result string, status int) {
	rLog := log.GetRequestLogger(c)
	l := make(map[string]interface{})
	//l["_fullPath"] = c.FullPath()  //reqUri可以取代
	l["operUrl"] = reqUri
	l["operIp"] = clientIP
	//用于定位ip所在城市
	l["operLocation"] = iputils.GetLocation(clientIP, config.ApplicationConfig.AmpKey)
	userId, _, _ := auth.Auth.GetUserId(c)
	l["userId"] = userId
	l["requestMethod"] = c.Request.Method
	l["operParam"] = body
	l["userAgent"] = c.Request.UserAgent()
	l["operTime"] = time.Now()
	l["jsonResult"] = result
	l["latencyTime"] = latencyTime.String()
	l["statusCode"] = statusCode
	l["status"] = strconv.Itoa(status)
	q := runtime.RuntimeConfig.GetMemoryQueue(c.Request.Host)
	message, err := runtime.RuntimeConfig.GetStreamMessage("", global.OperateLog, l)
	if err != nil {
		rLog.Errorf("GetStreamMessage error, %s", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			rLog.Errorf("Append message error, %s", err.Error())
		}
	}
}
