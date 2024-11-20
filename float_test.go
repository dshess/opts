package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatOption(t *testing.T) {
	{
		staySeven := 7.0
		wantEleven := 7.0
		args := []string{
			"--want-eleven", "11",
			"left",
		}
		ret, err := NewOpts().
			FloatOption("stay-seven", &staySeven).
			FloatOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7.0, staySeven)
			assert.Equal(t, 11.0, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		staySeven := 7.0
		wantEleven := 7.0
		args := []string{
			"--want-eleven=11",
			"left",
		}
		ret, err := NewOpts().
			FloatOption("stay-seven", &staySeven).
			FloatOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7.0, staySeven)
			assert.Equal(t, 11.0, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		wantEleven := 7.0
		args := []string{
			"--want-eleven",
		}
		ret, err := NewOpts().
			FloatOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "missing")
		}
		assert.Equal(t, []string{"--want-eleven"}, ret)
	}

	{
		wantEleven := 7.0
		args := []string{
			"--nowant-eleven",
			"left",
		}
		ret, err := NewOpts().
			FloatOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "nowant-eleven")
		}
		assert.Equal(t, []string{"--nowant-eleven", "left"}, ret)
	}

	{
		wantEleven := 7.0
		args := []string{
			"--want-eleven", "twelve",
		}
		ret, err := NewOpts().
			FloatOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "invalid syntax")
		}
		assert.Equal(t, []string{"--want-eleven", "twelve"}, ret)
	}
}

func TestOptionalFloatOption(t *testing.T) {
	{
		staySeven := 7.0
		wantEleven := 7.0
		args := []string{
			"--want-eleven", "11",
			"left",
		}
		ret, err := NewOpts().
			OptionalFloatOption("stay-seven", &staySeven, 11.0).
			OptionalFloatOption("want-eleven", &wantEleven, 7.0).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7.0, staySeven)
			assert.Equal(t, 11.0, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		staySeven := 7.0
		wantEleven := 7.0
		args := []string{
			"--want-eleven=11.0",
			"left",
		}
		ret, err := NewOpts().
			OptionalFloatOption("stay-seven", &staySeven, 11.0).
			OptionalFloatOption("want-eleven", &wantEleven, 7.0).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7.0, staySeven)
			assert.Equal(t, 11.0, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		staySeven := 7.0
		wantEleven := 7.0
		args := []string{
			"--want-eleven",
		}
		ret, err := NewOpts().
			OptionalFloatOption("stay-seven", &staySeven, 11.0).
			OptionalFloatOption("want-eleven", &wantEleven, 11.0).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7.0, staySeven)
			assert.Equal(t, 11.0, wantEleven)
			assert.Empty(t, ret)
		}
	}

	{
		wantEleven := 7.0
		args := []string{
			"--nowant-eleven",
			"left",
		}
		ret, err := NewOpts().
			OptionalFloatOption("want-eleven", &wantEleven, 7.0).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "nowant-eleven")
		}
		assert.Equal(t, []string{"--nowant-eleven", "left"}, ret)
	}

	// TODO: Should an optional part kick in if the next parameter does
	// not parse?
	if false {
		wantEleven := 7.0
		args := []string{
			"--want-eleven",
			"left",
		}
		ret, err := NewOpts().
			OptionalFloatOption("want-eleven", &wantEleven, 11.0).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Equal(t, 11.0, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}
}

func TestFloatArrayOption(t *testing.T) {
	{
		stayEmpty := []float64{}
		sevenEleven := []float64{}
		args := []string{
			"--seven-eleven", "7.0",
			"--seven-eleven", "11.0",
			"left",
		}
		ret, err := NewOpts().
			FloatArrayOption("stay-empty", &stayEmpty).
			FloatArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Empty(t, stayEmpty)
			assert.Equal(t, []float64{7.0, 11.0}, sevenEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		sevenEleven := []float64{}
		args := []string{
			"--seven-eleven",
		}
		ret, err := NewOpts().
			FloatArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "missing")
		}
		assert.Equal(t, []string{"--seven-eleven"}, ret)
	}

	{
		sevenEleven := []float64{}
		args := []string{
			"--noseven-eleven",
			"left",
		}
		ret, err := NewOpts().
			FloatArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "noseven-eleven")
		}
		assert.Equal(t, []string{"--noseven-eleven", "left"}, ret)
	}
}
