package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntOption(t *testing.T) {
	{
		staySeven := 7
		wantEleven := 7
		args := []string{
			"--want-eleven", "11",
			"left",
		}
		ret, err := NewOpts().
			IntOption("stay-seven", &staySeven).
			IntOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7, staySeven)
			assert.Equal(t, 11, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		staySeven := 7
		wantEleven := 7
		args := []string{
			"--want-eleven=11",
			"left",
		}
		ret, err := NewOpts().
			IntOption("stay-seven", &staySeven).
			IntOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7, staySeven)
			assert.Equal(t, 11, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		wantEleven := 7
		args := []string{
			"--want-eleven",
		}
		ret, err := NewOpts().
			IntOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "missing")
		}
		assert.Equal(t, []string{"--want-eleven"}, ret)
	}

	{
		wantEleven := 7
		args := []string{
			"--nowant-eleven",
			"left",
		}
		ret, err := NewOpts().
			IntOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "nowant-eleven")
		}
		assert.Equal(t, []string{"--nowant-eleven", "left"}, ret)
	}

	{
		wantEleven := 7
		args := []string{
			"--want-eleven", "twelve",
		}
		ret, err := NewOpts().
			IntOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "invalid syntax")
		}
		assert.Equal(t, []string{"--want-eleven", "twelve"}, ret)
	}
}

func TestOptionalIntOption(t *testing.T) {
	{
		staySeven := 7
		wantEleven := 7
		args := []string{
			"--want-eleven", "11",
			"left",
		}
		ret, err := NewOpts().
			OptionalIntOption("stay-seven", &staySeven, 11).
			OptionalIntOption("want-eleven", &wantEleven, 7).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7, staySeven)
			assert.Equal(t, 11, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		staySeven := 7
		wantEleven := 7
		args := []string{
			"--want-eleven=11",
			"left",
		}
		ret, err := NewOpts().
			OptionalIntOption("stay-seven", &staySeven, 11).
			OptionalIntOption("want-eleven", &wantEleven, 7).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7, staySeven)
			assert.Equal(t, 11, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		staySeven := 7
		wantEleven := 7
		args := []string{
			"--want-eleven",
		}
		ret, err := NewOpts().
			OptionalIntOption("stay-seven", &staySeven, 11).
			OptionalIntOption("want-eleven", &wantEleven, 11).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Equal(t, 7, staySeven)
			assert.Equal(t, 11, wantEleven)
			assert.Empty(t, ret)
		}
	}

	{
		wantEleven := 7
		args := []string{
			"--nowant-eleven",
			"left",
		}
		ret, err := NewOpts().
			OptionalIntOption("want-eleven", &wantEleven, 7).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "nowant-eleven")
		}
		assert.Equal(t, []string{"--nowant-eleven", "left"}, ret)
	}

	// TODO: Should an optional part kick in if the next parameter does
	// not parse?
	if false {
		wantEleven := 7
		args := []string{
			"--want-eleven",
			"left",
		}
		ret, err := NewOpts().
			OptionalIntOption("want-eleven", &wantEleven, 11).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Equal(t, 11, wantEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}
}

func TestIntArrayOption(t *testing.T) {
	{
		stayEmpty := []int{}
		sevenEleven := []int{}
		args := []string{
			"--seven-eleven", "7",
			"--seven-eleven", "11",
			"left",
		}
		ret, err := NewOpts().
			IntArrayOption("stay-empty", &stayEmpty).
			IntArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.Nil(t, err) {
			assert.Empty(t, stayEmpty)
			assert.Equal(t, []int{7, 11}, sevenEleven)
			assert.Equal(t, []string{"left"}, ret)
		}
	}

	{
		sevenEleven := []int{}
		args := []string{
			"--seven-eleven",
		}
		ret, err := NewOpts().
			IntArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "missing")
		}
		assert.Equal(t, []string{"--seven-eleven"}, ret)
	}

	{
		sevenEleven := []int{}
		args := []string{
			"--noseven-eleven",
			"left",
		}
		ret, err := NewOpts().
			IntArrayOption("seven-eleven", &sevenEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "noseven-eleven")
		}
		assert.Equal(t, []string{"--noseven-eleven", "left"}, ret)
	}

	{
		wantEleven := []int{}
		args := []string{
			"--want-eleven", "twelve",
		}
		ret, err := NewOpts().
			IntArrayOption("want-eleven", &wantEleven).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "invalid syntax")
		}
		assert.Equal(t, []string{"--want-eleven", "twelve"}, ret)
	}
}
