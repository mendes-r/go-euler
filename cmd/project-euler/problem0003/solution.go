package problem0003

import (
	"math"

	"github.com/mendes-r/go-euler/pkg/utils/printer"
)

var primes []int = []int{2}

func Solution() {
	result := 0
	limit := 600851475143
	count := int(math.Sqrt(float64(limit)))

	for number := 3; number < count; number++ {
		if isPrime(number) && isFactor(limit, number) {
			result = number
		}
	}

	printer.PrintToCli(3, result)
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
