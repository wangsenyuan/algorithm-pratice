package p1863

func subsetXORSum(nums []int) int {
	n := len(nums)

	tot := 1 << n
	var res int
	for mask := 1; mask < tot; mask++ {
		var xor int
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				xor ^= nums[i]
			}
		}
		res += xor
	}
	return res
}
