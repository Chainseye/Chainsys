package conf

import (
	. "chainseye.com/config"
)

var (
	INFO string
)

func InitConfigValue() {
	INFO = Config.String("info::info")
}
