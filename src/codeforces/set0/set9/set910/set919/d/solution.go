package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	s := readString(reader)

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges, s)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func solve(n int, edges [][]int, s string) int {
	g := NewGraph(n, len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
	}

	dp := make([][]int, n)
	vis := make([]int, n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		if vis[u] == 2 {
			// already processed
			return true
		}
		if vis[u] == 1 {
			// a loop
			return false
		}
		dp[u] = make([]int, 26)

		vis[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !dfs(v) {
				return false
			}
			for j := 0; j < 26; j++ {
				dp[u][j] = max(dp[u][j], dp[v][j])
			}
		}
		vis[u]++
		dp[u][int(s[u]-'a')]++
		return true
	}
	var ans int
	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			if !dfs(i) {
				return -1
			}
		}
		for j := 0; j < 26; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	return ans
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
