package p2081

func kMirror(k int, n int) int64 {
	// abcdcba k <= 9, and n <= 30
	K := int64(k)
	var res int64
	var cnt int

	bruteForce := func(d int) {
		// 10000
		start := power(10, d-1)
		end := start * 10
		B := start
		for x := start; cnt < n && x < end; x++ {
			num := x*B + reverse(x/10)
			if isKMirror(num, K) {
				res += num
				cnt++
			}
		}

		B *= 10
		for x := start; cnt < n && x < end; x++ {
			num := x*B + reverse(x)
			if isKMirror(num, K) {
				res += num
				cnt++
			}
		}
	}

	for d := 1; cnt < n; d++ {
		bruteForce(d)
	}
	return res
}

func power(a, b int) int64 {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = R * A
		}
		A = A * A
		b >>= 1
	}
	return R
}

func reverse(num int64) int64 {
	var res int64
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}

func isKMirror(num int64, k int64) bool {
	mun := num
	var tmp int64
	for num > 0 {
		r := num % k
		tmp = tmp*k + r
		num /= k
	}

	return mun == tmp
}
