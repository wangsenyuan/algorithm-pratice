package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, C, D := readThreeNums(reader)
	S := readNNums(reader, N)
	E := make([][]int, N-1)
	for i := 0; i < N-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	profit, steps := solve(N, C, D, E, S)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", profit))
	for _, step := range steps {
		buf.WriteString(step)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(N int, C int, D int, E [][]int, S []int) (int, []string) {
	g := NewGraph(N+1, 2*N)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	child := make([]int, N+1)
	sibling := make([]int, N+1)

	dp := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, C+1)
		for j := 0; j <= C; j++ {
			dp[i][j] = make([]int, D+1)
			for d := 0; d <= D; d++ {
				dp[i][j][d] = -1
			}
		}
	}

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sibling[v] = child[u]
				child[u] = v
				dfs1(u, v)
			}
		}
	}

	dfs1(0, 1)

	var dfs2 func(p, u int, c int, d int) int

	dfs2 = func(p, u int, c int, d int) int {
		if u == 0 {
			if c > 0 && d > 0 {
				return max(0, S[p-1])
			}
			return 0
		}
		if dp[u][c][d] < 0 {
			// no chef for u
			dp[u][c][d] = dfs2(p, sibling[u], c, d)
			// transfer c1 d1 to node u
			for c1 := 1; c1 <= c; c1++ {
				for d1 := 0; d1 < d; d1++ {
					dp[u][c][d] = max(dp[u][c][d], dfs2(u, child[u], c1, d1)+dfs2(p, sibling[u], c-c1, d-d1-1))
				}
			}
		}
		return dp[u][c][d]
	}

	best := dfs2(1, child[1], C, D)
	var steps []string
	var dfs3 func(p int, u int, c int, d int)
	dfs3 = func(p int, u int, c int, d int) {
		if u == 0 {
			if c > 0 && d > 0 && S[p-1] > 0 {
				steps = append(steps, fmt.Sprintf("build %d", p))
			}
			return
		}
		cur := dfs2(p, u, c, d)
		if cur == dfs2(p, sibling[u], c, d) {
			// no transfer to child of u
			dfs3(p, sibling[u], c, d)
			return
		}
		for c1 := 1; c1 <= c; c1++ {
			for d1 := 0; d1 < d; d1++ {
				tmp := dfs2(u, child[u], c1, d1) + dfs2(p, sibling[u], c-c1, d-d1-1)
				if tmp == cur {
					// transfer to child c1 & d1
					steps = append(steps, fmt.Sprintf("transfer %d %d %d", p, u, c1))
					dfs3(u, child[u], c1, d1)
					dfs3(p, sibling[u], c-c1, d-d1-1)
					return
				}
			}
		}
	}

	dfs3(1, child[1], C, D)

	for len(steps) < D {
		steps = append(steps, "nothing")
	}

	return best, steps
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}
