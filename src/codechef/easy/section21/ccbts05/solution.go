package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q, R := readThreeNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 3)
		}
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, R, E, Q)
		for i := 0; i < q; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}
	fmt.Print(buf.String())
}

const H = 20

func solve(n, R int, E [][]int, Q [][]int) []int64 {
	g := NewGraph(n, len(E)*2+10)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, e[2])
		g.AddEdge(v, u, e[2])
	}
	D := make([]int, n)
	P := make([][]int, n)

	sum := make([]int64, n)
	var dfs func(p, u int)
	dfs = func(p, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			j := P[u][i-1]
			if j >= 0 {
				P[u][i] = P[j][i-1]
			} else {
				P[u][i] = -1
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				D[v] = D[u] + 1
				sum[v] = sum[u] + int64(g.cost[i])
				dfs(u, v)
			}
		}
	}

	dfs(-1, R-1)

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		// D[u] >= D[v]
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

	get := func(u, v int) int64 {
		p := lca(u, v)
		return sum[u] - sum[p] + sum[v] - sum[p]
	}

	ans := make([]int64, len(Q))

	for i := 0; i < len(Q); i++ {
		u, v := Q[i][0], Q[i][1]
		u--
		v--
		ans[i] = get(u, v)
	}
	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cost  []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	cost := make([]int, e)
	return &Graph{nodes, next, to, cost, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = w
}
