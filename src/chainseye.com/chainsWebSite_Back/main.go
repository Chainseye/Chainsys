package main

import (
	"log"
	"net/http"

	"chainseye.com/chainsWebSite_Back/controller"
	"chainseye.com/chainsWebSite_Back/router"
)

func main() {
	http.HandleFunc("/", controller.Index)

	if err := http.ListenAndServe(":8099", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	router.SetRouter()
}
