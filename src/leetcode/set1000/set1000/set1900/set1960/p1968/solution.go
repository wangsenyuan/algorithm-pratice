package p1968

const MOD = 1000000007

func minNonZeroProduct(p int) int {
	if p == 1 {
		return 1
	}
	x := (uint64(1<<p) - 1) % MOD
	y := (uint64(1<<p) - 2) % MOD
	z := uint64(1<<uint64(p-1) - 1)
	return int((x * pow(y, z)) % MOD)
}

func pow(a, b uint64) uint64 {
	var r uint64 = 1

	for b != 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}
