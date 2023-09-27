package global

import (
	"ez/pkg/database"
	"ez/pkg/jsonwt"
	"ez/pkg/server"
)

var Config GlobalConfig

type GlobalConfig struct {
	HttpServer server.HttpServerConfig

	MysqlDb database.MysqlConfig

	RedisDb database.RedisConfig

	Jwt jsonwt.JwtConfig
}
