package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
)

var brackets = []uint8{'(', ')', '{', '}', '[', ']', '<', '>'}
var points = [][]int{{')', 3}, {']', 57}, {'}', 1197}, {'>', 25137}}

func main() {
	input := helpers.ReadInputFile("cmd/10/1/input.txt")

	score := calcScore(input)

	fmt.Printf("Score: %v\n", score)
}

func calcScore(input []string) int {
	result := 0

	for _, str := range input {
		_, _, found := isCorrupted(str)

		result += getPoints(found)
	}

	return result
}

func isCorrupted(source string) (corrupted bool, expected, found uint8) {
	chunks := make([]uint8, 0)

	for i := 0; i < len(source); i++ {
		if isOpeningBracket(source[i]) {
			chunks = append(chunks, source[i])
			continue
		}

		// eg opposite to { != ...
		if getOppositeBracket(chunks[len(chunks)-1]) != source[i] {
			//fmt.Printf("Expected %c, but found %c instead.\n", getOppositeBracket(chunks[len(chunks)-1]), source[i])
			return true, getOppositeBracket(chunks[len(chunks)-1]), source[i]
		}

		chunks = append(chunks[:len(chunks)-1], chunks[len(chunks):]...)

	}

	return false, '0', '0'
}

func getPoints(bracket uint8) int {
	for _, p := range points {
		if p[0] == int(bracket) {
			return p[1]
		}
	}

	return 0
}

func isOpeningBracket(c uint8) bool {
	return !isClosingBracket(c)
}

func isClosingBracket(c uint8) bool {
	result := false

	for i := 1; i < len(brackets); i += 2 {
		result = result || (brackets[i] == c)
	}

	return result
}

func getOppositeBracket(c uint8) uint8 {
	for i := 0; i < len(brackets); i++ {
		if c == brackets[i] {
			if i%2 == 0 {
				return brackets[i+1]
			}

			return brackets[i-1]
		}
	}

	panic("No bracket found")
}
