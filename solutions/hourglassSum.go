package solutions

func HourglassSum(arr [][]int32) int32 {
	arrLen := len(arr) - 2
	max := int32(-500)
	sum := make([]int32, 0)
	for col := 0; col < arrLen; col++ {
		for row := 0; row < arrLen; row++ {
			sum = append(sum, arr[row][col]+arr[row][col+1]+arr[row][col+2]+
				arr[row+1][col+1]+
				arr[row+2][col]+arr[row+2][col+1]+arr[row+2][col+2])
		}
	}

	for _, v := range sum {
		if v > max {
			max = v
		}
	}

	return max
}
