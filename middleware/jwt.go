package middleware

import (
	"ez/common"
	"ez/global"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		JwtAuthVerify(ctx)
	}
}

func JwtAuthVerify(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" || !strings.HasPrefix(token, "Bearer") {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    403,
			"message": "登录过期",
			"data":    nil,
		})
		ctx.Abort()
		return
	}

	//去掉Bearer
	token = token[7:]

	// 验证token，并存储在请求中
	var userData = common.UserJwt{}
	err := global.Config.Jwt.Jwt.ParseToken(token, &userData)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    403,
			"message": "登录过期",
			"data":    err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.Set("username", userData.Username)
	ctx.Set("userType", userData.UserType)
	ctx.Set("userId", userData.UserId)

	ctx.Next()
}
