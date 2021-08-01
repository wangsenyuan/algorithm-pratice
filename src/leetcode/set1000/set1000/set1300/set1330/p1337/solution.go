package p1337

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	n := len(mat)
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		rows[i] = []int{i, 0}
		for j := 0; j < len(mat[i]); j++ {
			rows[i][1] += mat[i][j]
			if mat[i][j] == 0 {
				break
			}
		}
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][1] < rows[j][1] || rows[i][1] == rows[j][1] && rows[i][0] < rows[j][0]
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = rows[i][0]
	}

	return res
}
