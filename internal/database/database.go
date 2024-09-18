package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxOpenConns = 30
	maxIdleTime  = time.Second * 239
)

// Connect to a database and verify with a ping.
func Connect(dataSourceName string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	config.MaxConns = maxOpenConns
	config.MaxConnIdleTime = maxIdleTime

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	conn, err := pgxpool.NewWithConfig(ctxWithTimeout, config)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctxWithTimeout); err != nil {
		return nil, err
	}

	return conn, nil
}
