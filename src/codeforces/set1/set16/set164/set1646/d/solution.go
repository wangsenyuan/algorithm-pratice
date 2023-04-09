package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}

	weight, sum, assign := solve(n, E)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d %d\n", weight, sum))

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", assign[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, E [][]int) (int, int, []int) {
	if n == 2 {
		return 2, 2, []int{1, 1}
	}
	g := NewGraph(n, 2*n)

	deg := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = []int{0, 0}
		}
	}
	// dp[u][c]
	var dfs func(p int, u int)
	// when u is 0, how to choose the edge i to v
	choice := make([]int, g.cur+1)

	dfs = func(p int, u int) {

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			dfs(u, v)
		}
		// dp[u][0][0] is the weights
		// when c = 0, no need to make u good, so assign u to 1
		//      c = 1, assign u to deg[u]
		dp[u][0][0] = 0
		dp[u][0][1] = 1
		dp[u][1][0] = 1
		dp[u][1][1] = deg[u]

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			choice[i] = 0

			if dp[v][0][0] < dp[v][1][0] ||
				dp[v][0][0] == dp[v][1][0] && dp[v][1][1] < dp[v][0][1] {
				choice[i] = 1
			}

			dp[u][0][0] += dp[v][choice[i]][0]
			dp[u][0][1] += dp[v][choice[i]][1]

			dp[u][1][0] += dp[v][0][0]
			dp[u][1][1] += dp[v][0][1]
		}
	}

	dfs(-1, 0)
	assign := make([]int, n)
	var dfs2 func(p int, u int, c int)
	dfs2 = func(p int, u int, c int) {
		if c == 0 {
			assign[u] = 1
		} else {
			assign[u] = deg[u]
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if c == 1 {
				dfs2(u, v, 0)
			} else {
				dfs2(u, v, choice[i])
			}
		}
	}

	weight := dp[0][0][0]
	sum := dp[0][0][1]

	c := 0

	if weight < dp[0][1][0] || weight == dp[0][1][0] && sum > dp[0][1][1] {
		weight = dp[0][1][0]
		sum = dp[0][1][1]
		c = 1
	}

	dfs2(-1, 0, c)

	return weight, sum, assign
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
