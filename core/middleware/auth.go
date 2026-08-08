package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth"
	"go-admin/core/middleware/auth/authdto"
)

func Auth() gin.HandlerFunc {
	return auth.Auth.AuthMiddlewareFunc()
}

// AuthOptional 可选登录：携带有效 token 时注入用户身份（同 Auth 的校验强度），
// 无 token 或 token 无效时放行为匿名访问。用于"游客可浏览、登录可看我的数据"的公开读接口。
func AuthOptional() gin.HandlerFunc {
	return auth.Auth.AuthOptionalMiddlewareFunc()
}

func AuthCheckRole() gin.HandlerFunc {
	return auth.Auth.AuthCheckRoleMiddlewareFunc()
}

// AdminOnly 仅允许超级管理员（admin 角色）访问
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString(authdto.RoleKey) != constant.RoleKeyAdmin {
			response.ErrorByHttpCode(c, http.StatusForbidden, baseLang.AdminOnlyErrCode,
				lang.MsgByCode(baseLang.AdminOnlyErrCode, lang.GetAcceptLanguage(c)))
			return
		}
		c.Next()
	}
}
