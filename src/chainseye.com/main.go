package main

import (
	"chainseye.com/controller"
	"chainseye.com/router"
)

const (
	//RouterPort 路由端口号
	RouterPort string = ":8009"
	//DBType 数据库类型
	DBType string = "mysql"
)

func main() {
	// 设置数据库
	controller.SetDB(DBType)
	// 设置路由
	router.SetRouter(RouterPort)
}
