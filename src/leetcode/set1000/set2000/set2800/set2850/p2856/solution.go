package p2856

func minLengthAfterRemovals(nums []int) int {
	// a < b < c
	// 这种情况下，应该优先使用消去 b和c
	n := len(nums)
	var maxCnt int
	for i := 0; i < n; {
		j := i
		for i < n && nums[i] == nums[j] {
			i++
		}
		maxCnt = max(maxCnt, i-j)
	}
	if maxCnt*2 > n {
		return 2*maxCnt - n
	}
	return n & 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
