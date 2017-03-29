package controller

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// 向客户端写入内容
	fmt.Fprintf(w, "Hello World! \n")
	TestFmt()
}
