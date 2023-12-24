package p2970

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)

	r := n - 1
	for r-1 >= 0 && nums[r-1] < nums[r] {
		r--
	}
	if r == 0 {
		return int64(1+n) * int64(n) / 2
	}

	var l int
	var res int

	for ; r < n; r++ {
		if nums[l] >= nums[r] {
			// l = 0, 要保留r，必须删除l
			res += l + 1
			continue
		}
		// nums[l] < nums[r] 在l的前面满足 nums[l-1] < nums[l]
		for nums[l+1] < nums[r] && nums[l] < nums[l+1] {
			l++
		}
		res += l + 2
	}

	for nums[l] < nums[l+1] {
		l++
	}

	res += l + 1

	res++

	return int64(res)
}
