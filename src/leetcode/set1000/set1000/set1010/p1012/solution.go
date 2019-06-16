package p1012

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	var res int

	var i uint

	for N != 0 {
		if N&1 == 0 {
			res |= (1 << i)
		}
		i++
		N >>= 1
	}

	return res
}
