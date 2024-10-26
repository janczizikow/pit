package models_test

import (
	"testing"

	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestSubmissionValid(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		submission models.Submission
		valid      bool
		errors     map[string]string
	}{
		{"is invalid when required fields are empty", models.Submission{}, false, map[string]string{"name": "is required", "class": "is required", "mode": "is required", "tier": "is required", "video": "is required", "duration": "is required"}},
		{"is valid when required fields are not empty", models.Submission{
			Name:     "Test",
			Class:    models.Barbarian,
			Tier:     1,
			Mode:     models.Softcore,
			Video:    "https://youtube.com",
			Duration: 899,
		}, true, map[string]string{}},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := validator.New()
			models.ValidateSubmission(v, &tt.submission)

			assert.Equal(t, tt.valid, v.Valid())

			if !tt.valid {
				assert.Equal(t, tt.errors, v.Errors)
			}
		})
	}
}
