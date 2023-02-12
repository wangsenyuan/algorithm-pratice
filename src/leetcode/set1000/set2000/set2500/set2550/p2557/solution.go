package p2557

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	n := len(nums)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			res += int64(nums[i])
			continue
		}
		var b int64 = 1
		for b <= int64(nums[j]) {
			b *= 10
		}
		// b > nums[j]
		res += int64(nums[i])*b + int64(nums[j])
	}

	return res
}
