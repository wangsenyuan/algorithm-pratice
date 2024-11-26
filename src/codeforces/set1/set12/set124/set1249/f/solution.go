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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(k, a, edges)
}

const inf = 1 << 60

func solve(k int, a []int, edges [][]int) int {
	n := len(a)

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+2)
	}

	fp := make([]int, k+2)

	merge := func(u int, v int) {
		copy(fp, dp[u])
		for i := 0; i <= k; i++ {
			// 如果只有子树v
			fp[i+1] = max(fp[i+1], dp[v][i])
		}
		for i := 1; i <= k; i++ {
			for j := k - i; j <= k; j++ {
				fp[min(i, j+1)] = max(fp[min(i, j+1)], dp[u][i]+dp[v][j])
			}
		}
		copy(dp[u], fp)
		for i := k; i > 0; i-- {
			dp[u][i] = max(dp[u][i], dp[u][i+1])
		}
	}

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		var sum int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sum += dp[v][k]
				merge(u, v)
			}
		}
		dp[u][0] = sum + a[u-1]
		dp[u][k+1] = sum
		for i := k; i >= 0; i-- {
			dp[u][i] = max(dp[u][i], dp[u][i+1])
		}
	}

	dfs(0, 1)

	return dp[1][0]
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
