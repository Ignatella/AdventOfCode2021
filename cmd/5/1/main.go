package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/cmd/5/internal/line"
	"github.com/Ignatella/AdventOfCode2021/cmd/5/internal/plane"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
)

/**
Assuming lines can be only horizontal or vertical
*/

func main() {
	input := helpers.ReadInputFile("cmd/5/1/input.txt")

	pl := plane.New()

	for _, p := range input {
		l := line.New(p)
		if l.IsHorizontalOrVertical() {
			pl.PutLineOnPlane(l)
		}
	}

	result := pl.TotalCrossNumber()

	fmt.Printf("Total cross number: %v\n", result)
}
