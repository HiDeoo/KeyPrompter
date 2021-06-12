package net

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func webSocketHandler(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		// TODO(HiDeoo)
		fmt.Println("Error when upgrading the connection to the WebSocket protocol.")
		return
	}

	client := &Client{conn: conn}

	go client.write()
	go client.read()
}
