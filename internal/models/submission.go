package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/janczizikow/pit/internal/validator"
)

// Modes
const (
	Softcore = "softcore"
	Hardcore = "hardcore"
)

// Classes
const (
	Barbarian   = "barbarian"
	Druid       = "druid"
	Necromancer = "necromancer"
	Rogue       = "rogue"
	Sorcerer    = "sorcerer"
)

type Submission struct {
	ID              int      `db:"id" json:"id"`
	Name            string   `db:"name" json:"name"`
	Class           string   `db:"class" json:"class"`
	Tier            int      `db:"tier" json:"tier"`
	Mode            string   `db:"mode" json:"mode"`
	Build           string   `db:"build" json:"build"`
	Video           string   `db:"video" json:"video"`
	Duration        duration `db:"-" json:"duration"`
	DurationSeconds int      `db:"duration" json:"-"`
	// Timestamps

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func ValidateSubmission(v *validator.Validator, submission *Submission) {
	v.Check(submission.Name != "", "name", "is required")
	v.Check(len(submission.Name) <= 255, "name", "must be not more than 255 bytes long")

	v.Check(submission.Class != "", "class", "is required")
	v.Check(validator.In(submission.Class, Barbarian, Druid, Necromancer, Rogue, Sorcerer), "class", "is invalid")

	v.Check(submission.Mode != "", "mode", "is required")
	v.Check(validator.In(submission.Mode, Softcore, Hardcore), "mode", "is invalid")

	v.Check(submission.Tier != 0, "tier", "is required")
	v.Check(submission.Tier > 0, "tier", "must be greater than zero")
	v.Check(submission.Tier <= 200, "tier", "must be a maximum of 200")

	v.Check(submission.Video != "", "video", "is required")

	v.Check(submission.Duration.Seconds() != 0, "duration", "is required")
	v.Check(submission.Duration.Minutes() <= 15, "duration", "must be a maximum of 15 minutes")
}

type duration struct {
	time.Duration
}

func (d duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
