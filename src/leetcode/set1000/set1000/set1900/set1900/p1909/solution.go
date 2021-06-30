package p1909

func canBeIncreasing(nums []int) bool {
	pivot := -1

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if pivot >= 0 {
				return false
			}
			pivot = i
		}
	}

	if pivot < 0 || pivot == 1 {
		return true
	}
	// nums[pivot-1] >= nums[pivot]
	a, b := true, true
	// a is for remove pivot - 1
	// b is for remove pivot
	for i := 0; i < len(nums); i++ {
		if a && i != pivot-1 {
			if i-1 == pivot-1 {
				a = nums[i] > nums[i-2]
			} else if i > 0 {
				a = nums[i] > nums[i-1]
			}
		}
		if b && i != pivot {
			if i-1 == pivot {
				b = nums[i] > nums[i-2]
			} else if i > 0 {
				b = nums[i] > nums[i-1]
			}
		}
	}
	return a || b
}
