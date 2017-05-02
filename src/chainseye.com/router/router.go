package router

import (
	"fmt"

	"chainseye.com/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

//SetRouter gin框架设置路由
func SetRouter() {

	//gin.SetMode(gin.ReleaseMode)

	ginRouter := gin.Default()
	//ginRouter := gin.New()

	ginRouter.Use(MiddleWare)

	ginRouter.Static("/static", "./src/chainseye.com/view/static")
	ginRouter.StaticFile("/favicon.ico", "./favicon.ico")
	ginRouter.LoadHTMLGlob("src/chainseye.com/view/templates/*")

	displayPage(ginRouter)
	makeAPI(ginRouter)
	startWebSocket(ginRouter)

	ginRouter.Run(":8099")
}

// Page
func displayPage(ginRouter *gin.Engine) {
	ginRouter.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "cWF_index.tmpl", nil)
	})

	ginRouter.GET("/chains_back/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "cWB_portal.tmpl", gin.H{
			"info": " Get Router ! ",
		})
	})
}

// API
func makeAPI(ginRouter *gin.Engine) {
	ginAPI := ginRouter.Group("/api", controller.APIMiddleWare)
	{
		ginAPI.GET("info", controller.GetBaseInfo)
		ginAPI.POST("checkPwd", controller.CheckPwd)
	}
}

// WebSocket
func startWebSocket(ginRouter *gin.Engine) {
	ginWebSocket := ginRouter.Group("/ws", controller.WebSocketMiddleWare)
	{
		ginWebSocket.GET("/ws", controller.WebSocketConnect)
	}
}

//MiddleWare 中间件
func MiddleWare(c *gin.Context) {
	fmt.Println("\n MiddleWare start !")
}
