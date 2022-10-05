package problem0007

func Solution() (int, string, float64) {
	description := "What is the 10 001st prime number?"
	limit := 10001

	primes := []int{2}
	count := 1
	number := 3

	for count < limit {
		if isPrime(number, primes) {
			primes = append(primes, number)
			count++
		}
		number++
	}

	return 7, description, float64(primes[len(primes)-1])
}

func isPrime(number int, primes []int) bool {
	for _, prime := range primes {
		if number%prime == 0 {
			return false
		}
	}
	return true
}
