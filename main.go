package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/terajari/idompet/api"
	database "github.com/terajari/idompet/database/sqlc"
	"github.com/terajari/idompet/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load .env")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Address)
	if err != nil {
		log.Fatal("cannot load server")
	}
}
