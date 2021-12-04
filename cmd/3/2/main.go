package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
)

type strArr []string
type gas int

const (
	oxygen gas = iota
	co2
)

func main() {
	input := strArr(helpers.ReadInputFile("cmd/3/2/input.txt"))

	oxygenGeneratorRate := input.filter(oxygen, 0)[0]
	co2ScrubberRate := input.filter(co2, 0)[0]

	println(oxygenGeneratorRate)
	println(co2ScrubberRate)

	oxygen, _ := strconv.ParseInt(oxygenGeneratorRate, 2, 64)
	co2, _ := strconv.ParseInt(co2ScrubberRate, 2, 64)

	fmt.Printf("Multiplication: %v", oxygen*co2)
}

func (arr strArr) filter(gasType gas, position int) strArr {
	result := make(strArr, 0)
	cb := 0 // common bit

	switch gasType {
	case oxygen:
		cb = mostCommonBit(arr, position)
	case co2:
		cb = leastCommonBit(arr, position)
	}

	for _, str := range arr {
		if strconv.Itoa(cb) == string(str[position]) {
			result = append(result, str)
			continue
		}

		if cb == -1 && gasType == oxygen && str[position] == '1' {
			result = append(result, str)
			continue
		}

		if cb == -1 && gasType == co2 && str[position] == '0' {
			result = append(result, str)
			continue
		}
	}

	if len(result) > 1 && position < len(arr[0])-1 {
		result = result.filter(gasType, position+1)
	}

	return result
}

func leastCommonBit(data []string, position int) int {
	mcb := mostCommonBit(data, position)

	switch mcb {
	case 1:
		return 0
	case 0:
		return 1
	}

	return mcb // equal
}

func mostCommonBit(data []string, position int) int {
	ones := 0
	for _, str := range data {
		if str[position] == '1' {
			ones++
		}
	}

	if ones*2 > len(data) {
		return 1
	}

	if ones*2 < len(data) {
		return 0
	}

	return -1 // equal
}
