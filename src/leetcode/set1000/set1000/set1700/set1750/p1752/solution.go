package p1752

func check(nums []int) bool {
	n := len(nums)
	var i int

	for i+1 < n && nums[i+1] >= nums[i] {
		i++
	}

	// i + 1 == n || nums[i+1] < nums[i]
	if i+1 == n {
		return true
	}

	i++
	for i+1 < n && nums[i+1] >= nums[i] {
		i++
	}
	if i+1 < n {
		return false
	}

	return nums[n-1] <= nums[0]
}
