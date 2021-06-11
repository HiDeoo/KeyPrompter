package net

import (
	"log"
	"net/http"
	"time"

	"github.com/HiDeoo/KeyPrompter/ui"
)

const Timeout = 10 * time.Second

func Serve() {
	addRouteHandlers()

	// TODO(HiDeoo)
	log.Fatal(http.ListenAndServe(":8484", nil))
}

func addRouteHandlers() {
	http.Handle("/", ui.AssetHandler())

	http.HandleFunc("/ws", webSocketHandler)
}
