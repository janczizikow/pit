package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/database"
)

var db *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	var err error
	db, err = database.Connect("postgres://postgres:postgres@localhost/pit_test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	ctx = context.Background()
	code := m.Run()
	db.Exec(ctx, "DELETE FROM submissions")
	db.Exec(ctx, "DELETE FROM seasons")
	db.Close()
	os.Exit(code)
}
