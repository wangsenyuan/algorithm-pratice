package p1893

import "sort"

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		a, b := ranges[i], ranges[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})

	if left < ranges[0][0] {
		return false
	}

	prev := ranges[0][0]

	for i := 0; i < len(ranges); i++ {
		cur := ranges[i]
		if cur[0] > prev+1 {
			return false
		}
		prev = max(prev, cur[1])
	}

	return prev >= right
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
