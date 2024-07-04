package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		s, _ := reader.ReadBytes('\n')
		var x int
		pos := readInt(s, 0, &x) + 1
		queries[i] = make([]int, x+1)
		queries[i][0] = x
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	res := solve(n, edges, queries)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dsu := NewDSU(n)

	dist := make([]int, n)
	track := make([]int, n)

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		track[u] = p
		far := u
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dist[v] = dist[u] + 1
				tmp := dfs(u, v)
				if dist[tmp] > dist[far] {
					far = tmp
				}
			}
		}

		return far
	}

	var union func(p int, u int)

	union = func(p int, u int) {
		if p >= 0 {
			dsu.Union(p, u)
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				union(u, v)
			}
		}
	}

	diamater := make([]int, n)

	process := func(u int) {
		dist[u] = 1
		first := dfs(-1, u)
		dist[first] = 1
		second := dfs(-1, first)
		mid := second
		for dist[mid] > (dist[second]+1)/2 {
			mid = track[mid]
		}
		diamater[mid] = dist[second]
		// 当直径 = 4时，半径 = 3， 当直径= 3， 半径 = 2
		// 这里算的节点数
		union(-1, mid)
	}

	for i := 0; i < n; i++ {
		if dist[i] == 0 {
			process(i)
		}
	}

	find := func(u int) int {
		return diamater[dsu.Find(u)] - 1
	}

	merge := func(u int, v int) {
		u = dsu.Find(u)
		v = dsu.Find(v)
		if u == v {
			return
		}
		ru := diamater[u] - (diamater[u]+1)/2 + 1
		rv := diamater[v] - (diamater[v]+1)/2 + 1
		nd := max(diamater[u], diamater[v], ru+rv)
		dsu.Union(u, v)
		u = dsu.Find(u)
		diamater[u] = nd
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			x := cur[1]
			ans = append(ans, find(x-1))
		} else {
			x, y := cur[1], cur[2]
			merge(x-1, y-1)
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

type DSU struct {
	arr  []int
	cnt  []int
	size int
}

func NewDSU(n int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = n
	return set
}

func (set *DSU) Find(u int) int {
	if set.arr[u] != u {
		set.arr[u] = set.Find(set.arr[u])
	}
	return set.arr[u]
}

func (set *DSU) Union(a, b int) bool {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return false
	}
	if set.cnt[a] < set.cnt[b] {
		a, b = b, a
	}
	set.cnt[a] += set.cnt[b]
	set.arr[b] = a
	set.size--
	return true
}
