package p1614

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m := len(rowSum)
	n := len(colSum)
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i, j := 0, 0; j < n; j++ {
		num := colSum[j]
		var tmp int
		for i < m && tmp+rowSum[i] <= num {
			res[i][j] = rowSum[i]
			tmp += rowSum[i]
			i++
		}
		if tmp < num {
			res[i][j] = num - tmp
			rowSum[i] -= res[i][j]
		}
	}

	return res
}
