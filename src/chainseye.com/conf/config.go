package conf

import (
	"errors"

	"github.com/beego/config"
)

//Config 导出Config
var Config config.Configer

//ConfigInit 读取conf文件
func ConfigInit(path ...string) error {
	configFile := "src/chainseye.com/config/config.conf"
	if len(path) > 0 && path[0] != "" {
		configFile = path[0]
	}
	var err error
	Config, err = config.NewConfig("ini", configFile)
	if err != nil {
		return errors.New("open config file " + configFile + " failed, err is " + err.Error())
	}
	return nil
}
