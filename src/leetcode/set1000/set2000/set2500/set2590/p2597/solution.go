package p2597

import "sort"

func beautifulSubsets(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)

	cnt := make([]int, 1001)
	// a = b - k
	// 如果选择了a， 就不能选择b

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 1
		}
		if nums[i]-k > 0 && cnt[nums[i]-k] > 0 {
			return dfs(i + 1)
		}
		res := dfs(i + 1)
		cnt[nums[i]]++
		res += dfs(i + 1)
		cnt[nums[i]]--
		return res
	}

	return dfs(0) - 1
}
