package p2845

func pivotInteger(n int) int {

	// x * (x + 1) / 2 = sum / 2
	// x * (x + 1) = sum
	for x := 1; x <= n; x++ {
		a := x * (1 + x) / 2
		b := (n - x + 1) * (x + n) / 2
		if a == b {
			return x
		}
	}
	return -1
}
