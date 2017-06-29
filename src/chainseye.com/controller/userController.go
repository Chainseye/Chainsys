package controller

import (
	"net/http"

	"chainseye.com/model"

	"github.com/gin-gonic/gin"
)

//GetUserInfoByUID 根据UID获取基本信息
func GetUserInfoByUID(context *gin.Context) {

	um := model.NewUserModel(model.DEFAULT_MODEL)
	ur, err := um.GetUserInfoByUID(1)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"retCode": 1001,
			"info":    err,
		})
		return
	}
	userData := map[string]interface{}{
		"mobile":   ur.Mobile,
		"password": ur.Password,
	}
	context.JSON(http.StatusOK, gin.H{
		"retCode": 0,
		"info":    userData,
	})
}
