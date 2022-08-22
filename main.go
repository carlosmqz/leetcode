package main

import (
	"fmt"

	"github.com/carlosmqz/leetcode/solutions"
)

func main() {
	//candles := solutions.GetNumberOfCandles()
	//timeConv := solutions.ConvertTime("10:59:59PM")
	//fmt.Println("Candles: ", candles)
	//fmt.Println(timeConv)
	//timeConv := solutions.ConvertTime("12:59:59AM")
	//fmt.Println("Candles: ", candles)
	//fmt.Println(timeConv)
	//fmt.Println(solutions.CircularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
	//fmt.Println(solutions.RepeatedString("afcfffaged", 962645758455))
	//fmt.Println(solutions.RepeatedString("aba", 10))
	//fmt.Println(solutions.JumpingOnClouds([]int32{0, 0, 0, 1, 0, 0}))    //3
	//fmt.Println(solutions.JumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})) //4
	//fmt.Println(solutions.JumpingOnClouds([]int32{0, 0, 0, 0, 0, 1, 0})) //3
	/*sample := [][]int32{{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}*/
	sample := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	fmt.Println(solutions.HourglassSum(sample))
}
