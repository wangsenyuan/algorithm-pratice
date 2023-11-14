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
		n, m := readTwoNums(reader)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 3)
		}
		res := solve(n, edges)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
	// 如果在不考虑方向的情况下，不连通，那答案是0
	// 否则肯定有结果
	// 然后二分就好了，不考虑方向，只要union-find即可

	m := len(edges)

	g := NewGraph(n, 2*m)

	vis := make([]int, n)

	var root int

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = 1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 0 {
				dfs(v)
			}
		}
		root = u
	}

	var dfsSize func(u int) int
	dfsSize = func(u int) int {
		vis[u]++
		sz := 1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] == 0 {
				sz += dfsSize(v)
			}
		}
		return sz
	}

	check := func(x int) bool {
		g.Reset()
		for i := 0; i < n; i++ {
			vis[i] = 0
			root = -1
		}
		for i := 0; i < m; i++ {
			u, v := edges[i][0]-1, edges[i][1]-1
			g.AddEdge(u, v)
			if edges[i][2] <= x {
				g.AddEdge(v, u)
			}
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				dfs(i)
			}
		}

		for i := 0; i < n; i++ {
			vis[i] = 0
		}

		return dfsSize(root) == n
	}

	l, r := 0, 1<<30

	if !check(r) {
		return -1
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) Reset() {
	for i := 0; i < len(g.node); i++ {
		g.node[i] = 0
	}
	for i := 0; i < len(g.next); i++ {
		g.next[i] = 0
	}
	g.cur = 0
}
