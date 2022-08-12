package solutions

import (
	"strings"
)

func RepeatedString(s string, n int64) int64 {
	numofS := int64(n / int64(len(s)))
	counter := int64(strings.Count(s, "a"))
	rest := n % int64(len(s))
	countRest := int64(strings.Count(s[:rest], "a"))
	if len(s) == 1 {
		return int64(strings.Count(s, "a")) * n
	}
	return numofS*counter + countRest
}
