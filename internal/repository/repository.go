package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	db          *sqlx.DB
	Submissions SubmissionsRepository
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db:          db,
		Submissions: NewSubmissionsRepository(db),
	}
}
