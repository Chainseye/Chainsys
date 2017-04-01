package router

import (
	"github.com/gin-gonic/gin"
)

//SetRouter gin框架设置路由
func SetRouter() {

	gin.SetMode(gin.DebugMode)

	ginRouter := gin.Default()

	makeAPI(ginRouter)

}

func makeAPI(ginRouter *gin.Engine) {

}
