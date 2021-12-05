package plane

import (
	"github.com/Ignatella/AdventOfCode2021/cmd/5/internal/line"
	"github.com/Ignatella/AdventOfCode2021/cmd/5/internal/point"
)

type Plane struct {
	points [][]*point.Point
}

func New() *Plane {
	return &Plane{}
}

func (p *Plane) TotalCrossNumber() (result int) {
	for _, row := range p.points {
		if row == nil {
			continue
		}

		for _, pt := range row {
			if pt == nil {
				continue
			}

			if pt.Cross > 1 {
				result++
			}
		}
	}

	return
}

func (p *Plane) PutLineOnPlane(l *line.Line) {
	points := l.GetPoints()
	if p.points == nil {
		p.points = make([][]*point.Point, 0)
	}

	for _, pt := range points {
		if len(p.points) < pt.Y+1 {
			p.points = append(p.points, make([][]*point.Point, pt.Y+1-len(p.points))...)
		}

		if p.points[pt.Y] == nil {
			p.points[pt.Y] = make([]*point.Point, 0)
		}

		if len(p.points[pt.Y]) < pt.X+1 {
			p.points[pt.Y] = append(p.points[pt.Y], make([]*point.Point, pt.X+1-len(p.points[pt.Y]))...)
		}

		pointOnPlane := p.points[pt.Y][pt.X]

		if pointOnPlane != nil {
			p.points[pt.Y][pt.X].Cross++
		}

		if pointOnPlane == nil {
			p.points[pt.Y][pt.X] = pt
		}
	}
}
