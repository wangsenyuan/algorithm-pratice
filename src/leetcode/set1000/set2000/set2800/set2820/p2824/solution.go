package p2824

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	var ans int
	n := len(nums)

	for i := 0; i < n; i++ {
		j := sort.Search(n, func(j int) bool {
			return nums[i]+nums[j] >= target
		})
		j--
		if j <= i {
			break
		}
		ans += j - i
	}
	return ans
}
