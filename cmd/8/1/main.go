package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strings"
)

type digit struct {
	value        int
	segmentCount int
}

func main() {
	input := helpers.ReadInputFile("cmd/8/1/input.txt")

	knownDigits :=
		[]digit{
			{1, 2},
			{4, 4},
			{7, 3},
			{8, 7}}

	countOfKnownDigits := calcCountOfKnownDigits(input, knownDigits)

	fmt.Printf("Resutl: %v\n", countOfKnownDigits)
}

func calcCountOfKnownDigits(input []string, knownDigits []digit) int {
	result := 0

	for _, in := range input {
		outputValues := strings.Split(in, " | ")[1]

		values := strings.Split(outputValues, " ")

		for _, v := range values {
			for _, d := range knownDigits {
				if len(v) == d.segmentCount {
					result++
				}
			}
		}
	}

	return result
}
