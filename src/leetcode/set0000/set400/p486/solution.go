package main

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))

}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			//dp[i][j] = sum[i][j] - min(dp[i+1][j], dp[i][j-1])
			dp[i][j] = sum[j+1] - sum[i] - min(dp[i+1][j], dp[i][j-1])
		}
	}
	x := dp[0][n-1]
	y := sum[n] - x
	return x >= y
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
