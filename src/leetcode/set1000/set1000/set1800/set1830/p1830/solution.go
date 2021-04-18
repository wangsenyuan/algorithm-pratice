package p1830

const MOD = 1000000007
const N = 3003

var F [N]int64
var I [N]int64

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}

	I[N-1] = pow(F[N-1], MOD-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}
}

func pow(a, b int64) int64 {
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return R
}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= I[r]
	res %= MOD
	return (res * I[n-r]) % MOD
}

func makeStringSorted(s string) int {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}

	count := func(x int, n int) int {
		var res int64 = 1
		for i := 0; i < 26; i++ {
			tmp := cnt[i]
			if i == x {
				tmp--
			}
			res *= nCr(n, tmp)
			n -= tmp
			res %= MOD
		}

		return int(res)
	}

	var res int
	n := len(s)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		for y := x - 1; y >= 0; y-- {
			if cnt[y] > 0 {
				res += count(y, n-1-i)
				res %= MOD
			}
		}
		cnt[x]--
	}
	return res
}
