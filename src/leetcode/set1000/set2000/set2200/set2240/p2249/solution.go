package p2249

const MAX_X = 200
const MAX_Y = 200

func countLatticePoints(circles [][]int) int {
	inCircles := func(x, y int) bool {
		for _, circle := range circles {
			if inCircle(x, y, circle) {
				return true
			}
		}
		return false
	}

	var res int
	for x := 0; x <= MAX_X; x++ {
		for y := 0; y <= MAX_Y; y++ {
			if inCircles(x, y) {
				res++
			}
		}
	}

	return res
}

func inCircle(x, y int, circle []int) bool {
	x0, y0, r := circle[0], circle[1], int64(circle[2])

	dx := int64(x - x0)
	dy := int64(y - y0)
	return dx*dx+dy*dy <= r*r
}
