package main

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/janczizikow/pit/internal/bot"
	"github.com/janczizikow/pit/internal/database"
	"github.com/janczizikow/pit/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	SENTRY_DSN := os.Getenv("SENTRY_DSN")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SSL_MODE := os.Getenv("DB_SSL_MODE")
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("starting server")
	if SENTRY_DSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         SENTRY_DSN,
			Environment: "production",
			// Set TracesSampleRate to 1.0 to capture 100% of transactions for tracing.
			// We recommend adjusting this value in production,
			TracesSampleRate: 0.2,
			EnableTracing:    true,
		}); err != nil {
			log.Error().Err(err).Msg("Sentry initialization failed")
		}
	} else {
		log.Info().Msg("SENTRY_DSN not set, skipping Sentry initialization")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, DB_SSL_MODE)

	log.Info().Msg("running migrations")
	err := database.RunMigrations(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("migrations failed")
	}

	log.Info().Msg("connecting to postgres DB")

	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to postgres DB")
	}
	defer db.Close()

	log.Info().Msg("connected to postgres DB")

	log.Info().Msg("starting discord bot")
	discord, err := bot.Start(db)
	if err != nil {
		log.Error().Err(err).Msg("failed to start discord bot")
	}
	defer discord.Close()

	log.Info().Msg("discord bot is running")

	server := server.New(db)
	err = server.Run()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to start the server")
	}
}
