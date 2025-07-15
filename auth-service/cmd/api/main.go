package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"auth-service.com/data"
	"github.com/gohugoio/hugo/tpl/data"
)

const WEB_PORT = 81

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", WEB_PORT),
		Handler: app.routes(),
	}

	log.Printf("Starting Authentication Service on port %d", WEB_PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Authentication Service NOT STARTED!")
		log.Panicln(err.Error())
	}
}
