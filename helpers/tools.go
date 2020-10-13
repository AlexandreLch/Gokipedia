package helpers

import (
	"fmt"
	"strconv"
)

// ParseUInt64 helper to avoid code repetition
func ParseUInt64(stringToParse string) (uint64, error) {
	intID, err := strconv.ParseInt(stringToParse, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse string to int")
	}
	return uint64(intID), nil
}
