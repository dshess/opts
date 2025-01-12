# opts

[![Go Reference](https://pkg.go.dev/badge/github.com/dshess/opts.svg)](https://pkg.go.dev/github.com/dshess/opts)
[![Go Report Card](https://goreportcard.com/badge/github.com/dshess/opts)](https://goreportcard.com/report/github.com/dshess/opts)

Command-line flag parsing in the style of Perl's Getopt::Long.

## Objectives

I found myself missing Perl's
[Getopt::Long](https://perldoc.perl.org/Getopt::Long) package.
So I set out to make something similar,
[github.com/dshess/getopt](https://pkg.go.dev/github.com/dshess/getopt).
getopt spent more time enforcing type safety than I cared for,
this package changes to a builder-chain style.

## Install

`go get github.com/dshess/opts`

Or just import it and let go guide you.

## Requirements

This was developed under go `1.23.2`.

## Documentation

https://pkg.go.dev/github.com/dshess/opts

## Usage

A translation of
[Getopt::Long#SYNOPSIS](https://perldoc.perl.org/Getopt::Long#SYNOPSIS):

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

Output:

	length:10, data:hello.world, verbose:true, rest:[rest]

## License

Licensed under the MIT License, see `LICENSE` file.
