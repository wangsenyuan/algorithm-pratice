package p1951

import "math"

func isThree(n int) bool {
	if n == 1 {
		return false
	}
	x := int(math.Sqrt(float64(n)))
	if x*x != n {
		return false
	}
	for y := 2; y < x; y++ {
		if n%y == 0 {
			return false
		}
	}
	return true
}
