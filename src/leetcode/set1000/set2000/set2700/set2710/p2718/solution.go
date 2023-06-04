package p2718

func matrixSumQueries(n int, queries [][]int) int64 {
	row := make([]int, n)
	col := make([]int, n)

	for i, cur := range queries {
		if cur[0] == 0 {
			row[cur[1]] = i
		} else {
			col[cur[1]] = i
		}
	}
	var row_sum int64
	var col_sum int64
	var res int64

	for i, cur := range queries {
		if cur[0] == 0 {
			if row[cur[1]] == i {
				res -= col_sum
				res += int64(n) * int64(cur[2])
				row_sum += int64(cur[2])
			}
		} else {
			if col[cur[1]] == i {
				res -= row_sum
				res += int64(n) * int64(cur[2])
				col_sum += int64(cur[2])
			}
		}
	}
	return res
}
