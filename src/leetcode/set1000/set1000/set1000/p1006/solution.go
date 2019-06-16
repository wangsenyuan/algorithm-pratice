package p1006

func clumsy(N int) int {

	n := N
	var res int
	sign := 1
	//three numbers in one group
	for n > 0 {
		if n > 3 {
			res += sign*n*(n-1)/(n-2) + (n - 3)
		} else if n == 1 {
			res += sign
		} else if n == 2 {
			res += sign * 2 / 1
		} else {
			// n == 3
			res += sign * 3 * 2 / 1
		}
		sign = -1
		n -= 4
	}
	return res
}
