package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
	"strings"
)

const birthTimer int = 7
const childShift int = 2
const numOfDays int = 80

type fish struct {
	birthAfter int
}

func main() {
	input := helpers.ReadInputFile("cmd/6/1/input.txt")
	f := createFishFromString(input[0])

	f = live(numOfDays, f)

	fmt.Printf("Result: %v", len(f))
}

func live(days int, _fish []*fish) []*fish {
	for i := 0; i < days; i++ {
		bornThisDay := make([]*fish, 0)

		for _, f := range _fish {
			bornFish := (*f).liveADay()
			if bornFish != nil {
				bornThisDay = append(bornThisDay, bornFish)
			}
		}

		_fish = append(_fish, bornThisDay...)
	}

	return _fish
}

func (f *fish) liveADay() *fish {
	if f.birthAfter == 0 {
		f.birthAfter = birthTimer - 1
		return &fish{birthAfter: birthTimer + childShift - 1}
	}

	f.birthAfter--

	return nil
}

func createFishFromString(input string) []*fish {
	nums := strings.Split(input, ",")
	result := make([]*fish, len(nums))

	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		result[i] = &fish{birthAfter: n}
	}

	return result
}
