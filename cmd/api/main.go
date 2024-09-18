package main

import (
	"fmt"
	"os"

	"github.com/janczizikow/pit/internal/database"
	"github.com/janczizikow/pit/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SSL_MODE := os.Getenv("DB_SSL_MODE")
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("starting server")
	log.Info().Msg("connecting to postgres DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL_MODE)
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to postgres DB")
	}
	defer db.Close()

	log.Info().Msg("connected to postgres DB")

	server := server.New(db)
	err = server.Run()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to start the server")
	}
}
