package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysqlConnect(database *MysqlConfig) (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", database.Username, database.Password, database.Url, database.Port, database.Database)

	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         255,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		DontSupportForShareClause: false,
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent), // 忽略慢sql日志
		PrepareStmt:            false,                                 // 关闭预加载
		SkipDefaultTransaction: true,                                  // 关闭gorm事务模式
	})

	if err != nil {
		LogFatal("Mysql", "数据库 %s 连接失败:%s", database.Name, err.Error())
	}

	LogInfo("Mysql", "数据库 %s 连接成功", database.Name)
	return
}

type MysqlConfig struct {
	Name     string // 自定义名称
	Username string // 用户名
	Password string // 密码
	Database string // 数据库
	Url      string // url地址
	Port     string // 端口
}
