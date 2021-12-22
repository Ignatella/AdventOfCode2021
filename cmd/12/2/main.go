package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"regexp"
	"strings"
)

type arr []string
type str string

func main() {
	input := helpers.ReadInputFile("cmd/12/2/input.txt")

	vertexes, adjMatrix := createAdjMatrix(input)

	paths := findAllPaths(vertexes, adjMatrix)

	fmt.Println(paths)
}

func findAllPaths(vertexes []string, adjMatrix [][]uint8) int {
	startIndex := arr(vertexes).indexOf("start")

	return walk(startIndex, "", vertexes, adjMatrix)
}

func isAnySmallVertexVisitedTwice(currentPath []string) bool {
	for _, v := range currentPath {
		occ := arr(currentPath).countOccurrence(v)

		if str(v).match("[a-z].*") && occ == 2 {
			return true
		}
	}

	return false
}

func walk(startIndex int, currentPath string, vertexes []string, adjMatrix [][]uint8) int {
	path := strings.Split(currentPath, "-")
	currentPath += "-" + vertexes[startIndex]

	if vertexes[startIndex] == "end" {
		fmt.Println(currentPath[1:])
		return 1
	}

	isNotAllowedVisitingCount := arr(path).contains(vertexes[startIndex])
	isSmallCave := str(vertexes[startIndex]).match("[a-z].*")
	oneSmallCaveIsAlreadyVisitedTwice := isAnySmallVertexVisitedTwice(path)
	notTheFirstStart := vertexes[startIndex] == "start" && arr(path).contains("start")

	if (isNotAllowedVisitingCount && isSmallCave && oneSmallCaveIsAlreadyVisitedTwice) || notTheFirstStart {
		return 0
	}

	pathsFound := 0

	for j, el := range adjMatrix[startIndex] {
		if el == 1 {
			pathsFound += walk(j, currentPath, vertexes, adjMatrix)
		}
	}

	return pathsFound
}

func createAdjMatrix(input []string) ([]string, [][]uint8) {
	in := arr(input)
	vertexes := findUniqueVertexes(input)
	adjMatrix := make([][]uint8, len(vertexes))

	for i := range vertexes {
		adjMatrix[i] = append(adjMatrix[i], make([]uint8, len(vertexes))...)
	}

	for i, v := range vertexes {
		for j := i + 1; j < len(vertexes); j++ {
			if in.contains(vertexes[j]+"-"+v) || in.contains(v+"-"+vertexes[j]) {
				adjMatrix[i][j] = 1
				adjMatrix[j][i] = 1
			}
		}
	}

	return vertexes, adjMatrix
}

func findUniqueVertexes(input []string) []string {
	vertexes := make([]string, 0)

	for _, row := range input {
		points := strings.Split(row, "-")

		if !arr(vertexes).contains(points[0]) {
			vertexes = append(vertexes, points[0])
		}

		if !arr(vertexes).contains(points[1]) {
			vertexes = append(vertexes, points[1])
		}
	}

	return vertexes
}

func (s str) match(pattern string) bool {
	m, _ := regexp.MatchString(pattern, string(s))

	return m
}

func (a arr) indexOf(el string) int {
	for i, e := range a {
		if e == el {
			return i
		}
	}

	return -1
}

func (a arr) countOccurrence(element string) int {
	occurrences := 0

	for _, el := range a {
		if el == element {
			occurrences++
		}
	}

	return occurrences
}

func (a arr) contains(element string) bool {

	for _, el := range a {
		if el == element {
			return true
		}
	}

	return false
}
