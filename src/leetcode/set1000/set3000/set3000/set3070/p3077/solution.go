package p3077

import "math"

const oo = 1 << 60

func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	f := make([][]int, k+1)
	f[0] = make([]int, n+1)
	for i := 1; i <= k; i++ {
		f[i] = make([]int, n+1)
		f[i][i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		for j := i; j <= n-k+i; j++ { // j 的上下界是因为其它子数组至少要选一个数
			mx = max(mx, f[i-1][j-1]-s[j-1]*w)
			f[i][j] = max(f[i][j-1], s[j]*w+mx)
		}
	}
	return int64(f[k][n])
}

func maximumStrength1(nums []int, k int) int64 {
	dp := make([]int, k+1)
	fp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = -oo
		fp[i] = -oo
	}
	dp[0] = 0
	fp[0] = 0
	// fp[j] = 以上当前数结尾的，且是第j个subarra时的最优值

	// dp[i][j] = max(dp[i-1][j-1] + j * nums[i], dp[i-1][j] + j * nums[i])
	for i, num := range nums {
		for j := min(k, i+1); j > 0; j-- {
			sign := 1
			if j&1 == 0 {
				sign = -1
			}
			fp[j] = max(fp[j], dp[j-1]) + sign*num*(k-j+1)
			dp[j] = max(dp[j], fp[j])
		}
	}

	return int64(dp[k])
}
