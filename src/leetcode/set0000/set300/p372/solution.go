package p372

//MOD = 1337
const MOD = 1337

func superPow(a int, b []int) int {
	c := 1

	for _, x := range b {
		c = (pow(c, 10) * pow(a, x)) % MOD
	}

	return c
}

func pow(a, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return a % MOD
	}

	return (pow(a, n/2) * pow(a, n-n/2)) % MOD
}
