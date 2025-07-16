package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"auth-service.com/data"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const WEB_PORT = 81

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func initEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnvs()

	dbConnection := connectToDB()
	if dbConnection == nil {
		log.Panicln("DB IS NOT CONNECTED...!!")
		return
	}

	app := Config{
		DB:     dbConnection,
		Models: data.New(dbConnection),
	}

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

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Panicln("*** Erorr with Open!" + err.Error())
		return nil, err
	}

	err2 := db.Ping()
	if err2 != nil {
		log.Panicln("*** Erorr with Ping Request!" + err.Error())
		return nil, err
	}

	return db, nil
}

var tryCount int64

func connectToDB() *sql.DB {
	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	log.Println("*** CONNECTION STRING IS -> " + connStr)

	for {
		connection, err := openDB(connStr)

		if err != nil {
			log.Println("Postgres not yet ready...!")

			tryCount++
			if tryCount > 10 {
				log.Panicln("Try to connect PostgreSQL 10 times and stop..!")
				return nil
			}

			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("Postgres is connected..!")
		return connection
	}
}
