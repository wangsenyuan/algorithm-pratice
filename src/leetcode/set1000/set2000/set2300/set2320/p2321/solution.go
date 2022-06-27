package p2321

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	n := len(nums1)

	a := solve(n, nums1, nums2)
	b := solve(n, nums2, nums1)

	return int(max(a, b))
}

func solve(n int, nums1 []int, nums2 []int) int64 {
	// 只能交换1次
	// sum1[r] - sum1[l-1] > sum2[r] - sum2[l-1]
	// sum1[r] - sum2[r] > sum1[l-1] - sum2[l-1]
	// 建nums1[l...r]替换到nums2[l...r]后，nums2所能达到的最大值
	var min_left int64
	var sum int64
	var sum2 int64
	var best int64
	for i := 0; i < n; i++ {
		sum2 += int64(nums2[i])
		sum += int64(nums1[i]) - int64(nums2[i])
		best = max(best, sum-min_left)
		min_left = min(min_left, sum)
	}

	return best + sum2
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
