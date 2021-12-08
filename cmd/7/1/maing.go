package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"math"
	"sort"
	"strconv"
	"strings"
)

/**
In this task it enough to find median of the positions to determine optimal position all over positions should be
	aligned to
*/

func main() {
	input := helpers.ReadInputFile("cmd/7/1/input.txt")[0]

	positions := parseInput(input)

	m := median(positions)
	cost := calculateCost(positions, m)

	fmt.Printf("Result: %v", cost)
}

func calculateCost(position []int, pos int) int {
	cost := 0
	for _, p := range position {
		cost += getDistance(p, pos)
	}

	return cost
}

func getDistance(pos1, pos2 int) int {
	return int(math.Abs(float64(pos1 - pos2)))
}

func median(position []int) int {
	sort.Ints(position)
	l := len(position)
	if l%2 == 0 {
		return (position[(l-1)/2] + position[(l+1)/2]) / 2
	}

	return position[l/2]
}

func parseInput(input string) []int {
	nums := strings.Split(input, ",")
	result := make([]int, len(nums))

	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		result[i] = n
	}

	return result
}
