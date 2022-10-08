package problem0010

var primes []int = []int{2}

func Solution() (int, string, float64) {
	description := "Find the sum of all the primes below two million."
	limit := 2000000

	result := 2

	for i := 3; i <= limit; i++ {
		if isPrime(i) {
			result = result + i
		}
	}

	return 10, description, float64(result)
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
