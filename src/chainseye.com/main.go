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

// 没什么用处，os.Open()报错
// func init() {
// 	//
// 	cfg, err := goconfig.LoadConfigFile("ini/config.ini")
// 	if err != nil {
// 		log.Println("=============>")
// 		log.Println(err)
// 		log.Println("<=============")
// 		return
// 	}
// 	str, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "info")
// 	fmt.Printf("%s\n", str)
// }

func main() {
	// 设置数据库
	controller.SetDB(DBType)
	// 设置路由
	router.SetRouter(RouterPort)
}
