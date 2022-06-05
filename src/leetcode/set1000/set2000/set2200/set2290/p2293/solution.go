package p2293

func minMaxGame(nums []int) int {
	n := len(nums)

	for n > 1 {
		for i := 0; i < n/2; i++ {
			if i&1 == 0 {
				nums[i] = min(nums[i*2], nums[i*2+1])
			} else {
				nums[i] = max(nums[i*2], nums[i*2+1])
			}
		}
		n /= 2
	}

	return nums[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
