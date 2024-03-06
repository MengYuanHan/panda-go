package response

import (
	"github.com/gin-gonic/gin"
)

// ResultMsg 返回值
func ResultMsg(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}
