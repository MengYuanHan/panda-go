package api

import (
	"Panda/application"
	"Panda/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserList @Summary 获取用户列表
// @Description 获取用户列表
// @Tags users
// @Accept  json
// @Produce  json
// @Param input query models.UsersParam true "query"
// @Success 200 {string} string "ok"
// @Router /users/query [get]
// +get /users/query
func GetUserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//token := ctx.GetHeader("token")
		//if len(token) == 0 {
		//	response.ResultMsg(ctx, http.StatusOK, 500, nil, "token is not empty")
		//}
		userInteractor := application.NewUserInteractor(ctx)
		result, err := userInteractor.Get(1001)
		if err != nil {
			response.ResultMsg(ctx, http.StatusInternalServerError, 500, nil, "[USERS]"+err.Error())
			return
		}
		response.ResultMsg(ctx, http.StatusOK, 200, result, "ok")
	}
}
