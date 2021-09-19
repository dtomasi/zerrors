package zerrors

import (
	"fmt"
	"golang.org/x/xerrors"
)

/**
NOTE: This file just exposes the API of xerrors for easier use
and to avoid both packages have to be imported.
 */

// Frame see: https://pkg.go.dev/golang.org/x/xerrors#Frame
type Frame = xerrors.Frame

// Printer see: https://pkg.go.dev/golang.org/x/xerrors#Printer
type Printer = xerrors.Printer

// Formatter see: https://pkg.go.dev/golang.org/x/xerrors#Formatter
type Formatter = xerrors.Formatter

// Wrapper see: https://pkg.go.dev/golang.org/x/xerrors#Wrapper
type Wrapper = xerrors.Wrapper

// As see: https://pkg.go.dev/golang.org/x/xerrors#As
func As(err error, target interface{}) bool {
	return xerrors.As(err, target)
}

// Is see: https://pkg.go.dev/golang.org/x/xerrors#Is
func Is(err error, target error) bool {
	return xerrors.Is(err, target)
}

// Unwrap see: https://pkg.go.dev/golang.org/x/xerrors#Unwrap
func Unwrap(err error) error {
	return xerrors.Unwrap(err)
}

// Opaque see: https://pkg.go.dev/golang.org/x/xerrors#Opaque
func Opaque(err error) error {
	return xerrors.Opaque(err)
}

// Errorf see: https://pkg.go.dev/golang.org/x/xerrors#Errorf
func Errorf(format string, a ...interface{}) error {
	return xerrors.Errorf(format, a...)
}

// Caller see: https://pkg.go.dev/golang.org/x/xerrors#Caller
func Caller(skip int) Frame {
	return xerrors.Caller(skip)
}

// FormatError see: https://pkg.go.dev/golang.org/x/xerrors#FormatError
func FormatError(f Formatter, s fmt.State, verb rune) {
	xerrors.FormatError(f, s, verb)
}
