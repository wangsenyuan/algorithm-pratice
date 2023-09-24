package p2866

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	// height[i] <= maxHeights[i]
	// 且存在某个i，是山峰，前面 <= height[i]，  height[i] >= 后面
	// 如果确定了位置i
	// 那么对于前面的来说 <= max_suf[j...i]
	var best int64

	for i := 0; i < n; i++ {
		tmp := int64(maxHeights[i])
		cur := maxHeights[i]
		for j := i - 1; j >= 0; j-- {
			cur = min(cur, maxHeights[j])
			tmp += int64(cur)
		}
		cur = maxHeights[i]
		for j := i + 1; j < n; j++ {
			cur = min(cur, maxHeights[j])
			tmp += int64(cur)
		}
		if best < tmp {
			best = tmp
		}
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
