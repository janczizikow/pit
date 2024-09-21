package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/models"
)

type ListSubmissionsParams struct {
	Limit   int
	Offset  int
	Class   string
	Mode    string
	OrderBy string
}

// SubmissionsRepository is the interface that a submissions repository should conform to.
type SubmissionsRepository interface {
	List(params ListSubmissionsParams) ([]*models.Submission, int, error)
	Create(submission *models.Submission) (*models.Submission, error)
}

type submissionsRepository struct {
	db *pgxpool.Pool
}

// NewSubmissionsRepository returns a new instance of a submissions repository.
func NewSubmissionsRepository(db *pgxpool.Pool) *submissionsRepository {
	return &submissionsRepository{db: db}
}

func (r *submissionsRepository) List(params ListSubmissionsParams) ([]*models.Submission, int, error) {
	count := 0
	submissions := make([]*models.Submission, 0)
	var query string
	var where string
	if params.Class != "" {
		where = fmt.Sprintf("WHERE class = '%s'", params.Class)
	}
	if params.Mode != "" {
		switch where {
		case "":
			where = fmt.Sprintf("WHERE mode = '%s'", params.Mode)
		default:
			where += fmt.Sprintf("AND mode = '%s'", params.Mode)
		}
	}
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
															SELECT DISTINCT ON (name, class) *
															FROM submissions
															ORDER BY name ASC, class ASC, %s
														) sub
												%s
												ORDER BY %s
												LIMIT $1 OFFSET $2;`, params.OrderBy, where, params.OrderBy)
	} else {
		query = fmt.Sprintf(`SELECT COUNT(*) OVER(),
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
									SELECT DISTINCT ON (name, class) *
									FROM submissions
									ORDER BY name ASC, class ASC, tier DESC, duration ASC
								) sub
						%s
						ORDER BY id DESC
						LIMIT $1 OFFSET $2;`, where)
	}
	var rows pgx.Rows
	var err error
	rows, err = r.db.Query(context.Background(), query, params.Limit, params.Offset)
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

func (r *submissionsRepository) Create(submission *models.Submission) (*models.Submission, error) {
	newSubmission := models.Submission{}
	query := `INSERT INTO submissions (name, class, tier, mode, build, video, duration)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
						RETURNING id,name,class,tier,mode,build,video,duration,created_at,updated_at;`
	err := r.db.QueryRow(
		context.Background(),
		query,
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
