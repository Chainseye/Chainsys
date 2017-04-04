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

	ginRouter.GET("/", func(context *gin.Context) {

		fmt.Println("\n router success !")

		context.JSON(http.StatusOK, gin.H{
			"title": "Gin Web Success !",
		})

	})

	ginRouter.Run(":8099")
}

//MiddleWare 中间件
func MiddleWare(c *gin.Context) {
	fmt.Println("\n MiddleWare start !")
}
