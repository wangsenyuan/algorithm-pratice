package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) []int {
	m := len(edges)

	type edge struct {
		id int
		u  int
		v  int
		w  int
	}

	es := make([]edge, m)
	for i, cur := range edges {
		es[i] = edge{i, cur[0] - 1, cur[1] - 1, cur[2]}
	}

	slices.SortFunc(es, func(a, b edge) int {
		return a.w - b.w
	})

	set := NewDSU(n)
	marked := make([]bool, m)

	g := NewGraph(n, 2*n)

	for _, e := range es {
		if set.Union(e.u, e.v) {
			marked[e.id] = true
			g.AddEdge(e.u, e.v, e.w)
			g.AddEdge(e.v, e.u, e.w)
		}
	}

	dep := make([]int, n)
	h := bits.Len(uint(n))
	fa := make([][]int, n)
	dp := make([][]int, n)
	for i := range n {
		fa[i] = make([]int, h)
		dp[i] = make([]int, h)
	}

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
			dp[u][i] = max(dp[u][i-1], dp[fa[u][i-1]][i-1])
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dp[v][0] = g.val[i]
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	find := func(u int, k int) int {
		var ans int
		for i := h - 1; i >= 0; i-- {
			if k&(1<<i) > 0 {
				ans = max(ans, dp[u][i])
				u = fa[u][i]
			}
		}
		return ans
	}

	var ans []int

	for i := 0; i < m; i++ {
		if !marked[i] {
			u, v := edges[i][0]-1, edges[i][1]-1
			p := lca(u, v)
			var tmp int
			if p != u {
				tmp = max(tmp, find(u, dep[u]-dep[p]))
			}
			if p != v {
				tmp = max(tmp, find(v, dep[v]-dep[p]))
			}
			ans = append(ans, tmp)
		}
	}

	return ans
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

func (set *DSU) Reset() {
	for i := 0; i < len(set.arr); i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	set.size = len(set.arr)
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
