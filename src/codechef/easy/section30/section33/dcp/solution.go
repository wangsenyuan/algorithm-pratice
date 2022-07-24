package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	Q := readNNums(reader, n)
	E := make([][]int, 2*m)

	for i := 0; i < 2*m; i += 2 {
		E[i] = readNNums(reader, 2)
		x := E[i][1]
		E[i+1] = readNNums(reader, 2*x)
	}

	res := solve(n, m, Q, E)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func solve(n int, m int, Q []int, E [][]int) []int {
	var edge_count int
	for i := 0; i < len(E); i += 2 {
		edge_count += E[i][1]
	}

	g := NewGraph(n, edge_count)
	for i := 0; i < len(E); i += 2 {
		c0 := E[i][0]
		c0--
		for j := 0; j < len(E[i+1]); j += 2 {
			w, c := E[i+1][j], E[i+1][j+1]
			c--
			g.AddEdge(c0, c, w)
		}
	}

	var ord []int

	var dfs func(u int)

	vis := make([]bool, n)

	dfs = func(u int) {
		vis[u] = true
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
			}
		}
		ord = append(ord, u)
	}

	R := make([]int, n)

	for u := 0; u < n; u++ {
		R[u] = Q[u]
		if !vis[u] {
			dfs(u)
		}
	}

	for i := n - 1; i >= 0; i-- {
		u := ord[i]
		if g.node[u] > 0 {
			for j := g.node[u]; j > 0; j = g.next[j] {
				v := g.to[j]
				w := g.cost[j]
				R[v] = modAdd(R[v], modMul(R[u], w))
			}
			R[u] = 0
		}
	}
	return R
}

type Graph struct {
	node []int
	next []int
	to   []int
	cost []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	node := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	cost := make([]int, e+1)
	g := Graph{node, next, to, cost, 0}
	return &g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = w
}
