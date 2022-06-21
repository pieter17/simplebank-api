package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pieter17/simplebank-api/api"
	db "github.com/pieter17/simplebank-api/db/sqlc"
	"github.com/pieter17/simplebank-api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot run server", err)
	}
}
