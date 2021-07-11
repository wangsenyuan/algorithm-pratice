package p1928

import "math"

func countTriples(n int) int {
	var res int
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			x := a*a + b*b
			r := int(math.Sqrt(float64(x)))
			if r*r == x && r <= n {
				res += 2
			}
		}
		x := 2 * a * a
		r := int(math.Sqrt(float64(x)))
		if r*r == x && r <= n {
			res++
		}
	}

	return res
}
