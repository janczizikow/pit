package models

import (
	"time"

	"github.com/janczizikow/pit/internal/validator"
)

type Season struct {
	ID    int        `db:"id" json:"id"`
	Name  string     `db:"name" json:"name"`
	Pit   bool       `db:"pit" json:"pit"`
	Start time.Time  `db:"start" json:"start"`
	End   *time.Time `db:"end" json:"end"`

	// Timestamps

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func ValidateSeason(v *validator.Validator, season *Season) {
	v.Check(season.Name != "", "name", "is required")
}
