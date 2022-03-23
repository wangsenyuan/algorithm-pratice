package p2210

func countHillValley(nums []int) int {
	n := len(nums)

	L := make([]int, n)

	for i := 1; i < n; i++ {
		if nums[i-1] != nums[i] {
			L[i] = i - 1
		} else {
			L[i] = L[i-1]
		}
	}

	var res int

	R := n - 1

	for i := n - 2; i >= 0; i-- {
		if nums[i+1] != nums[i] {
			R = i + 1
		} else {
			// R unchange
		}

		if L[i] != i && R != i {
			if nums[L[i]] < nums[i] && nums[R] < nums[i] || nums[L[i]] > nums[i] && nums[R] > nums[i] {
				res++
				if i > 0 && nums[i] == nums[i-1] {
					res--
				}
			}
		}
	}
	return res
}
