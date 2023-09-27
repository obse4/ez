package initial

import (
	"ez/global"
	"ez/model"
	"ez/pkg/database"
)

func InitMysql() {
	// 初始化连接
	database.InitMysqlConnect(&global.Config.MysqlDb)
	// 创建表
	global.Config.MysqlDb.Db.AutoMigrate(model.User{})
}
