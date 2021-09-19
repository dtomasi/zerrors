package zerrors

import (
	"fmt"
	"reflect"
)

// IsType Check if a given error is matching given type
// Example:
// 		zerrors.IsType(err, MyErrorTypeWithStringerInterface)
func IsType(err TypeAwareError, errType fmt.Stringer) bool {
	isComparable := reflect.TypeOf(err.Type()).Comparable()

	for {
		if isComparable && err.Type() == errType {
			return true
		}

		shouldConvert := Unwrap(err)
		if shouldConvert == nil {
			return false
		}

		// TODO: Check if this is suitable or has to be checked
		err = shouldConvert.(TypeAwareError) //nolint:forcetypeassert
	}
}

// Mask simply masks an error with no message.
func Mask(err error) error {
	return NewWithOpts("", WithWrappedError(err), WithSkipCallers(getCallerSkip(1)))
}

// New creates a new error with message.
func New(msg string) error {
	return NewWithOpts(msg)
}

// Newf creates a new error with formatted message.
func Newf(format string, a ...interface{}) error {
	return NewWithOpts(fmt.Sprintf(format, a...))
}

// NewWithOpts creates a new error with optional passed options
// see: ./options.go for available options
func NewWithOpts(msg string, opts ...ErrorOpt) error {
	// Create error with sensitive defaults
	e := &zError{
		msg:     msg,
		errType: GenericError,
		err:     nil,
		frame:   Caller(DefaultSkipCallers),
		ctx:     nil,
	}

	// Apply options
	for _, opt := range opts {
		opt(e)
	}

	return e
}
