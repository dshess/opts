package opts

import (
	"fmt"
	"os"
	"strings"
)

// TODO: Handle -- in the position of a required argument.  Lke
// "--integer-option", "--".  Getopt::Long takes "--" as the argument,
// which succeeds for string options, but for int or float will be a choice
// of how to fail.

// TODO: Right now, this is structured as a map of handler objects, which
// generate an array of commit objects.  On a different branch, I used a
// map of handler closures which generates an array of commit closures.
// That had less boilerplate, but each object carried a 4k context
// structure.

type Opts struct {
	// Error during construction.
	err error

	// <flag name> => <handler for that flag>
	handlers map[string]optHandler

	// Defer updates until after all options are processed.
	committers []optCommitter
}

// Generates the root structure for collecting argument descriptions.
func NewOpts() *Opts {
	return &Opts{
		err:        nil,
		handlers:   make(map[string]optHandler),
		committers: make([]optCommitter, 0, 10),
	}
}

func (oc *Opts) setError(err error) {
	if oc.err == nil {
		oc.err = err
	}
}

func (oc *Opts) addOption(name string, oh optHandler) *Opts {
	if _, ok := oc.handlers[name]; ok {
		oc.setError(fmt.Errorf("option %s already exists", name))
	} else {
		oc.handlers[name] = oh
	}
	return oc
}

func (oc *Opts) commit() {
	for _, c := range oc.committers {
		c.commit()
	}
}

func isNegatable(left, right string) bool {
	// Overthinking this :-).
	if len(left) >= 2 && len(right) >= 2 {
		if left[0:2] != "no" && right[0:2] != "no" {
			return false
		}
	}
	return left == negatedName(right) || right == negatedName(left)
}

type namedHandler struct {
	name    string
	handler optHandler
}

func (h namedHandler) checkConflict(o namedHandler) bool {
	// Common case, because conflicts aren't dynamic.
	if !h.handler.checkConflict(o.handler) {
		return false
	}

	// If these are a negatabe pair, there is no conflict.
	// TODO: I don't like this happening here.
	if isNegatable(h.name, o.name) {
		return false
	}
	return true
}

func (oc *Opts) checkConflicts() error {
	handlers := make([]namedHandler, 0, len(oc.handlers))
	for k, v := range oc.handlers {
		handlers = append(handlers, namedHandler{k, v})
	}

	// TODO: The N^2 is concerning.  One solution would be to have each
	// handler return uintptr(unsafe.Pointer(oh.option)) and sort/uniq
	// that.  Except unsafe.  The scope could be limited to smaller N
	// by binning by option pointer types.
	for i := len(handlers) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if handlers[i].checkConflict(handlers[j]) {
				return fmt.Errorf("%s and %s use the same pointer", handlers[i].name, handlers[j].name)
			}
		}
	}
	return nil
}

// Process the arguments using the collected config.  Any errors while
// building the Opts are immediately returned.  The Opts are also checked
// for pointer conflicts.
//
// Multiple same-named options is an error.
// Options refering to the same pointer output is an error.
// It is an error for args to contain a pattern which looks like an option but
// which is not defined.
// It is an error for a non-optional option to have no value.
func (oc *Opts) ProcessArgs(args []string) ([]string, error) {
	// Return any errors in construction.
	if oc.err != nil {
		return args, oc.err
	}

	// Check for duplicate targets.
	if err := oc.checkConflicts(); err != nil {
		return args, err
	}

	rest := args
	for len(rest) > 0 {
		name, ok := strings.CutPrefix(rest[0], "--")
		if !ok {
			break
		}
		rest = rest[1:]
		if len(name) == 0 {
			break
		}

		noneOrOne := strings.SplitN(name, "=", 2)
		// noneOrOne can't be empty?  But if it were, don't [0].
		if len(noneOrOne) > 0 {
			name = noneOrOne[0]
			noneOrOne = noneOrOne[1:]
		}

		h, ok := oc.handlers[name]
		if !ok {
			return args, fmt.Errorf("arg %s not recognized", name)
		}

		// TODO: Provide a way for handlers to vet the next arg.
		// For instance, --optional-integer followed by non-integer
		// text could yield the default and end processing.

		if h.getType() == optNoArg {
			// Nothing
		} else if len(noneOrOne) > 0 {
			// Nothing, already have an arg
		} else if len(rest) < 1 {
			if h.getType() == optRequiredArg {
				return args, fmt.Errorf("arg %s missing required argument", name)
			}
			// For optional, no more args is fine
		} else if h.getType() == optOptionalArg && strings.HasPrefix(rest[0], "--") {
			// Nothing, next arg looks flag-like
		} else {
			// This will treat the next arg as a value
			// unconditionally, even if it looks like an
			// option.
			noneOrOne = rest[0:1]
			rest = rest[1:]
		}

		c, err := h.handle(noneOrOne)
		if err != nil {
			return args, err
		}
		oc.committers = append(oc.committers, c)
	}

	// If we made it here without an error, commit the parsed arguments
	// to their pointers.
	oc.commit()
	return rest, nil
}

// Wrapper to pass [os.Args][1:] to [ProcessArgs].  On success [os.Args] is
// updated with the returned args.
func (oc *Opts) ProcessOSArgs() error {
	ret, err := oc.ProcessArgs(os.Args[1:])
	if err != nil {
		return err
	}

	replace := make([]string, 1, 1+len(ret))
	replace[0] = os.Args[0]
	os.Args = append(replace, ret...)
	return err
}
