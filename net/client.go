package net

import (
	"encoding/json"
	"time"

	"github.com/HiDeoo/KeyPrompter/cli"
	"github.com/gorilla/websocket"
)

const (
	pingInterval          = 50 * time.Second
	readDeadlineDuration  = 60 * time.Second
	writeDeadlineDuration = 10 * time.Second
)

type Client struct {
	conn *websocket.Conn
	pool *Pool
	send chan []byte
}

func (client *Client) read() {
	defer func() {
		client.pool.unregister <- client
		client.conn.Close()
	}()

	client.conn.SetReadDeadline(time.Now().Add(readDeadlineDuration))
	client.conn.SetPongHandler(func(appData string) error {
		client.conn.SetReadDeadline(time.Now().Add(readDeadlineDuration))
		return nil
	})

	for {
		_, p, err := client.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				cli.PrintServerError(err)
			}

			break
		}

		var message Message

		err = json.Unmarshal(p, &message)

		if err == nil {
			cli.PrintUIError(message.ErrorMessage.Message)
		}
	}
}

func (client *Client) write() {
	pinger := time.NewTicker(pingInterval)

	defer func() {
		pinger.Stop()
		client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeDeadlineDuration))

			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := client.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				return
			}

			writer.Write(message)

			if err := writer.Close(); err != nil {
				return
			}
		case <-pinger.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeDeadlineDuration))

			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
