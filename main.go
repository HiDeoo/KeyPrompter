package main

import (
	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/net"
)

func main() {
	net.Serve()

	keyboard.HandleEvents()
}
