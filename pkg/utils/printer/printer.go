package printer

import "fmt"

func PrintToCli(number int, message string) {
	fmt.Println("> Problem: ", number)
	fmt.Println(message)
	fmt.Println()
}
