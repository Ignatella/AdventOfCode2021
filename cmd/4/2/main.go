package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type field struct {
	value    int
	isMarked bool
}

type board struct {
	size   int
	fields [][]field
}

func main() {
	input, _ := os.ReadFile("cmd/4/2/input.txt")
	inputNums := make([]int, 0)
	rows := strings.Split(string(input), "\r\n\r\n")
	boards := make([]board, 0)

	// parse 'random' input values
	for _, num := range strings.Split(rows[0], ",") {
		n, _ := strconv.Atoi(num)
		inputNums = append(inputNums, n)
	}

	// create board for each row
	for i := 1; i < len(rows); i++ {
		boards = append(boards, createBoard(rows[i]))
	}

	winner, winnerNum, err := findLastWinner(inputNums, boards)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("Result: %v", winner.getScore(winnerNum))
}

func (b *board) getScore(winnerNum int) int {
	result := 0

	for _, row := range b.fields {
		for _, f := range row {
			if !f.isMarked {
				result += f.value
			}
		}
	}

	return result * winnerNum
}

func findLastWinner(inputNums []int, boards []board) (board, int, error) {
	winners := make([]bool, len(boards))
	winnersCount := 0

	for _, num := range inputNums {

		for i, b := range boards {
			if winners[i] {
				continue
			}

			b.markFields(num)

			if b.isWinner() {
				winners[i] = true
				winnersCount++

				if winnersCount == len(boards) {
					return b, num, nil
				}
			}
		}
	}

	return board{}, 0, errors.New("there is no winner")
}

func (b *board) isWinner() bool {
	for i := 0; i < len(b.fields); i++ {
		isRWinner := true
		isCRWinner := true

		for j := 0; j < len(b.fields); j++ {
			isRWinner = isRWinner && b.fields[i][j].isMarked
			isCRWinner = isCRWinner && b.fields[j][i].isMarked
		}

		if isRWinner || isCRWinner {
			return true
		}
	}

	return false
}

func (b *board) markFields(num int) {
	for i, row := range b.fields {
		for j, f := range row {
			if f.value == num {
				b.fields[i][j].isMarked = true
			}
		}
	}
}

func createBoard(seedSource string) board {
	rows := strings.Split(seedSource, "\r\n")
	boardSize := len(rows)

	fields := make([][]field, boardSize)

	for i, row := range rows {

		nums := strings.Split(row, " ")
		nums = removeEmptyStrings(nums)

		r := make([]field, boardSize)

		for j, n := range nums {
			v, _ := strconv.Atoi(n)

			r[j] = field{value: v}
		}

		fields[i] = r
	}

	return board{size: boardSize, fields: fields}
}

func removeEmptyStrings(arr []string) []string {
	arrLen := len(arr)

	for i := arrLen - 1; i > -1; i-- {
		if arr[i] == "" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	return arr
}
