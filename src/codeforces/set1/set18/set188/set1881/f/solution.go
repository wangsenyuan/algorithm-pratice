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
		n, k := readTwoNums(reader)
		a := readNNums(reader, k)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, a)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

type pair struct {
	first  int
	second int
}

func solve(n int, E [][]int, a []int) int {
	if len(a) == 1 {
		return 0
	}
	marked := make([]bool, n)
	for _, x := range a {
		marked[x-1] = true
	}

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	// 答案应该是这些marked的点组成的子树的，直径的中点？
	// 如果只有一个点，那么答案就是那个点（0）

	var dfs func(p int, u int, d int) pair

	dfs = func(p int, u int, d int) pair {
		var res pair

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				tmp := dfs(u, v, d+1)
				if tmp.first > 0 && res.first < tmp.first {
					res = tmp
				}
			}
		}
		if res.first == 0 && marked[u] {
			res = pair{d, u}
		}
		return res
	}

	first := dfs(-1, a[0], 0)
	second := dfs(-1, first.second, 0)

	return (second.first + 1) / 2
}

func solve1(n int, E [][]int, a []int) int {
	// 假设 dist[0] = 0到marked的最大距离
	// 假设u是0的其中一个children，
	// 要计算dist[u] 必须知道两个信息, 一个是u到它的子树中的marked的距离
	//  这个假设已知，第二个，就是通过0能得到的最远的距离，且不能在u中
	marked := make([]bool, n)
	for _, x := range a {
		marked[x-1] = true
	}

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	f := make([]int, n)

	dist := make([][]Pair, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		dist[u] = []Pair{{-1, -1}, {-1, -1}}

		if marked[u] {
			dist[u][0] = Pair{0, u}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			dfs(u, v)
			if dist[v][0].second >= 0 {
				// do have a marker in v
				if dist[v][0].first+1 > dist[u][0].first {
					dist[u][1] = dist[u][0]
					dist[u][0] = Pair{dist[v][0].first + 1, v}
				} else if dist[v][0].first+1 > dist[u][1].first {
					dist[u][1] = Pair{dist[v][0].first + 1, v}
				}
			}
		}
		f[u] = dist[u][0].first
	}

	dfs(0, 0)

	var dfs2 func(p int, u int, x int)

	dfs2 = func(p int, u int, x int) {
		if p >= 0 && x >= 0 {
			// 不需要restore
			if x+1 > dist[u][0].first {
				dist[u][0] = Pair{x + 1, p}
			} else if x+1 > dist[u][1].first {
				dist[u][1] = Pair{x + 1, p}
			}
		}

		f[u] = max(f[u], dist[u][0].first)

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			if dist[u][0].second == v {
				dfs2(u, v, dist[u][1].first)
			} else {
				dfs2(u, v, dist[u][0].first)
			}
		}
	}

	dfs2(-1, 0, -1)

	res := n
	for i := 0; i < n; i++ {
		res = min(res, f[i])
	}

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

type Pair struct {
	first  int
	second int
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
