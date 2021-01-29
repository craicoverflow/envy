package envy

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrNotFound defines when an environment variable does not exist
	ErrNotFound = errors.New("environment variable does not exist")
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

func notFoundError(key string, fn string) *EnvError {
	return &EnvError{fn, key, ErrNotFound}
}
