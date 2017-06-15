package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type hub struct {
	// 注册了的连接器
	connections map[*connection]bool

	// 从连接器中发入的信息
	broadcast chan []byte

	// 从连接器中注册请求
	register chan *connection

	// 从连接器中注销请求
	unregister chan *connection
}

var h = hub{
	broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}

type connection struct {
	// websocket 连接器
	ws *websocket.Conn

	// 发送信息的缓冲 channel
	send chan []byte
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		h.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

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

//RunWebSocketBroadcast 独立启动聊天室
func RunWebSocketBroadcast() {
	go h.run()
}

//WebSocketBroadcast 广播
func WebSocketBroadcast(context *gin.Context) {
	w := context.Writer
	r := context.Request
	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
