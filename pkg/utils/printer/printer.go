package printer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func PrintToCli(number int, input float64, result float64, duration string) {
	fmt.Println("> Problem: ", number)
	i := strconv.FormatFloat(input, 'f', -1, 64)
	fmt.Println(">    input: ", i)
	s := strconv.FormatFloat(result, 'f', -1, 64)
	fmt.Println(">    solution: ", s)
	fmt.Println(">    duration: ", duration)
	fmt.Println(".............")
	fmt.Println()
}

func PrintToFile(number int, input float64, result float64, duration string) {
	f, err := os.OpenFile("results/results.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)

	var sb strings.Builder
	defer f.Close()

	sb.Write([]byte(time.Now().String() + ", "))
	sb.Write([]byte(strconv.Itoa(number) + ", "))
	sb.Write([]byte(strconv.FormatFloat(input, 'f', -1, 64) + ", "))
	sb.Write([]byte(strconv.FormatFloat(result, 'f', -1, 64) + ", "))
	sb.Write([]byte(duration + "\n"))

	f.WriteString(sb.String())

	f.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
