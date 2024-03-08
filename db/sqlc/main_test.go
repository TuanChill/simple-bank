package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/simple-bank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../.")
	if err != nil {
		log.Fatal("Cannot load config::", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	fmt.Println(testDB)

	if err != nil {
		log.Fatal("Cannot connect db:: ", err)
	}
	defer testDB.Close()
	testQueries = New(testDB)

	os.Exit(m.Run())
}
