package p2133

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		vis := make(map[int]bool)
		for j := 0; j < n; j++ {
			if matrix[i][j] <= 0 || matrix[i][j] > n {
				return false
			}
			vis[matrix[i][j]] = true
		}
		if len(vis) != n {
			return false
		}
	}

	for j := 0; j < n; j++ {
		vis := make(map[int]bool)
		for i := 0; i < n; i++ {
			if matrix[i][j] <= 0 || matrix[i][j] > n {
				return false
			}
			vis[matrix[i][j]] = true
		}
		if len(vis) != n {
			return false
		}
	}
	return true
}
