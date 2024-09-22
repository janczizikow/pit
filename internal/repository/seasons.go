package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/models"
)

// SeasonsRepository is the interface that a seasons repository should conform to.
type SeasonsRepository interface {
	List() ([]*models.Season, int, error)
}

type seasonsRepository struct {
	db *pgxpool.Pool
}

// NewSeasonsRepository returns a new instance of a seasonsRepository.
func NewSeasonsRepository(db *pgxpool.Pool) SeasonsRepository {
	return &seasonsRepository{db: db}
}

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
