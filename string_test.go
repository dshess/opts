package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringOption(t *testing.T) {
	{
		stayCalm := "calm"
		wantChaos := "calm"
		args := []string{
			"--want-chaos", "chaos",
			"left",
		}
		ret, err := NewOpts().
			StringOption("stay-calm", &stayCalm).
			StringOption("want-chaos", &wantChaos).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, "calm", stayCalm)
			assert.Equal(t, "chaos", wantChaos)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		stayCalm := "calm"
		wantChaos := "calm"
		args := []string{
			"--want-chaos=chaos",
			"left",
		}
		ret, err := NewOpts().
			StringOption("stay-calm", &stayCalm).
			StringOption("want-chaos", &wantChaos).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, "calm", stayCalm)
			assert.Equal(t, "chaos", wantChaos)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		wantChaos := "calm"
		args := []string{
			"--want-chaos",
		}
		ret, err := NewOpts().
			StringOption("want-chaos", &wantChaos).
			ProcessArgs(args)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "missing")
		assert.Equal(t, []string{"--want-chaos"}, ret)
	}

	{
		wantChaos := "calm"
		args := []string{
			"--nowant-chaos",
			"left",
		}
		ret, err := NewOpts().
			StringOption("want-chaos", &wantChaos).
			ProcessArgs(args)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "nowant-chaos")
		assert.Equal(t, []string{"--nowant-chaos", "left"}, ret)
	}

	// String differs from the Int and Float case, because anything can
	// be a valid argument to a string option.
	//
	// TODO: Consider whether this is reasonable.  --opt=--other is a
	// reasonable workaround, but --opt --other is ambiguous.
	{
		wantChaos := "calm"
		args := []string{
			"--want-chaos", "--want-chaos",
			"left",
		}
		ret, err := NewOpts().
			StringOption("want-chaos", &wantChaos).
			ProcessArgs(args)
		assert.Nil(t, err)
		assert.Equal(t, "--want-chaos", wantChaos)
		assert.Equal(t, []string{"left"}, ret)
	}
}

func TestOptionalStringOption(t *testing.T) {
	{
		stayCalm := "calm"
		wantChaos := "calm"
		args := []string{
			"--want-chaos", "chaos",
			"left",
		}
		ret, err := NewOpts().
			OptionalStringOption("stay-calm", &stayCalm, "chaos").
			OptionalStringOption("want-chaos", &wantChaos, "calm").
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, "calm", stayCalm)
			assert.Equal(t, "chaos", wantChaos)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		stayCalm := "calm"
		wantChaos := "calm"
		args := []string{
			"--want-chaos=chaos",
			"left",
		}
		ret, err := NewOpts().
			OptionalStringOption("stay-calm", &stayCalm, "chaos").
			OptionalStringOption("want-chaos", &wantChaos, "calm").
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, "calm", stayCalm)
			assert.Equal(t, "chaos", wantChaos)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		stayCalm := "calm"
		wantChaos := "calm"
		args := []string{
			"--want-chaos",
		}
		ret, err := NewOpts().
			OptionalStringOption("stay-calm", &stayCalm, "chaos").
			OptionalStringOption("want-chaos", &wantChaos, "chaos").
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, "calm", stayCalm)
			assert.Equal(t, "chaos", wantChaos)
			assert.Empty(t, ret)
		}
	}

	{
		wantChaos := "calm"
		args := []string{
			"--nowant-chaos",
			"left",
		}
		ret, err := NewOpts().
			OptionalStringOption("want-chaos", &wantChaos, "chaos").
			ProcessArgs(args)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "nowant-chaos")
		assert.Equal(t, []string{"--nowant-chaos", "left"}, ret)
	}
}

func TestStringArrayOption(t *testing.T) {
	{
		stayEmpty := []string{}
		sevenEleven := []string{}
		args := []string{
			"--seven-eleven", "seven",
			"--seven-eleven", "eleven",
			"left",
		}
		ret, err := NewOpts().
			StringArrayOption("stay-empty", &stayEmpty).
			StringArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Empty(t, stayEmpty)
			assert.Equal(t, []string{"seven", "eleven"}, sevenEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		sevenEleven := []string{}
		args := []string{
			"--seven-eleven",
		}
		ret, err := NewOpts().
			StringArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "missing")
		assert.Equal(t, []string{"--seven-eleven"}, ret)
	}

	{
		sevenEleven := []string{}
		args := []string{
			"--noseven-eleven",
			"left",
		}
		ret, err := NewOpts().
			StringArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "noseven-eleven")
		assert.Equal(t, []string{"--noseven-eleven", "left"}, ret)
	}
}
