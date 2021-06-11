package main

import (
	"github.com/HiDeoo/KeyPrompter/keyboard"
	"github.com/HiDeoo/KeyPrompter/web"
)

func main() {
	web.ServeUI()

	keyboard.HandleEvents()
}
