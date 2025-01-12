package p3420

import "math/bits"

func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)
	h := bits.Len(uint(n + 1))

	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	stack := make([]int, n)
	var top int

	// dp[i][j] 表示从i开始，到第1 << j个parent，满足非递减所需要的最少操作数
	dp := make([][]int, n+1)
	fa := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		fa[i] = make([]int, h)
		for j := 0; j < h; j++ {
			fa[i][j] = n
		}
		dp[i] = make([]int, h)

	}

	for i := n - 1; i >= 0; i-- {
		for top > 0 && nums[stack[top-1]] <= nums[i] {
			top--
		}
		if top > 0 {
			fa[i][0] = stack[top-1]
		}

		stack[top] = i
		top++
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][0] = (fa[i][0]-i)*nums[i] - (sum[fa[i][0]] - sum[i])
		for j := 1; j < h; j++ {
			fa[i][j] = fa[fa[i][j-1]][j-1]
			dp[i][j] = dp[i][j-1] + dp[fa[i][j-1]][j-1]
		}
	}

	var res int

	for i := 0; i < n; i++ {
		kk := k
		u := i
		for j := h - 1; j >= 0; j-- {
			if dp[u][j] <= kk {
				kk -= dp[u][j]
				u = fa[u][j]
			}
		}
		// dp[u][0] < kk
		u = search(u, fa[u][0], func(j int) bool {
			tmp := (j-u+1)*nums[u] - (sum[j+1] - sum[u])
			return tmp > kk
		})

		res += u - i
	}

	return int64(res)
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
