package problem0001

func Solution(input float64) (int, float64) {
	result := 3
	limit := int(input) // 1000

	for count := 4; count < limit; count++ {
		if count%3 == 0 {
			result += count
		} else if count%5 == 0 {
			result += count
		}
	}

	return 1, float64(result)
}
