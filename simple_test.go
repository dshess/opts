package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleOption(t *testing.T) {
	wantTrue := false
	stayFalse := false

	args := []string{
		"--want-true",
		"left",
	}
	ret, err := NewOpts().
		SimpleOption("want-true", &wantTrue).
		SimpleOption("stay-false", &stayFalse).
		ProcessArgs(args)
	require.Nil(t, err)
	assert.True(t, wantTrue)
	assert.False(t, stayFalse)
	assert.Equal(t, []string{"left"}, ret)
}

func TestNegatableOption(t *testing.T) {
	wantTrue := false
	wantFalse := true
	stayFalse := false
	stayTrue := true

	args := []string{
		"--want-true",
		"--nowant-false",
		"left",
	}
	ret, err := NewOpts().
		NegatableOption("want-true", &wantTrue).
		NegatableOption("want-false", &wantFalse).
		NegatableOption("stay-false", &stayFalse).
		NegatableOption("stay-true", &stayTrue).
		ProcessArgs(args)
	require.Nil(t, err)
	assert.True(t, wantTrue)
	assert.False(t, wantFalse)
	assert.False(t, stayFalse)
	assert.True(t, stayTrue)
	assert.Equal(t, []string{"left"}, ret)
}
