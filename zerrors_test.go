package zerrors_test

import (
	"fmt"
	z "github.com/dtomasi/zerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	// This can be generated using //go:generate with stringer.
	MyTestError      = z.ErrorType("MyTestError")
	MyOtherTestError = z.ErrorType("MyOtherTestError")
)

func helperAssertErrorIsType(t *testing.T, err error, errType fmt.Stringer) {
	t.Helper()

	// Check if e is type of error
	assert.Error(t, err)

	// Check if Type() returns the expected type
	assert.Equal(t, err.(z.TypeAwareError).Type(), errType)

	// Check if IsType is returning true for type
	assert.True(t, z.IsType(err.(z.TypeAwareError), errType))

	// Check if IsType from TypeAwareError is returning true for type
	assert.True(t, err.(z.TypeAwareError).IsType(errType))

	// Test a switch case scenario
	switch err.(z.TypeAwareError).Type() {
	case errType:
	// All good here
	default:
		t.Errorf("expected switch does not fall into default case for type %s", errType)
	}
}

func TestNewWithOptsWithoutOpts(t *testing.T) {
	err := z.NewWithOpts("test")
	assert.Equal(t, err.Error(), "test")

	// Test type functions
	helperAssertErrorIsType(t, err, z.GenericError)

	// Test Masked err
	helperAssertErrorIsType(t, z.Mask(err), z.GenericError)
}

func TestNewWithOptsIsType(t *testing.T) {
	err := z.NewWithOpts(
		"test", // The message
		z.WithType(MyTestError),
	)
	assert.Equal(t, err.Error(), "test")

	// Test type functions
	helperAssertErrorIsType(t, err, MyTestError)

	// Test false positive
	assert.False(t, z.IsType(err.(z.TypeAwareError), MyOtherTestError))
	assert.False(t, err.(z.TypeAwareError).IsType(MyOtherTestError))

	// Test a switch case scenario on TypeAwareError
	switch err.(z.TypeAwareError).Type() {
	case MyTestError:
	case MyOtherTestError:
		t.Errorf("error type %s is matching %s. This should not be possible", MyOtherTestError, MyTestError)
	default:
		t.Errorf("error type not detected for %s", MyTestError)
	}
}

func TestNew(t *testing.T) {
	err := z.New("test")
	assert.Equal(t, err.Error(), "test")

	// Test type functions
	helperAssertErrorIsType(t, err, z.GenericError)
}

func TestNewf(t *testing.T) {
	err := z.Newf("%s", "test")
	assert.Equal(t, err.Error(), "test")

	// Test type functions
	helperAssertErrorIsType(t, err, z.GenericError)
}
