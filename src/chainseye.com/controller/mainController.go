package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//APIMiddleWare API中间件
func APIMiddleWare(context *gin.Context) {
	fmt.Println("APIMiddleWare")
}

//GetBaseInfo 获取基本信息
func GetBaseInfo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"retCode": 0,
		"info":    "OK",
	})
}

//CheckPwd 检测Password
func CheckPwd(context *gin.Context) {
	// "retCode": 0,
	// "info": "OK",
}

//CheckErr 打印错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
