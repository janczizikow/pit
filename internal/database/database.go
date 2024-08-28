package database

import (
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns = 30
	maxIdleConns = 30
	MaxIdleTime  = 20 * time.Minute
)

// Connect to a database and verify with a ping.
func Connect(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(MaxIdleTime)

	return db, nil
}
