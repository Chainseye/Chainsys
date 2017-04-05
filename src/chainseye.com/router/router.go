package router

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

//SetRouter gin框架设置路由
func SetRouter() {

	gin.SetMode(gin.ReleaseMode)

	//ginRouter := gin.Default()
	ginRouter := gin.New()

	ginRouter.Use(MiddleWare)

	ginRouter.Static("./src/chainseye.com/view/static", "ginStatic")
	ginRouter.LoadHTMLGlob("./src/chainseye.com/view/templates/*")

	ginRouter.GET("/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "cWF_index.tmpl", nil)

	})

	ginRouter.GET("/chains_back/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "cWB_portal.tmpl", nil)

	})

	ginRouter.Run(":8099")
}

//MiddleWare 中间件
func MiddleWare(c *gin.Context) {
	fmt.Println("\n MiddleWare start !")
}
