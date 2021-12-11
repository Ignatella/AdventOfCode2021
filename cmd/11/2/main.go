package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
)

func main() {
	input := helpers.ReadInputFile("cmd/11/2/input.txt")
	grid := parseInput(input)

	day := findStepWhenAllFlashes(grid)
	printGrid(grid)

	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")

	fmt.Printf("%v\n", day)
}

func findStepWhenAllFlashes(grid [][]int) int {
	n := 0
	flashes := 0
	for ; flashes != len(grid)*len(grid); n++ {
		flashes = step(grid)
	}

	return n
}

func step(grid [][]int) int {
	flashes := 0
	for i, r := range grid {
		for j, _ := range r {
			if grid[i][j] < 10 {
				grid[i][j]++
			}

			if grid[i][j] == 10 {
				flashes += flash(i, j, grid)
			}
		}
	}

	for i, r := range grid {
		for j, _ := range r {
			if grid[i][j] > 9 {
				grid[i][j] = 0
			}
		}
	}

	return flashes
}

func flash(i, j int, grid [][]int) int {
	flashes := 1
	grid[i][j] = 11

	for x := i - 1; x <= i+1; x++ {
		if x < 0 || x > len(grid)-1 {
			continue
		}

		if j-1 > -1 {
			if grid[x][j-1] < 10 {
				grid[x][j-1]++
			}

			if grid[x][j-1] == 10 {
				flashes += flash(x, j-1, grid)
			}
		}

		if j+1 < len(grid) {
			if grid[x][j+1] < 10 {
				grid[x][j+1]++
			}

			if grid[x][j+1] == 10 {
				flashes += flash(x, j+1, grid)
			}
		}
	}

	if i-1 > -1 {
		if grid[i-1][j] < 10 {
			grid[i-1][j]++
		}

		if grid[i-1][j] == 10 {
			flashes += flash(i-1, j, grid)
		}
	}

	if i+1 < len(grid) {
		if grid[i+1][j] < 10 {
			grid[i+1][j]++
		}

		if grid[i+1][j] == 10 {
			flashes += flash(i+1, j, grid)
		}
	}

	return flashes
}

func printGrid(grid [][]int) {
	for _, r := range grid {
		for _, el := range r {
			fmt.Printf("%v ", el)

		}

		fmt.Printf("\n")
	}
}

func parseInput(input []string) [][]int {
	grid := make([][]int, 0)
	for i, row := range input {
		grid = append(grid, make([][]int, 1)...)
		grid[i] = make([]int, len(row))

		for j, col := range row {
			el, _ := strconv.Atoi(string(col))
			grid[i][j] = el
		}
	}

	return grid
}
