package main

import "math"

func judgeSquareSum(c int) bool {
	x := sqrt(c)

	for a := 0; a <= x; a++ {
		b := sqrt(c - a*a)
		if a*a+b*b == c {
			return true
		}
	}

	return false
}
func sqrt(c int) int {
	return int(math.Sqrt(float64(c)))
}
