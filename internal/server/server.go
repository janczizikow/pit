package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	zlog "github.com/rs/zerolog/log"
)

type Server struct {
	db *pgxpool.Pool
}

// New instantiates a new server.
func New(db *pgxpool.Pool) *Server {
	return &Server{db: db}
}

// Run starts a HTTP server listening for connections.
func (s *Server) Run() error {
	srv := http.Server{
		Addr:         ":8080",
		Handler:      Router(s),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     log.New(zlog.With().Str("level", "error").Logger(), "", 0),
	}
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		signal := <-quit

		zlog.Info().Str("signal", signal.String()).Msg("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}
