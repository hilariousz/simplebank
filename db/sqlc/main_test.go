package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/hilariousz/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries // Queries for testing
var testDB *sql.DB       // Database for testing

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
