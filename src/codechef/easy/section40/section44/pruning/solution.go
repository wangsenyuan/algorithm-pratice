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
		n, d := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		res := solve(d, n, E)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(d int, n int, E [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	var dfs0 func(p int, u int)

	in := make([]int, n)
	out := make([]int, n)

	var time int

	dfs0 = func(p int, u int) {
		in[u] = time
		time++

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs0(u, v)
			}
		}

		out[u] = time
	}

	dfs0(-1, 0)

	isAnc := func(u int, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	que := make([]int, n)
	bfs := func(c int) {
		for i := 0; i < n; i++ {
			dist[c][i] = -1
		}
		dist[c][c] = 0
		var front, end int
		que[end] = c
		end++
		for front < end {
			u := que[front]
			front++
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[c][v] < 0 {
					dist[c][v] = dist[c][u] + g.val[i]
					que[end] = v
					end++
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		bfs(i)
	}

	K := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}

		for c := 0; c < n; c++ {
			// c is ancestor or u ( or c == u)
			for i := g.node[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if p != v {
					if isAnc(v, c) {
						// c is sub-tree of v
						if dist[c][u] <= d {
							dp[u][c] += dp[v][c] + g.val[i]
						} else {
							dp[u][c] += dp[v][c]
						}
						continue
					}
					if dist[c][v] <= d {
						dp[u][c] += max(K[v], dp[v][c]+g.val[i])
					} else {
						dp[u][c] += K[v]
					}
				}
			}
			K[u] = max(K[u], dp[u][c])
		}
	}

	dfs(-1, 0)

	return K[0]
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
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
