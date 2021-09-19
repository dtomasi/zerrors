package zerrors_test

import (
	"errors"
	"fmt"
	z "github.com/dtomasi/zerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func functionThatReturnsInitialError() error {
	return z.New("initial error")
}

func functionThatWrapsError() error {
	err := functionThatReturnsInitialError()
	if err != nil {
		return z.Wrap(functionThatReturnsInitialError(), "Wrap error")
	}

	return nil
}

func functionThatWrapsErrorWithFormat() error {
	err := functionThatReturnsInitialError()
	if err != nil {
		return z.Wrapf(functionThatReturnsInitialError(), "Wrapf %s", "error")
	}

	return nil
}

func functionThatWrapsErrorUsingPtr() (err error) {
	defer z.WrapPtr(&err, "WrapPtr error")

	err = functionThatReturnsInitialError()
	if err != nil {
		return err
	}

	return nil
}

func functionThatWrapsErrorUsingPtrWithFormat() (err error) {
	defer z.WrapPtrf(&err, "WrapPtrf %s", "error")

	err = functionThatReturnsInitialError()
	if err != nil {
		return err
	}

	return nil
}

func TestWrap_Integration(t *testing.T) {
	err := functionThatWrapsError()
	assert.Error(t, err)
	assert.Equal(t, "Wrap error: initial error", err.Error())
}

func TestWrapPtr_Integration(t *testing.T) {
	err := functionThatWrapsErrorUsingPtr()
	assert.Error(t, err)
	assert.Equal(t, "WrapPtr error: initial error", err.Error())
}

func TestWrapf_Integration(t *testing.T) {
	err := functionThatWrapsErrorWithFormat()
	assert.Error(t, err)
	assert.Equal(t, "Wrapf error: initial error", err.Error())
}

func TestWrapPtrf_Integration(t *testing.T) {
	err := functionThatWrapsErrorUsingPtrWithFormat()
	assert.Error(t, err)
	assert.Equal(t, "WrapPtrf error: initial error", err.Error())
}

func TestWrapWithOpts_WithType(t *testing.T) {
	errorType := z.ErrorType("foo")
	baseErr := functionThatWrapsError()

	// Test wrap with type
	err := z.WrapWithOpts(baseErr, "outer", z.WithType(errorType))
	assert.Error(t, err)
	helperAssertErrorIsType(t, err, errorType)
}

func TestWrapWithOpts_WithWrappedError_panics(t *testing.T) {
	// Test wrap with passed error should panic
	assert.Panics(t, func() {
		_ = z.WrapWithOpts(errors.New(""), "outer", z.WithWrappedError(errors.New(""))) //nolint:goerr113
	})
}

func TestWrapWithOpts_WithSkipCallers_overwrites(t *testing.T) {
	baseErr := functionThatWrapsError()

	// Test wrap with type
	err := z.WrapWithOpts(baseErr, "outer", z.WithSkipCallers(0))
	assert.Error(t, err)
	output := fmt.Sprintf("%+v", err)
	// If the output contains the function name then skip callers is overwritten
	assert.Contains(t, output, "WithSkipCallers")
}
