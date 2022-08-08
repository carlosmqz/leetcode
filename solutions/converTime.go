package solutions

import (
	"fmt"
	"strings"
)

func convertTime(s string) string {
	entryTime := strings.Split(s, ":")
	AMPM := s[:2]

	fmt.Println(entryTime)
	fmt.Println(AMPM)

	return ""
}
