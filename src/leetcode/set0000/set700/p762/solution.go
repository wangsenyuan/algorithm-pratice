package p762

func countPrimeSetBits(L int, R int) int {
	primes := make([]bool, 32)
	primes[2] = true
	primes[3] = true
	primes[5] = true
	primes[7] = true
	primes[11] = true
	primes[13] = true
	primes[17] = true
	primes[19] = true
	primes[23] = true
	primes[29] = true
	primes[31] = true

	var ones int

	for x := L; x > 0; x = x >> 1 {
		if x&1 == 1 {
			ones++
		}
	}

	var res int
	if primes[ones] {
		res = 1
	}
	for x := L; x < R; x++ {
		y := x
		for y&1 == 1 {
			ones--
			y >>= 1
		}

		ones++

		if primes[ones] {
			res++
		}
	}
	return res
}
