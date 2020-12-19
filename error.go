package envy

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrNotFound defines when an environment variable does not exist
	ErrNotFound = errors.New("environment variable does not exist")
	// ErrSyntax defines when a value does not have the correct syntax for the target type
	ErrSyntax = errors.New("invalid syntax")
	// ErrRange defines when a value has an invalid range
	ErrRange = errors.New("invalid range")
)

// EnvError records a failed conversion
type EnvError struct {
	Func string // The failing function
	Key  string // The environment variable key
	Err  error  // the reason the conversion failed (e.g. ErrNotExists, ErrSyntax, etc.)
}

func (e *EnvError) Error() string {
	return fmt.Sprintf("envy.%v: parsing %v: %v", e.Func, strconv.Quote(e.Key), e.Err.Error())
}

func (e *EnvError) Unwrap() error { return e.Err }

func syntaxError(key string, fn string) *EnvError {
	return &EnvError{fn, key, ErrSyntax}
}

func notFoundError(key string, fn string) *EnvError {
	return &EnvError{fn, key, ErrNotFound}
}

func rangeError(key string, fn string) *EnvError {
	return &EnvError{fn, key, ErrRange}
}

func baseError(fn, str string, base int) *EnvError {
	return &EnvError{fn, str, errors.New("invalid base " + strconv.Itoa(base))}
}

func bitSizeError(fn, str string, bitSize int) *EnvError {
	return &EnvError{fn, str, errors.New("invalid bit size " + strconv.Itoa(bitSize))}
}
