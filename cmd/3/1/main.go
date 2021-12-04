package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
)

func main() {
	input := helpers.ReadInputFile("cmd/3/1/input.txt")

	s := make([]int, len(input[0]))

	/**
	Calculation of the most frequent bit (1 or 0) in each column
	*/
	for _, num := range input {
		for i, c := range num {
			bit, _ := strconv.Atoi(string(c))
			s[i] += bit
		}
	}

	gammaRate := 0
	epsilonRate := 0

	/**
	Calculation of gammaRate & epsilonRate
	*/
	for i, v := range s {
		if 2*v >= len(input) {
			gammaRate = (1 << (len(s) - i - 1)) | gammaRate
		} else {
			epsilonRate = (1 << (len(s) - i - 1)) | epsilonRate
		}
	}

	fmt.Printf("%b\n", gammaRate)
	fmt.Printf("%b\n", epsilonRate)

	fmt.Printf("Multiplication: %d\n", gammaRate*epsilonRate)
}
