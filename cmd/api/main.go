package main

import (
	"os"

	"github.com/janczizikow/pit/internal/database"
	"github.com/janczizikow/pit/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	DSN := os.Getenv("DB_DSN")
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("starting server")
	log.Info().Msg("connecting to postgres DB")

	db, err := database.Connect(DSN)
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
