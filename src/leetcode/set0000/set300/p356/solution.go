package main

import "math"

func isReflected(points [][]int) bool {
	if len(points) == 0 {
		return true
	}

	a, b := math.MaxInt32, math.MinInt32

	type P struct {
		x, y int
	}
	cache := make(map[P]bool)

	for _, point := range points {
		if point[0] < a {
			a = point[0]
		}

		if point[0] > b {
			b = point[0]
		}
		p := P{point[0], point[1]}
		cache[p] = true
	}

	x := a + b
	for _, point := range points {
		var p P
		if point[0]*2 < x {
			//at left
			d := x - point[0]*2
			p = P{point[0] + d, point[1]}
		} else if point[0]*2 > x {
			//at right
			d := point[0]*2 - x
			p = P{point[0] - d, point[1]}
		} else {
			//at the line
			p = P{point[0], point[1]}
		}

		if !cache[p] {
			return false
		}
	}

	return true
}
