package main

import (
	"github.com/mendes-r/go-euler/pkg/utils/printer"

	"github.com/mendes-r/go-euler/cmd/project-euler/problem0001"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0002"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0003"
)

type problem interface {
	printSolution()
}

type solution func(input float64) (number int, result float64)

type problem1input struct {
	input    float64
	solution solution
}

func (p problem1input) printSolution() {
	input := p.input
	number, result := p.solution(input)
	printer.PrintToCli(number, input, result)
}

func main() {
	p0001 := problem1input{1000, problem0001.Solution}
	p0002 := problem1input{4000000, problem0002.Solution}
	p0003 := problem1input{600851475143, problem0003.Solution}

	problems := []problem{p0001, p0002, p0003}

	for _, problem := range problems {
		problem.printSolution()
	}
}
