package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
)

func main() {
	input := helpers.ReadInputFile("cmd/9/1/input.txt")

	parsedInput := parseInput(input)

	risk := calculateRisk(parsedInput)

	fmt.Printf("Result: %v\n", risk)
}

func calculateRisk(heights [][]int) int {
	risk := 0

	for i, r := range heights {
		for j, v := range r {
			if isLowPoint(i, j, heights) {
				risk += 1 + v
			}
		}
	}

	return risk
}

func isLowPoint(i, j int, heights [][]int) bool {
	isLow := true

	if i > 0 {
		isLow = isLow && (heights[i][j] < heights[i-1][j])
	}

	if i < len(heights)-1 {
		isLow = isLow && (heights[i][j] < heights[i+1][j])
	}

	if j > 0 {
		isLow = isLow && (heights[i][j] < heights[i][j-1])
	}

	if j < len(heights[0])-1 {
		isLow = isLow && (heights[i][j] < heights[i][j+1])
	}

	return isLow
}

func parseInput(input []string) [][]int {
	result := make([][]int, len(input))

	for i, r := range input {
		result[i] = make([]int, len(r))

		for j, c := range r {
			n, _ := strconv.Atoi(string(c))
			result[i][j] = n
		}
	}

	return result
}
