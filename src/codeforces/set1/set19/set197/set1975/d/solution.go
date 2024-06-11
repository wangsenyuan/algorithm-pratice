package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, 2)
		edges := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}

		res := solve(n, edges, P)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

func solve(n int, edges [][]int, P []int) int {
	g := NewGraph(n, len(edges)*2)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	var res int
	mid := P[0] - 1

	que := make([]int, n)
	dist := make([]int, n)
	trace := make([]int, n)

	bfs := func(x int) {
		for i := 0; i < n; i++ {
			dist[i] = -1
			trace[i] = -1
		}
		var front, tail int

		dist[x] = 0
		que[front] = x
		front++

		for tail < front {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					trace[v] = u
					que[front] = v
					front++
				}
			}
		}
	}

	if P[0] != P[1] {
		bfs(P[0] - 1)
		mid = P[1] - 1
		for dist[mid] != dist[P[1]-1]/2 {
			mid = trace[mid]
		}
		res = dist[P[1]-1] - dist[mid]
	}
	bfs(mid)

	res += 2 * (n - 1)

	var far int
	for i := 0; i < n; i++ {
		if dist[i] > dist[far] {
			far = i
		}
	}
	res -= dist[far]

	return res
}

func solve1(n int, edges [][]int, P []int) int {
	g := NewGraph(n, len(edges)*2)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	var res int
	mid := P[0] - 1
	if P[0] != P[1] {
		que := make([]int, n)
		var front, tail int

		dist := make([]int, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[P[0]-1] = 0
		que[front] = P[0] - 1
		front++
		trace := make([]int, n)

		for tail < front {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					trace[v] = u
					que[front] = v
					front++
				}
			}
		}
		mid = P[1] - 1
		for dist[mid] != dist[P[1]-1]/2 {
			mid = trace[mid]
		}
		res = dist[P[1]-1] - dist[mid]
	}

	// mid is blue now, 然后就是把所有的都涂成blue
	// 在涂完一课子树前，不应该出来，否则还的回去
	// dp[u][0] = sum(dp[v]) + 2 * cnt(direct chilren edge) 进去再出来
	// dp[u][1] = 表示它进入它不必回到u
	// 但是最后一次，不需要回来，所以应该减去最长的那条路径
	dp := make([][]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		dp[u] = make([]int, 2)
		mv := -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				dp[u][0] += dp[v][0] + 2
				if mv < 0 || dp[v][0]-dp[v][1] > dp[mv][0]-dp[mv][1] {
					mv = v
				}
			}
		}

		if mv >= 0 {
			// dp[mv][0] + 2 - 1 - dp[mv][1] 是最小的
			dp[u][1] = dp[u][0] - (dp[mv][0] + 2) + 1 + dp[mv][1]
		}
	}

	dfs(-1, mid)

	res += dp[mid][1]

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
