package initial

import (
	"ez/controller/admin"
	"ez/global"
	"ez/middleware"
	"ez/pkg/server"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	s := server.NewHttpServer(&server.HttpServerConfig{
		Port: global.Config.HttpServer.Port,
		Mode: global.Config.HttpServer.Mode,
	})

	s.Router().GET("health", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	s.Router().POST("upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("files")
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("file upload failed: %s", err.Error()))
			return
		}

		// Save the uploaded file
		filePath := "uploads/" + file.Filename
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("file save failed: %s", err.Error()))
			return
		}

		ctx.String(http.StatusOK, "File uploaded successfully")
	})

	// 注册swagger
	server.RegisterSwagger(s.Router())

	// 注册路由
	registerAdminRouter(s.Router())

	s.Init()
}

func registerAdminRouter(r *gin.Engine) {
	{
		adminUserRouter := r.Group("admin/user", middleware.JwtMiddleware(), middleware.Guard(1))
		adminUserRouter.POST("create", admin.UserAdminCreate)
		adminUserRouter.POST("remove", admin.UserAdminRemove)
		adminUserRouter.POST("edit", admin.UserAdminEdit)
		adminUserRouter.POST("find/paging", admin.UserAdminFindPaging)
	}
}
