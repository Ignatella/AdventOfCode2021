package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/cmd/15/internal/graph"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"math"
	"os"
	"strconv"
)

const file string = "cmd/15/2/input.txt"
const scaleFactor int = 5

func main() {
	scaleUp(scaleFactor)
	input := helpers.ReadInputFile(file)
	g := parseInput(input)
	risks := dijkstra(g[0][0], g)
	scaleDown(scaleFactor)

	fmt.Printf("%v\n", risks[len(risks)-1][len(risks[0])-1])
}

func dijkstra(start graph.Node, g graph.Graph) [][]int {
	risks := make([][]int, len(g))

	for i := range risks {
		risks[i] = make([]int, len(g[i]))
		for j := range risks[i] {
			risks[i][j] = math.MaxInt
			g[i][j].Previous = nil
		}
	}

	risks[start.I][start.J] = 0

	nSet := g.ToNodeSet()

	for !nSet.Empty() {
		u := nSet.GetNodeWithMinRisk(risks)

		nSet.Remove(u)

		if u == g[len(g)-1][len(g[0])-1] {
			break
		}

		for _, v := range g.GetNeighbors(u).Intersect(nSet) {
			alt := risks[u.I][u.J] + v.Risk

			if alt > -1 && alt < risks[v.I][v.J] {
				risks[v.I][v.J] = alt
				v.Previous = &u
			}
		}
	}

	return risks
}

func parseInput(input []string) graph.Graph {
	g := make(graph.Graph, len(input))

	for i, r := range input {
		g[i] = make([]graph.Node, len(r))

		for j, c := range r {
			risk, _ := strconv.Atoi(string(c))
			g[i][j] = graph.NewNode(i, j, risk)
		}
	}

	return g
}

// region Scale

func scaleDown(factor int) {
	input := helpers.ReadInputFile(file)

	initL := len(input) / factor

	input = input[0:initL]

	for i, str := range input {
		input[i] = str[0:initL]
	}

	writeStrings(input, file)
}

func scaleUp(factor int) {
	input := helpers.ReadInputFile(file)

	l := len(input)

	for f := 1; f < factor; f++ {
		for i := f * l; i < f*l+l; i++ {
			input = append(input, "")

			for j := 0; j < l; j++ {
				r, _ := strconv.Atoi(string(input[i-l][j]))
				r++

				if r > 9 {
					r = 1
				}

				input[i] += strconv.Itoa(r)
			}
		}
	}

	for f := 1; f < factor; f++ {
		for j := f * l; j < f*l+l; j++ {

			for i := 0; i < len(input); i++ {
				r, _ := strconv.Atoi(string(input[i][j-l]))
				r++

				if r > 9 {
					r = 1
				}

				input[i] += strconv.Itoa(r)
			}
		}
	}

	writeStrings(input, file)
}

func writeStrings(src []string, fileName string) {
	_ = os.Truncate(file, 0)
	f, err := os.OpenFile(file, os.O_WRONLY, 0644)

	for i, line := range src {
		if i == len(src)-1 {
			_, _ = f.WriteString(line)
		} else {
			_, _ = f.WriteString(line + "\r\n")
		}
	}

	if err != nil {
		panic(err)
	}
}

// endregion
