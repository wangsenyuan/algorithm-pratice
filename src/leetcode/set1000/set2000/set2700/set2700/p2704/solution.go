package p2704

import "sort"

func maxStrength(nums []int) int64 {
	n := len(nums)
	if n == 1 {
		return int64(nums[0])
	}

	var neg []int
	var pos []int
	var zero bool

	for _, num := range nums {
		if num > 0 {
			pos = append(pos, num)
		} else if num < 0 {
			neg = append(neg, -num)
		} else {
			zero = true
		}
	}

	if len(pos) == 0 && len(neg) <= 1 {
		if zero {
			return 0
		}
		return int64(neg[0])
	}
	var res int64 = 1

	for _, x := range pos {
		res *= int64(x)
	}

	if len(neg)%2 == 0 {
		for _, x := range neg {
			res *= int64(x)
		}
	} else {
		sort.Ints(neg)
		for i := 1; i < len(neg); i++ {
			res *= int64(neg[i])
		}
	}

	return res
}
