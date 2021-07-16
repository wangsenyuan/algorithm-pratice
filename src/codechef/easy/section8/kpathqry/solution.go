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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		q := readNum(reader)
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			s, _ := reader.ReadBytes('\n')
			var k int
			pos := readInt(s, 0, &k)
			Q[i] = make([]int, k)
			for j := 0; j < k; j++ {
				pos = readInt(s, pos+1, &Q[i][j])
			}
		}
		res := solve(n, E, Q)
		for i := 0; i < len(res); i++ {
			if res[i] {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
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

func solve(n int, E [][]int, Q [][]int) []bool {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0], e[1]
		g.AddEdge(u-1, v-1)
		g.AddEdge(v-1, u-1)
	}

	in := make([]int, n)
	out := make([]int, n)
	H := make([][]int, n)
	for i := 0; i < n; i++ {
		H[i] = make([]int, 17)
	}
	var time int
	var dfs func(p, u int)
	dfs = func(p, u int) {
		in[u] = time
		time++

		H[u][0] = p

		for i := 1; i < 17; i++ {
			H[u][i] = H[H[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
			}
		}
		// make sure time unique for every node
		time++
		out[u] = time
	}

	dfs(0, 0)

	isAncestor := func(u, v int) bool {
		// u is ancestor of v?
		return in[u] <= in[v] && out[v] <= out[u]
	}

	P := make([]Pair, n)

	lca := func(u, v int) int {
		if isAncestor(u, v) {
			return u
		}
		if isAncestor(v, u) {
			return v
		}
		// D[u] >= D[v]
		for i := 16; i >= 0; i-- {
			if !isAncestor(H[u][i], v) {
				u = H[u][i]
			}
		}
		return H[u][0]
	}

	check := func(r int) bool {
		for i := 0; i < len(Q[r]); i++ {
			P[i] = Pair{out[Q[r][i]-1], Q[r][i] - 1}
		}
		sort.Sort(Pairs(P[:len(Q[r])]))
		a := P[0].second
		// sort by out time increasing
		var pivot int

		for pivot+1 < len(Q[r]) && isAncestor(P[pivot+1].second, P[pivot].second) {
			pivot++
		}

		if pivot == len(Q[r])-1 {
			return true
		}

		// first path until pivot
		pivot++

		b := P[pivot].second

		for pivot+1 < len(Q[r]) && isAncestor(P[pivot+1].second, P[pivot].second) {
			pivot++
		}

		if pivot < len(Q[r])-1 {
			return false
		}

		l := lca(a, b)

		return isAncestor(l, P[pivot].second)
	}

	ans := make([]bool, len(Q))

	for i := 0; i < len(Q); i++ {
		ans[i] = check(i)
	}
	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.nodes = make([]int, n)
	g.next = make([]int, e+3)
	g.to = make([]int, e+3)
	g.cur = 0
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
