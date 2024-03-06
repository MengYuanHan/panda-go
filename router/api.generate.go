package router

import (
	"Panda/api"
	"Panda/middleware"
	"github.com/gin-gonic/gin"
)

// Register 向 Gin 注册业务路由
func Register(eng *gin.Engine) *gin.Engine {
	eng.Use(middleware.CORSMiddleware())
	group := eng.Group("/api/")
	// package Panda/api
	// 查询数据字典
	group.GET("users/query", api.GetUserList())
	return eng
}
