package p3382

func constructTransformedArray(nums []int) []int {
	n := len(nums)

	result := make([]int, n)

	for i, x := range nums {
		if x == 0 {
			result[i] = x
		} else if x > 0 {
			x %= n
			j := (i + x) % n
			result[i] = nums[j]
		} else {
			x = -x
			x %= n
			j := (i - x + n) % n
			result[i] = nums[j]
		}
	}

	return result
}
