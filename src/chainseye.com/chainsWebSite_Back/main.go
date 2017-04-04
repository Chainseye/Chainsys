package main

import (
	"chainseye.com/chainsWebSite_Back/router"
)

func main() {
	//http.Handle("/", websocket.Handler(Echo))

	// if err := http.ListenAndServe(":9999", nil); err != nil {
	// 	log.Fatal("ListenAndServe:", err)
	// }
	router.SetRouter()
}
