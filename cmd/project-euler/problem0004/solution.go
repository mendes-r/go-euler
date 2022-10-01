package problem0004

import (
	"fmt"
	"math"
)

// Find the largest palindrome made from the product of two 3-digit numbers.

func Solution(input float64) (int, float64) {
	result := 0
	limit := input // 999

	for i := limit; i > 0; i-- {
		for j := limit; j > 0; j-- {
			temp := i * j
			if isPalindrome(temp) {

				fmt.Println("Found ", temp, i, j)

				if temp > result {
					result = temp
				}
			}
		}
		limit--
	}

	return 4, float64(result)
}

func isPalindrome(number int) bool {
	length := int(math.Log10(float64(number)) + 1)
	flag := false

	if length > 1 {
		base := 10
		flag = true
		up := int(length / 2)
		down := 1

		for flag && up > down {
			first := "placeholder"
			last := "placeholder"
			flag = (first == last)

			fmt.Println()
			up--
			down++
		}
	}

	return flag
}
