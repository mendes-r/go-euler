package problem0002

func Solution(input float64) (int, float64) {
	result := 0
	previous := 1
	fibonacci := 2
	limit := int(input) // 4000000

	for fibonacci < limit {
		if fibonacci%2 == 0 {
			result += fibonacci
		}
		temp := fibonacci
		fibonacci = fibonacci + previous
		previous = temp
	}

	return 2, float64(result)
}
