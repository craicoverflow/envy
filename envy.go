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
		return false, syntaxError(key, "ParseBool")
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

	if err == nil {
		return parsedVal, nil
	}

	if err.(*strconv.NumError).Err == strconv.ErrRange {
		err = &EnvError{
			Func: fnParseInt,
			Key:  key,
			Err:  ErrRange,
		}

		return parsedVal, err
	}

	if err != nil {
		err = &EnvError{
			Func: fnParseInt,
			Key:  key,
			Err:  err.(*strconv.NumError).Err,
		}
	}

	return parsedVal, err
}

// ParseFloat returns an environment variable parsed to a Boolean
func ParseFloat(key string, bitSize int) (float64, error) {
	val := os.Getenv(key)

	if val == "" {
		return 0, notFoundError(key, "ParseInt")
	}

	parsedVal, err := strconv.ParseFloat(val, bitSize)

	if err != nil {
		return 0, syntaxError(key, "ParseInt")
	}

	return parsedVal, err
}
