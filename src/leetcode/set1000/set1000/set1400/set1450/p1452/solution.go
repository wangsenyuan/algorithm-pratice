package p1452

import "math"

func numPoints(points [][]int, r int) int {
	if len(points) == 0 {
		return 0
	}
	L := float64(2 * r)
	n := len(points)

	R := float64(r)

	countInPoints := func(center []float64) int {
		var cnt int
		for i := 0; i < n; i++ {
			d := distance2(points[i], center)
			if d <= R {
				cnt++
			}
		}
		return cnt
	}

	var best int = 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}

			d := distance(points[i], points[j])
			if d > L {
				continue
			}

			center := getCenter(points[i], points[j], r)

			cnt := countInPoints(center)
			if cnt > best {
				best = cnt
			}
		}
	}

	return best
}

func distance(a, b []int) float64 {
	dx := float64(b[0] - a[0])
	dy := float64(b[1] - a[1])

	d := dx*dx + dy*dy

	return math.Sqrt(d)
}

func distance2(a []int, b []float64) float64 {
	dx := b[0] - float64(a[0])
	dy := b[1] - float64(a[1])

	d := dx*dx + dy*dy

	return math.Sqrt(d)
}

func getCenter(a, b []int, r int) []float64 {
	//
	x1, y1 := float64(a[0]), float64(a[1])
	x2, y2 := float64(b[0]), float64(b[1])
	rr := float64(r) * float64(r)

	q := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))

	x3 := (x1 + x2) / 2
	y3 := (y1 + y2) / 2
	basex := math.Sqrt(rr-((q/2)*(q/2))) * ((y1 - y2) / q)
	basey := math.Sqrt(rr-((q/2)*(q/2))) * ((x2 - x1) / q)
	return []float64{x3 + basex, y3 + basey}
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}

	if b >= a && b >= c {
		return b
	}
	return c
}
