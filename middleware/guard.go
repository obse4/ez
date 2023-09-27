package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限校验
// 在jwt写入userType后使用
func Guard(userType int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_type := ctx.GetInt("userType")

		if user_type == 0 || userType != user_type {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    403,
				"message": "权限校验失败",
				"data":    nil,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
