package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const LOG = 20

func solve(n int, E [][]int) int64 {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	H := make([]int, n)
	P := make([][]int, n)

	var dfs func(p int, u int) int

	adj := make([][]Pair, n)

	dfs = func(p int, u int) int {
		P[u] = make([]int, LOG)
		P[u][0] = p
		for i := 1; i < LOG; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt++
				D[v] = D[u] + 1
				H[u] = max(H[u], dfs(u, v))
			}
		}

		adj[u] = make([]Pair, cnt)
		id := 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				adj[u][id] = Pair{H[v], v}
				id++
			}
		}

		sort.Slice(adj[u], func(i, j int) bool {
			return adj[u][i].first < adj[u][j].first
		})

		H[u]++
		return H[u]
	}

	dfs(0, 0)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := LOG - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = P[u][i]
			}
		}

		if u == v {
			return u
		}

		for i := LOG - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	que := make([]int, n)
	var front, end int
	que[end] = 0
	end++
	for front < end {
		u := que[front]
		front++
		for _, x := range adj[u] {
			que[end] = x.second
			end++
		}
	}

	var res int64

	for i := 1; i < n; i++ {
		u := que[i-1]
		v := que[i]
		p := lca(u, v)
		res += int64(D[u] + D[v] - 2*D[p])
	}

	return res
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
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
