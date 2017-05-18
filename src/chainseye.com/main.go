package main

import (
	"chainseye.com/router"
)

func main() {
	// 设置路由
	router.SetRouter(":8099")
	// 设置websocket
	router.SetWebSocket(":8079")
}
