package zerrors

import (
	"context"
	"fmt"
)

// zError defines the error itself.
type zError struct {
	msg     string
	errType fmt.Stringer
	err     error
	frame   Frame
	ctx     context.Context
}

// Error implements error interface.
func (e *zError) Error() string {
	return fmt.Sprint(e)
}

// Unwrap implements Wrapper interface.
func (e *zError) Unwrap() error {
	return e.err
}

// Format implements fmt.Formatter interface.
func (e *zError) Format(s fmt.State, v rune) { FormatError(e, s, v) }

// FormatError implements Formatter.
func (e *zError) FormatError(p Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)

	return e.err
}

// Type implements TypeAwareError interface.
func (e *zError) Type() fmt.Stringer {
	return e.errType
}

// IsType implements TypeAwareError interface.
func (e *zError) IsType(errType fmt.Stringer) bool {
	return IsType(e, errType)
}

// Context implements ContextAwareError interface.
func (e *zError) Context() (ctx context.Context) {
	return WalkErrChain(e, func(err error) interface{} {
		_, ok := err.(*zError)
		if !ok {
			return nil
		}
		ctx = err.(*zError).ctx
		if ctx != nil {
			return ctx
		}

		return nil
	}).(context.Context)
}

// ContextValue implements ContextAwareError interface.
func (e *zError) ContextValue(key interface{}) interface{} {
	return WalkErrChain(e, func(err error) interface{} {
		_, ok := err.(*zError)
		if !ok {
			return nil
		}
		ctx := err.(*zError).ctx
		if ctx != nil {
			val := ctx.Value(key)
			if val != nil {
				return val
			}
		}

		return nil
	})
}
