package zerrors

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const (
	callerSkipWrapErrFuncs = 3
)

func setWrapOptions(opts []ErrorOpt) []ErrorOpt {
	optNamesString := ""
	for _, opt := range opts {
		optNamesString += runtime.FuncForPC(reflect.ValueOf(opt).Pointer()).Name()
	}

	if strings.Contains(optNamesString, "WithWrappedError") {
		panic("WithWrappedError is not allowed in a Wrapper function")
	}

	// TODO ... Check if caller skip needs to be adjusted here
	// This would change all in here ....
	if !strings.Contains(optNamesString, "WithSkipCallers") {
		opts = append(opts, WithSkipCallers(getCallerSkip(callerSkipWrapErrFuncs)))
	}

	return opts
}

// Wrapf wraps a given error with a new zError instance and allows format message string.
func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

// Wrap wraps a given error with a new zError instance.
func Wrap(err error, msg string) error {
	return WrapWithOpts(err, msg)
}

// WrapWithOpts wraps an error.
func WrapWithOpts(err error, msg string, opts ...ErrorOpt) error {
	opts = append(setWrapOptions(opts), WithWrappedError(err))

	return NewWithOpts(msg, opts...)
}

// WrapPtrf wraps a given error pointer with a new zError instance and allows format message string
// This is useful for defer used with named return types.
func WrapPtrf(errp *error, format string, args ...interface{}) {
	if *errp != nil {
		WrapPtr(errp, fmt.Sprintf(format, args...))
	}
}

// WrapPtr wraps a given error pointer with a new zError instance
// This is useful for defer used with named return types.
func WrapPtr(errp *error, msg string) {
	if *errp != nil {
		WrapPtrWithOpts(errp, msg)
	}
}

// WrapPtrWithOpts wraps an error pointer.
func WrapPtrWithOpts(errp *error, msg string, opts ...ErrorOpt) {
	if *errp != nil {
		opts = append(setWrapOptions(opts), WithWrappedError(*errp))
		*errp = NewWithOpts(msg, opts...)
	}
}
