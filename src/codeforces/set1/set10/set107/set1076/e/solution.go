package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(n, edges, queries)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, n*2)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	cmds := make([][]int, n)
	for i := 0; i < n; i++ {
		cmds[i] = make([]int, 0, 1)
	}

	for i, cur := range queries {
		v := cur[0] - 1
		cmds[v] = append(cmds[v], i)
	}

	// 高度
	tr := make(SegTree, 2*n)

	var dfs func(p int, u int, depth int)

	ans := make([]int, n)

	dfs = func(p int, u int, depth int) {
		for _, i := range cmds[u] {
			qr := queries[i]
			d, x := qr[1], qr[2]
			tr.Update(depth, min(depth+d+1, n), x, n)
		}

		ans[u] = tr.Get(depth, n)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v, depth+1)
			}
		}

		for _, i := range cmds[u] {
			qr := queries[i]
			d, x := qr[1], qr[2]
			tr.Update(depth, min(depth+d+1, n), -x, n)
		}
	}

	dfs(-1, 0, 0)

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type SegTree []int

func (tr SegTree) Update(l int, r int, v int, n int) {
	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			tr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			tr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (tr SegTree) Get(p int, n int) int {
	p += n
	var res int

	for p > 0 {
		res += tr[p]
		p >>= 1
	}
	return res
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
