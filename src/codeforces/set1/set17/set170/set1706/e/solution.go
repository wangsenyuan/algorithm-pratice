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
		n, m, q := readThreeNums(reader)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 2)
		}
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, edges, Q)
		for i := 0; i < q; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

const H = 20

func solve(n int, edges [][]int, Q [][]int) []int {
	set := NewUF(n)

	F := make([]int, n)

	groups := make([][]int, n)
	for i := 0; i < n; i++ {
		groups[i] = make([]int, 0, 1)
		groups[i] = append(groups[i], i)
	}

	for i, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		pu := set.Find(u)
		pv := set.Find(v)
		if pu == pv {
			continue
		}
		// will connect
		if set.sz[pu] < set.sz[pv] {
			pu, pv = pv, pu
		}
		// sz[pu] > sz[pv]
		for _, x := range groups[pv] {
			if x > 0 && set.Find(x-1) == pu {
				F[x] = i + 1
			}
			if x+1 < n && set.Find(x+1) == pu {
				F[x+1] = i + 1
			}
		}

		set.sz[pu] += set.sz[pv]
		set.arr[pv] = pu
		groups[pu] = append(groups[pu], groups[pv]...)
	}

	return findAns(n, F, Q)
}

func solve1(n int, edges [][]int, Q [][]int) []int {

	g := NewGraph(n, 2*(n-1))

	set := NewUF(n)

	for i, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		if set.Union(u, v) {
			g.AddEdge(u, v, i+1)
			g.AddEdge(v, u, i+1)
		}
	}

	D := make([]int, n)
	P := make([][]int, n)
	C := make([][]int, n)
	var dfs func(p, u, w int)

	dfs = func(p, u, w int) {
		P[u] = make([]int, H)
		C[u] = make([]int, H)
		P[u][0] = p
		C[u][0] = w
		for i := 1; i < H; i++ {
			x := P[u][i-1]
			P[u][i] = P[x][i-1]
			C[u][i] = max(C[u][i-1], C[x][i-1])
		}

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v, g.cost[i])
			}
		}
	}

	dfs(0, 0, 0)

	find_cost := func(u int, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		var res int
		for i := H - 1; i >= 0; i-- {
			if (D[u] - (1 << i)) >= D[v] {
				res = max(res, C[u][i])
				u = P[u][i]
			}
		}
		if u == v {
			return res
		}

		for i := H - 1; i >= 0; i-- {
			if P[u][i] != P[v][i] {
				res = max(res, C[u][i])
				res = max(res, C[v][i])
				u = P[u][i]
				v = P[v][i]
			}
		}
		return max(max(res, C[u][0]), C[v][0])
	}

	// F[i] = max_cost to connect i & i - 1
	F := make([]int, n)

	for i := 1; i < n; i++ {
		F[i] = find_cost(i-1, i)
	}

	return findAns(n, F, Q)
}

func findAns(n int, F []int, Q [][]int) []int {
	arr := make([]int, 2*n)

	copy(arr[n:], F)

	for i := n - 1; i > 0; i-- {
		arr[i] = max(arr[i*2], arr[i*2+1])
	}

	get := func(l, r int) int {
		l += n
		r += n
		var res int

		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	ans := make([]int, len(Q))
	for i := 0; i < len(Q); i++ {
		l, r := Q[i][0]-1, Q[i][1]
		ans[i] = get(l+1, r)
	}
	return ans
}

type Graph struct {
	node []int
	next []int
	to   []int
	cost []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	g.cost = make([]int, e+1)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
	g.cost[g.cur] = w
}

type UF struct {
	arr []int
	sz  []int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		sz[i]++
		arr[i] = i
	}

	return &UF{arr, sz}
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.sz[pa] < uf.sz[pb] {
		pa, pb = pb, pa
	}

	uf.sz[pa] += uf.sz[pb]
	uf.arr[pb] = pa
	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
