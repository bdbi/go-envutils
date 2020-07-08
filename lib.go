package envutils

import (
	"fmt"
	"os"
)

// MustGetenv ...
func MustGetenv(key string) (val string, err error) {
	val = os.Getenv(key)
	if len(val) == 0 {
		err = fmt.Errorf("environment key: %s not set", key)
	}
	return val, err
}

// Getenv ...
func Getenv(key, def string) (val string) {
	val = os.Getenv(key)
	if len(val) == 0 {
		val = def
	}
	return val
}
