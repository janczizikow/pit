package validator_test

import (
	"regexp"
	"testing"

	"github.com/janczizikow/pit/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestMatches(t *testing.T) {
	t.Parallel()

	r := regexp.MustCompile(`\d`)

	assert.Equal(t, true, validator.Matches(r, "0123456789"))
	assert.Equal(t, false, validator.Matches(r, "abc"))
}

func TestIn(t *testing.T) {
	t.Parallel()

	str := "abc"

	assert.Equal(t, true, validator.In(str, "abc", "def"))
	assert.Equal(t, false, validator.In(str, "xyz"))
}

func TestValid(t *testing.T) {
	t.Parallel()

	v := validator.New()

	assert.Equal(t, true, v.Valid())
	v.AddError("test", "must not be valid")
	assert.Equal(t, false, v.Valid())
}

func TestAddError(t *testing.T) {
	t.Parallel()

	v := validator.New()

	assert.Equal(t, make(map[string]string), v.Errors)
	v.AddError("test", "must not be valid")
	assert.Equal(t, map[string]string{"test": "must not be valid"}, v.Errors)
	v.AddError("test", "override")
	assert.Equal(t, map[string]string{"test": "must not be valid"}, v.Errors)
	assert.Equal(t, false, v.Valid())
}

func TestCheck(t *testing.T) {
	t.Parallel()

	v := validator.New()

	v.Check(true, "test", "must not be valid")
	assert.Equal(t, make(map[string]string), v.Errors)

	v.Check(false, "test", "must not be valid")
	assert.Equal(t, map[string]string{"test": "must not be valid"}, v.Errors)
	assert.Equal(t, false, v.Valid())
}
