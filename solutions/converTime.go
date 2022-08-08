package solutions

import (
	"strconv"
	"strings"
)

func ConvertTime(s string) string {
	entryTime := strings.Split(s, ":")
	AMPM := entryTime[2][2:]
	hours, hoursErr := strconv.Atoi(entryTime[0])
	if hoursErr != nil {
		panic(hoursErr)
	}
	if AMPM == "AM" {
		if hours == 12 {
			return strings.Join([]string{"00", entryTime[1], entryTime[2][:2]}, ":")
		} else if hours < 10 {
			return strings.Join([]string{"0" + strconv.Itoa(hours), entryTime[1], entryTime[2][:2]}, ":")
		}
	} else if AMPM == "PM" && hours < 12 {
		hours += 12
	}

	return strings.Join([]string{strconv.Itoa(hours), entryTime[1], entryTime[2][:2]}, ":")
}
