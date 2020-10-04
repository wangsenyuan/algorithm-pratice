package p1610

import (
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	n := len(points)
	aa := make([]float64, 0, n)
	var same int
	x0, y0 := location[0], location[1]

	for i := 0; i < n; i++ {
		p := points[i]
		x, y := p[0], p[1]
		if x == x0 && y == y0 {
			same++
			continue
		}
		dx, dy := float64(x-x0), float64(y-y0)
		aa = append(aa, math.Atan2(dy, dx))
	}

	sort.Float64s(aa)

	bb := make([]float64, 2*len(aa))
	copy(bb, aa)
	copy(bb[len(aa):], aa)

	for i := len(aa); i < len(bb); i++ {
		bb[i] += 2 * math.Pi
	}

	ap := math.Pi * float64(angle) / 180

	var best int

	for i := 0; i < len(aa); i++ {
		j := sort.Search(len(bb), func(j int) bool {
			return bb[j] > bb[i]+ap
		})
		best = max(best, j-i)
	}

	return best + same
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
