package router

import (
	"net/http"

	"chainseye.com/controller"

	"github.com/gin-gonic/gin"
)

var admins = gin.H{
	"liph":  gin.H{"name": "李培晖", "nickName": "LPH"},
	"admin": gin.H{"name": "admin", "nickName": "admin"},
}

//SetRouter gin框架设置路由
func SetRouter(port string) {
	gin.SetMode("debug")
	ginRouter := gin.Default()

	ginRouter.Static("/static", "./src/chainseye.com/static/")
	ginRouter.StaticFile("/favicon.ico", "./favicon.ico")
	ginRouter.LoadHTMLGlob("src/chainseye.com/static/html/*")

	displayPage(ginRouter)
	makeAPI(ginRouter)
	startWebSocket(ginRouter)
	controller.RunWebSocketBroadcast()
	startWsBroadcast(ginRouter)

	ginRouter.Run(port)
}

// Page
func displayPage(ginRouter *gin.Engine) {
	ginRouter.GET("/", func(context *gin.Context) {
		style := context.Request.FormValue("style")
		cssFileExists, _ := controller.PathExists("src/chainseye.com/static/skin/css/" + style + ".css")
		if !cssFileExists == true {
			style = "style"
		}
		context.HTML(http.StatusOK, "cWF_index.html", gin.H{
			"info":  "The index page !",
			"style": style,
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
	ginAPI := ginRouter.Group("/api/user", gin.BasicAuth(gin.Accounts{
		"liph":  "123123",
		"admin": "123123",
	}))
	{
		ginAPI.GET("info", controller.GetUserInfoByUID)
		ginAPI.GET("/auth", func(context *gin.Context) {
			name := context.MustGet(gin.AuthUserKey).(string)
			if nickName, ok := admins[name]; ok {
				context.JSON(http.StatusOK, gin.H{"name": name, "nickName": nickName})
			} else {
				context.JSON(http.StatusOK, gin.H{"name": name, "nickName": "NO NickName :("})
			}
		})
	}
}

// WebSocket
func startWebSocket(ginRouter *gin.Engine) {
	ginRouter.GET("/ws", controller.WebSocketConnect)
}

func startWsBroadcast(ginRouter *gin.Engine) {
	ginRouter.GET("/wsBroadcast", controller.WebSocketBroadcast)
}
