package p1011

var dd = []int{-1, 0, 1, 0, -1}

func numEnclaves(A [][]int) int {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}

	m := len(A)
	n := len(A[0])
	vis := make([]bool, m*n)
	que := make([]int, m*n)
	var end int
	for i := 0; i < m; i++ {
		if A[i][0] == 1 {
			que[end] = i * n
			vis[i*n] = true
			end++
		}
		if n > 1 && A[i][n-1] == 1 {
			que[end] = i*n + n - 1
			vis[i*n+n-1] = true
			end++
		}
	}

	for j := 1; j < n-1; j++ {
		if A[0][j] == 1 {
			que[end] = j
			vis[j] = true
			end++
		}
		if m > 1 && A[m-1][j] == 1 {
			que[end] = (m-1)*n + j
			vis[(m-1)*n+j] = true
			end++
		}
	}

	var begin int
	for begin < end {
		cur := que[begin]
		begin++
		x := cur / n
		y := cur % n
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && A[u][v] == 1 && !vis[u*n+v] {
				vis[u*n+v] = true
				que[end] = u*n + v
				end++
			}
		}
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		vis[x*n+y] = true
		var ans int
		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if u >= 0 && u < m && v >= 0 && v < n && A[u][v] == 1 && !vis[u*n+v] {
				ans += dfs(u, v)
			}
		}
		return ans + 1
	}

	var ans int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 && !vis[i*n+j] {
				ans += dfs(i, j)
			}
		}
	}

	return ans
}
