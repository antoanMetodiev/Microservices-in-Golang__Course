package main

import (
	"fmt"
	"log"
	"net/http"
)

const WEB_PORT = 80

type Config struct{}

func main() {
	app := Config{}

	// create http server:
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", WEB_PORT),
		Handler: app.routes(),
	}

	// start http server:
	log.Printf("Starting broker service on port %d\n", WEB_PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
		return
	}
}
