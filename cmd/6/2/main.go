package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
	"strings"
)

const birthTimer int = 7
const childShift int = 2
const numOfDays int = 256

type fish struct {
	birthAfter int
	x          uint64
}

func main() {
	input := helpers.ReadInputFile("cmd/6/2/input.txt")
	f := createFishFromString(birthTimer, input[0])

	f = live(numOfDays, f)

	var res uint64 = 0
	for _, fish := range f {
		res += fish.x
	}
	fmt.Printf("Result: %v", res)
}

func live(days int, _fish []*fish) []*fish {
	for i := 0; i < days; i++ {
		var bornThisDay *fish = nil

		for _, f := range _fish {
			bornFish := f.liveADay()

			if bornFish != nil && bornThisDay != nil {
				bornThisDay.x += bornFish.x
			} else if bornFish != nil {
				bornThisDay = bornFish
			}
		}

		if bornThisDay != nil {
			_fish = append(_fish, bornThisDay)
		}
	}

	return _fish
}

func (f *fish) liveADay() *fish {
	if f.birthAfter == 0 {
		f.birthAfter = birthTimer - 1
		return &fish{birthAfter: birthTimer + childShift - 1, x: f.x}
	}

	f.birthAfter--

	return nil
}

func createFishFromString(birthTimer int, input string) []*fish {
	result := make([]*fish, birthTimer)

	nums := strings.Split(input, ",")

	for _, num := range nums {
		n, _ := strconv.Atoi(num)

		if result[n] == nil {
			result[n] = &fish{birthAfter: n, x: 1}
		} else {
			result[n].x++
		}
	}

	for i := len(result) - 1; i > -1; i-- {
		if result[i] == nil {
			result = append(result[:i], result[i+1:]...)
		}
	}

	return result
}
