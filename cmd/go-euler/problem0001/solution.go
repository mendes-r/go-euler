package problem0001

import (
	"strconv"

	"github.com/mendes-r/go-euler/pkg/utils/printer"
)

func Solution() {
	result := 3

	for count := 4; count < 1000; count++ {
		if count%3 == 0 {
			result += count
		} else if count%5 == 0 {
			result += count
		}
	}

	printer.PrintToCli(1, strconv.Itoa(result))
}
