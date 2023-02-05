package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	edges := make([][]int, m)

	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}

	res := solve(n, edges)

	fmt.Println(res)
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

const H = 22

func solve(n int, edges [][]int) string {

	g := NewGraph(n, 2*(n-1))

	uf := NewUF(n)

	var es []Edge

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		if !uf.Union(u, v) {
			// cross edge
			es = append(es, Edge{u, v})
		} else {
			g.AddEdge(u, v)
			g.AddEdge(v, u)
		}
	}

	D := make([]int, n)
	P := make([][]int, n)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	lca := func(u int, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if (D[u] - (1 << i)) >= D[v] {
				u = P[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				u = P[u][i]
				v = P[v][i]
			}
		}
		return P[u][0]
	}

	jump := func(x int, t int) int {
		for i := 0; t > 0; t, i = t>>1, i+1 {
			if t&1 == 1 {
				x = P[x][i]
			}
		}
		return x
	}

	vis := make([]int, n)

	for i := 0; i < len(es); i++ {
		u, v := es[i].u, es[i].v

		if D[u] > D[v] {
			u, v = v, u
		}
		// u is higher
		w := lca(u, v)
		if u == w {
			vis[v]--
			vis[jump(v, D[v]-D[u]-1)]++
		} else {
			vis[0]++
			vis[u]--
			vis[v]--
		}
	}

	res := make([]byte, n)

	var push func(p int, u int)

	push = func(p int, u int) {
		if vis[u] == 0 {
			res[u] = '1'
		} else {
			res[u] = '0'
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				vis[v] += vis[u]
				push(u, v)
			}
		}
	}

	push(0, 0)

	return string(res)
}

type Edge struct {
	u int
	v int
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
	g.to = make([]int, e+1)
	g.next = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

type UF struct {
	arr   []int
	count []int
	size  int
}

func NewUF(n int) UF {
	arr := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		count[i]++
	}

	return UF{arr, count, n}
}

func (uf *UF) Find(a int) int {
	if uf.arr[a] != a {
		uf.arr[a] = uf.Find(uf.arr[a])
	}
	return uf.arr[a]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.count[pa] < uf.count[pb] {
		pa, pb = pb, pa
	}
	uf.count[pa] += uf.count[pb]
	uf.arr[pb] = pa
	uf.size--
	return true
}
