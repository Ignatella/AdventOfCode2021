package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"math"
	"sort"
	"strconv"
	"strings"
)

type costFunc func(pos1, pos2 int) int

func main() {
	input := helpers.ReadInputFile("cmd/7/2/input.txt")[0]

	positions := parseInput(input)

	minimalCost := calculateMinimalCost(positions, cost)

	fmt.Printf("Result: %v\n", minimalCost)
}

func calculateMinimalCost(positions []int, cost costFunc) int {
	minimalCost := math.MaxInt
	sort.Ints(positions)

	for i := positions[0]; i < positions[len(positions)-1]; i++ {

		currCost := calculateCost(positions, i, cost)

		if currCost < minimalCost {
			minimalCost = currCost
		}
	}

	return minimalCost
}

func calculateCost(positions []int, pos int, cost costFunc) int {
	result := 0
	for _, p := range positions {
		result += cost(p, pos)
	}

	return result
}

func cost(pos1, pos2 int) int {
	distance := getDistance(pos1, pos2)

	return ((1 + distance) * distance) / 2
}

func getDistance(pos1, pos2 int) int {
	return int(math.Abs(float64(pos1 - pos2)))
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
