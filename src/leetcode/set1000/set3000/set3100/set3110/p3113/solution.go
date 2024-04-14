package p3113

import "slices"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	// only x matters
	slices.SortFunc(points, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})

	var res int
	n := len(points)

	for i := 0; i < n; {
		j := i
		for i < n && points[i][0] <= points[j][0]+w {
			i++
		}
		res++
	}

	return res
}
