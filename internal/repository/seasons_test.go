package repository_test

import (
	"testing"

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
