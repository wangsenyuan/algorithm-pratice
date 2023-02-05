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

	var buf bytes.Buffer

	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	q := readNum(reader)
	queries := make([][]int, q)

	for i := 0; i < q; i++ {
		k := readNum(reader)
		queries[i] = readNNums(reader, k)
	}
	res := solve(n, edges, queries)
	for _, ans := range res {
		if ans {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(n int, edges [][]int, queries [][]int) []bool {

	g := NewGraph(n, 2*(n-1))

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)

	in := make([]int, n)
	out := make([]int, n)

	P := make([][]int, n)
	for i := 0; i < n; i++ {
		P[i] = make([]int, H)
	}

	var dfs func(p, u int)
	var time int
	dfs = func(p, u int) {
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}
		in[u] = time
		time++
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				D[v] = D[u] + 1
				dfs(u, v)
			}
		}
		out[u] = time
		time++
	}

	dfs(0, 0)

	isAncestor := func(u, v int) bool {
		return in[u] <= in[v] && out[v] <= out[u]
	}

	lca := func(u, v int) int {
		if D[u] < D[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if D[u]-(1<<i) >= D[v] {
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

	arr := make([]int, n)

	check := func(query []int) bool {
		if len(query) <= 2 {
			return true
		}
		m := len(query)

		for i := 0; i < m; i++ {
			query[i]--
		}
		sort.Slice(query, func(i, j int) bool {
			a := query[i]
			b := query[j]
			return D[a] < D[b]
		})

		var it int

		x := query[m-1]

		for i := m - 2; i >= 0; i-- {
			y := query[i]
			if isAncestor(y, x) {
				x = y
			} else {
				arr[it] = y
				it++
			}
		}
		if it == 0 {
			return true
		}
		// last one is x
		u := arr[0]

		for i := 1; i < it; i++ {
			v := arr[i]
			if !isAncestor(v, u) {
				return false
			}
			u = v
		}

		p1 := lca(arr[0], query[m-1])
		p2 := lca(u, x)

		return p1 == p2
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur)
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
