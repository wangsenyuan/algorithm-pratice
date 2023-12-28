package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)

	for tc > 0 {
		tc--
		var n int
		fmt.Fscan(in, &n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = make([]int, 2)
			for j := 0; j < 2; j++ {
				fmt.Fscan(in, &edges[i][j])
			}
		}
		res := solve(n, edges)
		fmt.Fprintln(out, res)
	}

}

func solve(n int, edges [][]int) int {
	if n == 1 {
		return 1
	}
	// 模拟肯定会超时
	// 如果edge[i]依赖与某个edge[j], 且 i < j
	// 那么显然处理i就是浪费时间
	// 也就是这里go back就增加一次readings
	g := NewGraph(n, len(edges)*2)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	res := 1

	var dfs func(p int, u int, x, y int)

	dfs = func(p int, u int, x, y int) {
		res = max(res, x)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if g.val[i] >= y {
					dfs(u, v, x, g.val[i])
				} else {
					dfs(u, v, x+1, g.val[i])
				}
			}
		}
	}

	dfs(-1, 0, 1, 0)

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
