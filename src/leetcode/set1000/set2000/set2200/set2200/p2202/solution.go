package p2202

func maximumTop(nums []int, k int) int {
	//如果要将nums[i]变成栈顶，至少需要i次操作
	n := len(nums)

	if n == 1 {
		if k%2 == 0 {
			return nums[0]
		}
		return -1
	}

	best := -1
	soFar := -1
	for i := 0; i < n && k >= 0; i++ {
		if k%2 == 0 {
			// remove this and add back
			// 加回一个数字 用掉 1次操作，
			// 然后两个一组不断的删掉加回， 被将nums[i] 做为栈顶
			best = max(best, nums[i])
		}

		if k > 0 {
			best = max(best, soFar)
		}

		soFar = max(soFar, nums[i])
		k--
	}

	if k > 0 {
		best = soFar
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
