package server

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

// 打印请求日志
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()

		// 执行时间
		accessTime := endTime.Sub(startTime)

		LogDebug("Server", "| %3d | %13v | %15s | %s | %s |", ctx.Writer.Status(),
			accessTime,
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.Request.RequestURI)

	}
}

// Recover
// 确保该中间件处于最上层
// 可以防止程序挂掉
func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			LogRecover("Server", "%v", r)
			debug.PrintStack()

			recoverByte, _ := json.Marshal(r)

			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": string(recoverByte),
				"data":    nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			ctx.Abort()
		}
	}()

	//加载完 defer recover，继续后续接口调用
	ctx.Next()
}

// 跨域访问中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
