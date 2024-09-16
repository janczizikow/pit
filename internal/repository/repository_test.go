package repository_test

import (
	"log"
	"os"
	"testing"

	"github.com/janczizikow/pit/internal/database"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	var err error
	db, err = database.Connect("postgres://postgres:postgres@localhost/pit_test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	db.MustExec("DELETE FROM submissions")
	db.Close()
	os.Exit(code)
}
