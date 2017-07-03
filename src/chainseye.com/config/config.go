package config

import (
	"errors"

	"github.com/beego/config"
)

var Config config.Configer

func Config_Init(path ...string) error {
	configFile := "ini/app.conf"
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
