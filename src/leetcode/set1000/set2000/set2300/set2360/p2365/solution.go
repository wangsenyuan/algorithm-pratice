package p2365

func minimumReplacement(nums []int) int64 {
	n := len(nums)
	var res int64

	cur := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] <= cur {
			cur = nums[i]
			continue
		}
		// nums[i] > cur
		// 建nums[i]分成a + b + c .., a <= cur, b <= cur, c <= cur
		x := (nums[i] + cur - 1) / cur
		res += int64(x - 1)
		r := nums[i] % cur
		if r == 0 {
			continue
		}
		// 2, 5, 5
		// 4, 4, 4
		// a * x + y == nums[i], and a <= cur
		cur = nums[i] / x
	}

	return res
}
