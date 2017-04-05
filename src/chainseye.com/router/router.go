package router

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

//SetRouter gin框架设置路由
func SetRouter() {

	//gin.SetMode(gin.ReleaseMode)

	ginRouter := gin.Default()
	//ginRouter := gin.New()

	ginRouter.Use(MiddleWare)

	//ginRouter.StaticFS("/index", http.Dir("."))
	//router.Static("/i", "/bin")

	ginRouter.Static("/index", "./src/chainseye.com/view/static/skin/css")
	ginRouter.LoadHTMLGlob("src/chainseye.com/view/templates/*")

	displayPage(ginRouter)
	//makeApi(ginRouter)

	ginRouter.Run(":8099")
}

func displayPage(ginRouter *gin.Engine) {
	ginRouter.GET("/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "cWF_index.tmpl", nil)

	})

	ginRouter.GET("/chains_back/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "cWB_portal.tmpl", gin.H{
			"info": "---  from router  ---",
		})

	})
}

// func makeApi(ginRouter *gin.Engine) {

// }

//MiddleWare 中间件
func MiddleWare(c *gin.Context) {
	fmt.Println("\n MiddleWare start !")
}
