package p1074

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])

	sum := make([][]int, m)

	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)

	}

	count := make([][]map[int]int, n)

	for i := 0; i < n; i++ {
		count[i] = make([]map[int]int, n)
		for j := 0; j < n; j++ {
			count[i][j] = make(map[int]int)
			count[i][j][0] = 1
		}
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i][j] = matrix[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
			s := sum[i][j]
			for k := j; k >= 0; k-- {
				var s1 int
				if k > 0 {
					s1 = sum[i][k-1]
				}
				s2 := s - s1 - target
				res += count[k][j][s2]
				count[k][j][s-s1]++
			}
		}
	}

	return res
}
