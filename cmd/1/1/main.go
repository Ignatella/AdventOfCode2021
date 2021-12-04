package main

import (
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
)

func main() {
	input := helpers.ReadInputFile("cmd/1/2/input.txt")

	result := 0

	for i := 1; i < len(input); i++ {
		currNum, _ := strconv.Atoi(input[i])
		prevNum, _ := strconv.Atoi(input[i-1])

		if currNum > prevNum {
			result++
		}
	}

	println(result)
}
