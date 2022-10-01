package problem0007

// What is the 10 001st prime number?

func Solution(input float64) (int, float64) {
	primes := []int{2}
	count := 1
	number := 3
	limit := int(input)

	for count < limit {
		if isPrime(number, primes) {
			primes = append(primes, number)
			count++
		}
		number++
	}

	return 7, float64(primes[len(primes)-1])
}

func isPrime(number int, primes []int) bool {
	for _, prime := range primes {
		if number%prime == 0 {
			return false
		}
	}
	return true
}
