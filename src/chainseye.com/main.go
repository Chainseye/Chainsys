package main

import (
	"fmt"

	"chainseye.com/conf"
	"chainseye.com/controller"
	"chainseye.com/router"
)

func init() {
	err := conf.ConfigInit("src/chainseye.com/config/config.conf")
	if err != nil {
		fmt.Printf("init config failed, err is %v\n", err)
	}
	conf.InitConfigValue()
}

func main() {
	// 设置数据库
	controller.SetDB(conf.DBTYPE)
	// 设置路由
	router.SetRouter(conf.ROUTERPORT)
}
