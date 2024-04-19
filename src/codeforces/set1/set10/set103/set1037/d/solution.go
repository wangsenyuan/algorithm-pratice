package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}

	ord := readNNums(reader, n)

	res := solve(n, edges, ord)

	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

func solve(n int, edges [][]int, ord []int) bool {
	if ord[0] != 1 {
		return false
	}

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	pos := make([]int, n+1)
	for i, j := range ord {
		pos[j] = i
	}

	level := make([][]int, n)

	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		level[d] = append(level[d], u)

		var children []int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				children = append(children, v)
			}
		}

		sort.Slice(children, func(i, j int) bool {
			return pos[children[i]] < pos[children[j]]
		})

		for _, v := range children {
			dfs(u, v, d+1)
		}
	}

	dfs(0, 1, 0)

	res := make([]int, 0, n)

	for d := 0; d < n && len(level[d]) > 0; d++ {
		res = append(res, level[d]...)
	}

	for i := 0; i < n; i++ {
		if res[i] != ord[i] {
			return false
		}
	}
	return true
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
