package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)


var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// if no error create a new connection 
	testQueries = New(testDB)

	os.Exit(m.Run())
} 