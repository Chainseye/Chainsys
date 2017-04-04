package router

import (
	"chainseye.com/chainsWebSite_Back/controller"

	"github.com/gin-gonic/gin"
)

//SetRouter gin框架设置路由
func SetRouter() {

	gin.SetMode(gin.DebugMode)

	ginRouter := gin.Default()

	makeAPI(ginRouter)

	ginRouter.Run(":8099")
}

func makeAPI(ginRouter *gin.Engine) {
	ginRouter.Get("/", controller.Index)
}
