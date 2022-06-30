package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 2)
	}
	X := make([]int, m)
	Y := make([]int, m)
	S := make([]string, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &X[i]) + 1
		pos = readInt(s, pos, &Y[i]) + 1
		S[i] = string(s[pos:])
	}
	res := solve(n, E, X, Y, S)

	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(res)
	}
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

func solve(n int, E [][]int, X []int, Y []int, S []string) string {

	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	P := make([]int, n)
	H := make([]int, n)
	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		P[u] = p
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				H[v] = H[u] + 1
				dfs1(u, v)
			}
		}
	}

	dfs1(0, 0)

	left := make([]int, n)
	right := make([]int, n)

	get_path := func(u, v int) []int {
		var l, r int

		for u != v {
			if H[u] > H[v] {
				left[l] = u
				l++
				u = P[u]
			} else {
				right[r] = v
				r++
				v = P[v]
			}
		}
		left[l] = u
		l++
		for r > 0 {
			left[l] = right[r-1]
			r--
			l++
		}
		res := make([]int, l)
		copy(res, left)
		return res
	}

	paths := make([][]int, len(X))
	opts := make([]Assign, n)

	for i := 0; i < n; i++ {
		opts[i] = Assign{'a', 'a', false}
	}

	m := len(X)

	for i := 0; i < m; i++ {
		paths[i] = get_path(X[i]-1, Y[i]-1)
		k := len(paths[i])
		for j := 0; j < len(paths[i]); j++ {
			opts[paths[i][j]] = Assign{S[i][j], S[i][k-1-j], true}
		}
	}
	var edges int
	for i := 0; i < m; i++ {
		k := len(paths[i])

		for j := 0; j < k; j++ {
			u := paths[i][j]
			c := S[i][j]
			rc := S[i][k-1-j]
			d := opts[u].first
			rd := opts[u].second
			if d != c {
				edges++
			}
			if d != rc {
				edges++
			}
			if rd != c {
				edges++
			}
			if rd != rc {
				edges++
			}
		}
	}

	nm := (n + m) * 2

	ga := NewGraph(nm, edges*2)
	gb := NewGraph(nm, edges*2)

	add_edge := func(u int, vx int, v int, vy int) {
		// (val[u] == vx) -> (val[v] == vy)
		// a -> b and !b -> !a
		ga.AddEdge(u*2+vx, v*2+vy)
		gb.AddEdge(v*2+vy, u*2+vx)
		ga.AddEdge(v*2+1-vy, u*2+1-vx)
		gb.AddEdge(u*2+1-vx, v*2+1-vy)
	}

	for i := 0; i < m; i++ {
		k := len(paths[i])

		for j := 0; j < k; j++ {
			u := paths[i][j]
			c := S[i][j]
			rc := S[i][k-1-j]
			d := opts[u].first
			rd := opts[u].second
			if d != c {
				add_edge(u, 0, n+i, 1)
			}
			if d != rc {
				add_edge(u, 0, n+i, 0)
			}
			if rd != c {
				add_edge(u, 1, n+i, 1)
			}
			if rd != rc {
				add_edge(u, 1, n+i, 0)
			}
		}
	}

	used := make([]bool, nm)
	ord := make([]int, 0, nm)

	var dfs2 func(u int)

	dfs2 = func(u int) {
		used[u] = true
		for i := ga.node[u]; i > 0; i = ga.next[i] {
			v := ga.to[i]
			if !used[v] {
				dfs2(v)
			}
		}
		ord = append(ord, u)
	}

	for u := 0; u < nm; u++ {
		if !used[u] {
			dfs2(u)
		}
	}

	comp := make([]int, nm)
	for i := 0; i < nm; i++ {
		comp[i] = -1
	}

	var dfs3 func(u int, c int)

	dfs3 = func(u int, c int) {
		comp[u] = c
		for i := gb.node[u]; i > 0; i = gb.next[i] {
			v := gb.to[i]
			if comp[v] < 0 {
				dfs3(v, c)
			}
		}
	}
	var ctr int
	for i := nm - 1; i >= 0; i-- {
		if comp[ord[i]] < 0 {
			dfs3(ord[i], ctr)
			ctr++
		}
	}

	for i := 0; i < nm; i++ {
		if comp[i] == comp[i^1] {
			return ""
		}
	}

	buf := make([]byte, n)

	for i := 0; i < 2*n; i += 2 {
		if !opts[i/2].assign {
			buf[i/2] = 'a'
		} else if comp[i] > comp[i^1] {
			buf[i/2] = opts[i/2].first
		} else {
			buf[i/2] = opts[i/2].second
		}
	}
	return string(buf)
}

type Assign struct {
	first  byte
	second byte
	assign bool
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
