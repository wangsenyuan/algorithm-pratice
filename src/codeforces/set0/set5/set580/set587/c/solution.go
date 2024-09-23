package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, q := readThreeNums(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	c := readNNums(reader, m)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(n, edges, c, queries)
	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d", len(res[i])))
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i][j]))
		}
		buf.WriteByte('\n')
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

const inf = 1 << 50

func solve1(n int, edges [][]int, c []int, queries [][]int) [][]int {
	vals := make([][]int32, n+1)

	for i := 0; i < len(c); i++ {
		vals[c[i]] = append(vals[c[i]], int32(i))
	}

	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	lg := bits.Len(uint(n + 1))

	fa := make([][]int32, n+1)
	fa[0] = make([]int32, lg)
	mn := make([][][]int32, n+1)
	mn[0] = make([][]int32, lg)

	dep := make([]int32, n+1)

	merge := func(a, b []int32) []int32 {
		// len(a) <= 10, 而且不会有重复的值
		res := make([]int32, min(len(a)+len(b), 10))

		for u, v := 0, 0; u+v < len(res); {
			if v == len(b) || u < len(a) && a[u] < b[v] {
				res[u+v] = a[u]
				u++
			} else {
				res[u+v] = b[v]
				v++
			}
		}

		return res
	}

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int32, lg)
		fa[u][0] = int32(p)
		for i := 1; i < lg; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		mn[u] = make([][]int32, lg)
		mn[u][0] = vals[u]

		for i := 1; i < lg; i++ {
			mn[u][i] = merge(mn[u][i-1], mn[fa[u][i-1]][i-1])
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 1)

	lca := func(u, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := lg - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
				u = int(fa[u][i])
			}
		}
		if u == v {
			return u
		}
		for i := lg - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = int(fa[u][i])
				v = int(fa[v][i])
			}
		}
		return int(fa[u][0])
	}

	get := func(u int, k int) []int32 {
		// 表示找出从u开始，长度为k的路径上的最小值
		var res []int32
		for i := lg - 1; i >= 0; i-- {
			if k&(1<<i) > 0 {
				res = merge(res, mn[u][i])
				u = int(fa[u][i])
			}
		}
		return res
	}

	find := func(cur []int) []int {
		u, v, a := cur[0], cur[1], cur[2]

		p := lca(u, v)

		var res []int32

		res = merge(res, mn[p][0])

		if u != p {
			res = merge(res, get(u, int(dep[u]-dep[p])))
		}

		if v != p {
			res = merge(res, get(v, int(dep[v]-dep[p])))
		}
		if a < len(res) {
			res = res[:a]
		}
		arr := make([]int, len(res))
		for i := 0; i < len(res); i++ {
			arr[i] = int(res[i] + 1)
		}
		return arr
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur)
	}

	return ans
}

func solve(n int, edges [][]int, c []int, queries [][]int) [][]int {
	vals := make([][]int, n)

	for i := 0; i < len(c); i++ {
		c[i]--
		vals[c[i]] = append(vals[c[i]], i)
	}

	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	sz := make([]int, n)
	big := make([]int, n)
	dep := make([]int, n)
	fa := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = p
		big[u] = -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
				if big[u] < 0 || sz[big[u]] < sz[v] {
					big[u] = v
				}
				sz[u] += sz[v]
			}
		}
		sz[u]++
	}

	dfs(-1, 0)

	head := make([]int, n)
	pos := make([]int, n)
	var it int
	var dfs2 func(p int, u int, x int)
	dfs2 = func(p int, u int, x int) {
		pos[u] = it
		it++

		head[u] = x

		if big[u] >= 0 {
			dfs2(u, big[u], x)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p || v == big[u] {
				continue
			}
			dfs2(u, v, v)
		}
	}

	dfs2(-1, 0, 0)

	set := NewSegTree(n)

	for i := 0; i < n; i++ {
		if len(vals[i]) > 0 {
			set.Update(pos[i], vals[i][0])
		}
	}

	lca := func(u int, v int) int {
		for head[u] != head[v] {
			if dep[head[u]] < dep[head[v]] {
				u, v = v, u
			}
			// dep[head[u]] >= dep[head[v]]
			// 移动那个深的节点
			u = fa[head[u]]
		}

		if dep[u] > dep[v] {
			return v
		}
		return u
	}

	id := make([]int, n)

	get := func(u int, p int) int {
		res := inf
		for head[u] != head[p] {
			res = min(res, set.Get(pos[head[u]], pos[u]+1))
			u = fa[head[u]]
		}
		// 现在在同一段了
		res = min(res, set.Get(pos[p], pos[u]+1))

		if res < inf {
			u := c[res]
			id[u]++
			if id[u] < len(vals[u]) {
				i := vals[u][id[u]]
				set.Update(pos[u], i)
			} else {
				set.Update(pos[u], inf)
			}
		}

		return res
	}

	handle := func(u int, p int, k int) []int {
		var res []int
		// 要在这个区间内找到k个最小的下标
		for i := 0; i < k; i++ {
			tmp := get(u, p)
			if tmp >= inf {
				break
			}
			res = append(res, tmp)
		}

		// now must reset back
		for i := 0; i < len(res); i++ {
			u := c[res[i]]
			id[u] = 0
			// vals[u][0] 不一定等于 res[i]
			if vals[u][0] == res[i] {
				set.Update(pos[u], vals[u][id[u]])
			}
		}

		return res
	}

	merge := func(a, b []int) []int {
		c := make([]int, len(a)+len(b))
		copy(c, a)
		copy(c[len(a):], b)
		sort.Ints(c)
		var m int
		for i := 1; i <= len(c); i++ {
			if i == len(c) || c[i-1] < c[i] {
				c[m] = c[i-1]
				m++
			}
		}
		return c[:m]
	}

	find := func(cur []int) []int {
		u, v, a := cur[0]-1, cur[1]-1, cur[2]
		p := lca(u, v)

		res1 := handle(u, p, a)
		res2 := handle(v, p, a)

		res := merge(res1, res2)

		if len(res) > a {
			res = res[:a]
		}

		for i := 0; i < len(res); i++ {
			res[i]++
		}

		return res
	}

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur)
	}

	return ans
}

func createArray(k int) []int {
	arr := make([]int, k)
	for i := 0; i < k; i++ {
		arr[i] = inf
	}
	return arr
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

type SegTree struct {
	size int
	arr  []int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = inf
	}
	for i := n - 1; i > 1; i-- {
		arr[i] = arr[2*i]
	}

	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = min(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := inf
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = min(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
