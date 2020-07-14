package p1513

import "math"

const LMT = 1e-7

var dd = []float64{-1, 0, 1, 0, -1}

func getMinDistSum(positions [][]int) float64 {
	n := len(positions)

	if n <= 1 {
		return 0
	}

	getDistance := func(x, y float64) float64 {
		var res float64

		for i := 0; i < len(positions); i++ {
			res += distance(positions[i], x, y)
		}

		return res
	}

	var xx int
	var yy int
	for i := 0; i < len(positions); i++ {
		xx += positions[i][0]
		yy += positions[i][1]
	}

	x := float64(xx) / float64(n)
	y := float64(yy) / float64(n)

	best := getDistance(x, y)

	var test float64 = 1000

	for test > LMT {
		var found bool

		for i := 0; i < 4; i++ {
			nx, ny := x+test*dd[i], y+test*dd[i+1]
			tmp := getDistance(nx, ny)

			if tmp < best {
				best = tmp
				found = true
				x, y = nx, ny
				break
			}
		}
		if !found {
			test /= 2
		}
	}

	return best
}

func distance(a []int, x, y float64) float64 {
	dx := float64(a[0]) - x
	dy := float64(a[1]) - y

	ds := dx*dx + dy*dy

	return math.Sqrt(ds)
}
