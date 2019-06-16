package p889

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	valid := func(x, y int) bool {
		return x >= 0 && x < R && y >= 0 && y < C
	}
	x, y := r0, c0
	dx, dy := 0, 1
	W, w := 1, 0

	T := R * C
	var i int
	res := make([][]int, T)
	for i < T {
		if valid(x, y) {
			res[i] = []int{x, y}
			i++
		}
		w++
		if w > W {
			if dy == 0 {
				// change when turn east or west
				W++
			}
			w = 1
			turn(&dx, &dy)
		}
		x += dx
		y += dy
	}
	return res
}

func turn(dx, dy *int) {
	if *dx == 0 && *dy == 1 {
		*dx, *dy = 1, 0
	} else if *dx == 1 && *dy == 0 {
		*dx, *dy = 0, -1
	} else if *dx == 0 && *dy == -1 {
		*dx, *dy = -1, 0
	} else {
		*dx, *dy = 0, 1
	}
}
