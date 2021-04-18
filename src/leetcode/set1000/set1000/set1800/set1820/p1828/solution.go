package p1828

import (
	"math"
	"sort"
)

func countPoints(points [][]int, queries [][]int) []int {
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[0] < b[0]
	})

	qs := make([]Query, len(queries))
	for i, cur := range queries {
		qs[i] = Query{cur[0], cur[1], cur[2], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		return a.x-a.r < b.x-b.r || a.x-a.r == b.x-b.r && a.x < b.x
	})

	res := make([]int, len(queries))
	var j, k int
	for _, cur := range qs {
		x, y, r := cur.x, cur.y, cur.r

		for k < len(points) && points[k][0] <= x+r {
			k++
		}

		for j < len(points) && points[j][0] < x-r {
			j++
		}
		for ii := j; ii < k; ii++ {
			u, v := points[ii][0], points[ii][1]
			if inCycle(x, y, r, u, v) {
				res[cur.id]++
			}
		}
	}
	return res
}

func inCycle(x, y, r int, u, v int) bool {
	dx := float64(x - u)
	dy := float64(y - v)
	R := float64(r)
	X := dx*dx + dy*dy
	Y := R * R
	return X <= Y || math.Abs(X-Y) <= 1e-7
}

type Query struct {
	x, y, r int
	id      int
}
