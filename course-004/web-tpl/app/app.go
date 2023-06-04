package app

import (
	"web-tpl/app/core/config"
)

var Config config.Model

func Init(prjHome string) error {
	// 找到两个配置文件路径
	return Config.LoadConfig(prjHome)

	// viper
}
