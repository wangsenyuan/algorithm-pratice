package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E, m)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, E [][]int, k int) int {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	dist := make([]int, n)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dist[v] = dist[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	var root int

	for i := 0; i < n; i++ {
		if dist[i] > dist[root] {
			root = i
		}
	}

	var time = -1
	eu := make([]int, n)
	open := make([]int, n)
	end := make([]int, n)
	P := make([]int, n)

	var dfs2 func(p, u int)
	dfs2 = func(p, u int) {
		time++
		eu[time] = u
		open[u] = time
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				P[v] = u
				dist[v] = dist[u] + 1
				dfs2(u, v)
			}
		}
		end[u] = time
	}
	dist[root] = 0
	dfs2(root, root)

	lt := NewLazyTree(n)

	for i := 0; i < n; i++ {
		lt.Update(open[i], open[i], dist[i])
	}

	set := make([]bool, n)
	set[root] = true
	lt.Update(open[root], open[root], -(n + 1))
	ans := 1
	k--
	for k > 0 {
		pair := lt.Query(0, n-1)
		u := eu[pair[0]]
		add := pair[1]
		ans++
		for cur := u; !set[cur]; cur = P[cur] {
			k--
			lt.Update(open[cur], open[cur], -(n + 1))
			set[cur] = true
			for i := g.nodes[cur]; i > 0; i = g.next[i] {
				v := g.to[i]
				if P[cur] == v || set[v] {
					continue
				}
				lt.Update(open[v], end[v], -add)
			}
			add--
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

type LazyTree struct {
	tree []int
	lazy []int
	ind  []int
	n    int
	m    int
}

func NewLazyTree(n int) *LazyTree {
	m := 1
	for m < n {
		m <<= 1
	}
	lt := new(LazyTree)

	lt.n = n
	lt.m = m
	lt.tree = make([]int, m<<1)
	lt.lazy = make([]int, m<<1)
	lt.ind = make([]int, m<<1)
	for i := m; i < len(lt.ind); i++ {
		lt.ind[i] = i - m
	}

	for i := m - 1; i > 0; i-- {
		if lt.tree[i*2] <= lt.tree[i*2+1] {
			lt.ind[i] = lt.ind[2*i]
		} else {
			lt.ind[i] = lt.ind[2*i+1]
		}
	}

	return lt
}

func (lt *LazyTree) push(i int) {
	if lt.lazy[i] != 0 {
		lt.tree[i] += lt.lazy[i]
		if i < lt.m {
			lt.lazy[i*2] += lt.lazy[i]
			lt.lazy[i*2+1] += lt.lazy[i]
		}
		lt.lazy[i] = 0
	}
}

func (lt *LazyTree) Query(l, r int) []int {
	var loop func(l, r int, ll, rr int, i int) []int
	loop = func(l, r int, ll, rr int, i int) []int {
		if l == ll && r == rr {
			return []int{lt.ind[i], lt.tree[i]}
		}
		mid := (ll + rr) / 2
		if r <= mid {
			return loop(l, r, ll, mid, i*2)
		}
		if mid < l {
			return loop(l, r, mid+1, rr, i*2+1)
		}

		a := loop(l, mid, ll, mid, i*2)
		b := loop(mid+1, r, mid+1, rr, i*2+1)

		if a[1] >= b[1] {
			return a
		}
		return b
	}

	return loop(l, r, 0, lt.m-1, 1)
}

func (lt *LazyTree) Update(l, r int, x int) {

	var loop func(l, r, ll, rr int, i int)

	loop = func(l, r, ll, rr int, i int) {
		lt.push(i)
		if l == ll && r == rr {
			lt.lazy[i] += x
			lt.push(i)
			return
		}
		mid := (ll + rr) / 2
		if r <= mid {
			loop(l, r, ll, mid, i*2)
			lt.push(2*i + 1)
		} else if mid < l {
			lt.push(2 * i)
			loop(l, r, mid+1, rr, i*2+1)
		} else {
			loop(l, mid, ll, mid, i*2)
			loop(mid+1, r, mid+1, rr, i*2+1)
		}

		lt.tree[i] = max(lt.tree[2*i], lt.tree[2*i+1])
		if lt.tree[i] == lt.tree[2*i] {
			lt.ind[i] = lt.ind[2*i]
		} else {
			lt.ind[i] = lt.ind[2*i+1]
		}
	}
	loop(l, r, 0, lt.m-1, 1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
