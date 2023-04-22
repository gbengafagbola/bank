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

func TestMain(m *testing.M){

	connection, err := sql.Open(dbDriver, dbSource)
	// if error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// if no error create a new connection 
	testQueries = New(connection)

	os.Exit(m.Run())
}