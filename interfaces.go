package zerrors

import (
	"context"
	"fmt"
)

// TypeAwareError defines an error that is aware of its type.
type TypeAwareError interface {
	// implement error interface here
	error
	// Type returns the ErrorType
	Type() fmt.Stringer
	// IsType checks if error is of passed ErrorType
	IsType(errType fmt.Stringer) bool
}

// ContextAwareError defines an error that contains an context.
type ContextAwareError interface {
	// implement error interface here
	error
	// Context does a recursive lookup for context.Context and returns the first contexts in err´s chain
	Context() context.Context
	// ContextValue does a recursive lookup for a context value for given key. It returns nil if no value can be found.
	ContextValue(key interface{}) interface{}
}
