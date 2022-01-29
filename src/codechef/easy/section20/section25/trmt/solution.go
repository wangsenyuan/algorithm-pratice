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
		n, q := readTwoNums(reader)
		V := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, E, V, Q)
		for _, num := range res {
			buf.WriteString(fmt.Sprintf("%d\n", num))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const H = 22

func solve(n int, E [][]int, V []int, Q [][]int) []int64 {
	// sum(dist(x, k) - dist(y, k))
	// dist(x, k) - dist(y, k)
	// let k in the path x, p (p is lca(x, y))
	// then dist(x, k) = dist(x, p) - dist(k, p)
	// dist(y, k) = dist(y, p) + dist(k, p)
	// dist(x, k) - dist(y, k) = dist(x, p) - dist(y, p) - 2 * dist(k, p) when k on path (x, p)
	//                         = dist(x, p) - dist(y, p) + 2 * dist(k, p) when k on path (y, p)
	// when k on path (x, p) for one k
	// (dist(x, k) - dist(y, k)) * K = (dist(x, p) - dist(y, p) - 2 * (D(k) - D(p))) * K
	// D(x) * K - D(y) * K - 2 * D(k) * K + 2 * D(p) * K
	// D(x) - D(y) + 2 * D(p) = X
	// X * K - 2 * D(k) * K
	// when k on path (y, p), for each k
	// (dist(x, k) - dist(y, k))  * K = (dist(x, p) - dist(y, p) + 2 * (D(k) - D(p))) * K
	// (D(x) - D(y) + 2 * D(k) - 2 * D(p)) * K
	// let Y = D(x) - D(y) - 2 * D(p)
	// Y * K + 2 * D(k) * K
	g := NewGraph(n, 2*n)
	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	D := make([]int64, n)
	SK := make([]int64, n)
	DK := make([]int64, n)

	L := make([]int, n)

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

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			L[v] = L[u] + 1
			D[v] = D[u] + int64(g.weight[i])
			SK[v] = SK[u] + int64(V[v])
			DK[v] = DK[u] + int64(V[v])*D[v]
			dfs(u, v)
		}
	}

	SK[0] = int64(V[0])
	dfs(0, 0)

	lca := func(u, v int) int {
		if L[u] < L[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if L[u]-(1<<uint(i)) >= L[v] {
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

	get := func(x, y int) int64 {
		p := lca(x, y)
		var res1, res2 int64
		if x != p {
			res1 = (SK[x]-SK[p])*(D[x]-D[y]+2*D[p]) - 2*(DK[x]-DK[p])
		}
		if y != p {
			res2 = (SK[y]-SK[p])*(D[x]-D[y]-2*D[p]) + 2*(DK[y]-DK[p])
		}
		res3 := (D[x] - D[y]) * int64(V[p])
		return res1 + res2 + res3
	}

	ans := make([]int64, len(Q))

	for i, cur := range Q {
		x, y := cur[0], cur[1]
		x--
		y--
		ans[i] = get(x, y)
	}

	return ans
}

type Graph struct {
	nodes  []int
	next   []int
	to     []int
	weight []int
	cur    int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+2)
	g.to = make([]int, e+2)
	g.weight = make([]int, e+2)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.weight[g.cur] = w
}
