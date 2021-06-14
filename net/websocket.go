package net

import (
	"net/http"

	"github.com/HiDeoo/KeyPrompter/cli"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func webSocketHandler(rw http.ResponseWriter, r *http.Request, pool *Pool) {
	conn, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		cli.PrintServerError("Error when upgrading the connection to the WebSocket protocol.")
		return
	}

	client := &Client{conn: conn, pool: pool, send: make(chan []byte, 256)}
	client.pool.register <- client

	go client.write()
	go client.read()
}
