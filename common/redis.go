package common

import (
	"Panda/util"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Get 获取
func Get(key string) string {
	addressUrl := util.GetString("redis.address_url")
	password := util.GetString("redis.password")
	var rdb = redis.NewClient(&redis.Options{Addr: addressUrl, Password: password, DB: 0})
	res, err := rdb.Get(key).Result()
	if err != nil {
		fmt.Println("获取数据失败:", err)
	}
	return res
}

// Set 设置
func Set(key string, val string, t time.Duration) string {
	addressUrl := util.GetString("redis.address_url")
	password := util.GetString("redis.password")
	var rdb = redis.NewClient(&redis.Options{Addr: addressUrl, Password: password, DB: 0})
	res, err := rdb.Set(key, val, t*time.Second).Result()
	if err != nil {
		fmt.Println("设置数据失败:", err)
	}
	return res
}
