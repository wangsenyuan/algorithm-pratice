package p1969

/**
给你一个下标从 1开始的二进制矩阵，其中0表示陆地，1表示水域。同时给你row 和col分别表示矩阵中行和列的数目。

一开始在第0天，整个矩阵都是陆地。但每一天都会有一块新陆地被水淹没变成水域。给你一个下标从1开始的二维数组cells，其中cells[i] = [ri, ci]表示在第i天，第ri行ci列（下标都是从 1开始）的陆地会变成 水域（也就是 0变成 1）。

你想知道从矩阵最 上面一行走到最 下面一行，且只经过陆地格子的 最后一天是哪一天。你可以从最上面一行的任意格子出发，到达最下面一行的任意格子。你只能沿着四个基本方向移动（也就是上下左右）。

请返回只经过陆地格子能从最 上面一行走到最 下面一行的 最后一天。


cells.length == row * col

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-day-where-you-can-still-cross
*/
func latestDayToCross(row int, col int, cells [][]int) int {
	// a union-find problem
	arr := make([]int, (row+2)*(col+2))
	cnt := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arr[i] = i
		cnt[i] = 1
	}

	var find func(p int) int

	find = func(p int) int {
		if arr[p] != p {
			arr[p] = find(arr[p])
		}
		return arr[p]
	}

	union := func(a, b int) {
		pa := find(a)
		pb := find(b)
		if pa == pb {
			return
		}
		if cnt[pa] < cnt[pb] {
			pa, pb = pb, pa
		}
		cnt[pa] += cnt[pb]
		arr[pb] = pa
	}
	fill := make([][]bool, row+2)
	for i := 0; i < len(fill); i++ {
		fill[i] = make([]bool, col+2)
	}

	for i := 0; i+1 < row+2; i++ {
		fill[i][0] = true
		fill[i+1][0] = true
		union(i*(col+2), (i+1)*(col+2))
		fill[i][col+1] = true
		fill[i+1][col+1] = true
		union(i*(col+2)+col+1, (i+1)*(col+2)+col+1)
	}
	X := 0
	Y := col + 1
	for i, cell := range cells {
		r, c := cell[0], cell[1]

		for a := -1; a <= 1; a++ {
			for b := -1; b <= 1; b++ {
				if a == 0 && b == 0 {
					continue
				}
				x, y := r+a, c+b
				if fill[x][y] {
					union(x*(col+2)+y, r*(col+2)+c)
					if find(X) == find(Y) {
						return i
					}
				}
			}
		}
		fill[r][c] = true
	}
	return row * col
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
