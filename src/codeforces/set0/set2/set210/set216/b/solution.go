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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(n int, enemies [][]int) int {
	g := NewGraph(n, 2*len(enemies))

	for _, e := range enemies {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	color := make([]int, n)
	bad := make([]bool, n)

	for i := 0; i < n; i++ {
		color[i] = -1
	}

	cnt := make([]int, 2)

	var dfs func(u int, c int)
	dfs = func(u int, c int) {
		if color[u] >= 0 {
			if c != color[u] {
				bad[u] = true
			}
			return
		}

		color[u] = c
		cnt[c]++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v, c^1)
		}

		if bad[u] {
			cnt[c]--
		}
	}

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -n
	}
	dp[0] = 0
	fp := make([]int, n+1)

	for i := 0; i < n; i++ {
		if color[i] < 0 {
			clear(cnt)
			dfs(i, 0)
			x, y := cnt[0], cnt[1]

			for j := 0; j <= n; j++ {
				fp[j] = -n
			}

			for j := 0; j+x <= n; j++ {
				fp[j+x] = max(fp[j+x], dp[j]+y)
			}
			for j := 0; j+y <= n; j++ {
				fp[j+y] = max(fp[j+y], dp[j]+x)
			}
			copy(dp, fp)
		}
	}

	var best int
	for j := 0; j <= n; j++ {
		best = max(best, min(j, dp[j]))
	}

	return n - 2*best
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
