package main

import (
	"fmt"

	"github.com/carlosmqz/leetcode/solutions"
)

func main() {
	//candles := solutions.GetNumberOfCandles()
	timeConv := solutions.ConvertTime("12:59:59AM")
	//fmt.Println("Candles: ", candles)
	fmt.Println(timeConv)
}
