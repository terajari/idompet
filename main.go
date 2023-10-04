package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/terajari/idompet/api"
	database "github.com/terajari/idompet/database/sqlc"
)

const (
	Driver = "postgres"
	Source = "postgresql://postgres:1234@localhost:5432/idompet?sslmode=disable"
)

func main() {
	conn, err := sql.Open(Driver, Source)
	if err != nil {
		log.Fatal(err)
	}
	store := database.NewStore(conn)
	server := api.NewServer(store)

	server.Router.Run()
}
