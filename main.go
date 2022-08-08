package main

import (
	"fmt"

	"github.com/carlosmqz/leetcode/solutions"
)

func main() {
	//candles := solutions.GetNumberOfCandles()
	timeConv := solutions.ConvertTime("01:00:01AM")
	//fmt.Println("Candles: ", candles)
	fmt.Println(timeConv)
}
