package point

import (
	"strconv"
	"strings"
)

type Point struct {
	X, Y  int
	Cross int
}

func New(x, y int) *Point {
	return &Point{X: x, Y: y, Cross: 1}
}

func FromString(src string) *Point {
	coordinates := strings.Split(src, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])

	return &Point{X: x, Y: y, Cross: 1}
}
