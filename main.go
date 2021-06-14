package main

import (
	"encoding/json"

	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/net"
)

func main() {
	pool := net.Serve()

	keyboard.HandleEvents(func(keyboardEvent keyboard.KeyboardEvent) {
		eventJson, err := json.Marshal(keyboardEvent)

		if err == nil {
			pool.Broadcast <- eventJson
		}
	})
}
