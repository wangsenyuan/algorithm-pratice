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
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))
			for _, e := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
			}
		}
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

func solve(n int, E [][]int) [][]int {
	g := NewGraph(n, 2*len(E))
	deg := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}
	// 首先节点u的deg[u] >= 4
	// 判断点是否在cycle上，且存在某个点 u deg[u] >= 4

	stack := make([]int, n)
	var top int
	low := make([]int, n)
	num := make([]int, n)
	for i := 0; i < n; i++ {
		low[i] = -1
		num[i] = -1
	}

	var seq int

	var dfs func(p int, u int)

	var comps [][]int

	dfs = func(p int, u int) {
		num[u] = seq
		seq++
		low[u] = num[u]
		stack[top] = u
		top++

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if num[v] < 0 {
				dfs(u, v)
				low[u] = min(low[u], low[v])
			} else if p != v {
				low[u] = min(low[u], num[v])
			}
		}

		if low[u] == num[u] {
			var cycle []int
			for top > 0 {
				v := stack[top-1]
				top--
				cycle = append(cycle, v)
				if u == v {
					break
				}
			}
			// one node not a cycle
			if len(cycle) > 2 {
				comps = append(comps, cycle)
			}
		}
	}

	for i := 0; i < n; i++ {
		if num[i] < 0 {
			dfs(-1, i)
		}
	}

	if len(comps) == 0 {
		return nil
	}

	var res [][]int

	construct := func(comp []int) bool {
		// 很容易可以找到u，
		// 以u的两个节点x,y，找它们中间的（不经过u的）最短路径，
		u := -1
		for _, x := range comp {
			if deg[x] >= 4 {
				u = x
				break
			}
		}

		if u < 0 {
			return false
		}

		in_comp := make([]bool, n)

		for _, x := range comp {
			in_comp[x] = true
		}

		g2 := NewGraph(n, 2*len(E))
		first := -1

		for _, x := range comp {
			if x == u {
				for i := g.nodes[u]; i > 0; i = g.next[i] {
					v := g.to[i]
					if in_comp[v] {
						first = v
						break
					}
				}
				continue
			}

			for i := g.nodes[x]; i > 0; i = g.next[i] {
				y := g.to[i]
				if y == u || !in_comp[y] {
					continue
				}
				g2.AddEdge(x, y)
			}
		}

		que := make([]int, n)
		dist := make([]int, n)
		prev := make([]int, n)

		for i := 0; i < n; i++ {
			dist[i] = -1
			prev[i] = -1
		}
		prev[first] = u
		dist[first] = 0

		var front, tail int
		que[front] = first
		front++
		for tail < front {
			v := que[tail]
			tail++

			for i := g2.nodes[v]; i > 0; i = g2.next[i] {
				w := g2.to[i]
				if in_comp[w] && dist[w] < 0 {
					dist[w] = dist[v] + 1
					prev[w] = v
					que[front] = w
					front++
				}
			}
		}

		second := -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			// 排除未到达的点和first
			if dist[v] <= 0 {
				continue
			}
			if second < 0 || dist[second] > dist[v] {
				second = v
			}
		}
		// second is the closet to first
		res = append(res, []int{u + 1, second + 1})
		vis := make([]bool, n)
		for prev[second] >= 0 {
			vis[second] = true
			res = append(res, []int{second + 1, prev[second] + 1})
			second = prev[second]
		}
		// add tails
		var cnt int

		for i := g.nodes[u]; i > 0 && cnt < 2; i = g.next[i] {
			v := g.to[i]
			if vis[v] {
				// already added
				continue
			}
			res = append(res, []int{u + 1, v + 1})
			cnt++
		}

		return true
	}

	for _, cur := range comps {
		tmp := construct(cur)
		if tmp {
			return res
		}
	}

	return nil
}

func min(a, b int) int {
	if a <= b {
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
