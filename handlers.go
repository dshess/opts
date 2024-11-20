package opts

import (
	"errors"
	"strconv"
)

// optCommitters are like simple closures called after options processing
// has completed without error.
type optCommitter interface {
	commit()
}

// Store a value at a pointer on commit.
type optSimpleCommitter[T any] struct {
	value  T
	option *T
}

func (o optSimpleCommitter[_]) commit() {
	*o.option = o.value
}

// Append value to an array on commit.
type optArrayCommitter[T any] struct {
	value  T
	option *[]T
}

func (o optArrayCommitter[_]) commit() {
	*o.option = append(*o.option, o.value)
}

type optType int

const (
	optNoArg optType = iota
	optOptionalArg
	optRequiredArg
)

// optHandler provides a hint as to how many arguments, and a handler to call
// with those arguments.  The handler generates an optCommitter to be called
// later.
type optHandler interface {
	getType() optType
	handle(args []string) (optCommitter, error)

	checkConflict(other optHandler) bool
	getPointer() any
}

// Conflicts if other's pointer can be cast to the same type as ptr, and
// they point to the same place.
func checkConflictInner[T any](ptr *T, other optHandler) bool {
	op, ok := other.getPointer().(*T)
	if !ok {
		return false
	}
	return ptr == op
}

type optBasicType interface {
	bool | int | float64 | string
}

// TODO: There must be a better way to do this, but my skill is not there.
func optParseValue[T optBasicType](arg string, tp *T) error {
	switch p := any(tp).(type) {
	case *int:
		i, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		*p = i
		return nil
	case *float64:
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return err
		}
		*p = f
		return nil
	case *string:
		*p = arg
		return nil
	case *bool:
		// SimpleOption() and NegatableOption() are optNoArg, so
		// this should never fire.
		return errors.New("unexpected arg to simple option")
	default:
		// All four optBasicType options are already handled, so
		// this should never fire.
		return errors.New("parse for unexpected option type")
	}
}

type optBaseHandler[T optBasicType] struct {
	t      optType
	option *T
	def    T
}

func (oh optBaseHandler[_]) getType() optType {
	return oh.t
}
func (oh optBaseHandler[T]) handle(args []string) (optCommitter, error) {
	v := oh.def
	if len(args) > 0 {
		err := optParseValue(args[0], &v)
		if err != nil {
			return nil, err
		}
	}
	c := optSimpleCommitter[T]{v, oh.option}
	return c, nil
}
func (oh optBaseHandler[_]) getPointer() any {
	return oh.option
}
func (oh optBaseHandler[_]) checkConflict(other optHandler) bool {
	return checkConflictInner(oh.option, other)
}

type optBaseArrayHandler[T optBasicType] struct {
	t      optType
	option *[]T
}

func (oh optBaseArrayHandler[_]) getType() optType {
	return oh.t
}
func (oh optBaseArrayHandler[T]) handle(args []string) (optCommitter, error) {
	var v T
	err := optParseValue(args[0], &v)
	if err != nil {
		return nil, err
	}
	c := optArrayCommitter[T]{v, oh.option}
	return c, nil
}
func (oh optBaseArrayHandler[_]) getPointer() any {
	return oh.option
}
func (oh optBaseArrayHandler[_]) checkConflict(other optHandler) bool {
	return checkConflictInner(oh.option, other)
}

// Increment a pointed-to value on commit.
type optCountingCommitter struct {
	option *int
}

func (o optCountingCommitter) commit() {
	*o.option++
}

// TODO: Removing counting options would allow dropping getPointer() from
// the optHandler interface.
type optCountingHandler struct {
	t      optType
	option *int
}

func (oh optCountingHandler) getType() optType {
	return oh.t
}
func (oh optCountingHandler) handle(args []string) (optCommitter, error) {
	c := optCountingCommitter{oh.option}
	return c, nil
}
func (oh optCountingHandler) getPointer() any {
	return oh.option
}
func (oh optCountingHandler) checkConflict(other optHandler) bool {
	return checkConflictInner(oh.option, other)
}
