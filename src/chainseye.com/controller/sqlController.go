package controller

import (
	"chainseye.com/conf"
	"chainseye.com/model"
)

//SetDB 设置数据库链接
func SetDB(dbType string) {
	switch dbType {
	case "mysql":
		model.RegisterDefaultDB(conf.DBADMIN + ":" + conf.DBPASSWORD + "@/" + conf.DBCONNECTDB + "?charset=utf8")
	}
}
