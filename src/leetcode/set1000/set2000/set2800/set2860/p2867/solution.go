package p2866

func maximumSumOfHeights(maxHeights []int) int64 {
	//n := len(maxHeights)
	// n = 1e5
	// height[i] <= maxHeights[i]
	// 且存在某个i，是山峰，前面 <= height[i]，  height[i] >= 后面
	// 如果确定了位置i
	// 那么对于前面的来说 <= max_suf[j...i]
	// 如果已经直到了顶峰在i处时的最优，能不能很容易计算处i+1处的呢？
	// 如果 height[i+1] < height[i] 答案不会更优的
	n := len(maxHeights)
	pref := make([]int64, n)
	// 在满足左边单调递增的情况下，到i为止，最大的sum是多少
	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && maxHeights[stack[p-1]] > maxHeights[i] {
			p--
		}
		if p == 0 {
			pref[i] = int64(i+1) * int64(maxHeights[i])
		} else {
			j := stack[p-1]
			pref[i] = int64(i-j)*int64(maxHeights[i]) + pref[j]
		}
		stack[p] = i
		p++
	}

	p = 0

	suf := make([]int64, n)

	for i := n - 1; i >= 0; i-- {
		for p > 0 && maxHeights[stack[p-1]] > maxHeights[i] {
			p--
		}
		if p == 0 {
			suf[i] = int64(n-i) * int64(maxHeights[i])
		} else {
			j := stack[p-1]
			suf[i] = int64(j-i)*int64(maxHeights[i]) + suf[j]
		}
		stack[p] = i
		p++
	}

	var best int64

	for i := 0; i < n; i++ {
		tmp := pref[i] + suf[i] - int64(maxHeights[i])
		best = max(best, tmp)
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
