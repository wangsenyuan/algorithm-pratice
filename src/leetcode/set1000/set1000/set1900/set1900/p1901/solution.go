package p1901

func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	l, r := 0, n

	for l < r {
		mid := (l + r) / 2

		x, y := 0, mid
		for i := 0; i < m; i++ {
			for k := -1; k <= 1; k++ {
				j := mid + k
				if j >= 0 && j < n && mat[i][j] > mat[x][y] {
					x, y = i, j
				}
			}
		}
		if y < mid {
			// max in left side
			r = mid
		} else if mid < y {
			// max in right side
			l = mid + 1
		} else {
			return []int{x, y}
		}
	}
	return nil
}
