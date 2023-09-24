package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	Driver = "postgres"
	Source = "postgresql://postgres:1234@localhost:5432/idompet?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(Driver, Source)
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
