package p1838

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	K := int64(k)
	var best int
	var j int
	var tot int64
	var cnt int
	// if ends at some num x, how many nums < x, can change to x in k steps
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		cnt++
		for j < i && tot > K {
			y := nums[j]
			tot -= int64(x - y)
			cnt--
			j++
		}
		if tot <= K {
			best = max(best, cnt)
		}
		if i+1 < len(nums) {
			tot += int64(cnt) * int64(nums[i+1]-x)
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
