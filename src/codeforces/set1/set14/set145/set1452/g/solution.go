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
	n := readNum(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	A := readNNums(reader, m)
	res := solve(n, E, A)
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, E [][]int, A []int) []int {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	used := make([]bool, n)
	sz := make([]int, n)
	par := make([]int, n)
	max_dist := make([]int, n)
	dist := make([][]int, n)

	var dfs_size func(p int, u int)

	dfs_size = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !used[v] {
				dfs_size(u, v)
				sz[u] += sz[v]
			}
		}
	}

	var dfs_centroid func(x int, p int, s int) int

	dfs_centroid = func(x int, p int, s int) int {
		ans := -1
		ok := true
		for i := g.nodes[x]; i > 0 && ok; i = g.next[i] {
			y := g.to[i]
			if used[y] {
				continue
			}
			if y != p {
				ok = sz[y]*2 <= s
			} else {
				ok = (s-sz[x])*2 <= s
			}
		}
		if ok {
			ans = x
		}
		for i := g.nodes[x]; i > 0; i = g.next[i] {
			y := g.to[i]
			if y != p && !used[y] {
				ans = max(ans, dfs_centroid(y, x, s))
			}
		}

		return ans
	}

	var dfs_dist func(x int, p int, d int, s int)

	dfs_dist = func(x int, p int, d int, s int) {
		if len(dist[x]) == 0 {
			dist[x] = make([]int, 0, 1)
		}
		dist[x] = append(dist[x], d)
		for i := g.nodes[x]; i > 0; i = g.next[i] {
			y := g.to[i]
			if y != p && !used[y] {
				dfs_dist(y, x, d+1, s)
			}
		}
		max_dist[s] = max(max_dist[s], d)
	}

	var decomposition func(v int) int

	decomposition = func(v int) int {
		dfs_size(v, v)
		c := dfs_centroid(v, v, sz[v])
		used[c] = true
		for i := g.nodes[c]; i > 0; i = g.next[i] {
			y := g.to[i]
			if !used[y] {
				par[decomposition(y)] = c
			}
		}
		used[c] = false
		dfs_dist(c, c, 0, c)

		return c
	}

	decomposition(0)

	val := make([][]int, n)

	for i := 0; i < n; i++ {
		val[i] = make([]int, max_dist[i]+1)
	}

	d := getDistance(n, A, g)

	for i := 0; i < n; i++ {
		if d[i] == 0 {
			continue
		}
		for j, c := 0, i; j < len(dist[i]); j, c = j+1, par[c] {
			dd := dist[i][j]
			if dd > d[i]-1 {
				continue
			}
			dd = d[i] - 1 - dd
			if dd >= len(val[c]) {
				dd = len(val[c]) - 1
			}
			val[c][dd] = max(val[c][dd], d[i])
		}
	}

	for i := 0; i < n; i++ {
		for j := max_dist[i]; j > 0; j-- {
			val[i][j-1] = max(val[i][j], val[i][j-1])
		}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if d[i] > 0 {
			c := i
			for _, dd := range dist[i] {
				ans[i] = max(ans[i], val[c][dd])
				c = par[c]
			}
		}
	}

	return ans
}

func solve1(n int, E [][]int, A []int) []int {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		g.AddEdge(e[0]-1, e[1]-1)
		g.AddEdge(e[1]-1, e[0]-1)
	}

	a := getDistance(n, A, g)

	h := make([]int, n)

	for i := 0; i < n; i++ {
		h[i] = -1
	}

	pcd := make([]int, n)

	var dfs func(v, s, p int, cd *int) int

	dfs = func(v, s, p int, cd *int) int {
		var sz = 1
		for i := g.nodes[v]; i > 0; i = g.next[i] {
			u := g.to[i]
			if h[u] < 0 && u != p {
				sz += dfs(u, s, v, cd)
			}
		}
		if *cd == -1 && (2*sz >= s || p == -1) {
			*cd = v
		}
		return sz
	}

	var build func(v, s, d, p int)

	build = func(v, s, d, p int) {
		var cd = -1
		dfs(v, s, -1, &cd)
		h[cd] = d
		pcd[cd] = p
		for i := g.nodes[cd]; i > 0; i = g.next[i] {
			u := g.to[i]
			if h[u] < 0 {
				build(u, s/2, d+1, cd)
			}
		}
	}
	build(0, n, 0, -1)

	nd := make([]int, n)

	st := make([][]int, n)

	for v := 0; v < n; v++ {
		u := v
		for u != -1 {
			if len(st[u]) == 0 {
				st[u] = make([]int, 0, 1)
			}
			st[u] = append(st[u], v)
			if pcd[u] != -1 {
				nd[pcd[u]] = max(nd[pcd[u]], nd[u]+1)
			}
			u = pcd[u]
		}
	}

	ord := make([]int, n)
	for i := 0; i < n; i++ {
		ord[i] = i
	}

	sort.Slice(ord, func(i, j int) bool {
		return nd[ord[i]] < nd[ord[j]]
	})

	cur := make([]bool, n)
	d := make([]int, n)
	vals := make([][]int, n)

	var calc func(p int, u int)

	calc = func(p int, u int) {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && cur[v] {
				d[v] = d[u] + 1
				calc(u, v)
			}
		}
	}

	dist := make([][]int, n)

	for _, v := range ord {
		for _, u := range st[v] {
			cur[u] = true
		}
		d[v] = 0
		calc(v, v)
		var mx int
		for _, u := range st[v] {
			mx = max(mx, d[u])
		}
		vals[v] = make([]int, mx+1)
		for _, u := range st[v] {
			vals[v][d[u]] = max(vals[v][d[u]], a[u])
		}
		for j := 0; j < mx; j++ {
			vals[v][j+1] = max(vals[v][j+1], vals[v][j])
		}

		for _, u := range st[v] {
			cur[u] = false
			if len(dist[u]) == 0 {
				dist[u] = make([]int, 0, 1)
			}
			dist[u] = append(dist[u], d[u])
		}
	}

	check := func(v int, x int) bool {
		for i, u := 0, v; u != -1; i, u = i+1, pcd[u] {
			if x-dist[v][i] >= 0 && vals[u][min(len(vals[u])-1, x-dist[v][i])] > x {
				return true
			}
		}
		return false
	}

	ans := make([]int, n)

	for v := 0; v < n; v++ {
		l, r := 0, n
		for l < r {
			mid := (l + r) / 2
			if check(v, mid) {
				l = mid + 1
			} else {
				r = mid
			}
		}
		ans[v] = l
	}

	return ans
}

func getDistance(n int, A []int, g *Graph) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	que := make([]int, n)
	var front, end int
	pushBack := func(u int) {
		que[end] = u
		end++
	}
	popFront := func() int {
		res := que[front]
		front++
		return res
	}

	for _, a := range A {
		a--
		pushBack(a)
		dist[a] = 0
	}

	for front < end {
		u := popFront()
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v] < 0 {
				dist[v] = dist[u] + 1
				pushBack(v)
			}
		}
	}

	return dist
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
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
	next := make([]int, e+1)
	to := make([]int, e+1)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
