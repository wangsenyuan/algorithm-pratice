package p740

import "sort"

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cnts := make(map[int]int)

	for _, num := range nums {
		cnts[num]++
	}

	keys := make([]int, len(cnts))
	sz := 0
	for num := range cnts {
		keys[sz] = num
		sz++
	}

	if sz == 1 {
		return keys[0] * cnts[keys[0]]
	}

	sort.Ints(keys)
	scores := make([]int, sz)
	for i := 0; i < sz; i++ {
		scores[i] = keys[i] * cnts[keys[i]]
	}

	a, b := 0, scores[0]
	for i := 1; i < sz; i++ {
		var c int
		if keys[i] > keys[i-1]+1 {
			c = scores[i] + b
		} else {
			if i > 1 {
				c = max(scores[i]+a, b)
			} else {
				c = max(scores[i], b)
			}
		}
		a, b = b, c
	}
	return max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
