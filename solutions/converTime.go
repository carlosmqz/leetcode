package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertTime(s string) string {
	entryTime := strings.Split(s, ":")
	AMPM := entryTime[2][2:]

	fmt.Println(entryTime[0])
	fmt.Println(AMPM)

	hours, hoursErr := strconv.Atoi(entryTime[0])
	if hoursErr != nil {
		panic(hoursErr)
	}
	if AMPM == "PM" {
		hours += 12
		if hours == 24 {
			return strings.Join([]string{"00", entryTime[1], entryTime[2][:2]}, ":")
		}
	}
	return strings.Join([]string{strconv.Itoa(hours), entryTime[1], entryTime[2][:2]}, ":")
}
