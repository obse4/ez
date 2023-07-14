package initial

import (
	"ez/pkg/server"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	s := server.NewHttpServer(&server.HttpServerConfig{
		Port: ":8080",
		Mode: "debug",
	})

	s.Router().GET("health", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	s.Init()
}
