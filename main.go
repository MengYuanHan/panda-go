package main

import (
	"Panda/common"
	"Panda/router"
	"Panda/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// main 入口
func main() {
	// InitConfig 初始化配置文件
	common.InitConfig()
	// InitDB 加载数据库
	common.NewDBConnection().DB()
	r := gin.Default()
	// gin设置一个信任ip
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// 向gin注册业务路由
	r = router.Register(r)
	// 获取服务端口
	port := util.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080

}
