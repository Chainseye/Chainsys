package main

import (
	"chainsWebSite_Back/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("start")
	http.HandleFunc("/", controller.Index)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("end")
}
