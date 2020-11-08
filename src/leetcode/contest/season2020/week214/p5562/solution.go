package p5562

import "sort"

func minDeletions(s string) int {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		cnt[x]++
	}
	nums := make([]int, 0, 26)
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			nums = append(nums, cnt[i])
		}
	}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		nums[i] -= i
	}

	var res int

	n := len(nums)

	sum := func(end int) int {
		var tot int
		for i := 0; i <= end; i++ {
			tot += nums[i] + i
		}
		return tot
	}

	// make sure nums non-desc

	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			if nums[i]+i >= nums[i]-nums[i+1] {
				tmp := nums[i] - nums[i+1]
				res += tmp
				nums[i] -= tmp
			} else {
				res += sum(i)
				break
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
