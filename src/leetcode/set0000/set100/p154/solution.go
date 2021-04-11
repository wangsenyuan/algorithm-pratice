package p154

func findMin(nums []int) int {
	// min is between [l, r)
	n := len(nums)

	if n == 1 || nums[0] < nums[n-1] {
		return nums[0]
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[r]
}
