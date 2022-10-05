package problem0004

import (
	"math"
)

func Solution() (int, string, float64) {
	description := "Find the largest palindrome made from the product of two 3-digit numbers."
	limit := 999

	result := 0

	for i := limit; i > 0; i-- {
		for j := limit; j > 0; j-- {
			temp := i * j
			if isPalindrome(temp) {
				if temp > result {
					result = temp
				}
			}
		}
		limit--
	}

	return 4, description, float64(result)
}

func isPalindrome(number int) bool {
	length := numbersLength(number)
	flag := false
	position := 1

	if length > 1 {
		flag = true

		for flag && position <= int(length/2) {
			rightDigit := getRightDigit(number, position)
			leftDigit := getLeftDigit(number, length-position)
			flag = (rightDigit == leftDigit)

			position++
		}
	}

	return flag
}

func numbersLength(number int) int {
	return int(math.Log10(float64(number)) + 1)
}

func getLeftDigit(number int, position int) int {
	power := math.Pow10(position)
	digit := number / int(power)
	return digit % 10
}

func getRightDigit(number int, position int) int {
	power := math.Pow10(position)
	digit := number % int(power)
	return digit / (int(power / 10))
}
