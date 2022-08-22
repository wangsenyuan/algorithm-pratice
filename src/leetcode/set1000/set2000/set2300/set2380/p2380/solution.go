package p2380

func secondsToRemoveOccurrences(s string) int {
	n := len(s)
	//s[j] == 1
	//对于s[i] = 0, 要到达 位置j, 那么它需要的移动的时间 + 等待的时间
	//0遇到一个1，可以向右移动1个位置，
	// 假设它要移动到j，需要移动 j  - i 个位置，那么它需要pass j - i 个1
	var res int
	var cnt0 int
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			cnt0++
		} else if cnt0 > 0 {
			res = max(res+1, cnt0)
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
