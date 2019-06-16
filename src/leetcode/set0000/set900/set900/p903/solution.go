package p903

const MAX_N = 210

var C [][]int64

const MOD = 1e9 + 7

func init() {
	C = make([][]int64, MAX_N)
	for i := 0; i < MAX_N; i++ {
		C[i] = make([]int64, MAX_N)
	}

	C[0][0] = 1
	for i := 1; i < MAX_N; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
			if C[i][j] >= MOD {
				C[i][j] -= MOD
			}
		}
	}
}

func numPermsDISequence(S string) int {
	n := len(S)
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 0; i < n; i++ {
		if S[i] == 'I' {
			var cur int64
			for j := 0; j < n-i; j++ {
				cur += dp[i][j]
				if cur >= MOD {
					cur -= MOD
				}
				dp[i+1][j] = cur
			}
		} else {
			var cur int64
			for j := n - i - 1; j >= 0; j-- {
				cur += dp[i][j+1]
				if cur >= MOD {
					cur -= MOD
				}
				dp[i+1][j] = cur
			}
		}
	}

	return int(dp[n][0])
}

func numPermsDISequence2(S string) int {
	N := len(S)

	dp := make([][]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int64, N)
	}

	var loop func(i, j int) int64

	loop = func(i, j int) int64 {
		if i >= j {
			return 1
		}
		if dp[i][j] == 0 {
			n := j - i + 2
			var ans int64
			if S[i] == 'I' {
				ans += loop(i+1, j)
			}
			if S[j] == 'D' {
				ans += loop(i, j-1)
			}

			for k := i + 1; k <= j; k++ {
				if S[k-1:k+1] == "DI" {
					ans += (((C[n-1][k-i] * loop(i, k-2)) % MOD) * loop(k+1, j)) % MOD
					ans %= MOD
				}
			}
			dp[i][j] = ans
		}

		return dp[i][j]
	}
	return int(loop(0, N-1))
}

func numPermsDISequence1(S string) int {
	n := len(S)
	A := make([]int, n+1)
	var p int
	var l int
	for i := 0; i < n; i++ {
		l++
		if S[i] == 'I' {
			A[p] = l
			p++
			l = 0
		}
	}
	A[p] = l + 1
	p++

	dp := make([]int64, p+1)

	dp[0] = 1
	var all int
	for i := 1; i <= p; i++ {
		all += A[i-1]
		var s int
		for j, sign := i-1, int64(1); j >= 0; j, sign = j-1, -sign {
			s += A[j]
			dp[i] += dp[j] * C[all][s] * sign
			dp[i] %= MOD
		}
		if dp[i] < 0 {
			dp[i] += MOD
		}
	}

	return int(dp[p])
}
