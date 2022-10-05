package problem0009

import (
	"math"
)

func Solution() (int, string, float64) {
	description := "There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find the product abc."
	limit := 1000

	flag := false
	a, b, c := -1, -1, -1

	for i := int(limit / 4); i >= 1 && !flag; i-- {

		a = i

		for j := int(limit/2) - 2; j > i && !flag; j-- {

			b = j
			c = limit - b - a

			flag = isPythagorean(a, b, c)
		}

	}

	return 9, description, float64(a * b * c)
}

func isPythagorean(a int, b int, c int) bool {
	return (math.Pow(float64(a), 2) + math.Pow(float64(b), 2)) == math.Pow(float64(c), 2)
}
