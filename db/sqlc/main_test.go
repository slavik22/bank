package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:pass@localhost:5432/testbank?sslmode=disable"
)

var testStore *Store

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}

	testStore = NewStore(conn)

	os.Exit(m.Run())
}
