package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//WebSocketMiddleWare 中间件
func WebSocketMiddleWare(context *gin.Context) {
	fmt.Println("WsMiddleWare")
}

//WebSocketConnect 链接WebSocket
func WebSocketConnect(context *gin.Context) {
	w := context.Writer
	r := context.Request
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade : %+v", err)
		return
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
