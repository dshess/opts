package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingOption(t *testing.T) {
	{
		stayZero := 0
		wantOne := 0
		wantThree := 0

		args := []string{
			"--want-one",
			"--want-three",
			"--want-three",
			"--want-three",
			"left",
		}
		ret, err := NewOpts().
			CountingOption("stay-zero", &stayZero).
			CountingOption("want-one", &wantOne).
			CountingOption("want-three", &wantThree).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 0, stayZero)
			assert.Equal(t, 1, wantOne)
			assert.Equal(t, 3, wantThree)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		wantOne := 0
		args := []string{
			"--nowant-one",
			"left",
		}
		ret, err := NewOpts().
			CountingOption("want-one", &wantOne).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "nowant-one")
			assert.Equal(t, []string{"--nowant-one", "left"}, ret)
		}
	}

	{
		wantOne := 0
		args := []string{
			"--want-one",
			"left",
		}
		ret, err := NewOpts().
			CountingOption("count-one", &wantOne).
			IntOption("want-one", &wantOne).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "same pointer")
			assert.Equal(t, []string{"--want-one", "left"}, ret)
		}
	}
}
