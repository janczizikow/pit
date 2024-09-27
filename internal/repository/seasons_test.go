package repository_test

import (
	"database/sql"
	"testing"

	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListSeasons(t *testing.T) {
	repo := repository.New(db)

	_, count, err := repo.Seasons.List()
	require.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestCurrentSeason(t *testing.T) {

	repo := repository.New(db)

	t.Run("returns error when no seasons exist", func(t *testing.T) {
		_, err := repo.Seasons.Current()
		require.Error(t, err)
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})

	t.Run("returns current season", func(t *testing.T) {
		_, err := repo.Seasons.Create(&models.Season{Name: "Test", End: nil})
		require.NoError(t, err)

		season, err := repo.Seasons.Current()
		require.NoError(t, err)
		assert.NotNil(t, season)
		t.Cleanup(func() {
			_, err = db.Exec(ctx, "DELETE FROM seasons WHERE id = $1", season.ID)
			require.NoError(t, err)
		})
	})
}
