package zerrors_test

import (
	z "github.com/dtomasi/zerrors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorType_String(t *testing.T) {
	testType := z.ErrorType("foo")
	assert.Equal(t, "foo", testType.String())
}
