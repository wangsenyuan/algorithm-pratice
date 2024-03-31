package p3099

func countAlternatingSubarrays(nums []int) int64 {
	// 要连续

	check := func(l int, r int) bool {
		if (r-l)&1 == 1 {
			return nums[r] != nums[l]
		}
		return nums[r] == nums[l]
	}
	var res int64
	n := len(nums)
	for i := 0; i < n; {
		j := i
		for i < n && check(j, i) {
			res += int64(i - j + 1)
			i++
		}
	}

	return res
}
