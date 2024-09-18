package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	db          *pgxpool.Pool
	Submissions SubmissionsRepository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:          db,
		Submissions: NewSubmissionsRepository(db),
	}
}
