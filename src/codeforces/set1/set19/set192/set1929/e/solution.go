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
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		m := readNum(reader)
		pairs := make([][]int, m)
		for i := 0; i < m; i++ {
			pairs[i] = readNNums(reader, 2)
		}
		res := solve(n, edges, pairs)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const H = 20

func solve(n int, edges [][]int, pairs [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	P := make([][]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}

		return P[u][0]
	}

	kth := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = P[u][i]
				k -= 1 << i
			}
		}
		return u
	}

	k := len(pairs)
	K := 1 << k
	dp := make([]int, K)

	for i := 0; i < K; i++ {
		dp[i] = k
	}

	check := func(u int, v int, a int, b int) bool {
		// D[a] >= D[v]
		p := lca(a, b)
		if D[p] >= D[v] {
			return false
		}
		// D[p] < D[v] (D[v] = D[u] + 1)
		x := kth(u, D[u]-D[p])
		if x != p {
			return false
		}
		if D[a] >= D[v] {
			y := kth(a, D[a]-D[v])
			if y == v {
				return true
			}
		}
		if D[b] >= D[v] {
			y := kth(b, D[b]-D[v])
			if y == v {
				return true
			}
		}
		return false
	}

	// len(vals) is small enough
	vals := make(map[int]int)

	process := func(u int, v int) {
		var flag int
		for i, cur := range pairs {
			a, b := cur[0]-1, cur[1]-1
			// u, v 是否在a, b 的路径上
			if check(u, v, a, b) {
				flag |= 1 << i
			}
		}
		if flag > 0 {
			vals[flag]++
			dp[flag] = 1
		}
	}

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		if p >= 0 {
			process(p, u)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
			}
		}
	}

	dfs2(-1, 0)

	vs := make([]int, 0, len(vals))

	for v := range vals {
		vs = append(vs, v)
	}

	for state := 1; state < K; state++ {
		for _, s := range vs {
			dp[state|s] = min(dp[state|s], 1+dp[state])
		}
	}
	return dp[K-1]
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
