package p1482

type TreeAncestor struct {
	dp [][]int
	h  int
	n  int
}

func Constructor(n int, parent []int) TreeAncestor {
	var h int

	for 1<<h <= n {
		h++
	}

	// 1 << h > n
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, h)
		dp[i][0] = parent[i]
	}

	for j := 1; j < h; j++ {
		for i := 0; i < n; i++ {
			p := dp[i][j-1]
			if p >= 0 {
				dp[i][j] = dp[p][j-1]
			} else {
				dp[i][j] = -1
			}
		}
	}

	return TreeAncestor{dp, h, n}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	u := node
	h := this.h
	dp := this.dp
	for i := h - 1; i >= 0 && u >= 0; i-- {
		if k&(1<<i) > 0 {
			u = dp[u][i]
		}
	}

	return u
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
