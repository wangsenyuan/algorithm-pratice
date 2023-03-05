package p2578

func coloredCells(n int) int64 {
	// f(n)
	// f(n) = f(n - 1) + 4 * n - 4
	// 1 + 4 + 8 + 12 + ...
	// f(n) = 4 * (1 + 2 + ... n) - 4 * n
	//      = 4 * (1 + n) * n / 2 - 4 * n
	//
	N := int64(n)
	res := 4*(1+N-1)*(N-1)/2 + 1
	return res
}
