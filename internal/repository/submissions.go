package repository

import (
	"database/sql"
	"fmt"

	"github.com/janczizikow/pit/internal/models"
	"github.com/jmoiron/sqlx"
)

// SubmissionsRepository is the interface that a submissions repository should conform to.
type SubmissionsRepository interface {
	List(limit, offset int, class, orderBy string) ([]*models.Submission, int, error)
	Create(submission *models.Submission) (*models.Submission, error)
}

type submissionsRepository struct {
	db *sqlx.DB
}

// NewSubmissionsRepository returns a new instance of a submissions repository.
func NewSubmissionsRepository(db *sqlx.DB) *submissionsRepository {
	return &submissionsRepository{db: db}
}

func (r *submissionsRepository) List(limit, offset int, class, orderBy string) ([]*models.Submission, int, error) {
	count := 0
	submissions := make([]*models.Submission, 0)
	var query string
	var where string
	if class != "" {
		where = "WHERE class = $3"
	}
	if orderBy != "" {
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
												LIMIT $1 OFFSET $2;`, orderBy, where, orderBy)
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
	var rows *sql.Rows
	var err error
	if class != "" {
		rows, err = r.db.Query(query, limit, offset, class)
	} else {
		rows, err = r.db.Query(query, limit, offset)
	}
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
						RETURNING *`
	err := r.db.QueryRowx(
		query,
		submission.Name,
		submission.Class,
		submission.Tier,
		submission.Mode,
		submission.Build,
		submission.Video,
		submission.Duration,
	).StructScan(&newSubmission)
	return &newSubmission, err
}
