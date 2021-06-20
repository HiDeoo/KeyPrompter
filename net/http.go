package net

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/HiDeoo/KeyPrompter/cli"
	"github.com/HiDeoo/KeyPrompter/ui"
)

const Timeout = 10 * time.Second

func Serve(port int, clientConfig *cli.ClientConfig) *Pool {
	pool := newPool()

	go pool.run()

	addRouteHandlers(pool, clientConfig)

	go func() {
		fmt.Printf("You can now view the KeyPrompter UI in the browser: %s%s.\n", cli.Green("http://localhost:"), cli.BoldGreen(port))

		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	}()

	return pool
}

func addRouteHandlers(pool *Pool, clientConfig *cli.ClientConfig) {
	http.Handle("/", ui.AssetHandler())

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		webSocketHandler(rw, r, pool, clientConfig)
	})
}
