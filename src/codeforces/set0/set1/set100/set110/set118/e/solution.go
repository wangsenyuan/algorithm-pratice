package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, _, res := process(reader)

	if len(res) == 0 {
		fmt.Println(0)
		return
	}
	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (n int, edges [][]int, res [][]int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	g := NewGraph(n+1, 2*len(edges))
	for i, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	dsc := make([]int, n+1)
	low := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dsc[i] = -1
		low[i] = -1
	}

	vis := make([]bool, n+1)
	var timer int

	var res [][]int

	var dfs1 func(u int) bool

	marked := make([]bool, len(edges))

	dfs1 = func(u int) bool {
		dsc[u] = timer
		low[u] = timer
		timer++
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			j := g.val[i]
			if marked[j] {
				continue
			}
			marked[j] = true
			res = append(res, []int{u, v})
			if !vis[v] {
				if !dfs1(v) || low[v] > dsc[u] {
					// 这是一条关键边
					return false
				}
				low[u] = min(low[u], low[v])
			} else {
				// 这是一条回边
				low[u] = min(low[u], dsc[v])
			}
		}
		return true
	}
	if !dfs1(1) {
		return nil
	}
	// 那剩下就是给每条边选择一个方向就可以了？

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
