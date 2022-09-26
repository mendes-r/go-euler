package problem0002

import (
	"github.com/mendes-r/go-euler/pkg/utils/printer"
)

func Solution() {
	result := 0
	previous := 1
	fibonacci := 2

	for fibonacci < 4000000 {
		if fibonacci%2 == 0 {
			result += fibonacci
		}
		temp := fibonacci
		fibonacci = fibonacci + previous
		previous = temp
	}

	printer.PrintToCli(2, result)
}
