package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/slavik22/bank/util"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
	conf, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatal("Cannot read env ", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}

	testStore = NewStore(conn)

	os.Exit(m.Run())
}
