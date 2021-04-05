package a

import "sort"

const MOD = 1000000007

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)

	var res int
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := sort.Search(i, func(j int) bool {
			return nums[j] > target-cur
		})
		res += j
		res %= MOD
	}
	return res
}
