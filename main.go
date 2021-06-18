package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/HiDeoo/KeyPrompter/cli"
	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/net"
	flags "github.com/jessevdk/go-flags"
)

var Version = "dev"

var opts struct {
	Config  string `short:"c" long:"config" description:"Optional client configuration file" value-name:"PATH"`
	Port    uint   `short:"p" long:"port" description:"Port used to run the web UI" required:"true" value-name:"PORT"`
	Version func() `short:"v" long:"version" description:"Print the KeyPrompter current version."`
}

func main() {
	opts.Version = func() {
		fmt.Printf("KeyPrompter version %s\n", Version)
		os.Exit(0)
	}

	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	var clientConfig *cli.ClientConfig

	if len(opts.Config) > 0 {
		clientConfig = cli.ReadConfig(opts.Config)
	} else {
		clientConfig = new(cli.ClientConfig)
	}

	pool := net.Serve(opts.Port, clientConfig)

	keyboard.HandleEvents(func(keyboardEvent keyboard.KeyboardEvent) {
		eventJson, err := json.Marshal(keyboardEvent)

		if err == nil {
			pool.Broadcast <- eventJson
		}
	})
}
