package p1759

const MOD = 1000000007

func countHomogenous(s string) int {
	var res int
	var j int
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[j] {
			l := i - j
			// l + l - 1 + l - 2 + ... + 1
			// (l + 1) * l / 2
			tmp := int64(l+1) * int64(l) / 2
			tmp = tmp % MOD
			modAdd(&res, int(tmp))
			j = i
		}
	}

	return res
}

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func pow(a, b int) int {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}
