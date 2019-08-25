package p1105

func minHeightShelves(books [][]int, shelf_width int) int {
	n := len(books)

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		h := books[i-1][1]
		w := books[i-1][0]
		dp[i] = h + dp[i-1]

		for j := i - 1; j > 0 && w+books[j-1][0] <= shelf_width; j-- {
			h = max(h, books[j-1][1])
			w += books[j-1][0]
			dp[i] = min(dp[i], h+dp[j-1])
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
