package problem0001

func Solution() (int, string, float64) {
	description := "Find the sum of all the multiples of 3 or 5 below 1000."
	limit := 1000

	result := 3

	for count := 4; count < limit; count++ {
		if count%3 == 0 {
			result += count
		} else if count%5 == 0 {
			result += count
		}
	}

	return 1, description, float64(result)
}
