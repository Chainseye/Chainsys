package controller

import (
	"fmt"
	"net/http"
)

//CheckErr 打印错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
