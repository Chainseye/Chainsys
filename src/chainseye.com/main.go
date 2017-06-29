package main

import (
	"chainseye.com/controller"
	"chainseye.com/router"
)

const (
	RouterPort string = ":8009"
	DBType     string = "mysql"
)

func main() {
	// 设置路由
	router.SetRouter(RouterPort)
	// 设置数据库
	controller.SetDB(DBType)
}
