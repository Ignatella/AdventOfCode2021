package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"sort"
	"strconv"
)

const basinsCount int = 3

type point struct {
	i, j      int
	depth     int
	isVisited bool
}

func main() {
	input := helpers.ReadInputFile("cmd/9/2/input.txt")

	parsedInput := parseInput(input)

	sizes := calculateLargestBasins(parsedInput)

	multiplication := 1
	for _, s := range sizes {
		multiplication *= s
	}

	fmt.Printf("Result: %v\n", multiplication)
}

func calculateLargestBasins(points [][]point) []int {
	basinSizes := make([]int, basinsCount)

	for _, r := range points {
		for _, v := range r {
			if isLowPoint(v, points) {
				size := calculateBasinSize(v, points)

				if size > basinSizes[0] {
					basinSizes[0] = size
					sort.Ints(basinSizes)
				}
			}
		}
	}

	return basinSizes
}

func calculateBasinSize(p point, points [][]point) int {
	size := 1

	points[p.i][p.j].isVisited = true

	if p.i > 0 && !points[p.i-1][p.j].isVisited && points[p.i-1][p.j].depth != 9 {
		size += calculateBasinSize(points[p.i-1][p.j], points)
	}

	if p.i < len(points)-1 && !points[p.i+1][p.j].isVisited && points[p.i+1][p.j].depth != 9 {
		size += calculateBasinSize(points[p.i+1][p.j], points)
	}

	if p.j > 0 && !points[p.i][p.j-1].isVisited && points[p.i][p.j-1].depth != 9 {
		size += calculateBasinSize(points[p.i][p.j-1], points)
	}

	if p.j < len(points[0])-1 && !points[p.i][p.j+1].isVisited && points[p.i][p.j+1].depth != 9 {
		size += calculateBasinSize(points[p.i][p.j+1], points)
	}

	return size
}

func isLowPoint(p point, points [][]point) bool {
	isLow := true

	if p.i > 0 {
		isLow = isLow && (points[p.i][p.j].depth < points[p.i-1][p.j].depth)
	}

	if p.i < len(points)-1 {
		isLow = isLow && (points[p.i][p.j].depth < points[p.i+1][p.j].depth)
	}

	if p.j > 0 {
		isLow = isLow && (points[p.i][p.j].depth < points[p.i][p.j-1].depth)
	}

	if p.j < len(points[0])-1 {
		isLow = isLow && (points[p.i][p.j].depth < points[p.i][p.j+1].depth)
	}

	return isLow
}

func parseInput(input []string) [][]point {
	result := make([][]point, len(input))

	for i, r := range input {
		result[i] = make([]point, len(r))

		for j, c := range r {
			n, _ := strconv.Atoi(string(c))
			result[i][j] = point{i: i, j: j, depth: n, isVisited: false}
		}
	}

	return result
}
