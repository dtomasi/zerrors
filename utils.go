package zerrors

import (
	"os"
	"strings"
)

// DefaultSkipCallers exposes the default value that should be used for caller frame skipping
// This can be adjusted within an init function.
var DefaultSkipCallers = 0 //nolint:gochecknoglobals

// getDefaultCallerSkipForWrap returns the caller skip default value to use based on running mode.
func getCallerSkip(base int) int {
	// This is used for provide same stacks while running in `go test`
	if strings.HasSuffix(os.Args[0], ".test") {
		return DefaultSkipCallers + base + 1
	}

	return DefaultSkipCallers + base
}

// WalkErrChain provides a simple way to walk e error chain using Unwrap function.
func WalkErrChain(err error, f func(err error) interface{}) interface{} {
	for {
		retVal := f(err)
		if retVal != nil {
			return retVal
		}

		if err = Unwrap(err); err == nil {
			return nil
		}
	}
}
