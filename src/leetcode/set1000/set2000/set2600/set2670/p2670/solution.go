package p2670

func distinctDifferenceArray(nums []int) []int {
	cnt := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		cnt[nums[i]]++
	}

	pref := make(map[int]int)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		pref[nums[i]]++
		cnt[nums[i]]--
		if cnt[nums[i]] == 0 {
			delete(cnt, nums[i])
		}
		res[i] = len(pref) - len(cnt)
	}

	return res
}
