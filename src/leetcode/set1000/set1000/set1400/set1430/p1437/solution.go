package p1437

func kLengthApart(nums []int, k int) bool {
	var prev int = -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		// nums[i] == 1
		if prev >= 0 && i-prev <= k {
			return false
		}
		prev = i
	}

	return true
}
