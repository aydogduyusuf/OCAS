package main

import (
	"database/sql"
	"group33/ocas/src/API"
	DB "group33/ocas/src/DB/sqlc"
	"group33/ocas/src/Helpers"
	"log"
)

func main() {
	config, err := Helpers.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect do db", err)
	}

	store := DB.NewStore(conn)

	server, err := API.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}
