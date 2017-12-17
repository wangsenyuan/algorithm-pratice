package p749

var dd = []int{-1, 0, 1, 0, -1}

func containVirus(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	colors := make([][]int, n)
	for i := 0; i < n; i++ {
		colors[i] = make([]int, m)
	}

	var detect func(x, y, c int) int

	detect = func(x, y, c int) int {
		if grid[x][y] == 0 {
			if colors[x][y] == c {
				return 0
			}
			colors[x][y] = c
			return 1
		}
		colors[x][y] = c
		tmp := 0
		for k := 0; k < 4; k++ {
			i, j := x+dd[k], y+dd[k+1]
			if i >= 0 && i < n && j >= 0 && j < m {
				if grid[i][j] == 0 {
					tmp += detect(i, j, c)
				} else if grid[i][j] == 1 && colors[i][j] == 0 {
					tmp += detect(i, j, c)
				}
			}
		}

		return tmp
	}

	buildWall := func(color int) int {
		var res int
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == 1 && colors[i][j] == color {
					for k := 0; k < 4; k++ {
						x, y := i+dd[k], j+dd[k+1]
						if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 0 {
							res++
						}
					}
					// to avoid future process again
					colors[i][j] = -color
				}
			}
		}
		return res
	}

	var affect func(x, y, c int)
	affect = func(x, y, c int) {
		colors[x][y] = 0
		for k := 0; k < 4; k++ {
			i, j := x+dd[k], y+dd[k+1]
			if i >= 0 && i < n && j >= 0 && j < m {
				if grid[i][j] == 0 {
					grid[i][j] = 1
					colors[i][j] = 0
				} else if grid[i][j] == 1 && colors[i][j] == c {
					affect(i, j, c)
				}
			}
		}
	}

	var ans int

	color := 1
	for {
		most, cur := 0, -1
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == 1 && colors[i][j] == 0 {
					tmp := detect(i, j, color)
					if most < tmp {
						most, cur = tmp, color
					}
					color++
				}
			}
		}

		if most == 0 {
			break
		}
		ans += buildWall(cur)

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == 1 && colors[i][j] > 0 {
					affect(i, j, colors[i][j])
				}
			}
		}
	}

	return ans
}
