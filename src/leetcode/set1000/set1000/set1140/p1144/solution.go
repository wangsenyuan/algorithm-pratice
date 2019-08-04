package p1144

func movesToMakeZigzag(nums []int) int {
	// first try a > b < c > d
	// then try a < b > c < d
	n := len(nums)
	var res1 int

	for i := 1; i < n; i += 2 {
		x := nums[i-1] - 1
		if i+1 < n {
			x = min(x, nums[i+1]-1)
		}
		if nums[i] > x {
			res1 += nums[i] - x
		}
	}

	var res2 int

	for i := 0; i < n; i += 2 {
		x := 1 << 20
		if i+1 < n {
			x = nums[i+1] - 1
		}
		if i-1 >= 0 {
			x = min(x, nums[i-1]-1)
		}
		if nums[i] > x {
			res2 += nums[i] - x
		}
	}

	return min(res1, res2)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
