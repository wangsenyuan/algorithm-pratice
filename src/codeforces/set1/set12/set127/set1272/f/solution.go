package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readString(reader)
	b := readString(reader)
	res := solve(a, b)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func solve(x string, y string) string {
	n := len(x)
	m := len(y)
	// level 肯定不能无限加，应该是两个最大的level

	k := max(n, m)

	// 弄个queue，做bfs
	type item struct {
		i int16
		j int16
		u int16
	}

	var inf = 2 * int16(n+m)

	dp := make([][][]int16, n+1)
	from := make([][][]item, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int16, m+1)
		from[i] = make([][]item, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int16, k+1)
			from[i][j] = make([]item, k+1)
			for u := 0; u <= k; u++ {
				dp[i][j][u] = inf
				from[i][j][u] = item{-1, -1, -1}
			}
		}
	}
	dp[0][0][0] = 0

	get := func(c byte) int16 {
		if c == '(' {
			return 1
		}
		return -1
	}

	var que []item
	que = append(que, item{0, 0, 0})

	var tail int
	for tail < len(que) {
		i, j, u := que[tail].i, que[tail].j, que[tail].u
		tail++

		for _, du := range []int16{1, -1} {
			ni, nj, nu := i, j, u+du
			if nu >= 0 && nu <= int16(k) {
				if ni < int16(n) && get(x[ni]) == du {
					ni++
				}
				if nj < int16(m) && get(y[nj]) == du {
					nj++
				}
				if dp[ni][nj][nu] == inf {
					dp[ni][nj][nu] = dp[i][j][u] + 1
					from[ni][nj][nu] = item{i, j, du}
					que = append(que, item{ni, nj, nu})
				}
			}
		}
	}

	ln := dp[n][m][0]

	buf := make([]byte, ln)

	for i, j, u := n, m, 0; ln > 0; ln-- {
		cur := from[i][j][u]
		if cur.u < 0 {
			buf[ln-1] = ')'
		} else {
			buf[ln-1] = '('
		}
		i = int(cur.i)
		j = int(cur.j)
		u -= int(cur.u)
	}

	return string(buf)
}
