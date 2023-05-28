package p2711

import "sort"

func maxIncreasingCells(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	type Item struct {
		x int
		y int
		v int
	}

	items := make([]Item, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			items = append(items, Item{i, j, mat[i][j]})
		}
	}

	row := make([]int, m)
	col := make([]int, n)

	sort.Slice(items, func(i, j int) bool {
		return items[i].v < items[j].v
	})

	buf := make([]int, m*n)

	for i := 0; i < len(items); {
		j := i
		for i < len(items) && items[i].v == items[j].v {
			it := items[i]
			tmp := max(row[it.x], col[it.y])
			buf[i-j] = tmp + 1
			i++
		}

		for k := j; k < i; k++ {
			it := items[k]
			x, y := it.x, it.y
			v := buf[k-j]
			row[x] = max(row[x], v)
			col[y] = max(col[y], v)
		}
	}
	var res int

	for i := 0; i < m; i++ {
		res = max(res, row[i])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}
