package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(reader, 2)
		}
		P := readNNums(reader, n)
		qs := make([][]int, q)
		for i := 0; i < q; i++ {
			qs[i] = readNNums(reader, 3)
		}
		res := solve(n, edges, P, qs)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
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

func solve(n int, edges [][]int, P []int, queries [][]int) []bool {
	g := NewGraph(n+1, 2*n)

	for _, e := range edges {
		g.Add(e[0], e[1])
		g.Add(e[1], e[0])
	}

	pos := make([]int, n+1)

	for i := 0; i < n; i++ {
		pos[P[i]] = i
	}

	qs := make([][]int, n+1)

	for i, cur := range queries {
		v := cur[2]
		if qs[v] == nil {
			qs[v] = make([]int, 0, 1)
		}
		qs[v] = append(qs[v], i)
	}

	ft := NewFenwickTree(n + 5)

	var dfs func(p int, u int)

	ans := make([]bool, len(queries))

	sum := make([]int, len(queries))

	dfs = func(p int, u int) {
		for _, i := range qs[u] {
			l, r := queries[i][0], queries[i][1]
			sum[i] = ft.queryRange(l-1, r-1)
		}

		ft.update(pos[u], 1)

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			dfs(u, v)
		}

		for _, i := range qs[u] {
			l, r := queries[i][0], queries[i][1]
			tmp := ft.queryRange(l-1, r-1)
			ans[i] = tmp > sum[i]
		}
	}

	dfs(0, 1)

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

func (g *Graph) Add(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

type FenwickTree []int

func NewFenwickTree(n int) FenwickTree {
	arr := make([]int, n+1)
	return arr
}

func (t FenwickTree) update(p int, v int) {
	for i := p + 1; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t FenwickTree) query(p int) int {
	var res int
	for i := p + 1; i > 0; i -= i & -i {
		res += t[i]
	}
	return res
}

func (t FenwickTree) queryRange(l int, r int) int {
	res := t.query(r)
	if l > 0 {
		res -= t.query(l - 1)
	}
	return res
}
