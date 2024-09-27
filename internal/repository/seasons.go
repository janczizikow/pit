package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/models"
)

// SeasonsRepository is the interface that a seasons repository should conform to.
type SeasonsRepository interface {
	// List returns a list of seasons.
	List() ([]*models.Season, int, error)
	// Current returns the current season.
	Current() (*models.Season, error)
	// Create creates a new season.
	Create(season *models.Season) (*models.Season, error)
}

type seasonsRepository struct {
	db *pgxpool.Pool
}

// NewSeasonsRepository returns a new instance of a seasonsRepository.
func NewSeasonsRepository(db *pgxpool.Pool) SeasonsRepository {
	return &seasonsRepository{db: db}
}

// List returns a list of seasons.
func (r *seasonsRepository) List() ([]*models.Season, int, error) {
	count := 0
	seasons := make([]*models.Season, 0)
	query := `SELECT COUNT(*) OVER(), id, seasons.name, pit, seasons.start, seasons.end, created_at, updated_at
						FROM seasons
						ORDER BY id DESC
						LIMIT 100;`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		season := models.Season{}
		err = rows.Scan(
			&count,
			&season.ID, &season.Name, &season.Pit,
			&season.Start, &season.End,
			&season.CreatedAt, &season.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		seasons = append(seasons, &season)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	return seasons, count, nil
}

// Current returns the current season.
func (r *seasonsRepository) Current() (*models.Season, error) {
	season := models.Season{}
	query := `SELECT id, "name", "start", "end", pit, created_at, updated_at
						FROM seasons
						WHERE seasons.end > NOW() OR seasons.end IS NULL
						ORDER BY id ASC
						LIMIT 1;`
	err := r.db.QueryRow(context.Background(), query).Scan(
		&season.ID, &season.Name, &season.Start, &season.End, &season.Pit,
		&season.CreatedAt, &season.UpdatedAt,
	)
	return &season, err
}

// Create creates a new season.
func (r *seasonsRepository) Create(season *models.Season) (*models.Season, error) {
	newSeason := models.Season{}
	query := `INSERT INTO seasons (name, pit, "start", "end")
						VALUES ($1, $2, $3, $4)
						RETURNING id,"name",pit,"start","end",created_at,updated_at;`
	err := r.db.QueryRow(
		context.Background(),
		query,
		season.Name,
		season.Pit,
		season.Start,
		season.End,
	).Scan(
		&newSeason.ID,
		&newSeason.Name,
		&newSeason.Pit,
		&newSeason.Start,
		&newSeason.End,
		&newSeason.CreatedAt,
		&newSeason.UpdatedAt,
	)
	return &newSeason, err
}
