package p1922

func countGoodNumbers(n int64) int {
	// n <= 10^^15
	// 第一位去2,4,6,8

	if n == 1 {
		return 5
	}

	even := (n + 1) / 2
	odd := n - even

	res := pow(5, even) * pow(4, odd) % MOD

	return int(res)
}

const MOD = 1000000007

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}
