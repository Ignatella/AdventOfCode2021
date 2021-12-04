package main

import (
	"fmt"
	"github.com/Ignatella/AdventOfCode2021/internal/helpers"
	"strconv"
	"strings"
)

type position struct {
	x, z int
}

func main() {
	input := helpers.ReadInputFile("cmd/2/1/input.txt")
	p := position{}
	for _, cmd := range input {
		p.updatePosition(cmd)
	}

	fmt.Printf("%v\n", p)
	fmt.Printf("Multiplication: %v\n", p.x*p.z)
}

func (p *position) updatePosition(cmd string) {
	cmdParts := strings.Split(cmd, " ")

	moveDirection := cmdParts[0]
	moveSize, _ := strconv.Atoi(cmdParts[1])

	switch moveDirection {
	case "forward":
		p.x += moveSize
	case "up":
		p.z -= moveSize
	case "down":
		p.z += moveSize
	}
}
