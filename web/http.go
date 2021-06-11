package web

import (
	"log"
	"net/http"
	"time"

	"github.com/HiDeoo/KeyPrompter/ui"
)

const Timeout = 10 * time.Second

func ServeUI() {
	uiHandler := ui.AssetHandler()

	server := &http.Server{
		Addr:         ":8484",
		Handler:      uiHandler,
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
	}

	log.Fatal(server.ListenAndServe())
}
