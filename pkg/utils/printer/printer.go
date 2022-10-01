package printer

import (
	"fmt"
	"strconv"
)

func PrintToCli(number int, input float64, result float64) {
	fmt.Println("> Problem: ", number)
	i := strconv.FormatFloat(input, 'f', -1, 64)
	fmt.Println(">    input: ", i)
	s := strconv.FormatFloat(result, 'f', -1, 64)
	fmt.Println(">    solution: ", s)
	fmt.Println(".............")
	fmt.Println()
}
