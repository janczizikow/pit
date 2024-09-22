package models_test

import (
	"testing"

	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestSeasonValid(t *testing.T) {
	cases := []struct {
		season models.Season
		valid  bool
		errors map[string]string
	}{
		{models.Season{}, false, map[string]string{"name": "is required"}},
		{models.Season{
			Name: "Test",
		}, true, map[string]string{}},
	}

	for _, tt := range cases {
		v := validator.New()
		models.ValidateSeason(v, &tt.season)

		assert.Equal(t, tt.valid, v.Valid())

		if !tt.valid {
			assert.Equal(t, tt.errors, v.Errors)
		}

	}
}
