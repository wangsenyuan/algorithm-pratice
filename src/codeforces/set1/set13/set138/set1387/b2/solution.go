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

	edges := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res, move := solve(n, edges)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", res))
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", move[i]))
	}
	buf.WriteByte('\n')

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

const inf = 1 << 30

func solve(n int, edges [][]int) (int, []int) {
	g := NewGraph(n+1, 2*len(edges))

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n+1)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
		sz[u]++
	}

	dfs(0, 1)

	var get_result func(p int, u int) int

	get_result = func(p int, u int) int {
		var res int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				res += min(sz[v], n-sz[v])
				res += get_result(u, v)
			}
		}
		return res
	}

	sum := get_result(0, 1)

	var find_centroid func(p int, u int) int

	find_centroid = func(p int, u int) int {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && sz[v]*2 >= n {
				return find_centroid(u, v)
			}
		}
		return u
	}

	c := find_centroid(0, 1)

	var ord []int

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		ord = append(ord, u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
			}
		}
	}

	dfs2(0, c)

	ans := make([]int, n+1)

	for i, u := range ord {
		ans[u] = ord[(i+n/2)%n]
	}

	return sum * 2, ans[1:]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}