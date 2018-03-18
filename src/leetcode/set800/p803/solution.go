package p803

func hitBricks(grid [][]int, hits [][]int) []int {
	n := len(grid)
	m := len(grid[0])

	res := make([]int, len(hits))

	dd := []int{-1, 0, 1, 0, -1}
	var check func(x, y, label int) bool
	checked := make([][]int, n)
	for i := 0; i < n; i++ {
		checked[i] = make([]int, m)
	}
	check = func(x, y int, label int) bool {
		checked[x][y] = label

		if grid[x][y] == 0 {
			return true
		}
		if x == 0 {
			return false
		}
		for k := 0; k < 4; k++ {
			a, b := x+dd[k], y+dd[k+1]
			if a >= 0 && a < n && b >= 0 && b < m && checked[a][b] != label {
				tmp := check(a, b, label)
				if !tmp {
					return false
				}
			}
		}
		return true
	}
	que := make([]int, n*m)

	fall := func(x, y int) int {
		grid[x][y] = 0

		head, tail := 0, 0
		que[tail] = x*m + y
		tail++
		for head < tail {
			tt := tail
			for head < tt {
				a, b := que[head]/m, que[head]%m
				head++
				for k := 0; k < 4; k++ {
					c, d := a+dd[k], b+dd[k+1]
					if c >= 0 && c < n && d >= 0 && d < m && grid[c][d] > 0 {
						grid[c][d] = 0
						que[tail] = c*m + d
						tail++
					}
				}
			}
		}
		return tail
	}
	var label int
	for i := 0; i < len(hits); i++ {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] > 0 {
			grid[x][y] = 0
			var removal int
			for k := 0; k < 4; k++ {
				a, b := x+dd[k], y+dd[k+1]
				if a >= 0 && a < n && b >= 0 && b < m {
					label++
					if grid[a][b] > 0 && check(a, b, label) {
						removal += fall(a, b)
					}
				}
			}
			res[i] = removal
		}
	}

	return res
}
