package initial

import (
	"ez/global"
	"ez/pkg/config"
)

func InitConfig() {
	config.InitConfig("/global", &global.Config)
}
