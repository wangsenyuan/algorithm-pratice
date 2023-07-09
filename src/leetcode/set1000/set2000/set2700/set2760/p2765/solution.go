package p2765

func alternatingSubarray(nums []int) int {
	var res int

	for i := 0; i+1 < len(nums); i++ {
		if nums[i+1]-nums[i] == 1 {
			diff := 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j]-nums[j-1] != diff {
					break
				}
				res = max(res, j-i+1)
				diff *= -1
			}
		}
	}

	if res <= 1 {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
