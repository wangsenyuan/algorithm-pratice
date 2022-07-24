package p2354

import "sort"

func countExcellentPairs(nums []int, k int) int64 {
	sort.Ints(nums)
	var n int
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] > nums[i-1] {
			nums[n] = nums[i-1]
			n++
		}
	}
	cnt := make([]int, 32)
	var res int64

	for i := 0; i < n; i++ {
		num := nums[i]
		d := countDigit(num)
		for i := 30; i >= 0; i-- {
			if i+d >= k {
				res += 2 * int64(cnt[i])
			}
		}
		if d+d >= k {
			res++
		}
		cnt[d]++
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func countDigit(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
