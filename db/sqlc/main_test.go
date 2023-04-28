package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secretpass@localhost:5432/bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	// if error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// if no error create a new connection 
	testQueries = New(testDB)

	os.Exit(m.Run())
}