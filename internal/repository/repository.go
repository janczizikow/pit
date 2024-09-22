package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	db                *pgxpool.Pool
	Seasons           SeasonsRepository
	SeasonSubmissions SeasonSubmissionsRepository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:                db,
		Seasons:           NewSeasonsRepository(db),
		SeasonSubmissions: NewSeasonSubmissionsRepository(db),
	}
}
