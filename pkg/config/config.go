package config

import (
	"os"

	"github.com/spf13/viper"
)

// file 根目录下文件名 例如：/global
// config 全局配置struct
func InitConfig(file string, config interface{}) {
	fileName := os.Getenv("CONFIG_MODE")
	if fileName == "" {
		fileName = "env"
	}
	wd, _ := os.Getwd()
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType("yml")
	v.AddConfigPath(wd + file)

	if err := v.ReadInConfig(); err != nil {
		LogFatal("Config", "config read error:%s", err.Error())
	}
	if err := v.Unmarshal(&config); err != nil {
		LogFatal("Config", "unmarshal json error:%s", err.Error())
	}
	LogInfo("Config", "%#v\n", config)
}
