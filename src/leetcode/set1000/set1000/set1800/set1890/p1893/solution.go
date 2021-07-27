package p1893

import "sort"

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		a, b := ranges[i], ranges[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})

	for i := 0; i < len(ranges); i++ {
		if ranges[i][1] < left {
			continue
		}
		// ranges[i][1] >= left
		// left < ranges[i][0]
		if left < ranges[i][0] {
			return false
		}

		if ranges[i][0] > right {
			return true
		}

		// left >= ranges[i][0] and left <= ranges[i][1] && ranges[i][0] <= right
		left = ranges[i][1] + 1
	}

	return left > right
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
