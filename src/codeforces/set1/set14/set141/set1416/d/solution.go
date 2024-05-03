package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, q := readThreeNums(reader)
	p := readNNums(reader, n)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	cmds := make([][]int, q)
	for i := 0; i < q; i++ {
		cmds[i] = readNNums(reader, 2)
	}

	res := solve(n, p, edges, cmds)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, p []int, edges [][]int, cmds [][]int) []int {
	rem := make([]bool, len(edges))
	for i := 0; i < len(cmds); i++ {
		if cmds[i][0] == 2 {
			u := cmds[i][1] - 1
			rem[u] = true
		}
	}
	m := len(cmds)
	set := NewDSU(n + m + len(edges))

	g := NewGraph(n+m+len(edges), n+len(edges)*2+m*2)
	id := n

	for i, e := range edges {
		if !rem[i] {
			u, v := e[0]-1, e[1]-1
			u = set.Find(u)
			v = set.Find(v)
			if u != v {
				set.arr[u] = v
				g.AddEdge(v, u)
			}
		}
	}

	fa := make([]int, len(cmds))

	for i := len(cmds) - 1; i >= 0; i-- {
		j := cmds[i][1] - 1
		if cmds[i][0] == 1 {
			fa[i] = set.Find(j)
		} else {
			e := edges[j]
			u, v := e[0]-1, e[1]-1
			u = set.Find(u)
			v = set.Find(v)
			if u != v {
				set.arr[u] = id
				g.AddEdge(id, u)
				set.arr[v] = id
				g.AddEdge(id, v)
				id++
			}
		}
	}
	sz := make([]int, id)
	in := make([]int, id)
	for i := 0; i < id; i++ {
		in[i] = -1
	}
	var time int

	var dfs func(u int)

	dfs = func(u int) {
		in[u] = time
		time++
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs(v)
			sz[u] += sz[v]
		}
	}

	for i := id - 1; i >= 0; i-- {
		if in[i] < 0 {
			dfs(i)
		}
	}

	tr := NewSegTree(id)

	rev := make([]int, n+1)

	for i := 0; i < n; i++ {
		tr.Update(in[i], p[i])
		rev[p[i]] = i
	}

	var res []int

	for i, cur := range cmds {
		//j := cur[1] - 1
		if cur[0] == 1 {
			// query
			u := fa[i]
			v := tr.Get(in[u], in[u]+sz[u])
			res = append(res, v)
			if v > 0 {
				j := rev[v]
				tr.Update(in[j], 0)
			}
		}
	}

	return res
}

func reverse(res []int) {
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
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

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = max(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := 0
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = max(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
