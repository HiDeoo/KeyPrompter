package main

import (
	"encoding/json"
	"os"

	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/net"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Port uint `short:"p" long:"port" description:"Port used to run the web UI" required:"true"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	pool := net.Serve(opts.Port)

	keyboard.HandleEvents(func(keyboardEvent keyboard.KeyboardEvent) {
		eventJson, err := json.Marshal(keyboardEvent)

		if err == nil {
			pool.Broadcast <- eventJson
		}
	})
}
