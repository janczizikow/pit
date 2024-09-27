package models_test

import (
	"testing"

	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestSeasonValid(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		season models.Season
		valid  bool
		errors map[string]string
	}{
		{"is invalid when name is empty", models.Season{}, false, map[string]string{"name": "is required"}},
		{
			"is valid when name is not empty",
			models.Season{
				Name: "Test",
			}, true, map[string]string{}},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := validator.New()
			models.ValidateSeason(v, &tt.season)

			assert.Equal(t, tt.valid, v.Valid())

			if !tt.valid {
				assert.Equal(t, tt.errors, v.Errors)
			}
		})
	}
}
