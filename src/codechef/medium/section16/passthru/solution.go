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
		n := readNum(reader)
		A := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}

		res := solve(n, E, A)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(n int, E [][]int, A []int) []int64 {

	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	in := make([]int, n)
	out := make([]int, n)
	D := make([]int, n)
	P := make([][]int, n)

	var dfs func(p int, u int)
	var time int

	verts := make(map[int][]int)

	addVert := func(c int, u int) {
		if verts[c] == nil {
			verts[c] = make([]int, 0, 1)
		}
		verts[c] = append(verts[c], u)
	}

	dfs = func(p int, u int) {
		in[u] = time
		time++
		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		addVert(A[u], u)

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

	ans := make([]int64, n)

	update := func(u int, p int, c int) {
		ans[u] += int64(c)
		ans[p] -= int64(c)
	}

	stack := make([]int, n)

	for c, vs := range verts {
		sort.Slice(vs, func(i, j int) bool {
			return in[vs[i]] < in[vs[j]]
		})

		k := len(vs)

		for i := 0; i+1 < k; i++ {
			p := lca(vs[i], vs[i+1])
			vs = append(vs, p)
		}

		vs = unique(vs)

		sort.Slice(vs, func(i, j int) bool {
			return in[vs[i]] < in[vs[j]]
		})

		var id int

		for _, v := range vs {
			for id > 0 {
				u := stack[id-1]
				if out[u] >= out[v] && u != v {
					// u is parent of v
					break
				}
				id--
			}
			if id > 0 {
				p := stack[id-1]
				update(v, p, c)
			}

			stack[id] = v
			id++
		}
	}

	var dfs1 func(p int, u int)

	dfs1 = func(p int, u int) {
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v)
				ans[u] += ans[v]
			}
		}
	}

	dfs1(-1, 0)

	res := make([]int64, len(E))

	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if in[u] > in[v] {
			v = u
		}
		res[i] = ans[v]
	}

	return res
}

func unique(arr []int) []int {
	sort.Ints(arr)
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
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
