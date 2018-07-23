package p840

func numMagicSquaresInside(grid [][]int) int {
	n := len(grid)
	if n < 3 {
		return 0
	}

	m := len(grid[0])
	if m < 3 {
		return 0
	}

	var ans int
	for i := 0; i+3 <= n; i++ {
		for j := 0; j+3 <= m; j++ {
			can := true
			var sum int
			row := make([]int, 3)
			col := make([]int, 3)
			var d1 int
			var d2 int
			for a := 0; a < 3 && can; a++ {
				for b := 0; b < 3 && can; b++ {
					if grid[i+a][j+b] == 0 || grid[i+a][j+b] > 9 {
						can = false
					} else {
						sum += grid[i+a][j+b]
						row[a] += grid[i+a][j+b]
						col[b] += grid[i+a][j+b]
						if a == b {
							d1 += grid[i+a][j+b]
						}
						if a+b == 2 {
							d2 += grid[i+a][j+b]
						}
					}
				}
			}
			if can && sum == 45 &&
				row[0] == 15 && row[1] == 15 && row[2] == 15 &&
				col[0] == 15 && col[1] == 15 && col[2] == 15 &&
				d1 == 15 && d2 == 15 {
				ans++
			}
		}
	}

	return ans
}
