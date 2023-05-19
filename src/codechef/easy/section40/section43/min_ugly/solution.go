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

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			k := readNum(reader)
			Q[i] = readNNums(reader, k)
		}
		res := solve(n, E, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const H = 18
const inf = 1 << 30

func solve(n int, E [][]int, Q [][]int) []int {
	g := buildGraph(n, E)

	in := make([]int, n)
	out := make([]int, n)
	D := make([]int, n)
	fa := make([][]int, n)

	var dfs func(p int, u int)
	var time int
	dfs = func(p int, u int) {
		in[u] = time
		time++
		fa[u] = make([]int, H)
		fa[u][0] = p
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
		out[u] = time
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	stack := make([]int, n)

	active := NewGraph(n, 2*n)
	// 因为时棵树，所以可以用bfs，不用优先队列
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	que := make([]int, n)
	farthest := func(root int) (int, int) {
		var front, end int
		que[end] = root
		end++
		dist[root] = 0
		far := root
		for front < end {
			u := que[front]
			front++
			if dist[u] > dist[far] {
				far = u
			}
			for i := active.node[u]; i > 0; i = active.next[i] {
				v := active.to[i]
				if dist[v] == -1 {
					dist[v] = dist[u] + active.val[i]
					que[end] = v
					end++
				}
			}
		}
		res := dist[far]
		for i := 0; i < end; i++ {
			dist[que[i]] = -1
		}

		return far, res
	}

	find := func(verts []int) int {
		sort.Slice(verts, func(i, j int) bool {
			return in[verts[i]] < in[verts[j]]
		})
		k := len(verts)
		for i := 0; i+1 < k; i++ {
			verts = append(verts, lca(verts[i], verts[i+1]))
		}

		sort.Slice(verts, func(i, j int) bool {
			return in[verts[i]] < in[verts[j]]
		})

		verts = unique(verts)
		var p int
		for _, x := range verts {
			for p > 0 {
				u := stack[p-1]
				if out[u] >= out[x] && u != x {
					// u is parent of x
					break
				}
				p--
			}
			if p > 0 {
				u := stack[p-1]
				active.AddEdge(u, x, D[x]-D[u])
				active.AddEdge(x, u, D[x]-D[u])
			}
			stack[p] = x
			p++
		}
		first, _ := farthest(verts[0])
		_, diam := farthest(first)
		for _, x := range verts {
			active.node[x] = 0
		}
		active.cur = 0
		return (diam + 1) / 2
	}

	ans := make([]int, len(Q))
	for i, cur := range Q {
		for j := 0; j < len(cur); j++ {
			cur[j]--
		}
		ans[i] = find(cur)
	}

	return ans
}

func unique(arr []int) []int {
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] != arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func buildGraph(n int, E [][]int) *Graph {
	g := NewGraph(n, len(E)*2)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, 1)
		g.AddEdge(v, u, 1)
	}
	return g
}

type Graph struct {
	node []int
	next []int
	to   []int
	val  []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.val = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
