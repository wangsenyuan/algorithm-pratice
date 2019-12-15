package p1289

func minFallingPathSum(arr [][]int) int {
	n := len(arr)

	if n == 1 {
		return arr[0][0]
	}

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = arr[0][i]
	}

	for i := 1; i < n; i++ {
		first := -1
		second := -1

		for j := 0; j < n; j++ {
			if first < 0 || dp[j] < dp[first] {
				second = first
				first = j
			} else if second < 0 || dp[j] < dp[second] {
				second = j
			}
		}
		a := dp[first]
		b := dp[second]

		for j := 0; j < n; j++ {
			if j == first {
				dp[j] = b + arr[i][j]
			} else {
				dp[j] = a + arr[i][j]
			}
		}
	}

	res := dp[0]
	for j := 1; j < n; j++ {
		if dp[j] < res {
			res = dp[j]
		}
	}
	return res
}
