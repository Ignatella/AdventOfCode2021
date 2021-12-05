package line

import (
	"github.com/Ignatella/AdventOfCode2021/cmd/5/internal/point"
	"strings"
)

type Line struct {
	p1, p2 *point.Point
}

func New(src string) *Line {
	points := strings.Split(src, " -> ")
	p1 := point.FromString(points[0])
	p2 := point.FromString(points[1])

	return &Line{p1: p1, p2: p2}
}

func (l *Line) IsHorizontalOrVertical() (result bool) {
	result = false

	result = result || (l.p1.X == l.p2.X)
	result = result || (l.p1.Y == l.p2.Y)

	return
}

func (l *Line) GetPoints() (result []*point.Point) {

	/**
	y = kx + b
	from the task: slope can be only:
		0 + pi/2 * k, k = 0, 1, 2...
		pi/4 + pi/2 * k, k = 0, 1, 2...

	P.S. Assuming there is no lines built from the 2 identical points eg: (1, 1) (1, 1)
	*/

	dy := l.p2.Y - l.p1.Y
	dx := l.p2.X - l.p1.X

	// region Vertical line
	if dx == 0 {
		for y := l.p1.Y; y <= l.p2.Y; y++ {
			result = append(result, point.New(l.p1.X, y))
		}

		for y := l.p2.Y; y <= l.p1.Y; y++ {
			result = append(result, point.New(l.p1.X, y))
		}

		return
	}
	// endregion

	k := dy / dx
	b := l.p1.Y - k*l.p1.X

	// region All other lines
	if l.p2.X > l.p1.X {
		for x := l.p1.X; x <= l.p2.X; x++ {
			py := k*x + b
			result = append(result, point.New(x, py))
		}
	}

	if l.p2.X < l.p1.X {
		for x := l.p2.X; x <= l.p1.X; x++ {
			py := k*x + b
			result = append(result, point.New(x, py))
		}
	}
	//endregion

	return
}
