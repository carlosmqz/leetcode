package main

import (
	"fmt"

	"github.com/carlosmqz/leetcode/solutions"
)

func main() {
	//candles := solutions.GetNumberOfCandles()
	timeConv := solutions.ConvertTime("10:00:01PM")
	//fmt.Println("Candles: ", candles)
	fmt.Println(timeConv)
}
