package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"sort"
	"strconv"
	"strings"
)

type dots []dot

type dot struct {
	x, y int
}

type fold struct {
	c            int
	isHorizontal bool
}

func main() {
	input := helpers.ReadInputFile("cmd/13/2/input.txt")
	dts, folds := parseInput(input)
	for _, f := range folds {
		dts.fold(f)
		dts.Unique()
	}
	dts.print()
}

func (dts dots) print() {
	sort.Sort(dts)
	mX, mY := dts.MaxXY()

	for i := 0; i <= mY; i++ {
		for j := 0; j <= mX; j++ {
			if dts.Contains(j, i) {
				fmt.Printf("# ")
			} else {
				fmt.Print(". ")
			}

		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func (dts dots) fold(f fold) {
	for i, d := range dts {
		if f.isHorizontal {
			if d.y > f.c {
				dts[i].y -= 2 * (d.y - f.c)
			}
		} else {
			if d.x > f.c {
				dts[i].x -= 2 * (d.x - f.c)
			}
		}
	}
}

func parseInput(input []string) (dots, []fold) {
	dots := make([]dot, 0)
	folds := make([]fold, 0)

	i := 0

	// region Dots
	for ; i < len(input) && input[i] != ""; i++ {
		coords := strings.Split(input[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		dots = append(dots, dot{x: x, y: y})
	}
	// endregion

	i++

	// region Folds
	for ; i < len(input); i++ {
		f := strings.Split(strings.Split(input[i], "fold along ")[1], "=")
		c, _ := strconv.Atoi(f[1])

		folds = append(folds, fold{c: c, isHorizontal: f[0] == "y"})
	}
	// endregion

	return dots, folds
}

// region Dots helpers

func (dts *dots) Unique() {
	sort.Sort(dts)

	for i := 0; i < len(*dts)-1; i++ {
		if (*dts)[i] == (*dts)[i+1] {
			*dts = append((*dts)[:i+1], (*dts)[i+2:]...)
			i--
		}
	}
}

func (dts dots) MaxXY() (maxX int, maxY int) {
	for _, d := range dts {
		if d.x > maxX {
			maxX = d.x
		}

		if d.y > maxY {
			maxY = d.y
		}
	}

	return
}

func (dts dots) Contains(x, y int) bool {
	contains := false

	for _, d := range dts {
		contains = contains || (d.x == x && d.y == y)
	}

	return contains
}

func (dts dots) Swap(i, j int) {
	dts[i], dts[j] = dts[j], dts[i]
}

func (dts dots) Len() int {
	return len([]dot(dts))
}

func (dts dots) Less(i, j int) bool {
	if dts[i].x == dts[j].x {
		return dts[i].y < dts[j].y
	}

	return dts[i].x < dts[j].x
}

func (dts dots) Equal(i, j int) bool {
	return dts[i].x == dts[j].x && dts[i].y == dts[j].y
}

// endregion
