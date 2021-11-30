package p2091

func minimumDeletions(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	a, b := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[a] {
			a = i
		}
		if nums[i] < nums[b] {
			b = i
		}
	}
	n := len(nums)

	if a > b {
		a, b = b, a
	}

	// from both ends
	res := a + 1 + (n - b)

	if n-a < res {
		// from right, delete to a
		res = n - a
	}
	if b+1 < res {
		res = b + 1
	}
	return res
}
