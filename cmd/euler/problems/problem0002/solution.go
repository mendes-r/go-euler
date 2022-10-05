package problem0002

func Solution() (int, string, float64) {
	description := "By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms."
	limit := 4000000

	result := 0
	previous := 1
	fibonacci := 2

	for fibonacci < limit {
		if fibonacci%2 == 0 {
			result += fibonacci
		}
		temp := fibonacci
		fibonacci = fibonacci + previous
		previous = temp
	}

	return 2, description, float64(result)
}
