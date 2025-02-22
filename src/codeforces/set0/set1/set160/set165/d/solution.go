package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var tp int
		pos := readInt(s, 0, &tp) + 1
		if tp == 1 || tp == 2 {
			queries[i] = make([]int, 2)
		} else {
			queries[i] = make([]int, 3)
		}
		queries[i][0] = tp
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

	g := NewGraph(n, 2*n)
	deg := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
		deg[u]++
		deg[v]++
	}

	x := slices.Max(deg)

	var root int
	for i, v := range deg {
		if v == x {
			root = i
			break
		}
	}

	pos := make([]int, n)
	dep := make([]int, n)
	head := make([]int, n)
	head[root] = -1
	var timer int
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				if p < 0 {
					// not u
					head[v] = v
				} else {
					head[v] = head[u]
				}
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
		pos[u] = timer
		timer++
	}
	// 每条路径肯定是连在一起的
	dfs(-1, root)

	tr := NewSegTree(n, 0, add)

	find := func(u int, v int) int {
		if u == v {
			return 0
		}
		if dep[u] < dep[v] {
			u, v = v, u
		}

		if v == root || head[u] == head[v] {
			// 它们在一条线上
			second := pos[v]
			if v == root {
				second = pos[head[u]] + 1
			}
			cnt := tr.Get(pos[u], second)
			if cnt > 0 {
				// 被隔断了
				return -1
			}
			return dep[u] - dep[v]
		}
		cnt1 := tr.Get(pos[u], pos[head[u]]+1)
		cnt2 := tr.Get(pos[v], pos[head[v]]+1)
		if cnt1 > 0 || cnt2 > 0 {
			return -1
		}
		return dep[u] + dep[v]
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 || cur[0] == 2 {
			// paint black
			edge := edges[cur[1]-1]
			u, v := edge[0]-1, edge[1]-1
			if dep[u] > dep[v] {
				v = u
			}
			// u is the parent, so update v
			tr.Update(pos[v], cur[0]-1)
		} else {
			u, v := cur[1]-1, cur[2]-1
			// 计算它们中间有没有1
			ans = append(ans, find(u, v))
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

func add(a, b int) int {
	return a + b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
