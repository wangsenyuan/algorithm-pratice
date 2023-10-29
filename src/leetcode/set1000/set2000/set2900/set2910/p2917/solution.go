package p2917

func minSum(nums1 []int, nums2 []int) int64 {
	cnt := make([]int, 2)
	sum := make([]int64, 2)
	for _, num := range nums1 {
		if num == 0 {
			cnt[0]++
		}
		sum[0] += int64(num)
	}

	for _, num := range nums2 {
		if num == 0 {
			cnt[1]++
		}
		sum[1] += int64(num)
	}

	if sum[0] == sum[1] {
		// already equal
		if (cnt[0] == 0) != (cnt[1] == 0) {
			return -1
		}
		return sum[0] + int64(max(cnt[0], cnt[1]))
	}
	if cnt[0] == 0 {
		// 也就是无法增加sum[0]
		if cnt[1] == 0 || sum[1]+int64(cnt[1]) > sum[0] {
			return -1
		}
		return sum[0]
	}
	if cnt[1] == 0 {
		if sum[0]+int64(cnt[0]) > sum[1] {
			return -1
		}
		return sum[1]
	}
	// both not zero
	ans := sum[0] + int64(cnt[0])
	if ans < sum[1]+int64(cnt[1]) {
		ans = sum[1] + int64(cnt[1])
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
