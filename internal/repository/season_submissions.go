package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/models"
)

// SeasonSubmissionsRepository is the interface that a season submissions repository should conform to.
type SeasonSubmissionsRepository interface {
	List(params ListSubmissionsParams) ([]*models.Submission, int, error)
	Create(submission *models.Submission) (*models.Submission, error)
}

type seasonSubmissionsRepository struct {
	db *pgxpool.Pool
}

// NewSeasonSubmissionsRepository returns a new instance of a season submissions repository.
func NewSeasonSubmissionsRepository(db *pgxpool.Pool) SeasonSubmissionsRepository {
	return &seasonSubmissionsRepository{db: db}
}

func (r *seasonSubmissionsRepository) List(params ListSubmissionsParams) ([]*models.Submission, int, error) {
	count := 0
	submissions := make([]*models.Submission, 0)

	return submissions, count, nil
}

func (r *seasonSubmissionsRepository) Create(submission *models.Submission) (*models.Submission, error) {
	newSubmission := models.Submission{}
	query := `INSERT INTO submissions (season_id, name, class, tier, mode, build, video, duration)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
						RETURNING id,name,class,tier,mode,build,video,duration,created_at,updated_at;`
	err := r.db.QueryRow(
		context.Background(),
		query,
		submission.SeasonId,
		submission.Name,
		submission.Class,
		submission.Tier,
		submission.Mode,
		submission.Build,
		submission.Video,
		submission.Duration,
	).Scan(
		&newSubmission.ID,
		&newSubmission.Name,
		&newSubmission.Class,
		&newSubmission.Tier,
		&newSubmission.Mode,
		&newSubmission.Build,
		&newSubmission.Video,
		&newSubmission.Duration,
		&newSubmission.CreatedAt,
		&newSubmission.UpdatedAt,
	)
	return &newSubmission, err
}
