package p5561

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	best := 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			nums[i] = nums[i>>1]
		} else {
			nums[i] = nums[i>>1] + nums[i>>1+1]
		}
		if nums[i] > best {
			best = nums[i]
		}
	}
	return best
}
