package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"math"
	"strings"
)

func main() {
	input := helpers.ReadInputFile("cmd/14/1/input.txt")
	str := input[0]
	rules := parseRules(input[2:])
	result := steps(10, str, rules)

	min, max := getMinMaxOccurrence(result)

	fmt.Printf("Result: %v\n", max-min)
}

func getMinMaxOccurrence(str string) (uint64, uint64) {
	occurrences := calcCharOccurrence(str)

	var min uint64 = math.MaxUint64
	var max uint64 = 0

	for _, k := range occurrences {
		if k > max {
			max = k
		}

		if k < min {
			min = k
		}
	}

	return min, max
}

func calcCharOccurrence(str string) map[uint8]uint64 {
	result := make(map[uint8]uint64, 0)
	for _, c := range str {
		result[uint8(c)]++
	}

	return result
}

func steps(stepCount int, source string, rules map[string]uint8) string {
	for ; stepCount > 0; stepCount-- {
		source = step(source, rules)
	}

	return source
}

func step(source string, rules map[string]uint8) string {
	result := ""

	for i := 0; i < len(source)-1; i++ {
		result += string(source[i])
		el := string(source[i]) + string(source[i+1])
		if rules[el] != 0 {
			result += string(rules[el])
		}
	}

	result += string(source[len(source)-1])

	return result
}

func parseRules(input []string) map[string]uint8 {
	rules := make(map[string]uint8, 0)

	for _, r := range input {
		parsed := strings.Split(r, " -> ")
		rules[parsed[0]] = parsed[1][0]
	}

	return rules
}
