package p2478

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func beautifulPartitions(s string, k int, minLength int) int {
	n := len(s)

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		var tot int
		for j := minLength * i; j <= n; j++ {
			j0 := j - minLength + 1
			if isPrime(s[j0-1]) {
				tot = modAdd(tot, dp[i-1][j0-1])
			}
			if !isPrime(s[j-1]) {
				dp[i][j] = tot
			}
		}
	}

	return dp[k][n]
}

func beautifulPartitions2(s string, k int, minLength int) int {
	n := len(s)

	if !isPrime(s[0]) || isPrime(s[n-1]) {
		return 0
	}

	if k == 1 {
		return 1
	}

	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, n)
	}

	var r int
	for i := minLength; i <= n-minLength; i++ {
		if isPrime(s[i-1]) || !isPrime(s[i]) {
			continue
		}
		for j := 0; j < k-1; j++ {
			for x := r + 1; x <= i; x++ {
				dp[j][x] = dp[j][r]
			}
			if j > 0 {
				dp[j][i] = modAdd(dp[j][i], dp[j-1][i-minLength])
			} else {
				dp[j][i] = modAdd(dp[j][i], 1)
			}
		}
		r = i
	}

	return dp[k-2][r]
}

func beautifulPartitions1(s string, k int, minLength int) int {
	n := len(s)
	// n <= 1000
	// dp[i] = ans endint at i
	jump := make([][]int, n)
	for i := 0; i < n; i++ {
		jump[i] = make([]int, 0, 1)
		if isPrime(s[i]) {
			for j := i + minLength - 1; j+1 <= n; j++ {
				if !isPrime(s[j]) && (j+1 == n || isPrime(s[j+1])) {
					jump[i] = append(jump[i], j+1)
				}
			}
		}

	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	dp[n][0] = 1

	for i := n - 1; i >= 0; i-- {
		for _, ii := range jump[i] {
			for j := 1; j <= k; j++ {
				dp[i][j] = modAdd(dp[ii][j-1], dp[i][j])
			}
		}
	}

	return dp[0][k]
}

func isPrime(x byte) bool {
	return x == '2' || x == '3' || x == '5' || x == '7'
}
