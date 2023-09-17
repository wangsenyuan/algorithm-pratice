package p2858

func minimumRightShifts(nums []int) int {
	var i int
	n := len(nums)
	for i+1 < n && nums[i+1] > nums[i] {
		i++
	}
	if i+1 == n {
		return 0
	}
	// nums[i+1] < nums[i]
	j := i + 1
	for j+1 < n && nums[j+1] > nums[j] {
		j++
	}
	if j+1 < n || nums[j] > nums[0] {
		return -1
	}
	return n - i - 1
}
