package d

const MOD = 1000000007

const N = 26*5 + 1

var F [N]int64
var IF [N]int64

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = int64(i) * F[i-1]
		F[i] %= MOD
	}
	IF[N-1] = pow(F[N-1], MOD-2)
	for i := N - 2; i >= 0; i-- {
		IF[i] = (int64(i+1) * IF[i+1]) % MOD
	}
}

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
			r %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}
	return r
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF[n-r]
	res %= MOD
	res *= IF[r]
	res %= MOD
	return int(res)
}

func keyboard(k int, n int) int {
	dp := make([][]int, 27)
	for i := 0; i <= 26; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][n] = 1

	for i := 1; i <= 26; i++ {
		for j := 0; j <= n; j++ {
			for p := 0; p <= k; p++ {
				if j < p {
					continue
				}
				tmp := int64(dp[i-1][j]) * int64(nCr(j, p))
				tmp %= MOD
				dp[i][j-p] += int(tmp)
				dp[i][j-p] %= MOD
			}
		}
	}

	return dp[26][0]
}
