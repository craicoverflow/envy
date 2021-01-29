package envy

import (
	"os"
	"strconv"
)

// Get returns an environment variable as a string
func Get(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", ErrNotFound
	}

	return val, nil
}

// ParseBool returns an environment variable parsed to a Boolean
func ParseBool(key string) (bool, error) {
	val := os.Getenv(key)

	if val == "" {
		return false, notFoundError(key, "ParseBool")
	}

	parsedVal, err := strconv.ParseBool(val)
	if err != nil {
		return false, &EnvError{"ParseBool", key, err}
	}

	return parsedVal, err
}

// ParseInt returns an environment variable parsed to a Boolean
func ParseInt(key string, base int, bitSize int) (int64, error) {
	fnParseInt := "ParseInt"

	val := os.Getenv(key)

	if val == "" {
		return 0, notFoundError(key, fnParseInt)
	}

	parsedVal, err := strconv.ParseInt(val, base, bitSize)

	if err != nil {
		return 0, &EnvError{fnParseInt, key, err}
	}

	return parsedVal, err
}

// ParseFloat returns an environment variable parsed to a Boolean
func ParseFloat(key string, bitSize int) (float64, error) {
	fnParseFloat := "ParseFloat"

	val := os.Getenv(key)

	if val == "" {
		return 0, notFoundError(key, fnParseFloat)
	}

	parsedVal, err := strconv.ParseFloat(val, bitSize)

	if err != nil {
		return 0, &EnvError{fnParseFloat, key, err}
	}

	return parsedVal, err
}
