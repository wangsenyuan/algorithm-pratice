package p1808

const MOD = 1000000007

// primeFactors <= 1e9
// some number num, its prime divisor, count <= primeFactors
// and good factor is all (distinct) prime divisor's product
// and max the count
func maxNiceDivisors(primeFactors int) int {
	// pow(x, a) * pow(x-1, b)
	// a * x + (x - 1) * b = tot
	// (a + b) * x - b = tot
	// (a + b) * x = tot + b
	if primeFactors == 1 {
		// self
		return 1
	}
	// 2 个 3 优于 3 个2
	a := primeFactors / 3
	b := primeFactors % 3
	res := int64(pow(3, a))

	if b == 0 {
		return int(res)
	}

	if b == 2 {
		// 8 = 2 * 3 + 2
		// a = 2
		// b = 2
		res *= 2
		res %= MOD
		return int(res)
	}

	// 0 is good, 2 is good too, only 1 need to borrow
	a--
	b = 4
	res = int64(pow(3, a))
	res = (res * 4) % MOD
	return int(res)
}

func pow(a, b int) int {
	A := int64(a)
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(res)
}
