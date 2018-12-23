package p963

import "math"

func minAreaFreeRect(points [][]int) float64 {
	n := len(points)
	xx := make(map[int]map[int]bool)

	for _, p := range points {
		x, y := p[0], p[1]
		if _, found := xx[x]; !found {
			xx[x] = make(map[int]bool)
		}
		xx[x][y] = true
	}

	best := math.MaxFloat64

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				p1, p2, p3 := points[i], points[j], points[k]
				if valid(p1, p2, p3) {
					p4 := findPoint4(p1, p2, p3)
					if x, foundx := xx[p4[0]]; foundx {
						if _, foundy := x[p4[1]]; foundy {
							best = math.Min(best, area(p1, p2, p3))
						}
					}
				}

				if valid(p2, p1, p3) {
					p4 := findPoint4(p2, p1, p3)
					if x, foundx := xx[p4[0]]; foundx {
						if _, foundy := x[p4[1]]; foundy {
							best = math.Min(best, area(p2, p1, p3))
						}
					}
				}

				if valid(p3, p2, p1) {
					p4 := findPoint4(p3, p2, p1)
					if x, foundx := xx[p4[0]]; foundx {
						if _, foundy := x[p4[1]]; foundy {
							best = math.Min(best, area(p3, p2, p1))
						}
					}
				}
			}
		}
	}

	if best == math.MaxFloat64 {
		return 0
	}
	return best
}

func valid(p1, p2, p3 []int) bool {
	x1 := p2[0] - p1[0]
	y1 := p2[1] - p1[1]
	x2 := p3[0] - p1[0]
	y2 := p3[1] - p1[1]
	return x1*x2+y1*y2 == 0
}

func findPoint4(p1, p2, p3 []int) []int {
	return []int{p3[0] - p1[0] + p2[0], p3[1] - p1[1] + p2[1]}
}

func area(p1, p2, p3 []int) float64 {
	x1 := float64(p2[0] - p1[0])
	y1 := float64(p2[1] - p1[1])
	x2 := float64(p3[0] - p1[0])
	y2 := float64(p3[1] - p1[1])
	return math.Sqrt(x1*x1+y1*y1) * math.Sqrt(x2*x2+y2*y2)
}
