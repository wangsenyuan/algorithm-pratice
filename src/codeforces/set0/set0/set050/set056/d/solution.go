package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, _, ops, pos, chs := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(ops)))
	for i, op := range ops {
		buf.WriteString(op)
		buf.WriteString(fmt.Sprintf(" %d", pos[i]))
		if op != "DELETE" {
			buf.WriteString(fmt.Sprintf(" %c", chs[i]))
		}
		buf.WriteByte('\n')
	}

	buf.WriteTo(os.Stdout)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (s string, x string, ops []string, pos []int, chs []byte) {
	s = readString(reader)
	x = readString(reader)
	ops, pos, chs = solve(s, x)
	return
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(s string, t string) (ops []string, pos []int, chs []byte) {
	n := len(s)
	m := len(t)

	dp := make([][]int, n+1)
	fp := make([][]pair, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, m+1)
		fp[i] = make([]pair, m+1)
		for j := range m + 1 {
			dp[i][j] = inf
			fp[i][j] = pair{-1, -1}
		}
	}

	update := func(x int, y int, u int, v int, d int) {
		if dp[u][v]+d < dp[x][y] {
			dp[x][y] = dp[u][v] + d
			fp[x][y] = pair{u, v}
		}
	}

	dp[n][m] = 0
	for i := n; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			if i == n && j == m {
				continue
			}
			if i+1 <= n && j+1 <= m {
				update(i, j, i+1, j+1, check(s[i] != t[j]))
			}
			if i+1 <= n {
				update(i, j, i+1, j, 1)
			}
			if j+1 <= m {
				update(i, j, i, j+1, 1)
			}
		}
	}

	var add int
	var r, c int
	for r < n || c < m {
		nr, nc := fp[r][c].first, fp[r][c].second
		if r+1 == nr && c+1 == nc {
			if dp[r][c] == dp[nr][nc]+1 {
				// replease
				ops = append(ops, "REPLACE")
				pos = append(pos, r+add+1)
				chs = append(chs, t[c])
			}
		} else if r+1 == nr {
			ops = append(ops, "DELETE")
			pos = append(pos, r+add+1)
			add--
			chs = append(chs, '-')
		} else {
			ops = append(ops, "INSERT")
			pos = append(pos, r+add+1)
			add++
			chs = append(chs, t[c])
		}

		r = nr
		c = nc
	}

	return
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type BIT []int

func (bit BIT) update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}
