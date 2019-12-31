package p1301

const MOD = 1000000007

func pathsWithMaxScore(board []string) []int {
	n := len(board)

	cnt := make([][]int, n)
	num := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, n)
		num[i] = make([]int, n)
		vis[i] = make([]bool, n)

		for j := 0; j < n; j++ {
			num[i][j] = -1
		}
	}
	num[n-1][n-1] = 0
	cnt[n-1][n-1] = 1
	vis[n-1][n-1] = true

	update := func(x, y int, i, j int) {
		vis[x][y] = true
		tmp := num[i][j]
		if board[x][y] != 'E' {
			tmp += int(board[x][y] - '0')
		}
		if tmp > num[x][y] {
			cnt[x][y] = cnt[i][j]
			num[x][y] = tmp
		} else if tmp == num[x][y] {
			modAdd(&cnt[x][y], cnt[i][j])
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if !vis[i][j] {
				continue
			}
			if j > 0 && board[i][j-1] != 'X' {
				update(i, j-1, i, j)
			}
			if i > 0 && board[i-1][j] != 'X' {
				update(i-1, j, i, j)
			}

			if i > 0 && j > 0 && board[i-1][j-1] != 'X' {
				update(i-1, j-1, i, j)
			}
		}
	}
	if !vis[0][0] {
		return []int{0, 0}
	}

	return []int{num[0][0], cnt[0][0]}
}

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}
