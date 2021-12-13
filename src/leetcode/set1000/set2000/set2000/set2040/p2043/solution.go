package p2043

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	// n <= 16
	var res int
	for _, num := range nums {
		res |= num
	}
	var cnt int
	for mask := 1; mask < 1<<n; mask++ {
		var tmp int

		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				tmp |= nums[i]
			}
		}

		if tmp == res {
			cnt++
		}
	}
	return cnt
}
