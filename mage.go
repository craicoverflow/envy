//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Runs unit tests
func Test() error {
	return sh.Run("go", "test", "-v", "./...")
}

// Format all Go files
func Format() error {
	return sh.Run("go", "fmt", "./...")
}
