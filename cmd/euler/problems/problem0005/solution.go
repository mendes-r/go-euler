package problem0005

func Solution() (int, string, float64) {
	description := "What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?"
	limit := 20

	flag := true
	count := limit
	result := -1

	for flag {
		multiple := true
		for i := limit; i > int(limit/2) && multiple; i-- {
			multiple = isMultiple(count, i)
		}
		if multiple {
			result = count
			flag = false
		}
		count++
	}

	return 5, description, float64(result)
}

func isMultiple(number int, divider int) bool {
	return number%divider == 0
}
