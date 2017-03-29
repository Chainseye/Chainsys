package controller

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// 向客户端写入内容
	fmt.Println("////////////////")
	fmt.Fprintf(w, "Hello World! \n")
	fmt.Println("aaaa")
	Mysql_Driver_test()
}
