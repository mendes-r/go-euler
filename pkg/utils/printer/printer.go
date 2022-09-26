package printer

import (
	"fmt"
	"strconv"
)

func PrintToCli(number int, result int) {
	fmt.Println("> Problem: ", number)
	fmt.Println(strconv.Itoa(result))
	fmt.Println()
}
