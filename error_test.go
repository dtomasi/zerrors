package zerrors_test

import (
	"context"
	"fmt"
	z "github.com/dtomasi/zerrors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
	"testing"
)

func TestZError_Error(t *testing.T) {
	err := z.New("test")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "test")
}

func TestZError_Unwrap(t *testing.T) {
	errInner := z.New("inner")
	assert.Error(t, errInner)
	assert.Equal(t, errInner.Error(), "inner")

	// Test with NewWithOpts
	errWithWrappedError := z.NewWithOpts("outer", z.WithWrappedError(errInner))
	assert.Error(t, errWithWrappedError)
	assert.Equal(t, errInner, errWithWrappedError.(z.Wrapper).Unwrap())

	// Test with fmt.Errorf
	fmtErrfWrapper := fmt.Errorf("%w: %s", errInner, "outer")
	assert.Equal(t, errInner, fmtErrfWrapper.(z.Wrapper).Unwrap())

	// Test with xerrors.Errorf
	xerrorsErrfWrapper := xerrors.Errorf("%w: %s", errInner, "outer")
	assert.Equal(t, errInner, xerrorsErrfWrapper.(xerrors.Wrapper).Unwrap())
}

func TestZError_TypeAwareError(t *testing.T) {
	const errType = z.ErrorType("test")
	err := z.NewWithOpts("test", z.WithType(errType))
	assert.Error(t, err)
	assert.Equal(t, err.(z.TypeAwareError).Type(), errType)
	assert.True(t, err.(z.TypeAwareError).IsType(errType))
}

func TestZError_ContextAwareError(t *testing.T) {
	var (
		innerContext          = context.Background()
		innerContextTestKey   = "foo"
		innerContextTestValue = "bar"
	)

	innerContext = context.WithValue(innerContext, innerContextTestKey, innerContextTestValue) //nolint:revive,staticcheck
	errInnerWithContext := z.NewWithOpts("inner", z.WithContext(innerContext))

	// Test basic Context()
	assert.Implements(t, (*context.Context)(nil), errInnerWithContext.(z.ContextAwareError).Context())

	// Test basic ContextValue()
	assert.Equal(t, innerContextTestValue, errInnerWithContext.(z.ContextAwareError).ContextValue(innerContextTestKey))

	// Test empty outer wrapper returns inner context
	emptyOuter := z.Mask(errInnerWithContext)

	// Run same tests against empty outer wrapper
	assert.Implements(t, (*context.Context)(nil), emptyOuter.(z.ContextAwareError).Context())
	assert.Equal(t, innerContextTestValue, emptyOuter.(z.ContextAwareError).ContextValue(innerContextTestKey))

	// Test with no ContextAwareError in chain
	errMiddleNoZError := xerrors.Errorf("%w: %s", errInnerWithContext, "middle")
	outerWithEmptyContext := z.NewWithOpts(
		"outer",
		z.WithContext(context.TODO()),
		z.WithWrappedError(errMiddleNoZError),
	)

	// This should return the empty context from outer error as it is a valid context.Context type
	assert.Implements(t, (*context.Context)(nil), outerWithEmptyContext.(z.ContextAwareError).Context())

	// This should return the inner value from inner context as the outer context does not
	// have any value matching given key
	assert.Equal(t,
		innerContextTestValue,
		outerWithEmptyContext.(z.ContextAwareError).ContextValue(innerContextTestKey),
	)

	// Testing with ContextValue
	outerContextTestKey := "bar" // NOTE: We use the same key as in inner error. But expect to get first values here
	outerContextTestValue := "baz"
	outerWithContextValue := z.NewWithOpts(
		"outer",
		z.WithContextValue(outerContextTestKey, outerContextTestValue),
		z.WithWrappedError(errMiddleNoZError),
	)

	assert.Equal(t,
		outerContextTestValue,
		outerWithContextValue.(z.ContextAwareError).ContextValue(outerContextTestKey),
	)
	assert.Equal(t,
		innerContextTestValue,
		outerWithContextValue.(z.ContextAwareError).ContextValue(innerContextTestKey),
	)
}
