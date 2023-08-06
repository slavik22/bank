package main

import (
	"database/sql"
	"github.com/slavik22/bank/api"
	db "github.com/slavik22/bank/db/sqlc"
	"github.com/slavik22/bank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("fatal error config file: %w", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot open db connection ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server ", err)
	}
}
