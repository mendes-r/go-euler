package problem0005

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func Solution(input float64) (int, float64) {
	flag := true
	count := int(input)
	result := -1

	for flag {
		multiple := true
		for i := int(input); i > int(input/2) && multiple; i-- {
			multiple = isMultiple(count, i)
		}
		if multiple {
			result = count
			flag = false
		}
		count++
	}

	return 5, float64(result)
}

func isMultiple(number int, divider int) bool {
	return number%divider == 0
}
