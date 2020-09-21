package p1592

func isPrintable(targetGrid [][]int) bool {
	m := len(targetGrid)
	n := len(targetGrid[0])
	checkColor := func() bool {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if targetGrid[i][j] > 0 {
					return false
				}
			}
		}
		return true
	}

	dim := make(map[int][4]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			color := targetGrid[i][j]
			a, b, c, d := m, n, 0, 0
			if v, found := dim[color]; found {
				a, b, c, d = v[0], v[1], v[2], v[3]
			}
			a = min(a, i)
			b = min(b, j)
			c = max(c, i)
			d = max(d, j)
			dim[color] = [...]int{a, b, c, d}
		}
	}

	findRect := func() (int, int, int) {
		mem := make(map[int]bool)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				color := targetGrid[i][j]
				if color == 0 || mem[color] {
					continue
				}
				//the top-left corner
				mem[color] = true
				v := dim[color]
				a, b, c, d := v[0], v[1], v[2], v[3]
				//i == a, j == b
				var ok = true
				for ii := a; ii <= c && ok; ii++ {
					for jj := b; jj <= d && ok; jj++ {
						if targetGrid[ii][jj] != 0 && targetGrid[ii][jj] != color {
							ok = false
						}
					}
				}
				if ok {
					return i, j, color
				}
			}
		}
		return -1, -1, -1
	}

	colorRect := func(x, y, color int) {
		v := dim[color]
		a, b, c, d := v[0], v[1], v[2], v[3]
		//i == a, j == b
		for ii := a; ii <= c; ii++ {
			for jj := b; jj <= d; jj++ {
				targetGrid[ii][jj] = 0
			}
		}
	}

	for !checkColor() {
		// find some rect that has same color
		x, y, c := findRect()
		if x < 0 {
			return false
		}
		colorRect(x, y, c)
	}

	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
