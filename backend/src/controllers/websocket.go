package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
	return true
	},
}

var wsConnections []*websocket.Conn

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wsConnections = append(wsConnections, conn)
	//_ = conn.WriteMessage(websocket.TextMessage, []byte("Hello there!"))
}
