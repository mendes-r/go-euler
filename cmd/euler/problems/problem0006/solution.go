package problem0006

import (
	"math"
)

func Solution() (int, string, float64) {
	description := "Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum."
	limit := 100

	result := math.Abs(sumOfSquares(float64(limit)) - squareOfSums(float64(limit)))

	return 6, description, result
}

func sumOfSquares(input float64) float64 {
	var result float64 = 0

	for i := float64(1); i <= input; i++ {
		result = result + math.Pow(i, 2)
	}

	return result
}

func squareOfSums(input float64) float64 {
	var result float64 = 0

	for i := float64(1); i <= input; i++ {
		result = result + i
	}

	return math.Pow(result, 2)
}
