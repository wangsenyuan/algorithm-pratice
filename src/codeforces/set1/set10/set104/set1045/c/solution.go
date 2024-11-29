package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
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

func process(reader *bufio.Reader) []int {
	n, m, k := readThreeNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	queries := make([][]int, k)
	for i := range k {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	g := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make(map[int]bool)
	}

	for _, cur := range edges {
		u, v := cur[0], cur[1]
		g[u][v] = true
		g[v][u] = true
	}

	que := make([]int, n)
	fa := make([]int, n+1)

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = -1
	}
	bfs := func(x int) {
		var head, tail int

		fa[x] = 0
		que[head] = x
		head++
		dist[x] = 0

		for tail < head {
			u := que[tail]
			tail++
			for v := range g[u] {
				if dist[v] < 0 {
					fa[v] = u
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
	}

	bfs(1)

	// 找到那棵树
	tr := NewGraph(n+1, n)

	for i := 1; i <= n; i++ {
		if fa[i] > 0 {
			tr.AddEdge(fa[i], i)
		}
	}

	lca, kth := tr.prepare(1)

	query := func(u int, v int) int {
		if g[u][v] {
			return 1
		}
		// 节点x, 和它的父节点 y, 有可能在同一个分组内
		// 所以节点x和它的邻居节点，也可能在一个分组内
		if tr.dep[v] < tr.dep[u] {
			u, v = v, u
		}
		// v is deeper
		p := lca(u, v)
		if u == p {
			// 任何情况下，这个都是最短路径
			return tr.dep[v] - tr.dep[u]
		}
		l := kth(v, tr.dep[v]-tr.dep[p]-1)
		r := kth(u, tr.dep[u]-tr.dep[p]-1)

		if g[l][r] {
			return 1 + tr.dep[u] - tr.dep[r] + tr.dep[v] - tr.dep[l]
		}
		return tr.dep[u] + tr.dep[v] - 2*tr.dep[p]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		u, v := cur[0], cur[1]
		ans[i] = query(u, v)
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	dep   []int
	fa    [][]int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.dep = make([]int, n)
	g.fa = make([][]int, n)
	e++
	g.next = make([]int, e)
	g.to = make([]int, e)

	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

func (g *Graph) prepare(root int) (lca func(int, int) int, kth func(int, int) int) {
	n := len(g.nodes)

	h := bits.Len(uint(n))

	dep := g.dep
	fa := g.fa

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(root, root)

	lca = func(u int, v int) int {
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

	kth = func(u int, k int) int {
		for i := h - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = fa[u][i]
			}
		}
		return u
	}

	return
}
