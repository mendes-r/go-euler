package main

import (
	"time"

	"github.com/mendes-r/go-euler/pkg/utils/printer"

	"github.com/mendes-r/go-euler/cmd/project-euler/problem0001"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0002"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0003"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0004"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0005"
	"github.com/mendes-r/go-euler/cmd/project-euler/problem0006"
)

type problem interface {
	printSolution()
}

type solution func(input float64) (number int, result float64)

type problem1input struct {
	input    float64
	solution solution
	run      bool
}

func (p problem1input) printSolution() {
	if p.run {
		input := p.input
		start := time.Now()
		number, result := p.solution(input)
		duration := time.Since(start)
		printer.PrintToCli(number, input, result, duration.String())
	}
}

func main() {
	p0001 := problem1input{1000, problem0001.Solution, false}
	p0002 := problem1input{4000000, problem0002.Solution, false}
	p0003 := problem1input{600851475143, problem0003.Solution, false}
	p0004 := problem1input{999, problem0004.Solution, false}
	p0005 := problem1input{20, problem0005.Solution, false}
	p0006 := problem1input{100, problem0006.Solution, true}

	problems := []problem{p0001, p0002, p0003, p0004, p0005, p0006}

	for _, problem := range problems {
		problem.printSolution()
	}
}
