package opts

// Add simple flag, --<name> will set *option to true.
func (oc *Opts) SimpleOption(name string, option *bool) *Opts {
	return oc.addOption(name, optBaseHandler[bool]{
		t:      optNoArg,
		option: option,
		def:    true,
	})
}

func negatedName(name string) string {
	return "no" + name
}

// Add negatable option, --<name> will set *option to true, --no<name> will
// set *option to false.
func (oc *Opts) NegatableOption(name string, option *bool) *Opts {
	return oc.addOption(name, optBaseHandler[bool]{
		t:      optNoArg,
		option: option,
		def:    true,
	}).addOption(negatedName(name), optBaseHandler[bool]{
		t:      optNoArg,
		option: option,
		def:    false,
	})
}

// Add counting option, every occurance of --<name> increments *option.
func (oc *Opts) CountingOption(name string, option *int) *Opts {
	return oc.addOption(name, optCountingHandler{
		t:      optNoArg,
		option: option,
	})
}

// Add integer option, --<name>=val or --<name> val will set *option to
// val.
func (oc *Opts) IntOption(name string, option *int) *Opts {
	return oc.addOption(name, optBaseHandler[int]{
		t:      optRequiredArg,
		option: option,
	})
}

// Add optional integer option, --<name>=val or --<name> val will set
// *option to val, while --<name> alone will set *option to def.  "Alone"
// means that --<name> is followed immediately by another option, or --, or
// the end of arguments.
func (oc *Opts) OptionalIntOption(name string, option *int, def int) *Opts {
	return oc.addOption(name, optBaseHandler[int]{
		t:      optOptionalArg,
		option: option,
		def:    def,
	})
}

// Add integer array option, every --<name>=val or --<name> val will append
// val to *option.
func (oc *Opts) IntArrayOption(name string, option *[]int) *Opts {
	return oc.addOption(name, optBaseArrayHandler[int]{
		t:      optRequiredArg,
		option: option,
	})
}

// Add float option, --<name>=val or --<name> val will set *option to val.
func (oc *Opts) FloatOption(name string, option *float64) *Opts {
	return oc.addOption(name, optBaseHandler[float64]{
		t:      optRequiredArg,
		option: option,
	})
}

// Add optional float option, --<name>=val or --<name> val will set *option
// to val, while --<name> alone will set *option to def.  "Alone" means
// that --<name> is followed immediately by another option, or --, or the
// end of arguments.
func (oc *Opts) OptionalFloatOption(name string, option *float64, def float64) *Opts {
	return oc.addOption(name, optBaseHandler[float64]{
		t:      optOptionalArg,
		option: option,
		def:    def,
	})
}

// Add float array option, every --<name>=val or --<name> val will append
// val to *option.
func (oc *Opts) FloatArrayOption(name string, option *[]float64) *Opts {
	return oc.addOption(name, optBaseArrayHandler[float64]{
		t:      optRequiredArg,
		option: option,
	})
}

// Add string option, --<name>=val or --<name> val will set *option to val.
func (oc *Opts) StringOption(name string, option *string) *Opts {
	return oc.addOption(name, optBaseHandler[string]{
		t:      optRequiredArg,
		option: option,
	})
}

// Add optional string option, --<name>=val or --<name> val will set
// *option to val, while --<name> alone will set *option to def.  "Alone"
// means that --<name> is followed immediately by another option, or --, or
// the end of arguments.
func (oc *Opts) OptionalStringOption(name string, option *string, def string) *Opts {
	return oc.addOption(name, optBaseHandler[string]{
		t:      optOptionalArg,
		option: option,
		def:    def,
	})
}

// Add string array option, every --<name>=val or --<name> val will append
// val to *option.
func (oc *Opts) StringArrayOption(name string, option *[]string) *Opts {
	return oc.addOption(name, optBaseArrayHandler[string]{
		t:      optRequiredArg,
		option: option,
	})
}

// TODO:
// func (oc *Opts) CustomOption(name string, func(...)) *Opts
// func (oc *Opts) OptionalCustomOption(name string, func(...)) *Opts
//
// This would allow a caller-defined option processing.  Maybe it's better
// to just go with a stronger package in that case?  I would be nervous
// about how to consume multiple arguments in a clean way.

// TODO:
// func (oc *Opts) AKA(name string) *Opts
// func (oc *Opts) Short(ch string) *Opts
//
// Used like:
//     IntOption("length", &length).AKA("len").Short("l").
// This would key off the previous option.
//
// func (oc *Opts) Alias(name string, aliases ...string) *Opts
//
// This style would only need to follow the option somewhere, and could
// perhaps configure more bits of it.
//
// func (oc *Opts) IntOption(name string, option *int, cfg ...any) *Opts
//
// This might also work, but the config would end up being wordy because
// you'd have to use the package name.  Compare:
//
//     IntOption("name", &intPtr, opts.AKA("other"), opts.Short("n")).
//     IntOption("name", &intPtr).AKA("other").Short("n").
//     IntOption("name", &intPtr).Alias("name", "other", "n").
//
// I mean, the first is not the worst thing in the world, I suppose.  But
// if it were named dshessgoopts, it would be the worst thing in the world.
// I guess opts.AKA("other").Short("n") would also make sense.
//
// Going too far down this rabbit hole might lead to:
//
//    Option("name", "other", "n").Int(&intPtr).
