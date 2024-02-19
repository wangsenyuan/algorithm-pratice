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
		a := readNNums(reader, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		res := solve(n, edges, a)
		s := fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]
		buf.WriteString(fmt.Sprintf("%s\n", s))
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 62

func solve(n int, edges [][]int, a []int) []int {
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	ans := make([]int, n)
	sz := make([]int, n)

	dp := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				dfs(u, v)
				sz[u] += sz[v]
				dp[u] += dp[v]
				dp[u] += sz[v] * (a[u] ^ a[v])
			}
		}
		sz[u]++
	}

	dfs(-1, 0)

	ans[0] = dp[0]

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		if p >= 0 {
			dp[u] += dp[p]
			dp[u] += sz[p] * (a[p] ^ a[u])
			ans[u] = dp[u]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sz[u] = n - sz[v]
				ou := dp[u]
				dp[u] -= dp[v]
				dp[u] -= sz[v] * (a[u] ^ a[v])
				dfs2(u, v)

				dp[u] = ou
				sz[u] = n
			}
		}
	}

	dfs2(-1, 0)

	return ans
}

func solve1(n int, edges [][]int, a []int) []int {
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	ans := make([]int, n)

	// dp[i][x] = 将i子树中的都设置位0/1时的最小操作数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	sz := make([]int, n)

	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		for j := 0; j < 2; j++ {
			dp[u][j] = 0
		}
		sz[u] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, d)
				sz[u] += sz[v]
				for j := 0; j < 2; j++ {
					dp[u][j] += dp[v][j]
				}
			}
		}
		x := (a[u] >> d) & 1
		sz[u]++
		dp[u][1^x] = dp[u][x] + sz[u]
	}

	var dfs2 func(p int, u int, d int)

	dfs2 = func(p int, u int, d int) {
		x := (a[u] >> d) & 1
		sz[u] = n

		if p >= 0 {
			dp[u][x] += dp[p][x]
			dp[u][1^x] = dp[u][x] + sz[u]
			ans[u] += min(dp[u][0], dp[u][1]) * (1 << d)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				o0, o1, o2 := dp[u][0], dp[u][1], sz[u]
				sz[u] -= sz[v]
				dp[u][x] -= dp[v][x]
				dp[u][1^x] = dp[u][x] + sz[u]

				dfs2(u, v, d)

				sz[u] = o2
				dp[u][0], dp[u][1] = o0, o1
			}
		}
	}

	process := func(d int) {
		dfs(-1, 0, d)
		ans[0] += min(dp[0][0], dp[0][1]) * (1 << d)
		dfs2(-1, 0, d)
	}

	for d := 0; d < 20; d++ {
		process(d)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
