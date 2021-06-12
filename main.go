package main

import (
	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/net"
)

func main() {
	pool := net.Serve()

	keyboard.HandleEvents(func() {
		pool.Broadcast <- []byte("EVENT")
	})
}
