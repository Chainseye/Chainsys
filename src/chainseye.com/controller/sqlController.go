package controller

import (
	"chainseye.com/model"
)

// 数据库信息
const (
	DBAdmin     string = "root"
	DBPassword  string = "a123123"
	DBConnectDB string = "GoLang_DB"
)

//SetDB 设置数据库链接
func SetDB(dbType string) {
	switch dbType {
	case "mysql":
		model.RegisterDefaultDB(DBAdmin + ":" + DBPassword + "@/" + DBConnectDB + "?charset=utf8")
	}
}
