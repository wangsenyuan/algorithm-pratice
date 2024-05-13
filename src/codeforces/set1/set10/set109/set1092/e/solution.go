package main

import (
	"bufio"
	"bytes"
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

	d, res := solve(n, edges)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", d))
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(n int, edges [][]int) (int, [][]int) {
	g := NewGraph(n, 2*n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	vis := make([]bool, n)

	var cur []int

	dist := make([]int, n)

	var dfs func(p int, u int) int
	dfs = func(p int, u int) int {
		vis[u] = true
		cur = append(cur, u)
		var h int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dist[v] = dist[u] + 1
				h = max(h, dfs(u, v))
			}
		}
		return h + 1
	}

	var dfs2 func(p int, u int, expect int, x int) int

	dfs2 = func(p int, u int, expect int, x int) int {
		if u == x {
			return -1
		}
		res := -2
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			tmp := dfs2(u, v, expect, x)
			if tmp >= 0 {
				// already found a answer
				return tmp
			}
			if tmp == -1 && dist[v] == expect/2 {
				return v
			}
			res = max(res, tmp)
		}

		return res
	}

	var arr []Pair
	var mh int
	for i := 0; i < n; i++ {
		if !vis[i] {
			cur = make([]int, 0, 1)
			h := dfs(-1, i)
			first := -1
			for _, u := range cur {
				if dist[u] == h-1 {
					first = u
					break
				}
			}
			cur = make([]int, 0, 1)
			dist[first] = 0
			h = dfs(-1, first)

			var second int

			for _, u := range cur {
				if dist[u] == h-1 {
					second = u
					break
				}
			}
			if first == second {
				arr = append(arr, Pair{h, first})
			} else {
				center := dfs2(-1, first, h, second)
				arr = append(arr, Pair{h, center})
			}

			mh = max(mh, h)
		}
	}

	best := -1
	for i, p := range arr {
		if p.first == mh {
			best = i
			break
		}
	}

	var res [][]int
	for i, p := range arr {
		if i == best {
			continue
		}
		res = append(res, []int{p.second + 1, arr[best].second + 1})
	}

	return getDiameter(n, edges, res), res
}

func getDiameter(n int, edges [][]int, extra [][]int) int {
	if len(edges)+len(extra) != n-1 {
		panic("not a valid answer")
	}

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	for _, e := range extra {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	sz := make([]int, n)
	dist := make([]int, n)
	var dfs func(p, u int)
	dfs = func(p, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dist[v] = dist[u] + 1
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}
	dfs(-1, 0)
	if sz[0] != n {
		panic("not a valid tree")
	}

	var first int

	for i := 0; i < n; i++ {
		if dist[i] > dist[first] {
			first = i
		}
	}

	dist[first] = 0
	dfs(-1, first)

	var second int
	for i := 0; i < n; i++ {
		if dist[i] > dist[second] {
			second = i
		}
	}
	return dist[second]
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
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
