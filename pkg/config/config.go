package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig(config interface{}) {
	fileName := os.Getenv("CONFIG_MODE")
	if fileName == "" {
		fileName = "env"
	}
	wd, _ := os.Getwd()
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType("yml")
	v.AddConfigPath(wd + "/config")

	if err := v.ReadInConfig(); err != nil {
		LogFatal("Config", "config read error:%s", err.Error())
	}
	if err := v.Unmarshal(&config); err != nil {
		LogFatal("Config", "unmarshal json error:%s", err.Error())
	}
}
