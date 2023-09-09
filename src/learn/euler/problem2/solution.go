package problem2

func EvenFibNumSum(N int) int64 {
	a, b := 1, 2
	var sum int64 = 2
	for b < N {
		c := a + b
		if c%2 == 0 {
			sum += int64(c)
		}
		a, b = b, c
	}

	return sum
}
