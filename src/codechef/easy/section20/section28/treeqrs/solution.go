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
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 5)
	}
	res := solve(n, E, Q)

	var buf bytes.Buffer

	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d\n", num))
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

const H = 16

func solve(n int, E [][]int, Q [][]int) []int64 {

	g := NewGraph(n, 2*(n-1))

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	P := make([][]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
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

	find := func(n int64, m int64, a int64, b int64) int64 {
		var lo, hi int64 = 1, n
		for lo < hi {
			mid := (lo + hi) / 2
			if a*mid >= b*(n+m+1-mid) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		lo--
		return lo*(lo+1)/2*a + m*(m+1)/2*a + (n+m-lo)*(n+m-lo+1)/2*b - m*(m+1)/2*b
	}

	dist := func(u, v int) int64 {
		return int64(D[u] + D[v] - 2*D[lca(u, v)])
	}

	process := func(r int, a int64, b int64, x int, y int) int64 {
		dx := dist(r, x)
		dy := dist(r, y)
		tmp := dist(x, y)
		tmp = dx + dy - tmp
		tmp /= 2
		dx -= tmp
		dy -= tmp
		ans := a * (dx*(dx+1)/2 + dy*(dy+1)/2)
		ans = min(ans, find(dx, dy, a, b))
		ans = min(ans, find(dy, dx, a, b))
		return ans
	}

	ans := make([]int64, len(Q))

	for i, cur := range Q {
		r, a, b, x, y := cur[0], cur[1], cur[2], cur[3], cur[4]
		r--
		x--
		y--
		ans[i] = process(r, int64(a), int64(b), x, y)
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
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
