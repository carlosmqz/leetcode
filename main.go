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
	fmt.Println(solutions.JumpingOnClouds([]int32{0, 0, 0, 1, 0, 0}))    //3
	fmt.Println(solutions.JumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})) //4
	fmt.Println(solutions.JumpingOnClouds([]int32{0, 0, 0, 0, 0, 1, 0})) //3

}
