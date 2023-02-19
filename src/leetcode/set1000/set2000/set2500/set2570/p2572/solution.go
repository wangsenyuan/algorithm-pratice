package p2572

func minOperations(n int) int {
	// 连续的1 111 + 1 => 1111
	if n == 0 {
		return 0
	}
	var arr []int

	for n > 0 {
		arr = append(arr, n&1)
		n >>= 1
	}
	reverse(arr)
	m := len(arr)
	sum := make([]int, m+1)

	for i := 0; i < m; i++ {
		sum[i+1] = arr[i] + sum[i]
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := i; j < m; j++ {
			cnt := sum[j+1] - sum[i]
			dp[i][j] = min(cnt, j-i+1-cnt+2)
		}
	}
	for j := 0; j < m; j++ {
		for i := j - 1; i >= 0; i-- {
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
		}
	}
	return dp[0][m-1]
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
