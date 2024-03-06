package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/simple-bank/api"
	db "github.com/simple-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:admin@localhost:5432/simple-bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect db:: ", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server:: ", err)
	}
}
