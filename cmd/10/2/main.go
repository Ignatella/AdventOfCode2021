package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"sort"
)

var brackets = []uint8{'(', ')', '{', '}', '[', ']', '<', '>'}
var points = [][]int{{')', 1}, {']', 2}, {'}', 3}, {'>', 4}}

func main() {
	input := helpers.ReadInputFile("cmd/10/2/input.txt")

	score := calcScore(input)

	fmt.Printf("Middle score: %v\n", score)
}

func calcScore(input []string) int {
	scores := make([]int, 0)

	for _, str := range input {
		compBrackets := getComplementaryBrackets(str)
		score := 0

		for _, c := range compBrackets {
			score = score*5 + getPoints(c)
		}

		if score != 0 {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2] // according to considerations
}

func getComplementaryBrackets(source string) []uint8 {
	chunks := make([]uint8, 0)

	for i := 0; i < len(source); i++ {
		if isOpeningBracket(source[i]) {
			chunks = append(chunks, source[i])
			continue
		}

		// if isCorrupted
		if getOppositeBracket(chunks[len(chunks)-1]) != source[i] {
			return make([]uint8, 0)
		}

		chunks = append(chunks[:len(chunks)-1], chunks[len(chunks):]...)

	}

	// calculate complementary part
	for i := len(chunks) - 1; i > -1; i-- {
		chunks = append(chunks, getOppositeBracket(chunks[i]))
		chunks = append(chunks[:i], chunks[i+1:]...)
	}

	return chunks
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

	panic("No bracket found\n")
}
