package p2551

const MOD = 1e9 + 7

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func monkeyMove(n int) int {
	res := pow(2, n) - 2
	if res < 0 {
		res += MOD
	}
	return res
}
