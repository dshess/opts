/*
Package opts implements command-line flag parsing in the style of Perl's
[Getopt::Long], using a chained builder pattern.

# Usage

A translation of [Getopt::Long#SYNOPSIS]:

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

# Command line flag syntax

opts currently only handles --option style of options, no groups of
single-chararacter options.  Bare -- ends option processing.  Boolean
options can be negatable or simple, with no parameters (so --option or
--nooption).  Options with parameters can be --option=value or --option
value.  Optional options deliver the provided default if --option is seen
with no next argument, or where the next argument itself looks like another
option, or where the next argument is --.

# Why not flag package?

No reason, I'm not the flag police.  Mostly with [flag] I was frustrated by
having bits and pieces distributed around my code, and having it all
plugged together by some sort of global.  Using [flag.FlagSet] gets around
the global (I think), but then feels heavy to me.

# Why not other flag packages?

No reason, I'm still not the flag police.  Mostly the same frustrations as
flag package, often the packages feel like they are solving a pretty
substantial problem, and my needs have been generally more light-weight.

# What is the real reason?

You got me.  I initially built [github.com/dshess/getopt], which was pretty
similar to Getopt::Long.  But once I had worked through the fun puzzles of
making it more-or-less type-safe, I realized that using functions would be
cleaner than reflection.  So I parked that code and rewrote it.

[Getopt::Long]: https://perldoc.perl.org/Getopt::Long
[Getopt::Long#SYNOPSIS]: https://perldoc.perl.org/Getopt::Long#SYNOPSIS
[github.com/dshess/getopt]: https://pkg.go.dev/github.com/dshess/getopt
*/
package opts
