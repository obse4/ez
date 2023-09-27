package initial

import (
	"ez/global"
	"ez/pkg/database"
)

func InitRedis() {
	// 初始化连接池
	database.InitRedisPool(&global.Config.RedisDb)
}
