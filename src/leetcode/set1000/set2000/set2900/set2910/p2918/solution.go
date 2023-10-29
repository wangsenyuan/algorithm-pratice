package p2918

func minIncrementOperations(nums []int, k int) int64 {
	var f0, f1, f2 int64
	for _, num := range nums {
		cur := f0 + max(0, int64(k-num))
		f0 = min(f1, cur)
		f1 = min(f2, cur)
		f2 = cur
	}
	return f0
}

func min(a, b int64) int64 {
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
