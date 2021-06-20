package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/HiDeoo/KeyPrompter/cli"
	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/net"
)

var Version = "development version"

func main() {
	config := flag.String("c", "", "Optional client configuration file")
	port := flag.Int("p", 8484, "Port used to run the web UI")
	version := flag.Bool("v", false, "Print the current KeyPrompter version")

	flag.Parse()

	if *version {
		fmt.Printf("KeyPrompter %s\n", Version)
		os.Exit(0)
	}

	var clientConfig *cli.ClientConfig

	if len(*config) > 0 {
		clientConfig = cli.ReadConfig(*config)
	} else {
		clientConfig = new(cli.ClientConfig)
	}

	pool := net.Serve(*port, clientConfig)

	keyboard.HandleEvents(func(keyboardEvent keyboard.KeyboardEvent) {
		eventJson, err := json.Marshal(keyboardEvent)

		if err == nil {
			pool.Broadcast <- eventJson
		}
	})
}
