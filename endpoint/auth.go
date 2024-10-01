package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/dto"
)

func GetUserInfo(ctx *gin.Context) *dto.UserInfo {
	authData, isExist := ctx.Get("authData")
	if !isExist {
		ctx.String(http.StatusUnauthorized, "failure")
		return nil
	}

	userInfo := authData.(map[string]interface{})

	return &dto.UserInfo{
		UserId: int64(userInfo["user_id"].(float64)),
		Phone:  int64(userInfo["phone"].(float64)),
	}
}
