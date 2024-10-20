package p3326

import "slices"

func minOperations(nums []int) int {
	ma := slices.Max(nums)
	lps := make([]int, ma+1)
	var primes []int
	for i := 2; i <= ma; i++ {
		if lps[i] == 0 {
			lps[i] = i
			primes = append(primes, i)
		}
		for _, x := range primes {
			if i*x > ma {
				break
			}
			lps[i*x] = x
			if i%x == 0 {
				break
			}
		}
	}
	// 一次操作后，就变成了 lps[i]
	n := len(nums)
	next := nums[n-1]
	var res int
	for i := n - 2; i >= 0; i-- {
		x := nums[i]
		if x > next {
			x = lps[x]
			res++
		}
		if x > next {
			return -1
		}
		next = x
	}
	return res
}
