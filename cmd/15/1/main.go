package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/cmd/15/internal/graph"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"math"
	"strconv"
)

func main() {
	input := helpers.ReadInputFile("cmd/15/1/input.txt")
	g := parseInput(input)
	risks := dijkstra(g[0][0], g)

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
