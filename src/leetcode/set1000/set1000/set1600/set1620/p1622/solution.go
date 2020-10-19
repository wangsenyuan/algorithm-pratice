package p1622

const MOD = 1000000007

const N = 1001

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
		IF[i] = int64(i+1) * IF[i+1]
		IF[i] %= MOD
	}
}

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res *= IF[r]
	res %= MOD
	res *= IF[n-r]
	return res % MOD
}

func numberOfSets(n int, k int) int {
	var res int64

	for i := k + 1; i <= 2*k; i++ {
		tmp := nCr(n, i)
		if i > 3 {
			tmp *= nCr(k-1, i-k-1)
		}
		res += tmp % MOD
		res %= MOD
	}
	return int(res)
}

func numberOfSets2(n int, k int) int {
	dp := make([][]int, n+1)
	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		sum[i] = make([]int, k+1)
		if i > 0 {
			dp[i][0] = 1
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			sum[i][j-1] = (sum[i-1][j-1] + dp[i-1][j-1]) % MOD
			dp[i][j] += (sum[i][j-1] + dp[i-1][j]) % MOD
			dp[i][j] %= MOD
		}
	}

	return dp[n][k]
}

func numberOfSets1(n int, k int) int {
	// nCr(n, k)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][0] = 1

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % MOD

			dp[i][j][1] = dp[i-1][j][1]

			if j > 0 {
				dp[i][j][1] += dp[i-1][j-1][0]
				dp[i][j][1] %= MOD
				dp[i][j][1] += dp[i-1][j-1][1]
				dp[i][j][1] %= MOD
			}
		}
	}

	res := (dp[n-1][k][0] + dp[n-1][k][1]) % MOD
	return res
}
