package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}
		res := solve(A)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A [][]int) []int {
	n := len(A)
	g := NewGraph(n, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				g.AddEdge(i, j)
			}
		}
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}

	var dfs func(u int, c int, cnt []int) bool

	dfs = func(u int, c int, cnt []int) bool {
		if color[u] >= 0 {
			return color[u] == c
		}
		color[u] = c
		cnt[c]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]

			if !dfs(v, c^1, cnt) {
				return false
			}
		}
		return true
	}

	type Pair struct {
		first  int
		second int
	}

	var comps []Pair

	for i := 0; i < n; i++ {
		if color[i] < 0 {
			cnt := []int{0, 0}
			if !dfs(i, 0, cnt) {
				return nil
			}
			comps = append(comps, Pair{i, cnt[0] - cnt[1]})
		}
	}

	m := len(comps)

	dp := make([][]bool, m)
	fp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n*2+10)
		fp[i] = make([]int, n*2+10)
	}
	dp[0][n+comps[0].second] = true
	dp[0][n-comps[0].second] = true
	fp[0][n-comps[0].second] = 1
	for i := 0; i+1 < m; i++ {
		for j := -n; j <= n; j++ {
			if dp[i][j+n] {
				dp[i+1][n+j-comps[i+1].second] = true
				dp[i+1][n+j+comps[i+1].second] = true
				fp[i+1][n+j-comps[i+1].second] = 1
			}
		}
	}

	var best int
	for !dp[m-1][n+best] {
		best++
	}

	ass := make([]int, m)

	for i := m - 1; i >= 0; i-- {
		ass[i] = fp[i][best+n]
		if ass[i] == 1 {
			best += comps[i].second
		} else {
			best -= comps[i].second
		}
	}

	vis := make([]bool, n)
	ans := make([]int, n)

	var dfs2 func(u int, c int)

	dfs2 = func(u int, c int) {
		ans[u] = c
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs2(v, c^1)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs2(comps[i].first, ass[i])
	}

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
