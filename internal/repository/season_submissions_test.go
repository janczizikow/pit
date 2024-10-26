package repository_test

import (
	"testing"
	"time"

	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateSubmission(t *testing.T) {
	repo := repository.New(db)
	newSubmission := &models.Submission{
		Name:     "Test",
		Class:    models.Sorcerer,
		Mode:     models.Softcore,
		Tier:     150,
		Duration: int((time.Duration(14) * time.Minute).Seconds()),
		Video:    "https://youtube.com",
	}

	submission, err := repo.SeasonSubmissions.Create(newSubmission)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := db.Exec(ctx, "DELETE FROM submissions WHERE id = $1", submission.ID)
		require.NoError(t, err)
	})

	newSubmission.ID = submission.ID
	newSubmission.CreatedAt = submission.CreatedAt
	newSubmission.UpdatedAt = submission.UpdatedAt
	assert.Equal(t, newSubmission, submission)
}

func TestListSeasonSubmissions(t *testing.T) {
	repo := repository.New(db)

	t.Run("returns empty list when no submissions exist", func(t *testing.T) {
		submissions, count, err := repo.SeasonSubmissions.List(1, repository.ListSubmissionsParams{})
		require.NoError(t, err)
		assert.Equal(t, 0, count)
		assert.Equal(t, 0, len(submissions))
	})

	t.Run("returns list of verified submissions for a given season", func(t *testing.T) {
		season, err := repo.Seasons.Create(&models.Season{Name: "Test", End: nil})
		require.NoError(t, err)
		t.Cleanup(func() {
			_, err := db.Exec(ctx, "DELETE FROM seasons WHERE id = $1", season.ID)
			require.NoError(t, err)
		})

		sub := models.Submission{
			Name:     "Test",
			Class:    models.Sorcerer,
			Mode:     models.Softcore,
			Tier:     150,
			Duration: int((time.Duration(14) * time.Minute).Seconds()),
			Video:    "https://youtube.com",
			SeasonId: &season.ID,
		}
		s, err := repo.SeasonSubmissions.Create(&sub)
		require.NoError(t, err)

		submissions, count, err := repo.SeasonSubmissions.List(season.ID, repository.ListSubmissionsParams{Limit: 10, Offset: 0})
		require.NoError(t, err)
		assert.Equal(t, 0, count)
		assert.Equal(t, 0, len(submissions))

		_, err = db.Exec(ctx, "UPDATE submissions SET verified = true WHERE id = $1", s.ID)
		require.NoError(t, err)
		t.Cleanup(func() {
			_, err := db.Exec(ctx, "DELETE FROM submissions WHERE id = $1", s.ID)
			require.NoError(t, err)
		})

		submissions, count, err = repo.SeasonSubmissions.List(season.ID, repository.ListSubmissionsParams{Limit: 10, Offset: 0})
		require.NoError(t, err)
		assert.Equal(t, 1, count)
		assert.Equal(t, 1, len(submissions))
		assert.Equal(t, sub.Name, submissions[0].Name)
	})
}
