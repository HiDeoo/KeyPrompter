package net

import (
	"log"
	"net/http"
	"time"

	"github.com/HiDeoo/KeyPrompter/ui"
)

const Timeout = 10 * time.Second

func Serve() *Pool {
	pool := newPool()

	go pool.run()

	addRouteHandlers(pool)

	go func() {
		// TODO(HiDeoo)
		log.Fatal(http.ListenAndServe(":8484", nil))
	}()

	return pool
}

func addRouteHandlers(pool *Pool) {
	http.Handle("/", ui.AssetHandler())

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		webSocketHandler(rw, r, pool)
	})
}
