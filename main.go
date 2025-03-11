package main

import (
	"database/sql"
	"log"

	"github.com/gdygd/simplebank/api"
	db "github.com/gdygd/simplebank/db/sqlc"
	"github.com/gdygd/simplebank/util"
	_ "github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connectd db..", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
