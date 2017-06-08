package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chainseye.com/controller"
)

//SetRouter gin框架设置路由
func SetRouter(port string) {
	ginRouter := gin.Default()

	ginRouter.Static("/static", "./src/chainseye.com/static/")
	ginRouter.StaticFile("/favicon.ico", "./favicon.ico")
	ginRouter.LoadHTMLGlob("src/chainseye.com/static/html/*")

	displayPage(ginRouter)
	makeAPI(ginRouter)

	ginRouter.Run(port)
}

//SetWebSocket 设置websocket
func SetWebSocket(port string) {
	ginRouter := gin.Default()

	startWebSocket(ginRouter)

	ginRouter.Run(port)
}

// Page
func displayPage(ginRouter *gin.Engine) {
	ginRouter.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "cWF_index.html", gin.H{
			"info": "The index page !",
		})
	})

	ginRouter.GET("/chains_back/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "cWB_portal.html", gin.H{
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
	ginRouter.GET("/ws", controller.WebSocketConnect)
}
