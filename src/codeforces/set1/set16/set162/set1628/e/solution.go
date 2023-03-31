package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	E := make([][]int, n-1)

	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 3)
	}

	Q := make([][]int, k)

	for i := 0; i < k; i++ {
		s, _ := reader.ReadBytes('\n')
		var opt int
		pos := readInt(s, 0, &opt)
		if opt == 1 || opt == 2 {
			var l, r int
			pos = readInt(s, pos+1, &l)
			readInt(s, pos+1, &r)
			Q[i] = []int{opt, l, r}
		} else {
			var x int
			readInt(s, pos+1, &x)
			Q[i] = []int{opt, x}
		}
	}

	res := solve(n, E, Q)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

func solve(n int, E [][]int, Q [][]int) []int {
	sort.Slice(E, func(i, j int) bool {
		return E[i][2] < E[j][2]
	})

	set := NewDSU(n * 2)
	g := NewGraph(n*2, 4*n)
	tot := n

	W := make([]int, n*2)
	for i := 0; i < len(W); i++ {
		W[i] = -1
	}
	for _, e := range E {
		u, v, w := e[0], e[1], e[2]
		u = set.Find(u)
		v = set.Find(v)
		tot++
		g.AddEdge(tot, u)
		g.AddEdge(tot, v)
		W[tot] = w
		set.arr[u] = tot
		set.arr[v] = tot
	}

	big := make([]int, 2*n)
	sz := make([]int, 2*n)
	dad := make([]int, 2*n)
	dep := make([]int, 2*n)
	var dfs1 func(u int)

	dfs1 = func(u int) {
		sz[u]++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dad[v] = u
			dep[v] = dep[u] + 1
			dfs1(v)
			sz[u] += sz[v]
			if big[u] == 0 || sz[big[u]] < sz[v] {
				big[u] = v
			}
		}
	}

	dfs1(tot)

	bl := make([]int, tot+1)

	var dfs2 func(u int, fa int)

	dfs2 = func(u int, fa int) {
		bl[u] = fa
		if big[u] > 0 {
			dfs2(big[u], fa)
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != big[u] {
				dfs2(v, v)
			}
		}
	}

	dfs2(tot, tot)

	lca := func(u, v int) int {
		if u == 0 || v == 0 {
			return u | v
		}
		for bl[u] != bl[v] {
			if dep[bl[u]] < dep[bl[v]] {
				u, v = v, u
			}
			// bl[u] is deeper
			u = dad[bl[u]]
		}
		if dep[u] > dep[v] {
			u, v = v, u
		}
		return u
	}

	type Node struct {
		who  int
		lazy int
		lca  int
	}

	tree := make([]*Node, 4*tot)
	for i := 0; i < len(tree); i++ {
		tree[i] = new(Node)
	}
	var build func(u int, l int, r int)

	build = func(u int, l int, r int) {
		if l == r {
			tree[u].who = l
			return
		}
		mid := (l + r) / 2
		build(u*2, l, mid)
		build(u*2+1, mid+1, r)
		tree[u].who = lca(tree[u*2].who, tree[u*2+1].who)
	}

	pull := func(u int) {
		tree[u].lca = lca(tree[u*2].lca, tree[u*2+1].lca)
	}

	pushLazy := func(u int, opt int) {
		if opt == 1 {
			tree[u].lca = tree[u].who
		} else {
			tree[u].lca = 0
		}
		tree[u].lazy = opt
	}

	push := func(u int) {
		if tree[u].lazy != 0 {
			pushLazy(u*2, tree[u].lazy)
			pushLazy(u*2+1, tree[u].lazy)
			tree[u].lazy = 0
		}
	}

	var update func(u int, l int, r int, L int, R int, v int)

	update = func(u int, l int, r int, L int, R int, v int) {
		if L <= l && r <= R {
			pushLazy(u, v)
			return
		}
		push(u)

		mid := (l + r) / 2

		if L <= mid {
			update(u*2, l, mid, L, R, v)
		}
		if mid < R {
			update(u*2+1, mid+1, r, L, R, v)
		}

		pull(u)
	}

	build(1, 1, tot)

	var ans []int

	for _, cur := range Q {
		opt := cur[0]
		if opt == 1 || opt == 2 {
			l, r := cur[1], cur[2]
			v := 1
			if opt == 2 {
				v = -1
			}
			update(1, 1, tot, l, r, v)
		} else {
			x := cur[1]
			ans = append(ans, W[lca(tree[1].lca, x)])
		}
	}

	return ans
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
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	s := new(DSU)
	s.arr = make([]int, n)
	s.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		s.arr[i] = i
		s.cnt[i]++
	}
	return s
}

func (s *DSU) Find(u int) int {
	if s.arr[u] != u {
		s.arr[u] = s.Find(s.arr[u])
	}
	return s.arr[u]
}
