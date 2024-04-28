package p3128

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			// 以(i, j)为顶点的正方形
			var cnt int
			if grid[i][j] == 'B' {
				cnt++
			}
			if grid[i+1][j] == 'B' {
				cnt++
			}
			if grid[i][j+1] == 'B' {
				cnt++
			}
			if grid[i+1][j+1] == 'B' {
				cnt++
			}
			if cnt != 2 {
				return true
			}
		}
	}
	return false
}
