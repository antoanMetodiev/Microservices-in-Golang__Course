package main

import (
	"fmt"
	"log"
	"net/http"
)

const WEB_PORT = 81

type Config struct{}

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", WEB_PORT),
		Handler: app.router(),
	}

	log.Printf("Starting Authentication Service on port %d", WEB_PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Authentication Service NOT STARTED!")
		log.Panicln(err.Error())
	}
}
