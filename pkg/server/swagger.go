package server

import (
	_ "ez/docs" // ! 注意此处需要更改，docs中有init函数

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 注意引入docs下的init函数
// 配合github.com/swaggo/swag使用
// go get -u github.com/swaggo/swag
// swag init
func RegisterSwagger(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) { c.PersistAuthorization = true }))

	return r
}
