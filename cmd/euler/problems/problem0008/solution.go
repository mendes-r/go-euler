package problem0008

import (
	"strconv"

	"github.com/mendes-r/go-euler/pkg/utils/reader"
)

func Solution() (int, string, float64) {
	description := "Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?"

	input := reader.File2Slice("./cmd/euler/problems/problem0008/input")
	size := len(input)

	result := 0

	for i := 0; i < size; i++ {

		temp := getProduct(i, 13, &input)
		if temp > result {
			result = temp
		}

	}

	return 8, description, float64(result)
}

func getProduct(start int, frame int, slice *[]string) int {
	last := start + frame
	numbers := *slice
	size := len(numbers)

	result := 1

	for i := start; i < last && last < size; i++ {
		number, err := strconv.Atoi(numbers[i])
		check(err)
		result = result * number
	}

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
