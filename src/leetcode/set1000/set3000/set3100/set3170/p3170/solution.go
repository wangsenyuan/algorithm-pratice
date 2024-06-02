package p3170

import "slices"

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return b[1] - a[1]
	})

	var res int
	if meetings[0][0] != 1 {
		res = meetings[0][0] - 1
	}
	r := meetings[0][1] + 1
	for i := 1; i < len(meetings); i++ {
		if r < meetings[i][0] {
			res += meetings[i][0] - r
			r = meetings[i][1] + 1
		} else {
			r = max(r, meetings[i][1]+1)
		}
	}
	if r <= days {
		res += days - r + 1
	}
	return res
}
