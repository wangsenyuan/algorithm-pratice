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
		a := readNNums(reader, n)
		roads := make([][]int, m)
		for i := 0; i < m; i++ {
			roads[i] = readNNums(reader, 2)
		}
		res := solve(a, roads)

		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
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

const oo = 1 << 60

func solve(a []int, roads [][]int) []int {
	n := len(a)

	m := len(roads)
	g := NewGraph(n, m)
	r := NewGraph(n, m)

	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		// direction
		g.AddEdge(u, v, 1)
		r.AddEdge(v, u, 1)
	}

	stack := make([]int, n)
	var top int

	comp := make([]int, n)
	for i := 0; i < n; i++ {
		comp[i] = -1
	}

	var dfs func(u int)
	dfs = func(u int) {
		comp[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if comp[v] < 0 {
				dfs(v)
			}
		}
		stack[top] = u
		top++
	}

	for i := 0; i < n; i++ {
		if comp[i] < 0 {
			dfs(i)
		}
	}
	sz := make([]int, n)
	sum := make([]int, n)

	var dfs2 func(u int, c int)
	dfs2 = func(u int, c int) {
		comp[u] = c
		sz[c]++
		sum[c] += a[u]
		for i := r.nodes[u]; i > 0; i = r.next[i] {
			v := r.to[i]
			if comp[v] < 0 {
				dfs2(v, c)
			}
		}
	}

	for i := 0; i < n; i++ {
		comp[i] = -1
	}

	var gid int

	for i := n - 1; i >= 0; i-- {
		if comp[stack[i]] < 0 {
			dfs2(stack[i], gid)
			gid++
		}
	}

	h := NewGraph(gid, m)
	deg := make([]int, gid)
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		if comp[u] != comp[v] {
			h.AddEdge(comp[u], comp[v], sz[comp[v]])
			deg[comp[v]]++
		}
	}

	items := make([]*Item, gid)

	pq := make([]int, gid)
	var front, tail int
	// gid is the start vertex
	for i := 0; i < gid; i++ {
		items[i] = new(Item)
		items[i].id = i
		items[i].v1 = sz[i]
		items[i].v2 = sum[i]

		if deg[i] == 0 {
			pq[front] = i
			front++
		}
	}

	res := []int{0, -1}

	update := func(h int, s int) {
		if h > res[0] || h == res[0] && s < res[1] {
			res[0] = h
			res[1] = s
		}
	}

	for tail < front {
		u := pq[tail]
		tail++
		it := items[u]

		if h.nodes[u] == 0 {
			update(it.v1, it.v2)
			continue
		}
		for i := h.nodes[u]; i > 0; i = h.next[i] {
			v := h.to[i]
			if items[v].v1 < it.v1+sz[v] || items[v].v1 == it.v1+sz[v] && items[v].v2 > it.v2+sum[v] {
				items[v].v1 = it.v1 + sz[v]
				items[v].v2 = it.v2 + sum[v]
			}
			deg[v]--
			if deg[v] == 0 {
				pq[front] = v
				front++
			}
		}
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}

// An Item is something we manage in a priority queue.
type Item struct {
	id    int
	v1    int
	v2    int
	index int
}
