package solutions

import (
	"fmt"
	"strings"
)

func ConvertTime(s string) string {
	entryTime := strings.Split(s, ":")
	AMPM := entryTime[2][2:]

	fmt.Println(entryTime[0])
	fmt.Println(AMPM)

	return ""
}
