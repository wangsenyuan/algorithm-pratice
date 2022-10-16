package p2438

func minimizeArrayValue(nums []int) int {
	n := len(nums)
	// 只能把数字往左边移动

	check := func(x int) bool {
		var add int64

		for i := n - 1; i >= 0; i-- {
			// make all less than x
			cur := add + int64(nums[i])

			if cur <= int64(x) {
				add = 0
				continue
			}
			add = cur - int64(x)
		}

		return add == 0
	}

	var left int
	var right int

	for _, num := range nums {
		right = max(right, num)
	}

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
