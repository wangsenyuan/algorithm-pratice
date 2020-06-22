package p1486

func xorOperation(n int, start int) int {
	var res int

	for i := 0; i < n; i++ {
		res ^= (start + 2*i)
	}

	return res
}
