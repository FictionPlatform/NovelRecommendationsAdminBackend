package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-admin/config/base/constant"
	"go-admin/core/middleware/auth"
	"go-admin/core/middleware/auth/authdto"
)

func Auth() gin.HandlerFunc {
	return auth.Auth.AuthMiddlewareFunc()
}

func AuthCheckRole() gin.HandlerFunc {
	return auth.Auth.AuthCheckRoleMiddlewareFunc()
}

// AdminOnly 仅允许超级管理员（admin 角色）访问
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString(authdto.RoleKey) != constant.RoleKeyAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "仅超级管理员可访问",
			})
			return
		}
		c.Next()
	}
}
