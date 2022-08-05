package solutions

func GetNumberOfCandles() int {
	candles := []int{3, 1, 2, 3}
	count := make(map[int]int)
	max := 0
	for _, v := range candles {
		count[v] += 1
		if count[v] > max {
			max = count[v]
		}
	}
	return max
}
