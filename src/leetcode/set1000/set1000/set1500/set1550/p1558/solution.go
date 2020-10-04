package p1558
func minOperations(nums []int) int {
	var res int

	n := len(nums)

	for {
		var stop = true
		for i := 0; i < n; i++ {
			if nums[i]&1 == 1 {
				res++
				nums[i]--
			}
			nums[i] >>= 1
			if nums[i] > 0 {
				stop = false
			}
		}

		if stop {
			break
		}
		res++
	}
	return res
}