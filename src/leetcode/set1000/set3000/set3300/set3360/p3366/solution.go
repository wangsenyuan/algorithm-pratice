package p3366

import "slices"

const inf = 1 << 30

func minArraySum(nums []int, k int, op1 int, op2 int) int {

	dp := make([][]int, op1+1)
	fp := make([][]int, op1+1)
	for i := 0; i <= op1; i++ {
		dp[i] = make([]int, op2+1)
		fp[i] = make([]int, op2+1)
		for j := 0; j <= op2; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	calc1 := func(num int) int {
		num = (num + 1) / 2
		if num < k {
			return inf
		}
		return num - k
	}

	calc2 := func(num int) int {
		if num < k {
			return inf
		}
		return (num - k + 1) / 2
	}

	for _, num := range nums {
		// copy to fp, if no op
		for i := 0; i <= op1; i++ {
			for j := 0; j <= op2; j++ {
				fp[i][j] = dp[i][j] + num
			}
		}
		for i := 0; i <= op1; i++ {
			for j := 0; j <= op2; j++ {
				if i+1 <= op1 {
					// 取ceil
					fp[i+1][j] = min(fp[i+1][j], dp[i][j]+(num+1)/2)
				}
				if j+1 <= op2 && num >= k {
					fp[i][j+1] = min(fp[i][j+1], dp[i][j]+num-k)
				}
				// 两个同时操作
				if i+1 <= op1 && j+1 <= op2 {
					fp[i+1][j+1] = min(fp[i+1][j+1], dp[i][j]+min(calc1(num), calc2(num)))
				}
			}
		}

		for i := 0; i <= op1; i++ {
			for j := 0; j <= op2; j++ {
				dp[i][j] = fp[i][j]
			}
		}
	}

	res := inf

	for i := 0; i <= op1; i++ {
		res = min(res, slices.Min(dp[i]))
	}
	return res
}
