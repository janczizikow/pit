package repository

import (
	"github.com/janczizikow/pit/internal/models"
	"github.com/jmoiron/sqlx"
)

// SubmissionsRepository is the interface that a submissions repository should conform to.
type SubmissionsRepository interface {
	List(limit, offset int) ([]*models.Submission, int, error)
	Create(submission *models.Submission) (*models.Submission, error)
}

type submissionsRepository struct {
	db *sqlx.DB
}

// NewSubmissionsRepository returns a new instance of a submissions repository.
func NewSubmissionsRepository(db *sqlx.DB) *submissionsRepository {
	return &submissionsRepository{db: db}
}

func (r *submissionsRepository) List(limit, offset int) ([]*models.Submission, int, error) {
	return nil, 0, nil
}

func (r *submissionsRepository) Create(submission *models.Submission) (*models.Submission, error) {
	newSubmission := models.Submission{}
	return &newSubmission, nil
}
