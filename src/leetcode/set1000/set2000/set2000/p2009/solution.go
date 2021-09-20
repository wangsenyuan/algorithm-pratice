package p2009

import "sort"

/**
给你一个整数数组 nums 。每一次操作中，你可以将 nums 中 任意 一个元素替换成 任意 整数。

如果 nums 满足以下条件，那么它是 连续的 ：

nums 中所有元素都是 互不相同 的。
nums 中 最大 元素与 最小 元素的差等于 nums.length - 1 。
比方说，nums = [4, 2, 5, 3] 是 连续的 ，但是 nums = [1, 2, 3, 5, 6] 不是连续的 。

请你返回使 nums 连续 的 最少 操作次数。

*/

func minOperations(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// n <= 100000
	// if we fix nums[i], then we want to get nums[i], nums[i]+1, nums[i] + 2, ... nums[i] + n - 1
	// we need to get the cnt in the range, answer would be n - cnt
	var ans int = n
	var cnt int
	var end int
	mem := make(map[int]int)
	for i := 0; i < n; i++ {
		// fix nums[i]
		for end < n && nums[end]-n < nums[i] {
			mem[nums[end]]++
			if mem[nums[end]] == 1 {
				cnt++
			}
			end++
		}
		ans = min(ans, n-cnt)
		mem[nums[i]]--
		if mem[nums[i]] == 0 {
			cnt--
			delete(mem, nums[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
