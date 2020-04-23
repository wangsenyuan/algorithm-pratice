package p0002

func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}

	if m+1 == n {
		return m & n
	}

	var mask = 1
	var res int

	for i := 1; i < 32; i++ {
		if m&(1<<i) > 0 && n&(1<<i) > 0 {
			x := m & mask
			y := mask - x
			if m+y >= n {
				res |= 1 << i
			}
		}
		mask |= 1 << (i)
	}

	return res
}
