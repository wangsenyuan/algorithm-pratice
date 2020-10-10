package skyscrapers

const N = 4

func SolvePuzzle(clues []int) [][]int {

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		mat[i] = make([]int, N)
	}

	checkCol := func(c int) bool {
		a, b := clues[c], clues[11-c]

		if a != 0 {
			var cnt = 1
			var x = mat[0][c]
			for i := 1; i < 4; i++ {
				if mat[i][c] > x {
					cnt++
					x = mat[i][c]
				}
			}
			if cnt != a {
				return false
			}
		}

		if b != 0 {
			var cnt = 1
			var x = mat[3][c]
			for i := 2; i >= 0; i-- {
				if mat[i][c] > x {
					cnt++
					x = mat[i][c]
				}
			}
			if cnt != b {
				return false
			}
		}
		return true
	}

	checkRow := func(r int) bool {
		b, a := clues[4+r], clues[15-r]

		if a != 0 {
			var cnt = 1
			var x = mat[r][0]
			for i := 1; i < 4; i++ {
				if mat[r][i] > x {
					cnt++
					x = mat[r][i]
				}
			}
			if cnt != a {
				return false
			}
		}

		if b != 0 {
			var cnt = 1
			var x = mat[r][3]
			for i := 2; i >= 0; i-- {
				if mat[r][i] > x {
					cnt++
					x = mat[r][i]
				}
			}
			if cnt != b {
				return false
			}
		}
		return true
	}

	var put func(i, j int, row int, col []int) bool

	put = func(i, j int, row int, col []int) bool {
		if i == 4 {
			return true
		}

		for x := 0; x < 4; x++ {
			if col[j]&(1<<uint(x)) > 0 || row&(1<<uint(x)) > 0 {
				continue
			}
			mat[i][j] = x + 1
			if i == 3 && !checkCol(j) {
				continue
			}
			if j == 3 && !checkRow(i) {
				continue
			}
			col[j] |= 1 << uint(x)
			var ni, nj, nrow int
			if j < 3 {
				ni = i
				nj = j + 1
				nrow = row | (1 << uint(x))
			} else {
				ni = i + 1
				nj = 0
				nrow = 0
			}

			ok := put(ni, nj, nrow, col)

			col[j] ^= 1 << uint(x)

			if ok {
				return true
			}
		}
		return false
	}
	col := make([]int, 4)
	put(0, 0, 0, col)
	return mat
}
