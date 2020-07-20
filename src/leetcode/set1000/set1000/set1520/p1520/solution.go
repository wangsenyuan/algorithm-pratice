package p1520

import (
	"bytes"
)

func maxNumOfSubstrings(s string) []string {
	last := make([]int, 26)
	first := make([]int, 26)
	for i := 0; i < len(last); i++ {
		last[i] = -1
		first[i] = -1
	}

	conn := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		conn[i] = make([]bool, 26)
	}

	n := len(s)

	for i := 0; i < n; i++ {
		j := ind(s[i])
		if first[j] < 0 {
			first[j] = i
		}

		for k := 0; k < 26; k++ {
			if j != k && first[j] < last[k] {
				conn[j][k] = true
			}
		}
		last[j] = i
	}

	var buf bytes.Buffer
	res := make([]string, 0, 26)

	writeResult := func(begin, end int) {
		if end < 0 {
			return
		}
		for k := begin; k <= end; k++ {
			buf.WriteByte(s[k])
		}
		res = append(res, buf.String())
		buf.Reset()
	}

	leaf := make([]bool, 26)

	for i := 0; i < 26; i++ {
		leaf[i] = true
		for j := 0; j < 26; j++ {
			if conn[i][j] {
				leaf[i] = false
			}
		}
	}

	comp := make([]int, 26)

	var dfs func(u, c int)

	dfs = func(u, c int) {
		comp[u] = c
		for v := 0; v < 26; v++ {
			if conn[u][v] && comp[v] != c {
				dfs(v, c)
			}
		}
	}

	for u := 0; u < 26; u++ {
		if first[u] < 0 || comp[u] > 0 {
			continue
		}
		dfs(u, u+1)
	}
	vis := make([]int, 26)
	var dfs2 func(u int) int
	dfs2 = func(u int) int {
		if vis[u] == 2 {
			return u
		}
		vis[u] = 2
		for v := 0; v < 26; v++ {
			if !conn[u][v] {
				continue
			}
			x := dfs2(v)
			if x >= 0 {
				return x
			}
		}
		vis[u] = 1
		return -1
	}
	que := make([]int, 26)
	bfs := func(u int) (int, int) {
		for i := 0; i < 26; i++ {
			vis[i] = 0
		}
		a, b := n, -1
		var front, end int
		que[end] = u
		vis[u] = 1
		end++
		for front < end {
			u := que[front]
			front++
			a = min(a, first[u])
			b = max(b, last[u])

			for v := 0; v < 26; v++ {
				if conn[u][v] && vis[v] == 0 {
					que[end] = v
					end++
					vis[v] = 1
				}
			}
		}
		return a, b
	}

	for c := 1; c <= 26; c++ {
		var hasLeaf bool
		var root = -1
		for u := 0; u < 26; u++ {
			if first[u] < 0 || comp[u] != c {
				continue
			}
			if leaf[u] {
				hasLeaf = true
				writeResult(first[u], last[u])
			}
			root = u
		}
		if hasLeaf || root < 0 {
			continue
		}
		for u := 0; u < 26; u++ {
			vis[u] = 0
		}
		// find the loop
		begin, end := bfs(dfs2(root))
		writeResult(begin, end)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func ind(b byte) int {
	return int(b - 'a')
}
