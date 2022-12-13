package common

import (
	"strconv"
	"strings"
)

// TrimSplit trims a string and splits the result by a separator
func TrimSplit(str string, prefix, separator string) []string {
	return strings.Split(strings.TrimPrefix(str, prefix), separator)
}

// ToUInt64s converts a slice of strings to a slice of uint64s
func ToUInt64s(str ...string) []uint64 {
	ints := make([]uint64, 0, len(str))
	for _, str := range str {
		v, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			ints = append(ints, v)
		}
	}
	return ints
}
