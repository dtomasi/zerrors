package zerrors

// ErrorType defines an easy to use fmt.Stringer compatible error type.
// Example:
// 		zerrors.ErrorType("MyNamedErrorType")
type ErrorType string

// String returns the ErrorType string.
// This function is required to implement fmt.Stringer interface.
func (t ErrorType) String() string {
	return string(t)
}

// GenericError is the default error type.
const GenericError = ErrorType("GenericError")
