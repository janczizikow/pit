package handlers_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/database"
)

var db *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	db, err = database.Connect("postgres://postgres:postgres@localhost/pit_test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	db.Exec(context.Background(), "DELETE FROM submissions")
	db.Close()
	os.Exit(code)
}
