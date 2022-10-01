package problem0003

import (
	"math"
)

var primes []int = []int{2}

func Solution(input float64) (int, float64) {
	result := 0
	limit := int(input) // 600851475143
	count := int(math.Sqrt(float64(limit)))

	for number := 3; number < count; number++ {
		if isPrime(number) && isFactor(limit, number) {
			result = number
		}
	}

	return 3, float64(result)
}

func isPrime(number int) bool {

	for _, prime := range primes {
		if number%prime == 0 {
			return false
		}
	}

	primes = append(primes, number)
	return true
}

func isFactor(numerator int, denominator int) bool {
	return numerator%denominator == 0
}
