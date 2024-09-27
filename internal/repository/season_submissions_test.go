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
	t.Parallel()

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
	newSubmission.ID = submission.ID
	newSubmission.CreatedAt = submission.CreatedAt
	newSubmission.UpdatedAt = submission.UpdatedAt
	assert.Equal(t, newSubmission, submission)
	t.Cleanup(func() {
		db.Exec(ctx, "DELETE FROM submissions WHERE id = $1", submission.ID)
	})
}
