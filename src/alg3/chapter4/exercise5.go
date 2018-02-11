package chapter4

func QuickSelect(nums []int, k int) int {
	s := LomutoPartition(nums)
	if s == k-1 {
		return nums[s]
	} else if s > k-1 {
		return QuickSelect(nums[:s], k)
	}
	return QuickSelect(nums[s+1:], k-1-s)
}

func LomutoPartition(nums []int) int {
	l, r, s := 0, len(nums), 0

	for i := l + 1; i < r; i++ {
		if nums[i] < nums[l] {
			s++
			nums[i], nums[s] = nums[s], nums[i]
		}
	}
	nums[s], nums[l] = nums[l], nums[s]
	return s
}
