package p059

func generateMatrix(n int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	bounds := []int{0, 0, n - 1, n - 1}

	for x := 1; x <= n*n; {
		for j := bounds[1]; j <= bounds[3]; j++ {
			mat[bounds[0]][j] = x
			x++
		}
		for i := bounds[0] + 1; i <= bounds[2]; i++ {
			mat[i][bounds[3]] = x
			x++
		}
		for j := bounds[3] - 1; j >= bounds[1]; j-- {
			mat[bounds[2]][j] = x
			x++
		}
		for i := bounds[2] - 1; i > bounds[0]; i-- {
			mat[i][bounds[1]] = x
			x++
		}
		bounds[0]++
		bounds[1]++
		bounds[2]--
		bounds[3]--
	}

	return mat
}
