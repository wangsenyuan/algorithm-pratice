package p790

const MOD = 1e9 + 7

func numTilings(N int) int {
	if N == 1 {
		return 1
	}
	if N == 2 {
		return 2
	}

	if N == 3 {
		return 5
	}
	a, b, c := 1, 2, 5

	for i := 4; i <= N; i++ {
		d := ((2*c)%MOD + a) % MOD
		a, b, c = b, c, d
	}

	return c
}
