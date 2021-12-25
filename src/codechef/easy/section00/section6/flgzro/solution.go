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
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		A := readNNums(reader, n)
		res := solve(n, A, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1000000007

func solve(n int, A []int, E [][]int) int {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	B := make([]int, n)
	copy(B, A)
	sort.Ints(B)
	var m int
	for i := 1; i <= n; i++ {
		if i == n || B[i] > B[i-1] {
			B[m] = B[i-1]
			m++
		}
	}
	B = B[:m]

	for i := 0; i < n; i++ {
		A[i] = sort.Search(m, func(j int) bool {
			return B[j] >= A[i]
		})
	}

	bit := make([]int64, m+1)
	update := func(p int, v int64) {
		p++
		for p <= m {
			bit[p] = modAdd(bit[p], v)
			p += p & -p
		}
	}

	get := func(p int) int64 {
		p++
		var res int64
		for p > 0 {
			res = modAdd(res, bit[p])
			p -= p & -p
		}
		return res
	}

	var dfs func(p, u int)

	var res int64

	dfs = func(p, u int) {
		cur := (get(A[u]) + MOD - get(A[u]-1)) % MOD
		cur = modAdd(get(m-1), MOD-cur)
		cur = modAdd(cur, 1)

		update(A[u], cur)

		var isLeaf = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			isLeaf = false
			dfs(u, v)
		}
		if isLeaf {
			res = modAdd(res, cur)
		}
		update(A[u], MOD-cur)
	}

	dfs(-1, 0)

	return int(res)
}

func solve1(n int, A []int, E [][]int) int {
	g := NewGraph(n, len(E)*2)

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	B := make([]int, n)
	copy(B, A)
	sort.Ints(B)
	var m int
	for i := 1; i <= n; i++ {
		if i == n || B[i] > B[i-1] {
			B[m] = B[i-1]
			m++
		}
	}
	B = B[:m]

	for i := 0; i < n; i++ {
		A[i] = sort.Search(m, func(j int) bool {
			return B[j] >= A[i]
		})
	}

	size := make([]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		size[u]++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				size[u] += size[v]
			}
		}
	}

	dfs(-1, 0)

	value := make([]int64, m)
	ways := make([]int64, n)
	var total int64

	var add func(p, u int, mul int64)

	add = func(p, u int, mul int64) {
		total = modAdd(total, modMul(ways[u], mul))
		value[A[u]] = modAdd(value[A[u]], modMul(ways[u], mul))
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				add(u, v, mul)
			}
		}
	}

	var dfs2 func(p, u int, keep bool)
	dfs2 = func(p, u int, keep bool) {
		big := -1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p {
				if big < 0 || size[v] > size[big] {
					big = v
				}
			}
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p && v != big {
				dfs2(u, v, false)
			}
		}

		if big != - 1 {
			dfs2(u, big, true)
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != p && v != big {
				add(u, v, 1)
			}
		}

		if big < 0 {
			// a leaf node
			ways[u] = 1
		} else {
			ways[u] = modAdd(total, MOD-value[A[u]])
		}

		total = modAdd(total, ways[u])
		value[A[u]] = modAdd(value[A[u]], ways[u])

		if !keep {
			add(p, u, -1)
		}
	}

	dfs2(-1, 0, true)

	var ans int64
	for i := 0; i < n; i++ {
		ans = modAdd(ans, ways[i])
	}

	return int(ans)
}

func modAdd(a, b int64) int64 {
	a += b
	a += MOD
	return a % MOD
}

func modMul(a, b int64) int64 {
	a *= b
	return a % MOD
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
