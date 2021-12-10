package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strings"
)

/** Indexes
   _0_
1 |   | 2
  |_3_|
4 |   | 5
   _6_
*/

type str string

type segmentNames struct {
	names [8]rune
}

func main() {
	input := helpers.ReadInputFile("cmd/8/2/input.txt")
	result := getSumOfInput(input)

	fmt.Printf("Result: %v\n", result)
}

func getSumOfInput(input []string) int {
	res := 0

	for _, in := range input {
		inputSDigits := strings.Split(in, " | ")[0]
		inputDigits := strings.Split(inputSDigits, " ")
		outputSDigits := strings.Split(in, " | ")[1]
		outputDigits := strings.Split(outputSDigits, " ")

		sNames := determineSegmentsName(inputDigits)

		num := 0
		x := 1000
		for _, d := range outputDigits {
			num += convertDigitToNum(d, sNames) * x
			x /= 10
		}

		res += num
	}

	return res
}

func convertDigitToNum(num string, s *segmentNames) int {
	if is0(num, s) {
		return 0
	}

	if is1(num, s) {
		return 1
	}

	if is2(num, s) {
		return 2
	}

	if is3(num, s) {
		return 3
	}

	if is4(num, s) {
		return 4
	}

	if is5(num, s) {
		return 5
	}

	if is6(num, s) {
		return 6
	}

	if is7(num, s) {
		return 7
	}

	if is8(num, s) {
		return 8
	}

	if is9(num, s) {
		return 9
	}

	return -1
}

func determineSegmentsName(digits []string) *segmentNames {
	result := segmentNames{names: [8]rune{}}
	segmentSums := [7][3]int32{{'a', 0, 0}, {'b', 0, 0}, {'c', 0, 0},
		{'d', 0, 0}, {'e', 0, 0}, {'f', 0, 0}, {'g', 0, 0}}

	// region Segment Sums
	for _, d := range digits {
		if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
			for _, c := range d {
				for i := 0; i < len(segmentSums); i++ {
					if segmentSums[i][0] == c {
						segmentSums[i][1]++
					}
				}
			}

			continue
		}

		for _, c := range d {
			for i := 0; i < len(segmentSums); i++ {
				if segmentSums[i][0] == c {
					segmentSums[i][2]++
				}
			}
		}
	}
	// endregion

	for _, r := range segmentSums {
		if r[1] == 2 && r[2] == 6 {
			result.names[0] = r[0]
		}

		if r[1] == 2 && r[2] == 4 {
			result.names[1] = r[0]
		}

		if r[1] == 4 && r[2] == 4 {
			result.names[2] = r[0]
		}

		if r[1] == 2 && r[2] == 5 {
			result.names[3] = r[0]
		}

		if r[1] == 1 && r[2] == 3 {
			result.names[4] = r[0]
		}

		if r[1] == 4 && r[2] == 5 {
			result.names[5] = r[0]
		}

		if r[1] == 1 && r[2] == 6 {
			result.names[6] = r[0]
		}
	}

	return &result
}

func (s *str) contains(char rune) bool {
	for _, c := range *s {
		if char == c {
			return true
		}
	}

	return false
}

// region isX
func is0(num string, s *segmentNames) bool {
	n := s.names
	st := str(num)

	if st.contains(n[0]) &&
		st.contains(n[1]) &&
		st.contains(n[2]) &&
		st.contains(n[3]) == false &&
		st.contains(n[4]) &&
		st.contains(n[5]) &&
		st.contains(n[6]) {
		return true
	}

	return false
}

func is1(num string, s *segmentNames) bool {
	if len(num) == 2 {
		return true
	}

	return false
}

func is2(num string, s *segmentNames) bool {
	n := s.names
	st := str(num)

	if st.contains(n[0]) &&
		st.contains(n[1]) == false &&
		st.contains(n[2]) &&
		st.contains(n[3]) &&
		st.contains(n[4]) &&
		st.contains(n[5]) == false &&
		st.contains(n[6]) {
		return true
	}

	return false
}

func is3(num string, s *segmentNames) bool {
	n := s.names
	st := str(num)

	if st.contains(n[0]) &&
		st.contains(n[1]) == false &&
		st.contains(n[2]) &&
		st.contains(n[3]) &&
		st.contains(n[4]) == false &&
		st.contains(n[5]) &&
		st.contains(n[6]) {
		return true
	}

	return false
}

func is4(num string, s *segmentNames) bool {
	if len(num) == 4 {
		return true
	}

	return false
}

func is5(num string, s *segmentNames) bool {
	n := s.names
	st := str(num)

	if st.contains(n[0]) &&
		st.contains(n[1]) &&
		st.contains(n[2]) == false &&
		st.contains(n[3]) &&
		st.contains(n[4]) == false &&
		st.contains(n[5]) &&
		st.contains(n[6]) {
		return true
	}

	return false
}

func is6(num string, s *segmentNames) bool {
	n := s.names
	st := str(num)

	if st.contains(n[0]) &&
		st.contains(n[1]) &&
		st.contains(n[2]) == false &&
		st.contains(n[3]) &&
		st.contains(n[4]) &&
		st.contains(n[5]) &&
		st.contains(n[6]) {
		return true
	}

	return false
}

func is7(num string, s *segmentNames) bool {
	if len(num) == 3 {
		return true
	}

	return false
}

func is8(num string, s *segmentNames) bool {
	if len(num) == 7 {
		return true
	}

	return false
}

func is9(num string, s *segmentNames) bool {
	n := s.names
	st := str(num)

	if st.contains(n[0]) &&
		st.contains(n[1]) &&
		st.contains(n[2]) &&
		st.contains(n[3]) &&
		st.contains(n[4]) == false &&
		st.contains(n[5]) &&
		st.contains(n[6]) {
		return true
	}

	return false
}

// endregion
