package opts

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmpty(t *testing.T) {
	args := []string{
		"this", "is", "a", "test",
	}
	ret, err := NewOpts().
		ProcessArgs(args)
	require.Nil(t, err)
	assert.Equal(t, args, ret)
}

func TestDashDash(t *testing.T) {
	wantTrue := false
	stayFalse := false

	args := []string{
		"--want-true",
		"--",
		"--stay-false",
		"left",
	}
	ret, err := NewOpts().
		SimpleOption("want-true", &wantTrue).
		SimpleOption("stay-false", &stayFalse).
		ProcessArgs(args)
	require.Nil(t, err)
	assert.True(t, wantTrue)
	assert.False(t, stayFalse)
	assert.Equal(t, []string{"--stay-false", "left"}, ret)
}

func TestDuplicateName(t *testing.T) {
	wantTrue := false
	otherTrue := false

	args := []string{
		"--want-true",
		"--",
		"left",
	}
	ret, err := NewOpts().
		SimpleOption("want-true", &wantTrue).
		SimpleOption("want-true", &otherTrue).
		ProcessArgs(args)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "already exists")
	}
	assert.False(t, wantTrue)
	assert.Equal(t, []string{"--want-true", "--", "left"}, ret)
}

func TestDuplicateTarget(t *testing.T) {
	{
		sameBool := false

		args := []string{
			"--want-true",
			"left",
		}
		_, err := NewOpts().
			SimpleOption("want-true", &sameBool).
			SimpleOption("stay-false", &sameBool).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "same pointer")
		}
	}

	{
		sameInt := 7

		args := []string{
			"--want-eleven=11",
			"--counting",
			"left",
		}
		_, err := NewOpts().
			IntOption("want-eleven", &sameInt).
			CountingOption("counting", &sameInt).
			ProcessArgs(args)
		if assert.NotNil(t, err) {
			assert.Contains(t, err.Error(), "same pointer")
		}
	}
}

func TestUnknown(t *testing.T) {
	wantTrue := false

	args := []string{
		"--want-true",
		"--want-false",
		"left",
	}
	ret, err := NewOpts().
		SimpleOption("want-true", &wantTrue).
		ProcessArgs(args)
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "not recognized")
		assert.False(t, wantTrue)
		assert.Equal(t, []string{"--want-true", "--want-false", "left"}, ret)
	}
}

func TestOSArgs(t *testing.T) {
	tmpArgs := os.Args
	defer func() {
		os.Args = tmpArgs
	}()

	os.Args = []string{
		"command",
		"--simple",
		"--string=aValue",
		"--string", "otherValue",
		"--optional", "--",
		"additional", "arguments",
	}
	simpleFlag := false
	stringArray := []string{}
	optionalValue := 17
	err := NewOpts().
		SimpleOption("simple", &simpleFlag).
		StringArrayOption("string", &stringArray).
		OptionalIntOption("optional", &optionalValue, 23).
		ProcessOSArgs()
	require.Nil(t, err)
	assert.True(t, simpleFlag)
	assert.Equal(t, []string{"aValue", "otherValue"}, stringArray)
	assert.Equal(t, 23, optionalValue)
	assert.Equal(t, []string{"command", "additional", "arguments"}, os.Args)
}

func TestOSArgsError(t *testing.T) {
	tmpArgs := os.Args
	defer func() {
		os.Args = tmpArgs
	}()

	os.Args = []string{
		"command",
		"--string",
	}
	stringValue := ""
	err := NewOpts().
		StringOption("string", &stringValue).
		ProcessOSArgs()
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "missing")
	}
	assert.Equal(t, []string{"command", "--string"}, os.Args)
}

func ExampleNewOpts() {
	args := []string{
		"--files=hello.world", "--length", "10", "--verbose", "rest",
	}

	data := "file.dat"
	length := 24
	var verbose bool
	rest, err := NewOpts().
		IntOption("length", &length).
		StringOption("files", &data).
		SimpleOption("verbose", &verbose).
		ProcessArgs(args)
	if err != nil {
		log.Fatal("Error in command-line arguments:", err)
	}
	fmt.Printf("length:%d, data:%s, verbose:%t, rest:%s\n",
		length, data, verbose, rest)
	// Output:
	// length:10, data:hello.world, verbose:true, rest:[rest]
}
