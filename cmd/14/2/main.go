package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"math"
	"strings"
)

func main() {
	input := helpers.ReadInputFile("cmd/14/2/input.txt")
	str := input[0]
	rules := parseRules(input[2:])
	result := steps(40, str, rules)

	min, max := getMinMaxOccurrenceCount(result)

	fmt.Printf("Result: %v\n", max-min)
}

func steps(steps int, str string, rules map[string]uint8) map[uint8]uint64 {
	occurrences := calcCharOccurrence(str)
	pairs := splitToPairs(str)

	for ; steps > 0; steps-- {
		newPairs := make(map[string]uint64)
		for k := range pairs {
			c := rules[k]

			occurrences[c] += pairs[k]

			newPairs[string(k[0])+string(c)] += pairs[k]
			newPairs[string(c)+string(k[1])] += pairs[k]
			delete(pairs, k)
		}

		pairs = sum2Maps(pairs, newPairs)
	}

	return occurrences
}

func sum2Maps(m1 map[string]uint64, m2 map[string]uint64) map[string]uint64 {
	res := make(map[string]uint64)

	for k := range m1 {
		res[k] += m1[k]
	}

	for k := range m2 {
		res[k] += m2[k]
	}

	return res
}

func getMinMaxOccurrenceCount(m map[uint8]uint64) (uint64, uint64) {

	var min uint64 = math.MaxUint64
	var max uint64 = 0

	for _, v := range m {
		if v > max {
			max = v
		}

		if v < min {
			min = v
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

func splitToPairs(str string) map[string]uint64 {
	pairs := make(map[string]uint64)

	for i := 0; i < len(str)-1; i++ {
		pairs[string(str[i])+string(str[i+1])]++
	}

	return pairs
}

func parseRules(input []string) map[string]uint8 {
	rules := make(map[string]uint8, 0)

	for _, r := range input {
		parsed := strings.Split(r, " -> ")
		rules[parsed[0]] = parsed[1][0]
	}

	return rules
}
