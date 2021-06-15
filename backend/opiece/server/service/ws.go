package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func NewWSUpgrader(c *gin.Context) *websocket.Conn {
	fmt.Println(c.Request.Header)
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	return conn
}
