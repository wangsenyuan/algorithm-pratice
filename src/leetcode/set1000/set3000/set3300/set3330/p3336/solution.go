package p3336

import "slices"

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func subsequencePairCount(nums []int) int {
	x := slices.Max(nums)
	n := len(nums)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, x+1)
		for j := 0; j <= x; j++ {
			dp[i][j] = make([]int, x+1)
			for k := 0; k <= x; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var f func(i int, j int, k int) (ans int)

	f = func(i int, j int, k int) (ans int) {
		if i == 0 {
			if j == k && j > 0 {
				return 1
			}
			return 0
		}
		if dp[i][j][k] >= 0 {
			return dp[i][j][k]
		}

		defer func() {
			dp[i][j][k] = ans
		}()

		ans = f(i-1, j, k)
		ans = add(ans, f(i-1, gcd(nums[i-1], j), k))
		ans = add(ans, f(i-1, j, gcd(nums[i-1], k)))
		return ans
	}

	return f(n, 0, 0)
}

func subsequencePairCount1(nums []int) int {
	x := slices.Max(nums)
	// n := len(nums)

	// 理解错了
	dp := make([][]int, x+1)
	fp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, x+1)
		fp[i] = make([]int, x+1)
	}
	dp[0][0] = 1

	for _, num := range nums {
		for a := 0; a <= x; a++ {
			for b := 0; b <= x; b++ {
				fp[gcd(a, num)][b] = add(fp[gcd(a, num)][b], dp[a][b])
				fp[a][gcd(b, num)] = add(fp[a][gcd(b, num)], dp[a][b])
				fp[a][b] = add(fp[a][b], dp[a][b])
			}
		}
		for a := 0; a <= x; a++ {
			for b := 0; b <= x; b++ {
				dp[a][b] = fp[a][b]
				fp[a][b] = 0
			}
		}
	}

	var res int

	for i := 1; i <= x; i++ {
		res = add(res, dp[i][i])
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
