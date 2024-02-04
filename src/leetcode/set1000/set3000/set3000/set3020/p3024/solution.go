package p3024

import "sort"

func numberOfPairs(points [][]int) int {
	n := len(points)

	// y = x + d, 按照d排序
	sort.Slice(points, func(i, j int) bool {
		a := points[i]
		b := points[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	var res int

	for i := 0; i < n; i++ {
		my := -(1 << 30)
		for j := i + 1; j < n; j++ {
			if points[j][1] <= points[i][1] && points[j][1] > my {
				my = points[j][1]
				res++
			}
		}
	}

	return res
}
