package zerrors

import (
	"context"
	"fmt"
)

// ErrorOpt defines the option function.
type ErrorOpt func(e *zError)

// WithType allows to pass a type (fmt.Stringer compatible) to the error.
func WithType(errType fmt.Stringer) ErrorOpt {
	return func(e *zError) {
		e.errType = errType
	}
}

// WithWrappedError allows to pass an error that should be wrapped.
func WithWrappedError(err error) ErrorOpt {
	return func(e *zError) {
		e.err = err
	}
}

// WithSkipCallers allows defining how many caller frames should be skipped.
// This only applies to output format `%+v` that prints the stacktrace.
func WithSkipCallers(skip int) ErrorOpt {
	return func(e *zError) {
		e.frame = Caller(1 + skip)
	}
}

// WithContext allows passing a context to the error.
func WithContext(ctx context.Context) ErrorOpt {
	return func(e *zError) {
		e.ctx = ctx
	}
}

// WithContextValue allows passing a key/value pair to the error, which is stored inside the context.Context
// NOTE: If no context is defined while this option is applied, it will create an empty context using context.TODO().
func WithContextValue(key interface{}, value interface{}) ErrorOpt {
	return func(e *zError) {
		if e.ctx == nil {
			e.ctx = context.TODO()
		}

		e.ctx = context.WithValue(e.ctx, key, value)
	}
}
