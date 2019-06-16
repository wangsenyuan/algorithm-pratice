package p1034

var dd = [...]int{-1, 0, 1, 0, -1}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	m := len(grid)
	n := len(grid[0])

	borders := make([][]bool, m)
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		borders[i] = make([]bool, n)
		vis[i] = make([]bool, n)
	}

	var findBorder func(x, y int)

	findBorder = func(x, y int) {
		if vis[x][y] {
			return
		}
		vis[x][y] = true

		var flag bool

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n {
				if grid[u][v] != grid[x][y] {
					// it is a border
					flag = true
				} else {
					findBorder(u, v)
				}
			} else {
				flag = true
			}
		}

		if flag {
			borders[x][y] = true
		}
	}

	findBorder(r0, c0)

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		copy(res[i], grid[i])
		for j := 0; j < n; j++ {
			if borders[i][j] {
				res[i][j] = color
			}
		}
	}

	return res
}
