package p2389

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	n := len(nums)
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + nums[i]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		j := search(n+1, func(j int) bool {
			return pref[j] > cur
		})
		ans[i] = j - 1
	}

	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
