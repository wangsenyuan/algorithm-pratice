package p1362

import "math"

func closestDivisors(num int) []int {
	a := find(num + 1)
	b := find(num + 2)
	if a[1]-a[0] <= b[1]-b[0] {
		return a
	}
	return b
}

func find(num int) []int {
	// find a * b == num
	x := int(math.Sqrt(float64(num)))

	for x > 1 && num%x != 0 {
		x--
	}
	y := num / x
	return []int{x, y}
}
