package p1284

func minFlips(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	bak := make([][]int, m)
	for i := 0; i < m; i++ {
		bak[i] = make([]int, n)
	}

	N := 1 << uint(m*n)

	que := make([]int, N)
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = -1
	}
	x := toNum(mat)
	dist[x] = 0
	var end int
	que[end] = x
	end++
	var front int

	for front < end {
		cur := que[front]
		front++

		nexts := flip(cur, bak)

		for _, next := range nexts {
			if dist[next] < 0 {
				dist[next] = dist[cur] + 1
				que[end] = next
				end++
			}
		}

	}

	return dist[0]
}

func flip(num int, mat [][]int) []int {
	// try every position
	toMat(num, mat)
	var res []int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			flipAt(mat, i, j)

			res = append(res, toNum(mat))

			flipAt(mat, i, j)
		}
	}

	return res
}

func flipAt(mat [][]int, i, j int) {
	flip := func(x, y int) {
		if x < 0 || y < 0 || x == len(mat) || y == len(mat[0]) {
			return
		}
		mat[x][y] = 1 - mat[x][y]
	}

	flip(i-1, j)
	flip(i, j+1)
	flip(i+1, j)
	flip(i, j-1)
	flip(i, j)
}

func toMat(num int, mat [][]int) {
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[i]) - 1; j >= 0; j-- {
			mat[i][j] = num & 1
			num >>= 1
		}
	}
}

func toNum(mat [][]int) int {
	var res int

	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[i]) - 1; j >= 0; j-- {
			res = (res << 1) | mat[i][j]
		}
	}

	return res
}
