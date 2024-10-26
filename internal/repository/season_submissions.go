package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/models"
)

type ListSubmissionsParams struct {
	Class   string
	Mode    string
	OrderBy string
	Limit   int
	Offset  int
}

// SeasonSubmissionsRepository is the interface that a season submissions repository should conform to.
type SeasonSubmissionsRepository interface {
	List(seasonId int, params ListSubmissionsParams) ([]*models.Submission, int, error)
	Create(submission *models.Submission) (*models.Submission, error)
}

type seasonSubmissionsRepository struct {
	db *pgxpool.Pool
}

// NewSeasonSubmissionsRepository returns a new instance of a season submissions repository.
func NewSeasonSubmissionsRepository(db *pgxpool.Pool) SeasonSubmissionsRepository {
	return &seasonSubmissionsRepository{db: db}
}

func (r *seasonSubmissionsRepository) List(seasonId int, params ListSubmissionsParams) ([]*models.Submission, int, error) {
	count := 0
	submissions := make([]*models.Submission, 0)
	var query string
	if params.OrderBy != "" {
		query = fmt.Sprintf(`SELECT
													COUNT(*) OVER(),
													id,
													"name",
													class,
													tier,
													mode,
													build,
													video,
													duration,
													created_at,
													updated_at
														FROM (
															SELECT DISTINCT ON (name, class, mode) *
															FROM submissions
															WHERE verified = true AND season_id = $1
															ORDER BY name ASC, class ASC, mode ASC, %s
														) sub
												WHERE ($2 = '' OR class = $2::class)
												AND ($3 = '' OR mode = $3::mode)
												ORDER BY %s
												LIMIT $4 OFFSET $5;`, params.OrderBy, params.OrderBy)
	} else {
		query = `SELECT COUNT(*) OVER(),
							id,
							"name",
							class,
							tier,
							mode,
							build,
							video,
							duration,
							created_at,
							updated_at
								FROM (
									SELECT DISTINCT ON (name, class, mode) *
									FROM submissions
									WHERE verified = true AND season_id = $1
									ORDER BY name ASC, class ASC, mode ASC, tier DESC, duration ASC
								) sub
						WHERE ($2 = '' OR class = $2::class)
						AND ($3 = '' OR mode = $3::mode)
						ORDER BY id DESC
						LIMIT $4 OFFSET $5;`
	}
	rows, err := r.db.Query(context.Background(), query, seasonId, params.Class, params.Mode, params.Limit, params.Offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		submission := models.Submission{}
		err = rows.Scan(
			&count,
			&submission.ID,
			&submission.Name,
			&submission.Class,
			&submission.Tier,
			&submission.Mode,
			&submission.Build,
			&submission.Video,
			&submission.Duration,
			&submission.CreatedAt,
			&submission.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		submissions = append(submissions, &submission)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
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
