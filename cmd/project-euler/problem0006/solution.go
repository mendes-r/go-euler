package problem0006

import "math"

// Find the difference between the sum of the squares of the first one hundred natural numbers
// and the square of the sum.

func Solution(input float64) (int, float64) {
	return 6, math.Abs(sumOfSquares(input) - squareOfSums(input))
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
