package p2320

const MOD = 1000000007

func countHousePlacements(n int) int {
	a, b := 1, 1

	for i := 1; i < n; i++ {
		// if place at i
		a, b = b, modAdd(a, b)
	}
	x := modAdd(a, b)
	return modMul(x, x)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}
