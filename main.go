package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/simple-bank/api"
	db "github.com/simple-bank/db/sqlc"
	"github.com/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config::", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect db:: ", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:: ", err)
	}
}
