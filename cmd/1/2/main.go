package main

import (
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
)

const slidingWindowSize int = 3

func main() {
	input := helpers.ReadInputFile("cmd/1/2/input.txt")

	result := 0

	for i := 0; i < len(input)-slidingWindowSize; i++ {
		firstNum, _ := strconv.Atoi(input[i])
		fourthNum, _ := strconv.Atoi(input[i+3])

		if fourthNum > firstNum {
			result++
		}
	}

	println(result)
}
