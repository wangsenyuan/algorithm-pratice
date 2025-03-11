package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur.first, cur.second))
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

func process(reader *bufio.Reader) []pair {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) []pair {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)

	}

	sz := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sz[u]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	var res []pair

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		dp := make([]bool, sz[u]+1)
		dp[0] = true
		var cur int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs2(u, v)
				cur += sz[v]
				for j := cur; j >= sz[v]; j-- {
					if !dp[j] {
						dp[j] = dp[j-sz[v]]
					}
				}
			}
		}

		// u不涂色的情况下
		for j := cur; j > 0; j-- {
			if dp[j] && n-1-j > 0 {
				// j表示一个完全包含一个颜色的分组
				res = append(res, pair{j, n - 1 - j})
				res = append(res, pair{n - 1 - j, j})
			}
		}
	}

	dfs2(-1, 0)

	slices.SortFunc(res, func(a, b pair) int {
		return a.first - b.first
	})

	// remove duplicates
	var it int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] != res[i-1] {
			res[it] = res[i-1]
			it++
		}
	}
	return res[:it]
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
