package p812

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var res int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]
				s := abs(x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))
				if s > res {
					res = s
				}
			}
		}
	}

	return 0.5 * float64(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
