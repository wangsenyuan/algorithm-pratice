package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, E, Q)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}
	fmt.Print(buf.String())
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

const H = 20

func solve(n int, E [][]int, Q [][]int) []int64 {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	down := make([]int64, n)
	sub := make([]int64, n)
	D := make([]int, n)
	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
	}
	var dfs func(p, u int)
	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		sub[u]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
				sub[u] += sub[v]
				down[u] += sub[v] + down[v]
			}
		}
	}

	dfs(0, 0)

	up := make([]int64, n)

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				up[v] = up[u] + down[u] - (sub[v] + down[v]) + int64(n) - sub[v]
				dfs2(u, v)
			}
		}
	}

	dfs2(0, 0)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<uint(i)) >= D[v] {
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

	dist := func(u, v int) int {
		return D[u] + D[v] - 2*D[lca(u, v)]
	}

	lift := func(u int, delta int) int {
		for i := H - 1; i >= 0; i-- {
			if (delta>>uint(i))&1 == 1 {
				u = P[u][i]
			}
		}
		return u
	}

	ans := make([]int64, len(Q))

	N := int64(n)

	for i, cur := range Q {
		u, v := cur[0], cur[1]
		u--
		v--
		if D[u] > D[v] {
			u, v = v, u
		}
		// D[u] < D[v]
		d := dist(u, v)
		ev := lift(v, (d-1)/2)
		eu := P[ev][0]
		//Edge (eu, ev) is removed, where u is reachable from eu and v is reachable from ev
		//node eu is parent of ev
		//sub[ev]*dist(u, ev) => distance from u to ev for each node in subtree of ev
		//dDown[ev] = sum of distances of nodes in subtree of ev from ev

		//((N-sub[ev])*dist(par, depth, eu, v) => distance of node v to node eu for each node not in subtree of ev
		//dUp[eu] + dDown[eu] => distance of eu to all nodes
		//-(sub[ev]+dDown[ev])) => distance of eu to all nodes in subtree of ev
		du := down[u] + up[u] - (sub[ev]*int64(dist(u, ev)) + +down[ev])
		dv := down[v] + up[v] - ((N-sub[ev])*int64(dist(eu, v)) + up[eu] + down[eu] - (sub[ev] + down[ev]))
		ans[i] = du + dv
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
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
