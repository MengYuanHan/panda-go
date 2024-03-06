package util

import "github.com/spf13/viper"

// GetString 获取配置文件
func GetString(key string) (value string) {
	return viper.GetString(key)
}
