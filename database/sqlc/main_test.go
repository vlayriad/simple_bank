package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"kaffein/simple_bank/util"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQuery *Queries
	testDB    *sql.DB
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connection to database ", err.Error())
	}
	testQuery = New(testDB)
	os.Exit(m.Run())
}
