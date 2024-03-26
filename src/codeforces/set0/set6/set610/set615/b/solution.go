package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges)

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

func solve(n int, edges [][]int) int {
	return solve1(n, edges, func(a, b int) bool {
		return a < b
	})
}

func solve1(n int, edges [][]int, cmp func(int, int) bool) int {
	g := NewGraph(n, len(edges))

	deg := make([]int, n)
	deg2 := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		if cmp(u, v) {
			g.AddEdge(u, v)
			deg[v]++
		} else {
			g.AddEdge(v, u)
			deg[u]++
		}
		deg2[u]++
		deg2[v]++
	}

	que := make([]int, n)
	var front, tail int
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			dp[i] = 1
			que[front] = i
			front++
		}
	}
	var best int
	for tail < front {
		u := que[tail]
		tail++
		best = max(best, dp[u]*deg2[u])

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			deg[v]--
			dp[v] = max(dp[v], dp[u]+1)
			if deg[v] == 0 {
				que[front] = v
				front++
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
