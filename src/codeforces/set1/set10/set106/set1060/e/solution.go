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

	res := solve(n, edges)

	fmt.Println(res)
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

func solve(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	var res int

	// res = sum of (dist(u, v) + 1) / 2
	freq := make([]int, 2)
	var dfs func(p int, u int, d int) int

	dfs = func(p int, u int, d int) int {
		freq[d]++
		sz := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				sz += dfs(u, v, d^1)
			}
		}
		res += sz * (n - sz)
		return sz
	}

	dfs(-1, 0, 0)

	res = (res + freq[0]*freq[1]) / 2

	return res
}

func solve1(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	dp := make([]int, n)
	sz := make([][]int, n)
	var dfs func(u int, v int)

	dfs = func(p int, u int) {
		sz[u] = make([]int, 2)
		sz[u][0]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				for j := 0; j < 2; j++ {
					sz[u][j] += sz[v][1^j]
				}
				dp[u] += dp[v] + sz[v][0]
			}
		}
	}

	dfs(-1, 0)

	var res int

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		if p >= 0 {
			dp[u] += dp[p] + sz[p][0]
			for j := 0; j < 2; j++ {
				sz[u][j] += sz[p][1^j]
			}
		}
		// 所有的距离 / 2， 但是奇数条边的节点+1
		res += dp[u]
		// u is the new root
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				du := dp[u]
				dp[u] -= dp[v] + sz[v][0]
				s0, s1 := sz[u][0], sz[u][1]
				for j := 0; j < 2; j++ {
					sz[u][j] -= sz[v][1^j]
				}

				dfs2(u, v)
				// restore
				dp[u] = du
				sz[u][0] = s0
				sz[u][1] = s1
			}
		}
	}

	dfs2(-1, 0)

	return res / 2
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
