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
	ws, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		// TODO(HiDeoo)
		fmt.Println("Error when upgrading connection")

		return
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			// TODO(HiDeoo)
			fmt.Println("Error when reading message")

			return
		}

		fmt.Println(string(p))
	}
}
