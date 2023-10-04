package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/terajari/idompet/util"

	_ "github.com/lib/pq"
)

var testQueries *Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load .env")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = NewStore(conn)
	os.Exit(m.Run())
}
