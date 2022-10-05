package problems

import (
	"time"

	"github.com/mendes-r/go-euler/pkg/utils/printer"

	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0001"
	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0002"
	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0003"
	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0004"
	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0005"
	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0006"
	"github.com/mendes-r/go-euler/cmd/euler/problems/problem0007"
)

type problem interface {
	runSolution()
}

type solution func() (number int, description string, result float64)

type aProblem struct {
	solution solution
}

func (p aProblem) runSolution() {
	start := time.Now()
	number, description, result := p.solution()
	duration := time.Since(start)
	printer.PrintToCli(number, description, result, duration.String())
	printer.PrintToFile(number, description, result, duration.String())
}

func Run(first int, last int) {

	problems := getProblems()
	size := len(problems)

	for i := first; i <= last && i <= size; i++ {
		problems[i-1].runSolution()
	}

}

func getProblems() []problem {

	p0001 := aProblem{problem0001.Solution}
	p0002 := aProblem{problem0002.Solution}
	p0003 := aProblem{problem0003.Solution}
	p0004 := aProblem{problem0004.Solution}
	p0005 := aProblem{problem0005.Solution}
	p0006 := aProblem{problem0006.Solution}
	p0007 := aProblem{problem0007.Solution}
	// p0008 := aProblem{problem0008.Solution}

	return []problem{p0001, p0002, p0003, p0004, p0005, p0006, p0007}
}
