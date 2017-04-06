package main

import (
	"chainseye.com/controller"
	"chainseye.com/router"
)

func main() {
	// 设置路由
	router.SetRouter()
	controller.StartWebSocket()
}
