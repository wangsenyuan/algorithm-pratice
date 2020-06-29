package p1497

import "sort"

const MOD = 1000000007

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)

	var left, right = 0, len(nums) - 1
	var res int64
	for left <= right {
		if nums[left]+nums[right] <= target {
			res += pow(2, right-left)
			res %= MOD
			left++
		} else {
			right--
		}
	}
	return int(res)
}

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}
