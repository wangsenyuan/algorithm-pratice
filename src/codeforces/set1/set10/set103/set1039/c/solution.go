package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m, k := readThreeNums(reader)
	c := readNNums(reader, n)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(k, c, edges)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(k int, c []int, edges [][]int) int {
	n := len(c)

	res := mul(pow(2, n), pow(2, k))

	groupBy := make(map[int][]int)

	for i, cur := range edges {
		u, v := cur[0]-1, cur[1]-1
		x := c[u] ^ c[v]
		if len(groupBy[x]) == 0 {
			groupBy[x] = make([]int, 0, 1)
		}
		groupBy[x] = append(groupBy[x], i)
	}

	m := len(edges)
	g := NewGraph(n, 2*m)

	marked := make([]bool, n)

	var dfs func(u int)
	dfs = func(u int) {
		if marked[u] {
			return
		}
		marked[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
		}
	}

	for _, vs := range groupBy {
		nodes := make(map[int]bool)

		for _, i := range vs {
			edge := edges[i]
			u, v := edge[0]-1, edge[1]-1
			g.AddEdge(u, v)
			g.AddEdge(v, u)
			nodes[u] = true
			nodes[v] = true
		}
		sz := len(nodes)
		var w int
		// 确实有点问题
		for u := range nodes {
			if !marked[u] {
				dfs(u)
				w++
			}
		}
		// 所有的方案数，选中使用x或者不使用, pow(2, sz)
		// 其中对于同一个连通块，如果全部选中或者全部不选中，是安全的
		tmp := sub(pow(2, sz), pow(2, w))
		tmp = mul(tmp, pow(2, n-sz))
		res = sub(res, tmp)

		for u := range nodes {
			marked[u] = false
			g.nodes[u] = 0
		}
		g.cur = 0
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
	e++
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
